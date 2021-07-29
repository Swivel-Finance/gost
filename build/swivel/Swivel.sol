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

  string constant public NAME = "Swivel Finance";
  string constant public VERSION = "2.0.0";
  bytes32 public immutable DOMAIN;
  address public immutable marketPlace;

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
  /// @param a Amount of volume (interest) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function initiateVaultFillingZcTokenInitiate(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o, c) {
    // Checks the side, and the amount compared to amount available
    require(a <= (o.premium - filled[o.key]), 'taker amount > available volume');

    filled[o.key] += a;
    
    // transfer tokens to this contract
    Erc20 uToken = Erc20(o.underlying);
    uToken.transferFrom(msg.sender, o.maker, a);
    uint256 principalFilled = (((a * 1e18) / o.premium) * o.principal) / 1e18;
    uToken.transferFrom(o.maker, address(this), principalFilled);

    MarketPlace mPlace = MarketPlace(marketPlace);
    address cTokenAddr = mPlace.cTokenAddress(o.underlying, o.maturity);
    uToken.approve(cTokenAddr, principalFilled); 
    require(CErc20(cTokenAddr).mint(principalFilled) == 0, 'minting CToken failed');
    // alert MarketPlace.
    require(mPlace.custodialInitiate(o.underlying, o.maturity, o.maker, msg.sender, principalFilled), 'custodial initiate failed');

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

    // transfer tokens to this contract
    Erc20 uToken = Erc20(o.underlying);
    uint256 premiumFilled = (((a * 1e18) / o.principal) * o.premium) / 1e18;
    uToken.transferFrom(o.maker, msg.sender, premiumFilled);
    uToken.transferFrom(msg.sender, address(this), a);
    
    MarketPlace mPlace = MarketPlace(marketPlace);
    address cTokenAddr = mPlace.cTokenAddress(o.underlying, o.maturity);

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
    
    // .interest is interest * ratio / 1e18 where ratio is (a * 1e18) / principal
    uint256 premiumFilled = (((a * 1e18) / o.principal) * o.premium) / 1e18;
    // transfer tokens to this contract
    Erc20(o.underlying).transferFrom(msg.sender, o.maker, (a - premiumFilled));
    // notify the marketplace...
    require(MarketPlace(marketPlace).p2pZcTokenExchange(o.underlying, o.maturity, o.maker, msg.sender, a), 'zcToken exchange failed');
            
    emit Initiate(o.key, o.maker, o.vault, o.exit, msg.sender, a, premiumFilled);
  }

  /// @notice Allows a user to initiate a Vault by filling an offline vault exit order
  /// @dev This method should pass (underlying, maturity, maker, sender, principalFilled) to MarketPlace.p2pVaultExchange
  /// @param o The order being filled
  /// @param a Amount of volume (interest) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function initiateVaultFillingVaultExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o, c) {
    // Checks the side, and the amount compared to amount available
    require(a <= (o.premium - filled[o.key]), 'taker amount > available volume');
    
    filled[o.key] += a;

    // transfer tokens to this contract
    Erc20(o.underlying).transferFrom(msg.sender, o.maker, a);
    // notify marketplace
    uint256 principalFilled = (((a * 1e18) / o.premium) * o.principal) / 1e18;
    require(MarketPlace(marketPlace).p2pVaultExchange(o.underlying, o.maturity, o.maker, msg.sender, principalFilled), 'vault exchange failed');

    emit Initiate(o.key, o.maker, o.vault, o.exit, msg.sender, a, principalFilled);
  }

  // ********* EXITING ***************

  /// @notice Allows a user to exit (sell) a currently held position to the marketplace.
  /// @param o Array of offline Swivel.Orders
  /// @param a Array of order volume (principal) amounts relative to passed orders
  /// @param c Components of a valid ECDSA signature
  function exit(Hash.Order[] calldata o, uint256[] calldata a, Sig.Components[] calldata c) public returns (bool) {
    for (uint256 i=0; i < o.length; i++) {
      // Determine whether the order being filled is an exit
      if (!o[i].exit) {
        // Determine whether the order being filled is a vault initiate or a zcToken initiate
          if (!o[i].vault) {
            // If filling a zcToken initiate with an exit, one is exiting zcTokens
            exitZcTokenFillingZcTokenInitiate(o[i], a[i], c[i]);
          } else {
            // If filling a vault initiate with an exit, one is exiting vault notional
            exitVaultFillingVaultInitiate(o[i], a[i], c[i]);
          }
      } else {
        // Determine whether the order being filled is a vault exit or zcToken exit
        if (!o[i].vault) {
          // If filling a zcToken exit with an exit, one is exiting vault
          exitVaultFillingZcTokenExit(o[i], a[i], c[i]);
        } else {
          // If filling a vault exit with an exit, one is exiting zcTokens
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
  function exitZcTokenFillingZcTokenInitiate(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o,c) {
    require(a <= (o.premium - filled[o.key]), 'taker amount > available volume');
    
    filled[o.key] += a;       

    // notify marketplace...
    uint256 principalFilled = (((a * 1e18) / o.premium) * o.principal) / 1e18;
    require(MarketPlace(marketPlace).p2pZcTokenExchange(o.underlying, o.maturity, msg.sender, o.maker, principalFilled), 'zcToken exchange failed');
    // Transfer underlying from initiating party to exiting party, minus the price the exit party pays for the exit (interest).
    Erc20(o.underlying).transferFrom(o.maker, msg.sender, (principalFilled - a));
    
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
        
    // market should transfer <a> notional from sender to maker
    require(MarketPlace(marketPlace).p2pVaultExchange(o.underlying, o.maturity, msg.sender, o.maker, a), 'vault exchange failed');
    uint256 premiumFilled = (((a * 1e18) / o.principal) * o.premium) / 1e18;
    Erc20(o.underlying).transferFrom(o.maker, msg.sender, premiumFilled);

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
    
    MarketPlace mPlace = MarketPlace(marketPlace);
    // alert MarketPlace...
    require(mPlace.custodialExit(o.underlying, o.maturity, o.maker, msg.sender, a), 'custodial exit failed');
    // transfer cost of interest coupon to floating party
    Erc20 uToken = Erc20(o.underlying);
    uint256 premiumFilled = (((a * 1e18) / o.principal) * o.premium) / 1e18;
    uToken.transferFrom(o.maker, msg.sender, premiumFilled);
    // redeem principal from compound now that coupon and zcb have been redeemed
    address cTokenAddr = mPlace.cTokenAddress(o.underlying, o.maturity);
    require((CErc20(cTokenAddr).redeemUnderlying(a) == 0), "compound redemption error");
    // Transfer principal back to fixed exit party now that the interest coupon and zcb have been redeemed
    uToken.transfer(o.maker, a);
    
    emit Exit(o.key, o.maker, o.vault, o.exit, msg.sender, a, premiumFilled);
  }

  /// @notice Allows a user to exit their zcTokens by filling an offline vault exit order
  /// @dev This method should pass (underlying, maturity, sender, maker, principalFilled) to MarketPlace.exitFillingExit
  /// @param o The order being filled
  /// @param a Amount of volume (interest) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function exitZcTokenFillingVaultExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o,c) {
    require(a <= (o.premium - filled[o.key]), 'taker amount > available volume');
    
    filled[o.key] += a;

    MarketPlace mPlace = MarketPlace(marketPlace);
    // inform MarketPlace what happened...
    uint256 principalFilled = (((a * 1e18) / o.premium) * o.principal) / 1e18;
    require(mPlace.custodialExit(o.underlying, o.maturity, msg.sender, o.maker, principalFilled), 'custodial exit failed');
    // Transfer cost of interest coupon to floating party
    Erc20 uToken = Erc20(o.underlying);
    uToken.transferFrom(msg.sender, o.maker, a);
    // Redeem principal from compound now that coupon and zcb have been redeemed
    address cTokenAddr = mPlace.cTokenAddress(o.underlying, o.maturity);
    require((CErc20(cTokenAddr).redeemUnderlying(principalFilled) == 0), "compound redemption Error");
    // Transfer principal back to fixed exit party now that the interest coupon and zcb have been redeemed
    uToken.transfer(msg.sender, principalFilled);
    
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

  /// @dev Agreements may only be Initiated if the Order is valid.
  /// @param o An offline Swivel.Order
  /// @param c Components of a valid ECDSA signature
  modifier valid(Hash.Order calldata o, Sig.Components calldata c) {
    require(!cancelled[o.key], 'order cancelled');
    require(o.expiry >= block.timestamp, 'order expired');
    require(o.maker == Sig.recover(Hash.message(DOMAIN, Hash.order(o)), c), 'invalid signature');
    _;
  }
}
