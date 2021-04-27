// SPDX-License-Identifier: UNLICENSED

// TODO update to 0.8.4 (or whatever latest is...)
pragma solidity 0.8.0;

contract Swivel {
  string constant public NAME = "Swivel Finance";
  string constant public VERSION = "2.0.0";
  address public admin;
  bytes32 public DOMAIN;

  /// @dev maps the key of an order to a boolean indicating if an order was cancelled
  mapping (bytes32 => bool) public cancelled;
    
  /// @dev maps the key of an order to an amount representing its taken volume
  mapping (bytes32 => uint256) public filled;  

  /// @notice Emitted on order cancellation
  event Cancel (bytes32 indexed key);

  constructor() {
    admin = msg.sender;
    DOMAIN = Hash.domain(NAME, VERSION, block.chainid, address(this));
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
        if (o[i].timeLock == false) {
          require(initiateVaultFillingZcTokenInitiate(o[i], a[i], c[i]));
        } else {
          require(initiateZcTokenFillingVaultInitiate(o[i], a[i], c[i]));
        }
      } else {
        if (o[i].timeLock == false) {
          require(initiateZcTokenFillingZcTokenExit(o[i], a[i], c[i]));
        } else {
          require(initiateVaultFillingVaultExit(o[i], a[i], c[i]));
        }
      }
    }

    return true;
  }

  /// @notice Allows a user to initiate zcToken? by filling an offline zcToken exit order
  /// @param o The order being filled
  /// @param a Amount of volume (principal) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function initiateZcTokenFillingZcTokenExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o, c) returns (bool) {
    // TODO

    return true;
  }

  /// @notice Allows a user to initiate a Vault by filling an offline vault exit order
  /// @param o The order being filled
  /// @param a Amount of volume (interest) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function initiateVaultFillingVaultExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) internal valid(o, c) returns (bool) {
    // TODO

    return true;
  }

  /// @notice Allows a user to initiate a Vault by filling an offline zcToken initiate order
  /// @param o The order being filled
  /// @param a Amount of volume (interest) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function initiateVaultFillingZcTokenInitiate(Hash.Order calldata o,uint256 a,Sig.Components calldata c) internal valid(o, c) returns (bool) {
    // TODO

    return true;
  }

  /// @notice Allows a user to initiate a zcToken _ by filling an offline vault initiate order
  /// @param o The order being filled
  /// @param o Amount of volume (principal) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function initiateZcTokenFillingVaultInitiate(Hash.Order calldata o, uint256 a, Sig.Components calldata c) public valid(o, c) returns (bool) {
    // TODO

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
          if (o[i].timeLock == false) {
            // If filling a zcToken initiate with an exit, one is exiting zcTokens
            require(exitZcTokenFillingZcTokenInitiateOrder(o[i], a[i], c[i]));
          } else {
            // If filling a vault initiate with an exit, one is exiting vault notional
            require(exitVaultFillingVaultInitiateOrder(o[i], a[i], c[i]));
          }
      } else {
        // Determine whether the order being filled is a vault exit or zcToken exit
        if (o[i].timeLock == false) {
          // If filling a zcToken exit with an exit, one is exiting vault
          require(exitVaultFillingZcTokenExitOrder(o[i], a[i], c[i]));
        } else {
          // If filling a vault exit with an exit, one is exiting zcTokens
          require(exitZcTokenFillingVaultExitOrder(o[i], a[i], c[i]));
        }   
      }   
    }

    return true;
  }
  
  /// @notice Allows a user to exit their Vault by filling an offline vault initiate order
  /// @param o The order being filled
  /// @param a Amount of volume (principal) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function exitVaultFillingVaultInitiate(Hash.Order calldata o, uint256 a, Sig.Components calldata c) valid(o,c) internal returns (bool) {
    // TODO

    return true;
  }

  /// @notice Allows a user to exit their Vault filling an offline zcToken exit order
  /// @param o The order being filled
  /// @param a Amount of volume (principal) being filled by the taker's exit
  /// @param c Components of a valid ECDSA signature
  function exitVaultFillingZcTokenExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) valid(o,c) internal returns (bool) {
    // TODO

    return true;
  }

  /// @notice Allows a user to exit their zcTokens by filling an offline vault exit order
  /// @param : o The order being filled
  /// @param : a Amount of volume (interest) being filled by the taker's exit
  /// @param : c Components of a valid ECDSA signature
  function exitZcTokenFillingVaultExit(Hash.Order calldata o, uint256 a, Sig.Components calldata c) valid(o,c) internal returns (bool) {
    // TODO

    return true;
  }

  /// @notice Allows a user to exit their zcTokens by filling an offline zcToken initiate order
  /// @param : o The order being filled
  /// @param : a Amount of volume (interest) being filled by the taker's exit
  /// @param : c Components of a valid ECDSA signature
  function exitZcTokenFillingZcTokenInitiate(Hash.Order calldata o, uint256 a, Sig.Components calldata c) valid(o,c) internal returns (bool) {
    // TODO

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
