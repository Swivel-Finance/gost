// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract MarketPlace {

  /// for any given deployed MarketPlace these are the possible principal tokens available
  enum Principals {
    Illuminate,
    Swivel,
    Yield,
    Element,
    Pendle,
    Tempus,
    Sense,
    Apwine,
  }

  /// markets are defined by a market pair which point to a fixed length array of principal token addresses.
  /// the principal tokens those addresses represent correspond to their Principals enum value, thus the
  /// array should be ordered in that way
  mapping (address => mapping (uint256 => address[8])) public markets;

  address public admin;

  /// adress of the deployed illuminate contract that my act as authorized within this contract
  address public illuminate;

  event Create(address indexed underlying, uint256 indexed maturity)

  constructor () {
    admin = msg.sender;
    // ...
  }

  /// @param u underlying
  /// @param m maturity
  /// @param t principal token addresses for this market
  /// @param n name for the illuminate token
  /// @param d decimals for the illuminate token
  function createMarket(address u, uint256 m, address[8] memory t, string calldata n, uint8 d) external authorized(admin) returns (bool) {
    require(markets[u][m][Principals.Illuminate] == address(0), 'market exists'); 
    // ...
    return true
  }
}
