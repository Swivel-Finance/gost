// SPDX-License-Identifier: GPL-3.0-or-later

pragma solidity 0.8.13;

contract Redeemer {

  // NOTE: the imported interfaces don't need to be named anything other than what they are:
  // ISwivel
  // IYield
  // etc...

  address public admin;
  address public marketPlace;

  event Redeem(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount);  

  constructor () {
    admin = msg.sender;
  }

  /// @dev redeem method signature for illuminate, tempus, apwine
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being redeemed
  /// @param m maturity of the market being redeemed
  /// @param o owner / tempus pool / apiwine vault ?
  function redeem(uint8 p, address u, uint256 m, address o) public returns (bool) {
    // ...
    return true;
  }

  /// @dev redeem method signature for swivel, yield, element, 
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being redeemed
  /// @param m maturity of the market being redeemed
  function redeem(uint8 p, address u, uint256 m) public returns (bool) {
    // ...
    return true;
  }

  /// @dev redeem method signature for pendle 
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being redeemed
  /// @param m maturity of the market being redeemed
  /// @param i pendle id ?
  function redeem(uint8 p, address u, uint256 m, bytes32 i) public returns (bool) {
    // ...
    return true;
  }

  /// @dev redeem method signature for sense
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being redeemed
  /// @param m maturity of the market being redeemed
  /// @param d sense wut (divider?) ?
  /// @param a sense wut (adapter?) ?
  function redeem(uint8 p, address u, uint256 m, address d, address a) public returns (bool) {
    // ...
    return true;
  }
}
