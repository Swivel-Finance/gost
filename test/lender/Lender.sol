// SPDX-License-Identifier: GPL-3.0-or-later

pragma solidity 0.8.13;

contract Lender {
  import './Interfaces.sol';
  import './Swivel.sol'; // library of swivel specific structs

  // TODO any reason to allow an admin to change these? (not use constant or immutable)
  address constant SWIVEL = '0xswivel'; // TODO the actual deployed addr (also pass to ctor and be immutable? why over just constant?)

  address public admin;
  address public marketPlace;

  event Lend(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 returned);

  constructor() {
    admin = msg.sender;
  }

  /// @dev lend method signature for both illuminate and yield
  /// @notice Can be called before maturity to lend to Illuminate / Yield ?
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param y yield pool ?
  /// @param a amount of underlying tokens to lend ?
  function lend(uint8 p, address u, uint256 m, address y, uint128 a) public returns (uint256) {
    // ... TODO use the principal given to differentiate if needed (may not in this case)

    // uses iyield token interface...
    IYieldToken yieldToken = IYieldToken(y);
    // yieldtoken.maturity(), yieldtoken.sellBase(...) etc...

    uint256 returned = 0;

    emit Lend(p, u, m, returned)
    return returned;
  }

  /// @dev lend method signature for swivel
  /// @notice can be called before maturity to lend to Swivel while minting Illuminate tokens
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param y yield pool
  /// @param o array of swivel orders being filled
  /// @param a array of amounts of underlying tokens lent to each order in the orders array
  /// @param s array of signatures for each order in the orders array
  function lend(uint8 p, address u, uint256 m, address y, Swivel.Order[] calldata o, uint256[] calldata a, Swivel.Components[] calldata s) public returns (uint256) {
    // ...

    // uses iswivel interface
    ISwivel(SWIVEL).initiate(o, a, s);
    
    // calls lend method for yield?
    // ... = lend(market.Principals.Yield, u, m, y, ...)

    uint256 returned = 0;

    emit Lend(p, u, m, returned)
    return returned;
  }

  /// @dev lend method signature for element and sense
  /// @notice can be called before maturity to lend to Element / Sense ?
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param e element / sense pool ?
  /// @param i element / sense pool id ?
  /// @param a amount ?
  /// @param r minimum amount to return ?
  /// @param d deadline ?
  function lend(uint8 p, address u, uint256 m, address e, bytes32 i, uint128 a, uint256 r, uint256 d) public returns (uint256) {
    // ... TODO use the principal given to differentiate. could be done in line or sent to specefied sub-methods
    uint256 returned = 0;

    emit Lend(p, u, m, returned);
    return returned;
  }

  /// @dev lend method signature for pendle
  /// @notice Can be called before maturity to lend to Pendle while minting Illuminate tokens
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param i pendle id ?
  /// @param a amount ?
  /// @param r minimum amount to return ?
  function lend(uint8 p, address u, uint256 m, bytes32 i, uint256 a, uint256 r) public returns (uint256) {
    // ...
    uint256 returned = 0;

    emit Lend(p, u, m, returned);
    return returned;
  }

  /// @dev lend method signature for tempus
  /// @notice Can be called before maturity to lend to Tempus while minting Illuminate tokens
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param t tempus pool ?
  /// @param amm tempus amm ?
  /// @param a amount ?
  /// @param r minimum amount to return ?
  /// @param d deadline ?
  function lend(uint8 p, address u, uint256 m, address t, address amm, uint256 a, uint256 r, uint256 d) public returns (uint256) {
    // ...
    uint256 returned = 0;

    emit Lend(p, u, m, returned);
    return returned;
  }

  /// @dev lend method signature for apwine
  /// @notice Can be called before maturity to lend to APWine while minting Illuminate tokens
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param w apwine pool ?
  /// @param i apwine pair id ?
  /// @param r minimum amount to return ?
  function lend(uint8 p, address u, uint256 m, address w, uint256 i, uint256 r) public returns (uint256) {
    // ...
    uint256 returned = 0;

    emit Lend(p, u, m, returned);
    return returned;
  }
}
