// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.4;

import "./Abstracts.sol";

contract VaultTracker {
  address public admin; // marketplace
  uint256 public maturity;
  bool public matured;
  uint256 public maturityRate;
  address public cTokenAddr;

  struct Vault {
    uint256 notional;
    uint256 redeemable;
    uint256 exchangeRate;
  }

  mapping(address => Vault) public vaults;

  /// @param m Maturity timestamp of the new market
  /// @param c cToken address associated with underlying for the new market
  constructor(uint256 m, address c) {
    admin = msg.sender;
    maturity = m;
    cTokenAddr = c;
  }

  /// @notice ...
  /// @param o Address that owns a vault
  /// @param a Amount of notional added
  function addNotional(address o, uint256 a) public onlyAdmin(admin) returns (bool) {
    uint256 exchangeRate = CErc20(cTokenAddr).exchangeRateCurrent();

    if (vaults[o].notional > 0) {
      uint256 yield;
      uint256 interest;

      // If market has matured, calculate marginal interest between the maturity rate and previous position exchange rate
      // Otherwise, calculate marginal exchange rate between current and previous exchange rate.
      if (matured == true) {
        // Calculate marginal interest
        yield = ((maturityRate * 1e26) / vaults[o].exchangeRate) - 1e26;
        interest = (yield * vaults[o].notional) / 1e26;
      }
      else {
        // Calculate marginal interest
        yield = ((exchangeRate * 1e26) / vaults[o].exchangeRate) - 1e26;
        interest = (yield * vaults[o].notional) / 1e26;
      }

      // Add interest and amount to position, reset cToken exchange rate
      vaults[o].redeemable += interest;
      vaults[o].notional += a;
      vaults[o].exchangeRate = exchangeRate;
    }
    else {
      vaults[o].notional = a;
      vaults[o].exchangeRate = exchangeRate;
    }
    return true;
  }

  /// @notice ...
  /// @param o Address that owns a vault
  /// @param a Amount of notional to remove
  function removeNotional(address o, uint256 a) public onlyAdmin(admin) returns (bool) {
    require(vaults[o].notional >= a, "amount exceeds vault balance");

    uint256 yield;
    uint256 interest;
    uint256 exchangeRate = CErc20(cTokenAddr).exchangeRateCurrent();

    // If market has matured, calculate marginal interest between the maturity rate and previous position exchange rate
    // Otherwise, calculate marginal exchange rate between current and previous exchange rate.
    if (matured == true) {
      // Calculate marginal interest
      yield = ((maturityRate * 1e26) / vaults[o].exchangeRate) - 1e26;
      interest = (yield * vaults[o].notional) / 1e26;
    }
    else {
      // Calculate marginal interest
      yield = ((exchangeRate * 1e26) / vaults[o].exchangeRate) - 1e26;
      interest = (yield * vaults[o].notional) / 1e26;
    }

    // Remove amount from position, Add interest to position, reset cToken exchange rate
    vaults[o].redeemable += interest;
    vaults[o].notional -= a;
    vaults[o].exchangeRate = exchangeRate;
    return true;
  }

  /// @notice ...
  /// @param o Address that owns a vault
  function redeemInterest(address o) external onlyAdmin(admin) returns (uint256) {
    uint256 redeemAmount = vaults[o].redeemable;

    uint256 yield;
    uint256 interest;
    uint256 exchangeRate = CErc20(cTokenAddr).exchangeRateCurrent();

    // If market has matured, calculate marginal interest between the maturity rate and previous position exchange rate
    // Otherwise, calculate marginal exchange rate between current and previous exchange rate.
    if (matured == true) {
      // Calculate marginal interest
      yield = ((maturityRate * 1e26) / vaults[o].exchangeRate) - 1e26;
      interest = (yield * vaults[o].notional) / 1e26;
    }
    else {
      // Calculate marginal interest
      yield = ((exchangeRate * 1e26) / vaults[o].exchangeRate) - 1e26;
      interest = (yield * vaults[o].notional) / 1e26;
    }
    // Add marginal interest to previously accrued redeemable interest
    redeemAmount += interest;

    vaults[o].exchangeRate = exchangeRate;
    vaults[o].redeemable = 0;

    return(redeemAmount);
  }

  /// @notice ...
  function matureVault() public returns (bool) {
    require(block.timestamp >= maturity, 'maturity has not been reached');
    matured = true;
    maturityRate = CErc20(cTokenAddr).exchangeRateCurrent();
    return true;
  }

  /// @notice ...
  /// @param f Owner of the amount
  /// @param t Recipient of the amount
  /// @param a Amount to transfer
  function transferNotional(address f, address t, uint256 a) external onlyAdmin(admin) returns (bool) {
    require(vaults[f].notional >= a, "amount exceeds available balance");

    uint256 yield;
    uint256 interest;
    uint256 exchangeRate = CErc20(cTokenAddr).exchangeRateCurrent();

    // If market has matured, calculate marginal interest between the maturity rate and previous position exchange rate
    // Otherwise, calculate marginal exchange rate between current and previous exchange rate.
    if (matured == true) {
      // Calculate marginal interest
      yield = ((maturityRate * 1e26) / vaults[f].exchangeRate) - 1e26;
      interest = (yield * vaults[f].notional) / 1e26;
    }
    else {
      // Calculate marginal interest
      yield = ((exchangeRate * 1e26) / vaults[f].exchangeRate) - 1e26;
      interest = (yield * vaults[f].notional) / 1e26;
    }

    // Remove amount from position, Add interest to position, reset cToken exchange rate
    vaults[f].redeemable += interest;
    vaults[f].notional -= a;
    vaults[f].exchangeRate = exchangeRate;

    // transfer notional to address "t", calculate interest if necessary
    if (vaults[t].notional > 0) {
      uint256 newVaultInterest;

      // If market has matured, calculate marginal interest between the maturity rate and previous position exchange rate
      // Otherwise, calculate marginal exchange rate between current and previous exchange rate.
      if (matured == true) {
        // Calculate marginal interest
        yield = ((maturityRate * 1e26) / vaults[t].exchangeRate) - 1e26;
        newVaultInterest = (yield * vaults[t].notional) / 1e26;
      }
      else {
        // Calculate marginal interest
        yield = ((exchangeRate * 1e26) / vaults[t].exchangeRate) - 1e26;
        newVaultInterest = (yield * vaults[t].notional) / 1e26;
      }

      // Add interest and amount to position, reset cToken exchange rate
      vaults[t].redeemable += newVaultInterest;
      vaults[t].notional += a;
      vaults[t].exchangeRate = exchangeRate;
    }
    else {
      vaults[t].notional += a;
      vaults[t].exchangeRate = exchangeRate;
    }

    return true;
  }

  /// @notice ...
  /// @param f Address of the sender
  /// @param t Address of the recipient
  /// @param a Amount to transfer
  function transferNotionalFrom(address f, address t, uint256 a) external onlyAdmin(admin) returns (bool) {
    require(removeNotional(f, a), 'remove notional failed');
    require(addNotional(t, a), 'add notional failed');
    return true;
  }

  /// @notice ...
  /// @param o Address that owns a vault
  function balancesOf(address o) public view returns (uint256, uint256) {
    return (vaults[o].notional, vaults[o].redeemable);
  }

  modifier onlyAdmin(address a) {
    require(msg.sender == a, 'sender must be admin');
    _;
  }
}
