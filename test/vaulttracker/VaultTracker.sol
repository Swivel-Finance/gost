// SPDX-License-Identifier: UNLICENSED

// TODO update to 0.8.4 (or whatever latest is...)

// NOTE the pattern [underlying, [maturity,], cToken, ...]

pragma solidity 0.8.0;

import "./Abstracts.sol";

contract VaultTracker {
  address public admin;
  uint256 public maturity;
  bool matured;
  uint256 maturityRate;
  address public cTokenAddr;

  struct Vault {
    uint256 notional;
    uint256 redeemable;
    uint256 exchangeRate;
  }

  mapping(address => Vault) public vaults;

  event RedeemInterest(address indexed owner, uint256 indexed amount);

  /// @param m Maturity timestamp of the new market
  /// @param c cToken address associated with underlying for the new market
  constructor(uint256 m, address c) {
    admin = msg.sender;
    maturity = m;
    cTokenAddr = c;
  }

  /// @notice ...
  /// @param o Address that owns a timelock in the vault
  /// @param a Amount to ...
  function addNotional(address o, uint256 a) public onlyAdmin(admin) returns (bool) {
    Vault memory vault = vaults[o];
    CErc20 cToken = CErc20(cTokenAddr);

    if (vault.notional > 0) {
      uint256 interest;

      // If market has matured, calculate marginal interest between the maturity rate and previous position exchange rate
      // Otherwise, calculate marginal exchange rate between current and previous exchange rate.
      if (matured == true) {
        // Calculate marginal interest
        uint256 yield = ((maturityRate * 1e26) / vault.exchangeRate) - 1e26;
        interest = (yield * vault.notional) / 1e26;
      }
      else {
        // Calculate marginal interest
        uint256 yield = ((cToken.exchangeRateCurrent() * 1e26) / vault.exchangeRate) - 1e26;
        interest = (yield * vault.notional) / 1e26;
      }

      // Add interest and amount to position, reset cToken exchange rate
      vault.redeemable += interest;
      vault.notional += a;
      vault.exchangeRate = cToken.exchangeRateCurrent();
      vaults[o] = vault;
    }
    else {
      vault.notional = a;
      vault.exchangeRate = cToken.exchangeRateCurrent();
      vaults[o] = vault;
    }
    return true;
  }

  /// @notice ...
  /// @param o Address that owns a timelock in the vault
  /// @param a Amount to ...
  function removeNotional(address o, uint256 a) public onlyAdmin(admin) returns (bool) {
    Vault memory vault = vaults[o];

    require(vault.notional >= a, "Amount exceeds vault balance");

    uint256 interest;
    CErc20 cToken = CErc20(cTokenAddr);

    // If market has matured, calculate marginal interest between the maturity rate and previous position exchange rate
    // Otherwise, calculate marginal exchange rate between current and previous exchange rate.
    if (matured == true) {
      // Calculate marginal interest
      uint256 yield = ((maturityRate * 1e26) / vault.exchangeRate) - 1e26;
      interest = (yield * vault.notional) / 1e26;
    }
    else {
      // Calculate marginal interest
      uint256 yield = ((cToken.exchangeRateCurrent() * 1e26) / vault.exchangeRate) - 1e26;
      interest = (yield * vault.notional) / 1e26;
    }

    // Remove amount from position, Add interest to position, reset cToken exchange rate
    vault.redeemable += interest;
    vault.notional -= a;
    vault.exchangeRate = cToken.exchangeRateCurrent();
    vaults[o] = vault;
    return true;
  }

  /// @notice ...
  /// @param o Address that owns a timelock in the vault
  function redeemInterest(address o) public onlyAdmin(admin) returns (uint256) {
    Vault memory vault = vaults[o];
    uint256 redeemAmount = vault.redeemable;

    uint256 interest;
    CErc20 cToken = CErc20(cTokenAddr);

    // If market has matured, calculate marginal interest between the maturity rate and previous position exchange rate
    // Otherwise, calculate marginal exchange rate between current and previous exchange rate.
    if (matured == true) {
      // Calculate marginal interest
      uint256 yield = ((maturityRate * 1e26) / vault.exchangeRate) - 1e26;
      interest = (yield * vault.notional) / 1e26;
    }
    else {
      // Calculate marginal interest
      uint256 yield = ((cToken.exchangeRateCurrent() * 1e26) / vault.exchangeRate) - 1e26;
      interest = (yield * vault.notional) / 1e26;
    }
    // Add marginal interest to previously accrued redeemable interest
    redeemAmount += interest;

    vault.exchangeRate = cToken.exchangeRateCurrent();
    vault.redeemable = 0;
    vaults[o] = vault;

    emit RedeemInterest(o, redeemAmount);
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
  /// @param t Address recipient of the amount
  /// @param a Amount to transfer
  function transfer(address t, uint256 a) public returns (bool) {
    Vault memory vault = vaults[msg.sender];

    require(vault.notional >= a, "Amount exceeds vault balance");

    uint256 interest;
    CErc20 cToken = CErc20(cTokenAddr);

    // If market has matured, calculate marginal interest between the maturity rate and previous position exchange rate
    // Otherwise, calculate marginal exchange rate between current and previous exchange rate.
    if (matured == true) {
      // Calculate marginal interest
      uint256 yield = ((maturityRate * 1e26) / vault.exchangeRate) - 1e26;
      interest = (yield * vault.notional) / 1e26;
    }
    else {
      // Calculate marginal interest
      uint256 yield = ((cToken.exchangeRateCurrent() * 1e26) / vault.exchangeRate) - 1e26;
      interest = (yield * vault.notional) / 1e26;
    }

    // Remove amount from position, Add interest to position, reset cToken exchange rate
    vault.redeemable += interest;
    vault.notional -= a;
    vault.exchangeRate = cToken.exchangeRateCurrent();
    vaults[msg.sender] = vault;

    // transfer notional to address "t", calculate interest if necessary
    Vault memory newVault = vaults[t];

    if (newVault.notional > 0) {
      uint256 newVaultInterest;

      // If market has matured, calculate marginal interest between the maturity rate and previous position exchange rate
      // Otherwise, calculate marginal exchange rate between current and previous exchange rate.
      if (matured == true) {
        // Calculate marginal interest
        uint256 yield = ((maturityRate * 1e26) / newVault.exchangeRate) - 1e26;
        newVaultInterest = (yield * newVault.notional) / 1e26;
      }
      else {
        // Calculate marginal interest
        uint256 yield = ((cToken.exchangeRateCurrent() * 1e26) / newVault.exchangeRate) - 1e26;
        newVaultInterest = (yield * newVault.notional) / 1e26;
      }

      // Add interest and amount to position, reset cToken exchange rate
      newVault.redeemable += newVaultInterest;
      newVault.notional += a;
      newVault.exchangeRate = cToken.exchangeRateCurrent();
      vaults[t] = newVault;
    }
    else {
      newVault.notional += a;
      newVault.exchangeRate = cToken.exchangeRateCurrent();
      vaults[t] = newVault;
    }
    return true;
  }

  /// @notice ...
  /// @param o Address that owns a timelock in the vault
  function balancesOf(address o) public view returns (uint256, uint256) {
    return (vaults[o].notional, vaults[o].redeemable);
  }

  modifier onlyAdmin(address a) {
    require(msg.sender == a, 'sender must be admin');
    _;
  }
}
