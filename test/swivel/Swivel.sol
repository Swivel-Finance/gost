// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.4;

import './Utils/Abstracts.sol';
import './Utils/Hash.sol';
import './Utils/Sig.sol';

contract Swivel {
  /// @dev maps the key of an order to a boolean indicating if an order was cancelled
  mapping (bytes32 => bool) public cancelled;
  /// @dev maps the key of an order to an amount representing its taken volume
  mapping (bytes32 => uint256) public filled;
  /// @dev maps a token address to a point in time, a hold, after which a withdrawal can be made
  mapping (address => uint256) public withdrawals;

  string constant public NAME = 'Swivel Finance';
  string constant public VERSION = '2.0.0';
  uint256 constant public HOLD = 259200; // obvs could be a smaller uint but packing?
  bytes32 public immutable domain;
  address public immutable marketPlace;
  address public admin;
  /// @dev holds the fee demoninators for [zcTokenInitiate, zcTokenExit, vaultInitiate, vaultExit]
  uint16[4] public fenominator;

  /// @notice Emitted on order cancellation
  event Cancel (bytes32 indexed key, bytes32 hash);
  /// @notice Emitted on any initiate*
  /// @dev filled is 'principalFilled' when (vault:false, exit:false) && (vault:true, exit:true)
  /// @dev filled is 'premiumFilled' when (vault:true, exit:false) && (vault:false, exit:true)
  event Initiate(bytes32 indexed key, bytes32 hash, address indexed maker, bool vault, bool exit, address indexed sender, uint256 amount, uint256 filled);
  /// @notice Emitted on any exit*
  /// @dev filled is 'principalFilled' when (vault:false, exit:false) && (vault:true, exit:true)
  /// @dev filled is 'premiumFilled' when (vault:true, exit:false) && (vault:false, exit:true)
  event Exit(bytes32 indexed key, bytes32 hash, address indexed maker, bool vault, bool exit, address indexed sender, uint256 amount, uint256 filled);
  /// @notice Emitted on token withdrawal scheduling
  /// @dev token is the address of the token scheduled for withdrawal
  /// @dev withdrawalTime is the timestamp at which the queued withdrawal will be possible
  event WithdrawalScheduled (address indexed token, uint256 hold);

  /// @param m deployed MarketPlace contract address
  constructor(address m) {
    admin = msg.sender;
    domain = Hash.domain(NAME, VERSION, block.chainid, address(this));
    marketPlace = m;
    fenominator = [200, 600, 400, 200];
  }

  // ********* INITIATING *************

  /// @notice Allows a user to initiate a position
  /// @param o Array of offline Swivel.Orders
  /// @param a Array of order volume (principal) amounts relative to passed orders
  /// @param c Array of Components from valid ECDSA signatures
  function initiate(Hash.Order[] calldata o, uint256[] calldata a, Sig.Components[] calldata c) external returns (bool) {
    unit256 length = o.length; 
    for (uint256 i; i < length; i++)
      Hash.Order memory order = o[i];
      if (!order.exit) {
        if (!order.vault) {
          initiateVaultFillingZcTokenInitiate(o[i], a[i], c[i]);
        } else {
          initiateZcTokenFillingVaultInitiate(o[i], a[i], c[i]);
        }
      } else {
        if (!order.vault) {
          initiateZcTokenFillingZcTokenExit(o[i], a[i], c[i]);
        } else {
          initiateVaultFillingVaultExit(o[i], a[i], c[i]);
        }
      }
    }

    return true;
  }

  /// @notice Allows a user to initiate a Vault by filling an offline zcToken initiate order
  /// @dev This method should pass (underlying, maturity, maker, sender, principalFilled) to MarketPlace.custodialInitiate
  /// @param o The order being filled
  /// @param a Amount of volume (premium) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function initiateVaultFillingZcTokenInitiate(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal {
    bytes32 hash = validOrderHash(o, c);
    // checks the side, and the amount compared to amount available
    require(a <= (o.premium - filled[hash]), 'taker amount > available volume');

    filled[hash] += a;

    uint256 principalFilled = a * o.principal / o.premium;
    uint256 fee = principalFilled / fenominator[2];

    Erc20 uToken = Erc20(o.underlying);
    uToken.transferFrom(msg.sender, o.maker, a);
    uToken.transferFrom(o.maker, address(this), principalFilled);

    MarketPlace mPlace = MarketPlace(marketPlace);
    // mint cTokens
    require(CErc20(mPlace.cTokenAddress(o.underlying, o.maturity)).mint(principalFilled) == 0, 'minting CToken failed');

    // alert MarketPlace.
    require(mPlace.custodialInitiate(o.underlying, o.maturity, o.maker, msg.sender, principalFilled), 'custodial initiate failed');
    // transfer fee in vault notional to swivel (from msg.sender)
    require(mPlace.transferVaultNotionalFee(o.underlying, o.maturity, msg.sender, fee), "notional fee transfer failed");

    emit Initiate(o.key, hash, o.maker, o.vault, o.exit, msg.sender, a, principalFilled);
  }

  /// @notice Allows a user to initiate a zcToken by filling an offline vault initiate order
  /// @dev This method should pass (underlying, maturity, sender, maker, a) to MarketPlace.custodialInitiate
  /// @param o The order being filled
  /// @param o Amount of volume (principal) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function initiateZcTokenFillingVaultInitiate(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal {
    bytes32 hash = validOrderHash(o, c);
    // checks the side, and the amount compared to amount available
    require((a <= o.principal - filled[hash]), 'taker amount > available volume');

    filled[hash] += a;

    uint256 premiumFilled = a * o.premium / o.principal;
    uint256 fee = premiumFilled / fenominator[0];

    Erc20 uToken = Erc20(o.underlying);
    uToken.transferFrom(o.maker, msg.sender, premiumFilled);
    // transfer principal + fee in underlying to swivel (from sender)
    uToken.transferFrom(msg.sender, address(this), (a + fee));
    
    MarketPlace mPlace = MarketPlace(marketPlace);
    // mint cTokens
    require(CErc20(mPlace.cTokenAddress(o.underlying, o.maturity)).mint(a) == 0, 'minting CToken Failed');
    // alert MarketPlace
    require(mPlace.custodialInitiate(o.underlying, o.maturity, msg.sender, o.maker, a), 'custodial initiate failed');

    emit Initiate(o.key, hash, o.maker, o.vault, o.exit, msg.sender, a, premiumFilled);
  }

  /// @notice Allows a user to initiate zcToken? by filling an offline zcToken exit order
  /// @dev This method should pass (underlying, maturity, maker, sender, a) to MarketPlace.p2pZcTokenExchange
  /// @param o The order being filled
  /// @param a Amount of volume (principal) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function initiateZcTokenFillingZcTokenExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal {
    bytes32 hash = validOrderHash(o, c);
    // checks the side, and the amount compared to amount available
    require(a <= ((o.principal - filled[hash])), 'taker amount > available volume');

    filled[hash] += a;

    // .interest is interest * ratio / 1e18 where ratio is (a * 1e18) / principal
    uint256 premiumFilled = a * o.premium / o.principal;
    uint256 fee = premiumFilled / fenominator[0];

    // transfer principal - the premium paid + fee in underliyng to swivel (from sender)
    Erc20(o.underlying).transferFrom(msg.sender, o.maker, ((a - premiumFilled) + fee));
    // notify the marketplace...
    MarketPlace(marketPlace).p2pZcTokenExchange(o.underlying, o.maturity, o.maker, msg.sender, a);
            
    emit Initiate(o.key, hash, o.maker, o.vault, o.exit, msg.sender, a, premiumFilled);
  }

  /// @notice Allows a user to initiate a Vault by filling an offline vault exit order
  /// @dev This method should pass (underlying, maturity, maker, sender, principalFilled) to MarketPlace.p2pVaultExchange
  /// @param o The order being filled
  /// @param a Amount of volume (interest) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function initiateVaultFillingVaultExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal {
    bytes32 hash = validOrderHash(o, c);
    // checks the side, and the amount compared to amount available
    require(a <= (o.premium - filled[hash]), 'taker amount > available volume');
    
    filled[hash] += a;

    Erc20(o.underlying).transferFrom(msg.sender, o.maker, a);

    uint256 principalFilled = a * o.principal / o.premium;
    uint256 fee = principalFilled / fenominator[2];

    // notify marketplace
    MarketPlace mPlace = MarketPlace(marketPlace);
    require(mPlace.p2pVaultExchange(o.underlying, o.maturity, o.maker, msg.sender, principalFilled), 'vault exchange failed');
    // transfer fee in vault notional to swivel (from msg.sender)
    require(mPlace.transferVaultNotionalFee(o.underlying, o.maturity, msg.sender, fee), "notional fee transfer failed");

    emit Initiate(o.key, hash, o.maker, o.vault, o.exit, msg.sender, a, principalFilled);
  }

  // ********* EXITING ***************

  /// @notice Allows a user to exit (sell) a currently held position to the marketplace.
  /// @param o Array of offline Swivel.Orders
  /// @param a Array of order volume (principal) amounts relative to passed orders
  /// @param c Components of a valid ECDSA signature
  function exit(Hash.Order[] calldata o, uint256[] calldata a, Sig.Components[] calldata c) external returns (bool) {
    unit256 length = o.length; 
    for (uint256 i; i < length; i++)
      Hash.Order memory order = o[i];
      // determine whether the order being filled is an exit
      if (!order.exit) {
        // determine whether the order being filled is a vault initiate or a zcToken initiate
          if (!order.vault) {
            // if filling a zcToken initiate with an exit, one is exiting zcTokens
            exitZcTokenFillingZcTokenInitiate(o[i], a[i], c[i]);
          } else {
            // if filling a vault initiate with an exit, one is exiting vault notional
            exitVaultFillingVaultInitiate(o[i], a[i], c[i]);
          }
      } else {
        // determine whether the order being filled is a vault exit or zcToken exit
        if (!order.vault) {
          // if filling a zcToken exit with an exit, one is exiting vault
          exitVaultFillingZcTokenExit(o[i], a[i], c[i]);
        } else {
          // if filling a vault exit with an exit, one is exiting zcTokens
          exitZcTokenFillingVaultExit(o[i], a[i], c[i]);
        }   
      }   
    }

    return true;
  }

  /// @notice Allows a user to exit their zcTokens by filling an offline zcToken initiate order
  /// @dev This method should pass (underlying, maturity, sender, maker, principalFilled) to MarketPlace.p2pZcTokenExchange
  /// @param o The order being filled
  /// @param a Amount of volume (interest) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function exitZcTokenFillingZcTokenInitiate(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal {
    bytes32 hash = validOrderHash(o, c);

    require(a <= (o.premium - filled[hash]), 'taker amount > available volume');
    
    filled[hash] += a;       

    uint256 principalFilled = a * o.principal / o.premium;
    uint256 fee = principalFilled / fenominator[1];

    Erc20 uToken = Erc20(o.underlying);
    // transfer underlying from initiating party to exiting party, minus the price the exit party pays for the exit (premium), and the fee.
    uToken.transferFrom(o.maker, msg.sender, principalFilled - a - fee);
    // notify marketplace...
    MarketPlace(marketPlace).p2pZcTokenExchange(o.underlying, o.maturity, msg.sender, o.maker, principalFilled);
    // transfer fee in underlying to swivel
    uToken.transferFrom(o.maker, address(this), fee);
    
    emit Exit(o.key, hash, o.maker, o.vault, o.exit, msg.sender, a, principalFilled);
  }
  
  /// @notice Allows a user to exit their Vault by filling an offline vault initiate order
  /// @dev This method should pass (underlying, maturity, sender, maker, a) to MarketPlace.p2pVaultExchange
  /// @param o The order being filled
  /// @param a Amount of volume (principal) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function exitVaultFillingVaultInitiate(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal {
    bytes32 hash = validOrderHash(o, c);

    require(a <= (o.principal - filled[hash]), 'taker amount > available volume');
    
    filled[hash] += a;
        
    uint256 premiumFilled = a * o.premium / o.principal;
    uint256 fee = premiumFilled / fenominator[3];

    Erc20 uToken = Erc20(o.underlying);
    // transfer premium minus fee from maker to sender
    uToken.transferFrom(o.maker, msg.sender, premiumFilled - fee);
    // market should transfer <a> notional from sender to maker
    require(MarketPlace(marketPlace).p2pVaultExchange(o.underlying, o.maturity, msg.sender, o.maker, a), 'vault exchange failed');
    // transfer fee in underlying to swivel from sender
    uToken.transferFrom(msg.sender, address(this), fee);

    emit Exit(o.key, hash, o.maker, o.vault, o.exit, msg.sender, a, premiumFilled);
  }

  /// @notice Allows a user to exit their Vault filling an offline zcToken exit order
  /// @dev This method should pass (underlying, maturity, maker, sender, a) to MarketPlace.exitFillingExit
  /// @param o The order being filled
  /// @param a Amount of volume (principal) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function exitVaultFillingZcTokenExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal {
    bytes32 hash = validOrderHash(o, c);

    require(a <= (o.principal - filled[hash]), 'taker amount > available volume');
    
    filled[hash] += a;

    uint256 premiumFilled = a * o.premium / o.principal;
    uint256 fee = premiumFilled / fenominator[3];
    
    MarketPlace mPlace = MarketPlace(marketPlace);
    // alert MarketPlace...
    mPlace.custodialExit(o.underlying, o.maturity, o.maker, msg.sender, a);

    // redeem principal from compound now that coupon and zcb have been burned
    address cTokenAddr = mPlace.cTokenAddress(o.underlying, o.maturity);
    require((CErc20(cTokenAddr).redeemUnderlying(a) == 0), "compound redemption error");

    Erc20 uToken = Erc20(o.underlying);
    // transfer principal-premium back to fixed exit party now that the interest coupon and zcb have been redeemed
    uToken.transfer(o.maker, a - premiumFilled);
    // transfer premium-fee to floating exit party
    uToken.transfer(msg.sender, premiumFilled - fee);

    emit Exit(o.key, hash, o.maker, o.vault, o.exit, msg.sender, a, premiumFilled);
  }

  /// @notice Allows a user to exit their zcTokens by filling an offline vault exit order
  /// @dev This method should pass (underlying, maturity, sender, maker, principalFilled) to MarketPlace.exitFillingExit
  /// @param o The order being filled
  /// @param a Amount of volume (interest) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function exitZcTokenFillingVaultExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal {
    bytes32 hash = validOrderHash(o, c);

    require(a <= (o.premium - filled[hash]), 'taker amount > available volume');
    
    filled[hash] += a;

    uint256 principalFilled = a * o.principal / o.premium;
    uint256 fee = principalFilled / fenominator[1];

    MarketPlace mPlace = MarketPlace(marketPlace);
    // inform MarketPlace what happened...
    mPlace.custodialExit(o.underlying, o.maturity, msg.sender, o.maker, principalFilled);

    // redeem principal from compound now that coupon and zcb have been burned
    address cTokenAddr = mPlace.cTokenAddress(o.underlying, o.maturity);
    require((CErc20(cTokenAddr).redeemUnderlying(principalFilled) == 0), "compound redemption error");

    Erc20 uToken = Erc20(o.underlying);
    // transfer principal-premium-fee back to fixed exit party now that the interest coupon and zcb have been redeemed
    uToken.transfer(msg.sender, principalFilled - a - fee);
    // transfer premium to floating exit party
    uToken.transfer(o.maker, a);

    emit Exit(o.key, hash, o.maker, o.vault, o.exit, msg.sender, a, principalFilled);
  }

  /// @notice Allows a user to cancel an order, preventing it from being filled in the future
  /// @param o An offline Swivel.Order
  /// @param c Components of a valid ECDSA signature
  function cancel(Hash.Order calldata o, Sig.Components calldata c) external returns (bool) {
    bytes32 hash = validOrderHash(o, c);

    require(msg.sender == o.maker, 'sender must be maker');

    cancelled[hash] = true;

    emit Cancel(o.key, hash);

    return true;
  }

  // ********* ADMINISTRATIVE ***************

  /// @notice Allows the admin to schedule the withdrawal of tokens
  /// @param e Address of token to withdraw
  function scheduleWithdrawal(address e) external onlyAdmin(admin) {
    uint256 when = block.timestamp + HOLD;
    withdrawals[e] = when;
    emit WithdrawalScheduled(e, when);
  }

  /// @notice Emergency function to block unplanned withdrawals
  /// @param e Address of token withdrawal to block
  function blockWithdrawal(address e) external onlyAdmin(admin) {
      withdrawals[e] = 0;
  }

  /// @notice Allows the admin to withdraw the given token, provided the holding period has been observed
  /// @param e Address of token to withdraw
  function withdraw(address e) external onlyAdmin(admin) {
    uint256 when = withdrawals[e];
    require (when != 0, 'no withdrawal scheduled');
    require (block.timestamp >= when, 'withdrawal still on hold');

    withdrawals[e] = 0;

    Erc20 token = Erc20(e);
    token.transfer(admin, token.balanceOf(address(this)));
  }

  /// @notice Allows the admin to set a new fee denominator
  /// @param t The index of the new fee denominator
  /// @param d The new fee denominator
  function setFee(uint16 t, uint16 d) external onlyAdmin(admin) returns (bool) {
    fenominator[t] = d;
    return true;
  }

  /// @notice Allows the admin to transfer ownership
  /// @param t Address of the new admin
  function transferAdmin(address t) external onlyAdmin(admin) returns (bool) {
    admin = t;
    return true;
  }

  /// @notice Allows the admin to mass approve all compound markets, saving marginal approvals
  /// @param u The underlying token address
  /// @param c The compound cToken address
  function approveCompound(address[] calldata u, address[] calldata c) external onlyAdmin(admin) returns (bool) {
    for (uint256 i; i < u.length; i++) {
      Erc20 uToken = Erc20(u[i]);
      uToken.approve(c[i], 2**256 - 1);
    }
    return true;
  }

  // ********* PROTOCOL UTILITY ***************

  /// @notice Allows users to deposit underlying and in the process split it into/mint 
  /// zcTokens and vault notional. Calls mPlace.mintZcTokenAddingNotional
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param a Amount of underlying being deposited
  function splitUnderlying(address u, uint256 m, uint256 a) external returns (bool) {
    Erc20 uToken = Erc20(u);
    uToken.transferFrom(msg.sender, address(this), a);
    MarketPlace mPlace = MarketPlace(marketPlace);
    require(CErc20(mPlace.cTokenAddress(u, m)).mint(a) == 0, 'minting CToken Failed');
    require(mPlace.mintZcTokenAddingNotional(u, m, msg.sender, a), 'mint ZcToken adding Notional failed');

    return true;
  }

  /// @notice Allows users deposit/burn 1-1 amounts of both zcTokens and vault notional,
  /// in the process "combining" the two, and redeeming underlying. Calls mPlace.burnZcTokenRemovingNotional.
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param a Amount of zcTokens being redeemed
  function combineTokens(address u, uint256 m, uint256 a) external returns (bool) {
    MarketPlace mPlace = MarketPlace(marketPlace);
    require(mPlace.burnZcTokenRemovingNotional(u, m, msg.sender, a), 'burn ZcToken removing Notional failed');
    address cTokenAddr = mPlace.cTokenAddress(u, m);
    require((CErc20(cTokenAddr).redeemUnderlying(a) == 0), "compound redemption error");
    Erc20(u).transfer(msg.sender, a);

    return true;
  }

  /// @notice Allows zcToken holders to redeem their tokens for underlying tokens after maturity has been reached (via MarketPlace).
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param a Amount of zcTokens being redeemed
  function redeemZcToken(address u, uint256 m, uint256 a) external returns (bool) {
    MarketPlace mPlace = MarketPlace(marketPlace);
    // call marketplace to determine the amount redeemed
    uint256 redeemed = mPlace.redeemZcToken(u, m, msg.sender, a);
    // redeem underlying from compound
    require(CErc20(mPlace.cTokenAddress(u, m)).redeemUnderlying(redeemed) == 0, 'compound redemption failed');
    // transfer underlying back to msg.sender
    Erc20(u).transfer(msg.sender, redeemed);

    return true;
  }

  /// @notice Allows Vault owners to redeem any currently accrued interest (via MarketPlace)
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  function redeemVaultInterest(address u, uint256 m) external returns (bool) {
    MarketPlace mPlace = MarketPlace(marketPlace);
    // call marketplace to determine the amount redeemed
    uint256 redeemed = mPlace.redeemVaultInterest(u, m, msg.sender);
    // redeem underlying from compound
    require(CErc20(mPlace.cTokenAddress(u, m)).redeemUnderlying(redeemed) == 0, 'compound redemption failed');
    // transfer underlying back to msg.sender
    Erc20(u).transfer(msg.sender, redeemed);

    return true;
  }

  /// @notice Varifies the validity of an order and it's signature.
  /// @param o An offline Swivel.Order
  /// @param c Components of a valid ECDSA signature
  /// @return the hashed order.
  function validOrderHash(Hash.Order calldata o, Sig.Components calldata c) internal view returns (bytes32) {
    bytes32 hash = Hash.order(o);

    require(!cancelled[hash], 'order cancelled');
    require(o.expiry >= block.timestamp, 'order expired');
    require(o.maker == Sig.recover(Hash.message(domain, hash), c), 'invalid signature');

    return hash;
  }

  modifier onlyAdmin(address a) {
    require(msg.sender == a, 'sender must be admin');
    _;
  }
}
