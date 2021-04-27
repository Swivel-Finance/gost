// SPDX-License-Identifier: UNLICENSED

// TODO update to 0.8.4 (or whatever latest is...)

// NOTE the pattern [underlying, [maturity,], cToken, ...]

pragma solidity 0.8.0;

contract Vault {
  // TODO check the visibilibty of all of these...
  address public admin; 
  address public underlying; 
  uint256 public maturity; 
  uint256 public maturityRate; 
  bool public matured;
  address public cTokenAddr;

  struct TimeLock {
    uint256 notional;
    uint256 redeemable;
    uint256 exchangeRate;
  }

  mapping(address => TimeLock) pubilc vaults;

  // TODO events?

  /// @param u Underlying token address associated with the new vaul
  /// @param m Maturity timestamp of the new market
  /// @param c cToken address associated with underlying for the new market 
  constructor(address u, uint256 m, address c) {
    admin = msg.sender;
    underlying = u;
    maturity = m;
    cTokenAddr = c;
  }

  /// @notice ...
  /// @param o Address that owns a timelock in the vault
  /// @param a Amount to ...
  function addNotional(address o, uint256 a) public restricted(admin) returns (bool) {
    locked memory TimeLock = vaults[o];
    if (locked.notional > 0 ) {
      // do stuff
    } else {
      // do other stuff
    }

    return true;
  }

  /// @notice ...
  /// @param o Address that owns a timelock in the vault
  /// @param a Amount to ...
  function removeNotional(address o, uint256 a) public restricted(admin) returns (bool) {
    // TODO

    return true;
  }

  /// @notice ...
  /// @param o Address that owns a timelock in the vault
  function redeemInterest(address o) public restricted(admin) returns (uint256) {
    // TODO
  }

  /// @notice ...
  function matureVault() public returns (bool) {
    require(block.timestamp >= maturity, 'maturity has not been reached');  
    matured = true;
    // NOTE removed the stored instance of the CErc20...
    CErc20 cToken = CErc20(cTokenAddr);
    maturityRate = cToken.exchangeRateCurrent();
    return true;
  }

  /// @notice ...
  /// @param t Address recipient of the amount
  /// @param a Amount to transfer
  function transfer(address t, uint256 a) public returns (bool) {
    // TODO
  }

  /// @notice ...
  /// @param o Address that owns a timelock in the vault
  function balancesOf(address o) public view returns (uint256, uint256) {
    return (vaults[o].notional, vaults[o].redeemable);
  }

  modifier restricted(address a) {
    require(msg.sender == a, 'sender must be admin')
    _;
  }
}
