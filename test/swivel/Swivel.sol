// SPDX-License-Identifier: UNLICENSED

// TODO update to 0.8.4 (or whatever latest is...)
pragma solidity 0.8.0;

import './Hash.sol';
import './Sig.sol';
import './Abstracts.sol';

contract Swivel {
  string constant public NAME = "Swivel Finance";
  string constant public VERSION = "2.0.0";
  bytes32 public DOMAIN;
  address public admin;
  address public marketPlaceAddr;

  /// @dev maps the key of an order to a boolean indicating if an order was cancelled
  mapping (bytes32 => bool) public cancelled;
    
  /// @dev maps the key of an order to an amount representing its taken volume
  mapping (bytes32 => uint256) public filled;  

  /// @notice Emitted on order cancellation
  event Cancel (bytes32 indexed key);

  /// @param m deployed MarketPlace contract address
  constructor(address m) {
    admin = msg.sender;
    DOMAIN = Hash.domain(NAME, VERSION, block.chainid, address(this));
    marketPlaceAddr = m;
  }

  // ********* INITIATING *************

  /// @notice Allows a user to initiate a position
  /// @param o Array of offline Swivel.Orders
  /// @param a Array of order volume (principal) amounts relative to passed orders
  /// @param c Array of Components from valid ECDSA signatures
  function initiate(Hash.Order[] calldata o, uint256[] calldata a, Sig.Components[] calldata c) public returns (bool) {
    for (uint256 i=0; i < o.length; i++) {
      // TODO explain the scenarios
      if (o[i].exit == false) {
        if (o[i].vault == false) {
          require(initiateVaultFillingZcTokenInitiate(o[i], a[i], c[i]));
        } else {
          require(initiateZcTokenFillingVaultInitiate(o[i], a[i], c[i]));
        }
      } else {
        if (o[i].vault == false) {
          require(initiateZcTokenFillingZcTokenExit(o[i], a[i], c[i]));
        } else {
          require(initiateVaultFillingVaultExit(o[i], a[i], c[i]));
        }
      }
    }

    return true;
  }

  /// @notice Allows a user to initiate a Vault by filling an offline zcToken initiate order
  /// @param o The order being filled
  /// @param a Amount of volume (interest) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function initiateVaultFillingZcTokenInitiate(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o, c) returns (bool) {
    // Checks the side, and the amount compared to amount available
    require(a <= (o.premium - filled[o.key]), 'taker amount > available volume');
    
    uint256 principalFilled = (((a * 1e18) / o.premium) * o.principal) / 1e18;
    
    // transfer tokens to this contract
    Erc20 uToken = Erc20(o.underlying);
    require(uToken.transferFrom(msg.sender, o.maker, a), 'transfer failed');
    require(uToken.transferFrom(o.maker, address(this), principalFilled), 'transfer failed');

    MarketPlace marketPlace = MarketPlace(marketPlaceAddr);
    address cTokenAddr = marketPlace.cTokenAddress(o.underlying, o.maturity);

    require(uToken.approve(cTokenAddr, principalFilled), 'underlying approval failed'); 
    require(CErc20(cTokenAddr).mint(principalFilled) == 0, 'minting CToken failed');
    // TODO assure sender, maker...
    require(marketPlace.mintZcTokenAddingNotional(o.underlying, o.maturity, o.maker, msg.sender, principalFilled), 'minting ZCToken failed');

    filled[o.key] += a;

    return true;
  }

  /// @notice Allows a user to initiate a zcToken _ by filling an offline vault initiate order
  /// @param o The order being filled
  /// @param o Amount of volume (principal) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function initiateZcTokenFillingVaultInitiate(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o, c) returns (bool) {
    // Checks the side, and the amount compared to amount available
    require((a <= o.principal - filled[o.key]), 'taker amount > available volume');

    uint256 premiumFilled = (((a * 1e18) / o.principal) * o.premium) / 1e18;

    // transfer tokens to this contract
    Erc20 uToken = Erc20(o.underlying);
    require(uToken.transferFrom(o.maker, msg.sender, premiumFilled), 'transfer failed');
    require(uToken.transferFrom(msg.sender, address(this), a), 'transfer failed');
    
    MarketPlace marketPlace = MarketPlace(marketPlaceAddr);
    address cTokenAddr = marketPlace.cTokenAddress(o.underlying, o.maturity);

    uToken.approve(cTokenAddr, a);
    require(CErc20(cTokenAddr).mint(a) == 0, 'minting CToken Failed');
    // TODO sender, maker order...
    require(marketPlace.mintZcTokenAddingNotional(o.underlying, o.maturity, o.maker, msg.sender, a), 'minting ZCToken, adding notional failed');
    
    filled[o.key] += a;

    return true;
  }

  /// @notice Allows a user to initiate zcToken? by filling an offline zcToken exit order
  /// @param o The order being filled
  /// @param a Amount of volume (principal) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function initiateZcTokenFillingZcTokenExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o, c) returns (bool) {
    // Checks the side, and the amount compared to amount available
    require(a <= ((o.principal - filled[o.key])), 'taker amount > available volume');
    
    // .interest is interest * ratio / 1e18 where ratio is (a * 1e18) / principal
    uint256 premiumFilled = (((a * 1e18) / o.principal) * o.premium) / 1e18;

    // transfer tokens to this contract
    require(Erc20(o.underlying).transferFrom(msg.sender, o.maker, (a - premiumFilled)), 'transfer failed');
    // the zctoken in this order's market should transfer <a> from sender to maker TODO assure order...
    require(MarketPlace(marketPlaceAddr).transferFromZcToken(o.underlying, o.maturity, o.maker, msg.sender, a), 'ZCToken transfer failed');
    
    filled[o.key] += a;
            
    return true;
  }

  /// @notice Allows a user to initiate a Vault by filling an offline vault exit order
  /// @param o The order being filled
  /// @param a Amount of volume (interest) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function initiateVaultFillingVaultExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o, c) returns (bool) {
    // Checks the side, and the amount compared to amount available
    require(a <= (o.premium - filled[o.key]), 'taker amount > available volume');
    
    uint256 principalFilled = (((a * 1e18) / o.premium) * o.principal) / 1e18;
 
    // transfer tokens to this contract
    require(Erc20(o.underlying).transferFrom(msg.sender, o.maker, a), 'transfer failed');
    // the notional in this order's market should transfer <principalFilled> from maker to sender TODO assure order...
    require(MarketPlace(marketPlaceAddr).transferFromNotional(o.underlying, o.maturity, o.maker, msg.sender, principalFilled), 'notional transfer failed');
    
    filled[o.key] += a;

    return true;
  }

  // ********* EXITING ***************

  /// @notice Allows a user to exit (sell) a currently held position to the marketplace.
  /// @param o Array of offline Swivel.Orders
  /// @param a Array of order volume (principal) amounts relative to passed orders
  /// @param c Components of a valid ECDSA signature
  function exit(Hash.Order[] calldata o, uint256[] calldata a, Sig.Components[] calldata c) public returns (bool) {
    for (uint256 i=0; i < o.length; i++) {
      // Determine whether the order being filled is an exit
      if (o[i].exit == false) {
        // Determine whether the order being filled is a vault initiate or a zcToken initiate
          if (o[i].vault == false) {
            // If filling a zcToken initiate with an exit, one is exiting zcTokens
            require(exitZcTokenFillingZcTokenInitiate(o[i], a[i], c[i]));
          } else {
            // If filling a vault initiate with an exit, one is exiting vault notional
            require(exitVaultFillingVaultInitiate(o[i], a[i], c[i]));
          }
      } else {
        // Determine whether the order being filled is a vault exit or zcToken exit
        if (o[i].vault == false) {
          // If filling a zcToken exit with an exit, one is exiting vault
          require(exitVaultFillingZcTokenExit(o[i], a[i], c[i]));
        } else {
          // If filling a vault exit with an exit, one is exiting zcTokens
          require(exitZcTokenFillingVaultExit(o[i], a[i], c[i]));
        }   
      }   
    }

    return true;
  }

  /// @notice Allows a user to exit their zcTokens by filling an offline zcToken initiate order
  /// @param : o The order being filled
  /// @param : a Amount of volume (interest) being filled by the taker's exit
  /// @param : c Components of a valid ECDSA signature
  function exitZcTokenFillingZcTokenInitiate(Hash.Order calldata o, uint256 a, Sig.Components calldata c) valid(o,c) internal returns (bool) {
    require(a <= (o.premium - filled[o.key]), 'taker amount > available volume');
    
    uint256 principalFilled = (((a * 1e18) / o.premium) * o.principal) / 1e18;
    // transferZcToken <premiumFilled> from sender to maker TODO assure sender, maker...
    require(MarketPlace(marketPlaceAddr).transferFromZcToken(o.underlying, o.maturity, o.maker, msg.sender, principalFilled), 'ZCToken transfer failed');
    // Transfer underlying from initiating party to exiting party, minus the price the exit party pays for the exit (interest).
    // TODO audit these revert messages...
    require(Erc20(o.underlying).transferFrom(o.maker, msg.sender, (principalFilled - a)), 'transfer failed');
    
    filled[o.key] += a;       
    
    // TODO events?
    return true;
  }
  
  /// @notice Allows a user to exit their Vault by filling an offline vault initiate order
  /// @param o The order being filled
  /// @param a Amount of volume (principal) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function exitVaultFillingVaultInitiate(Hash.Order calldata o, uint256 a, Sig.Components calldata c) valid(o,c) internal returns (bool) {
    require(a <= (o.principal - filled[o.key]), 'taker amount > available volume');
    
    uint256 premiumFilled = (((a * 1e18) / o.principal) * o.premium) / 1e18;
        
    // market should transfer <a> notional from sender to maker
    // NOTE the method naming is being addressed in a follow up PR
    // this will prevent errors with argument passing and such...
    require(MarketPlace(marketPlaceAddr).transferFromNotional(o.underlying, o.maturity, o.maker, msg.sender, a), 'notional transfer failed');
    require(Erc20(o.underlying).transferFrom(o.maker, msg.sender, premiumFilled), 'transfer failed');
    
    filled[o.key] += a;
    // TODO events?
    return true;
  }

  /// @notice Allows a user to exit their Vault filling an offline zcToken exit order
  /// @param o The order being filled
  /// @param a Amount of volume (principal) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function exitVaultFillingZcTokenExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) valid(o,c) internal returns (bool) {
    require(a <= (o.principal - filled[o.key]), 'taker amount > available volume');
    
    uint256 premiumFilled = (((a * 1e18) / o.principal) * o.premium) / 1e18;
    
    MarketPlace marketPlace = MarketPlace(marketPlaceAddr);
    // burn zcTokens for fixed exit party, burn interest coupon for floating exit party (burn maker, remove sender)
    require(marketPlace.burnZcTokenRemovingNotional(o.underlying, o.maturity, o.maker, msg.sender, a), 'notional transfer failed');
    // transfer cost of interest coupon to floating party
    Erc20 uToken = Erc20(o.underlying);
    require(uToken.transferFrom(o.maker, msg.sender, premiumFilled), 'transfer failed');
    // redeem principal from compound now that coupon and zcb have been redeemed
    address cTokenAddr = marketPlace.cTokenAddress(o.underlying, o.maturity);
    require((CErc20(cTokenAddr).redeemUnderlying(a) == 0), "compound redemption error");
    // Transfer principal back to fixed exit party now that the interest coupon and zcb have been redeemed
    require(uToken.transfer(o.maker, a), 'transfer failed');
    
    filled[o.key] += a;
    
    // TODO events?

    return true;
  }

  /// @notice Allows a user to exit their zcTokens by filling an offline vault exit order
  /// @param : o The order being filled
  /// @param : a Amount of volume (interest) being filled by the taker's exit
  /// @param : c Components of a valid ECDSA signature
  function exitZcTokenFillingVaultExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) valid(o,c) internal returns (bool) {
    require(a <= (o.premium - filled[o.key]));
    
    uint256 principalFilled = (((a * 1e18) / o.premium) * o.principal) / 1e18;
    
    MarketPlace marketPlace = MarketPlace(marketPlaceAddr);
    // Burn zcTokens for fixed exit party... (burn sender, remove maker)
    require(marketPlace.burnZcTokenRemovingNotional(o.underlying, o.maturity, o.maker, msg.sender, a), 'notional transfer failed');
    // Transfer cost of interest coupon to floating party
    Erc20 uToken = Erc20(o.underlying);
    require(uToken.transferFrom(msg.sender, o.maker, a), 'transfer failed');
    // Redeem principal from compound now that coupon and zcb have been redeemed
    address cTokenAddr = marketPlace.cTokenAddress(o.underlying, o.maturity);
    require((CErc20(cTokenAddr).redeemUnderlying(principalFilled) == 0), "compound redemption Error");
    // Transfer principal back to fixed exit party now that the interest coupon and zcb have been redeemed
    require(uToken.transfer(msg.sender, principalFilled), 'transfer failed');
    
    filled[o.key] += a;
    
    // TODO events

    return true;
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
    require(cancelled[o.key] == false, 'order cancelled');
    require(o.expiry >= block.timestamp, 'order expired');
    require(o.maker == Sig.recover(Hash.message(DOMAIN, Hash.order(o)), c), 'invalid signature');
    _;
  }
}
