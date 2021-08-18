// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.4;

import './Abstracts.sol';
import './Hash.sol';
import './Sig.sol';

contract Swivel {
  /// @dev maps the key of an order to a boolean indicating if an order was cancelled
  mapping (bytes32 => bool) public cancelled;
  /// @dev maps the key of an order to an amount representing its taken volume
  mapping (bytes32 => uint256) public filled;

  string constant public NAME = 'Swivel Finance';
  string constant public VERSION = '2.0.0';
  bytes32 public immutable DOMAIN;
  address public immutable marketPlace;
  address public immutable admin;
  uint16[] public feeDenominator;
  //0 = zcTokenInitiate Fee Denominator
  //1 = zcTokenExit Fee Denominator
  //2 = VaultInitiate Fee Denominator
  //3 = VaultExit Fee Denominator

  /// @notice Emitted on order cancellation
  event Cancel (bytes32 indexed key);
  /// @notice Emitted on any initiate*
  /// @dev filled is 'principalFilled' when (vault:false, exit:false) && (vault:true, exit:true)
  /// @dev filled is 'premiumFilled' when (vault:true, exit:false) && (vault:false, exit:true)
  event Initiate(bytes32 indexed key, address indexed maker, bool vault, bool exit, address indexed sender, uint256 amount, uint256 filled);
  /// @notice Emitted on any exit*
  /// @dev filled is 'principalFilled' when (vault:false, exit:false) && (vault:true, exit:true)
  /// @dev filled is 'premiumFilled' when (vault:true, exit:false) && (vault:false, exit:true)
  event Exit(bytes32 indexed key, address indexed maker, bool vault, bool exit, address indexed sender, uint256 amount, uint256 filled);

  /// @param m deployed MarketPlace contract address
  constructor(address m) {
    DOMAIN = Hash.domain(NAME, VERSION, block.chainid, address(this));
    marketPlace = m;
    admin = msg.sender;
    feeDenominator = [200,600,400,200];
  }


  // ********* ADMIN FUNCTIONS *************

  // @notice Allows the admin to set a new fee denominator
  // @param t The type of order 
  // @param f The new fee denominator
  function setFee(uint8 t, uint16 d) external onlyAdmin(admin) {
      feeDenominator[t] = d;
  }

  // ********* INITIATING *************

  /// @notice Allows a user to initiate a position
  /// @param o Array of offline Swivel.Orders
  /// @param a Array of order volume (principal) amounts relative to passed orders
  /// @param c Array of Components from valid ECDSA signatures
  function initiate(Hash.Order[] calldata o, uint256[] calldata a, Sig.Components[] calldata c) public returns (bool) {
    for (uint256 i=0; i < o.length; i++) {
      // TODO explain the scenarios
      if (!o[i].exit) {
        if (!o[i].vault) {
          initiateVaultFillingZcTokenInitiate(o[i], a[i], c[i]);
        } else {
          initiateZcTokenFillingVaultInitiate(o[i], a[i], c[i]);
        }
      } else {
        if (!o[i].vault) {
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
  function initiateVaultFillingZcTokenInitiate(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o, c) {
    // Checks the side, and the amount compared to amount available
    require(a <= (o.premium - filled[o.key]), 'taker amount > available volume');

    filled[o.key] += a;
    
    Erc20 uToken = Erc20(o.underlying);
    uint256 principalFilled = (((a * 1e18) / o.premium) * o.principal) / 1e18;
    uint256 fee = ((principalFilled * 1e18) / feeDenominator[2]) / 1e18;
    
    uToken.transferFrom(msg.sender, o.maker, a);
    uToken.transferFrom(o.maker, address(this), principalFilled);

    MarketPlace mPlace = MarketPlace(marketPlace);
    address cTokenAddr = mPlace.cTokenAddress(o.underlying, o.maturity);
    // mint cTokens
    uToken.approve(cTokenAddr, principalFilled); 
    require(CErc20(cTokenAddr).mint(principalFilled) == 0, 'minting CToken failed');
    
    // alert MarketPlace.
    require(mPlace.custodialInitiate(o.underlying, o.maturity, o.maker, msg.sender, principalFilled), 'custodial initiate failed');
    // transfer fee in vault notional to swivel (from msg.sender)
    require(mPlace.transferVaultNotionalFee(o.underlying, o.maturity, msg.sender, address(this), fee), "notional fee transfer failed");
    
    emit Initiate(o.key, o.maker, o.vault, o.exit, msg.sender, a, principalFilled);
  }

  /// @notice Allows a user to initiate a zcToken by filling an offline vault initiate order
  /// @dev This method should pass (underlying, maturity, sender, maker, a) to MarketPlace.custodialInitiate
  /// @param o The order being filled
  /// @param o Amount of volume (principal) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function initiateZcTokenFillingVaultInitiate(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o, c) {
    // Checks the side, and the amount compared to amount available
    require((a <= o.principal - filled[o.key]), 'taker amount > available volume');

    filled[o.key] += a;

    Erc20 uToken = Erc20(o.underlying);
    uint256 premiumFilled = (((a * 1e18) / o.principal) * o.premium) / 1e18;
    uint256 fee = ((premiumFilled * 1e18) / feeDenominator[0]) / 1e18;
    
    uToken.transferFrom(o.maker, msg.sender, premiumFilled);
    // transfer principal + fee in underlying to swivel (from msg.sender)
    uToken.transferFrom(msg.sender, address(this), (a+fee));
    
    MarketPlace mPlace = MarketPlace(marketPlace);
    address cTokenAddr = mPlace.cTokenAddress(o.underlying, o.maturity);
    // mint cTokens
    uToken.approve(cTokenAddr, a);
    require(CErc20(cTokenAddr).mint(a) == 0, 'minting CToken Failed');
    
    // alert MarketPlace
    require(mPlace.custodialInitiate(o.underlying, o.maturity, msg.sender, o.maker, a), 'custodial initiate failed');
    
    emit Initiate(o.key, o.maker, o.vault, o.exit, msg.sender, a, premiumFilled);
  }

  /// @notice Allows a user to initiate zcToken? by filling an offline zcToken exit order
  /// @dev This method should pass (underlying, maturity, maker, sender, a) to MarketPlace.p2pZcTokenExchange
  /// @param o The order being filled
  /// @param a Amount of volume (principal) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function initiateZcTokenFillingZcTokenExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o, c) {
    // Checks the side, and the amount compared to amount available
    require(a <= ((o.principal - filled[o.key])), 'taker amount > available volume');

    filled[o.key] += a;
    
    Erc20 uToken = Erc20(o.underlying);
    uint256 premiumFilled = (((a * 1e18) / o.principal) * o.premium) / 1e18;
    uint256 fee = ((premiumFilled * 1e18) / feeDenominator[0]) / 1e18;
    
    // transfer principal - the premium paid + fee in underlying to swivel (from msg.sender)
    uToken.transferFrom(msg.sender, o.maker, ((a - premiumFilled)+fee));
    // notify the marketplace...
    require(MarketPlace(marketPlace).p2pZcTokenExchange(o.underlying, o.maturity, o.maker, msg.sender, a), 'zcToken exchange failed');
    
    emit Initiate(o.key, o.maker, o.vault, o.exit, msg.sender, a, premiumFilled);
  }

  /// @notice Allows a user to initiate a Vault by filling an offline vault exit order
  /// @dev This method should pass (underlying, maturity, maker, sender, principalFilled) to MarketPlace.p2pVaultExchange
  /// @param o The order being filled
  /// @param a Amount of volume (premium) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function initiateVaultFillingVaultExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o, c) {
    // Checks the side, and the amount compared to amount available
    require(a <= (o.premium - filled[o.key]), 'taker amount > available volume');
    
    filled[o.key] += a;
    
    uint256 principalFilled = (((a * 1e18) / o.premium) * o.principal) / 1e18;
    uint256 fee = ((principalFilled * 1e18) / feeDenominator[2]) / 1e18;

    MarketPlace mPlace = MarketPlace(marketPlace);

    Erc20(o.underlying).transferFrom(msg.sender, o.maker, a);
    // notify marketplace
    require(MarketPlace(marketPlace).p2pVaultExchange(o.underlying, o.maturity, o.maker, msg.sender, principalFilled), 'vault exchange failed');

    // transfer fee in vault notional to swivel (from msg.sender)
    require(mPlace.transferVaultNotionalFee(o.underlying, o.maturity, msg.sender, address(this), fee), "notional fee transfer failed");
    
    emit Initiate(o.key, o.maker, o.vault, o.exit, msg.sender, a, principalFilled);
  }

  // ********* EXITING ***************

  /// @notice Allows a user to exit (sell) a currently held position to the marketplace.
  /// @param o Array of offline Swivel.Orders
  /// @param a Array of order volume (principal) amounts relative to passed orders
  /// @param c Components of a valid ECDSA signature
  function exit(Hash.Order[] calldata o, uint256[] calldata a, Sig.Components[] calldata c) public returns (bool) {
    for (uint256 i=0; i < o.length; i++) {
      // determine whether the order being filled is an exit
      if (!o[i].exit) {
        // determine whether the order being filled is a vault initiate or a zcToken initiate
          if (!o[i].vault) {
            // if filling a zcToken initiate with an exit, one is exiting zcTokens
            exitZcTokenFillingZcTokenInitiate(o[i], a[i], c[i]);
          } else {
            // if filling a vault initiate with an exit, one is exiting vault notional
            exitVaultFillingVaultInitiate(o[i], a[i], c[i]);
          }
      } else {
        // determine whether the order being filled is a vault exit or zcToken exit
        if (!o[i].vault) {
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
  /// @param a Amount of volume (premium) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function exitZcTokenFillingZcTokenInitiate(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o,c) {
    require(a <= (o.premium - filled[o.key]), 'taker amount > available volume');
    
    filled[o.key] += a;   
    
    Erc20 uToken = Erc20(o.underlying);
    uint256 principalFilled = (((a * 1e18) / o.premium) * o.principal) / 1e18;
    uint256 fee = ((principalFilled * 1e18) / feeDenominator[1]) / 1e18;
    
    // transfer underlying from initiating party to exiting party, minus the price the exit party pays for the exit (premium), and the fee.
    uToken.transferFrom(o.maker, msg.sender, (principalFilled - a - fee));
    // notify marketplace
    require(MarketPlace(marketPlace).p2pZcTokenExchange(o.underlying, o.maturity, msg.sender, o.maker, principalFilled), 'zcToken exchange failed');
    
    // transfer fee in underlying to swivel (with msg.sender's change)
    uToken.transferFrom(o.maker, address(this), fee);
    
    emit Exit(o.key, o.maker, o.vault, o.exit, msg.sender, a, principalFilled);
  }
  
  /// @notice Allows a user to exit their Vault by filling an offline vault initiate order
  /// @dev This method should pass (underlying, maturity, sender, maker, a) to MarketPlace.p2pVaultExchange
  /// @param o The order being filled
  /// @param a Amount of volume (principal) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function exitVaultFillingVaultInitiate(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o,c) {
    require(a <= (o.principal - filled[o.key]), 'taker amount > available volume');
    
    filled[o.key] += a;
    
    Erc20 uToken = Erc20(o.underlying);
    uint256 premiumFilled = (((a * 1e18) / o.principal) * o.premium) / 1e18;
    uint256 fee = ((premiumFilled * 1e18) / feeDenominator[3]) / 1e18;

    // transfer premium from o.maker to msg.sender
    uToken.transferFrom(o.maker, msg.sender, premiumFilled-fee);

    // market should transfer <a> notional from sender to maker
    require(MarketPlace(marketPlace).p2pVaultExchange(o.underlying, o.maturity, msg.sender, o.maker, a), 'vault exchange failed');

    // transfer fee in underlying to swivel (from msg.sender)
    uToken.transferFrom(msg.sender, address(this), fee);

    emit Exit(o.key, o.maker, o.vault, o.exit, msg.sender, a, premiumFilled);
  }

  /// @notice Allows a user to exit their Vault filling an offline zcToken exit order
  /// @dev This method should pass (underlying, maturity, maker, sender, a) to MarketPlace.exitFillingExit
  /// @param o The order being filled
  /// @param a Amount of volume (principal) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function exitVaultFillingZcTokenExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o,c) {
    require(a <= (o.principal - filled[o.key]), 'taker amount > available volume');
    
    filled[o.key] += a;
    
    Erc20 uToken = Erc20(o.underlying);
    uint256 premiumFilled = (((a * 1e18) / o.principal) * o.premium) / 1e18;
    uint256 fee = ((premiumFilled * 1e18) / feeDenominator[3]) / 1e18;
    
    MarketPlace mPlace = MarketPlace(marketPlace);
    // alert MarketPlace...
    require(mPlace.custodialExit(o.underlying, o.maturity, o.maker, msg.sender, a), 'custodial exit failed');
    
    // redeem principal from compound now that coupon and zcb have been burnt
    address cTokenAddr = mPlace.cTokenAddress(o.underlying, o.maturity);
    require((CErc20(cTokenAddr).redeemUnderlying(a) == 0), 'compound redemption error');
    //1000
    // transfer principal-premium back to fixed exit party now that the interest coupon and zcb have been redeemed
    uToken.transfer(o.maker, (a-premiumFilled));
    // transfer premium-fee to floating exit party
    uToken.transfer(msg.sender, (premiumFilled-fee));
    
    emit Exit(o.key, o.maker, o.vault, o.exit, msg.sender, a, premiumFilled);
  }

  /// @notice Allows a user to exit their zcTokens by filling an offline vault exit order
  /// @dev This method should pass (underlying, maturity, sender, maker, principalFilled) to MarketPlace.exitFillingExit
  /// @param o The order being filled
  /// @param a Amount of volume (premium) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function exitZcTokenFillingVaultExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o,c) {
    require(a <= (o.premium - filled[o.key]), 'taker amount > available volume');
    
    filled[o.key] += a;

    Erc20 uToken = Erc20(o.underlying);
    uint256 principalFilled = (((a * 1e18) / o.premium) * o.principal) / 1e18;
    uint256 fee = ((principalFilled * 1e18) / feeDenominator[1]) / 1e18;
    
    MarketPlace mPlace = MarketPlace(marketPlace);
    // alert MarketPlace...
    require(mPlace.custodialExit(o.underlying, o.maturity, msg.sender, o.maker, principalFilled), 'custodial exit failed');

    // redeem principal from compound now that coupon and zcb have been burnt
    address cTokenAddr = mPlace.cTokenAddress(o.underlying, o.maturity);
    require((CErc20(cTokenAddr).redeemUnderlying(principalFilled) == 0), 'compound redemption Error');

    // transfer principal-premium-fee back to fixed exit party now that the interest coupon and zcb have been redeemed
    uToken.transfer(msg.sender, ((principalFilled-a)-fee));
    // transfer premium to floating exit party
    uToken.transfer(o.maker, (a));

    emit Exit(o.key, o.maker, o.vault, o.exit, msg.sender, a, principalFilled);
  }

  /// @notice Allows a user to cancel an order, preventing it from being filled in the future
  /// @param o An offline Swivel.Order
  /// @param c Components of a valid ECDSA signature
  function cancel(Hash.Order calldata o, Sig.Components calldata c) public returns (bool) {
    require(o.maker == Sig.recover(Hash.message(DOMAIN, Hash.order(o)), c), 'invalid signature');
    cancelled[o.key] = true;

    emit Cancel(o.key);

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
    address cTokenAddr = mPlace.cTokenAddress(u, m);
    uToken.approve(cTokenAddr, a);
    require(CErc20(cTokenAddr).mint(a) == 0, 'minting CToken Failed');
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
    require((CErc20(cTokenAddr).redeemUnderlying(a) == 0), 'compound redemption error');
    Erc20(u).transfer(msg.sender, a);

    return true;
  }

  /// @dev Agreements may only be Initiated if the Order is valid.
  /// @param o An offline Swivel.Order
  /// @param c Components of a valid ECDSA signature
  modifier valid(Hash.Order calldata o, Sig.Components calldata c) {
    require(!cancelled[o.key], 'order cancelled');
    require(o.expiry >= block.timestamp, 'order expired');
    require(o.maker == Sig.recover(Hash.message(DOMAIN, Hash.order(o)), c), 'invalid signature');
    _;
  }

  modifier onlyAdmin(address a) {
    require(msg.sender == admin, 'sender must be admin');
    _;
  }
}
