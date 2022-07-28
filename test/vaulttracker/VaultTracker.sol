// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.13;

import './Compounding.sol';

contract VaultTracker {
  /// @notice A single custom error capable of indicating a wide range of detected errors by providing
  /// an error code value whose string representation is documented <here>, and any possible other values
  /// that are pertinent to the error.
  error Exception(uint8, uint256, uint256, address, address);

  struct Vault {
    uint256 notional;
    uint256 redeemable;
    uint256 exchangeRate;
  }

  mapping(address => Vault) public vaults;

  address public immutable admin;
  address public immutable cTokenAddr;
  address public immutable swivel;
  uint256 public immutable maturity;
  uint256 public maturityRate;
  uint8 public immutable protocol;

  /// @param m Maturity timestamp associated with this vault
  /// @param c Compounding Token address associated with this vault
  /// @param s Address of the deployed swivel contract
  constructor(uint8 p, uint256 m, address c, address s) {
    admin = msg.sender;
    protocol = p;
    maturity = m;
    cTokenAddr = c;
    swivel = s;

    // instantiate swivel's vault (unblocking transferNotionalFee)
    vaults[s] = Vault({
      notional: 0,
      redeemable: 0,
      exchangeRate: Compounding.exchangeRate(p, c)
    });
  }

  /// @notice Adds notional to a given address
  /// @param o Address that owns a vault
  /// @param a Amount of notional added
  function addNotional(address o, uint256 a) external authorized(admin) returns (bool) {
    uint256 exchangeRate = Compounding.exchangeRate(protocol, cTokenAddr);

    Vault memory vlt = vaults[o];

    if (vlt.notional > 0) {
      uint256 yield;

      // if market has matured, calculate marginal interest between the maturity rate and previous position exchange rate
      // otherwise, calculate marginal exchange rate between current and previous exchange rate.
      if (maturityRate > 0) { // Calculate marginal interest
        yield = ((maturityRate * 1e26) / vlt.exchangeRate) - 1e26;
      } else {
        yield = ((exchangeRate * 1e26) / vlt.exchangeRate) - 1e26;
      }

      uint256 interest = (yield * vlt.notional) / 1e26;
      // add interest and amount to position, reset cToken exchange rate
      vlt.redeemable += interest;
      vlt.notional += a;
    } else {
      vlt.notional = a;
    }

    vlt.exchangeRate = exchangeRate;
    vaults[o] = vlt;

    return true;
  }

  /// @notice Removes notional from a given address
  /// @param o Address that owns a vault
  /// @param a Amount of notional to remove
  function removeNotional(address o, uint256 a) external authorized(admin) returns (bool) {

    Vault memory vlt = vaults[o];

    if (a > vlt.notional) {
      revert Exception(31, a, vlt.notional, address(0), address(0));
    }

    uint256 exchangeRate = Compounding.exchangeRate(protocol, cTokenAddr);

    // if market has matured, calculate marginal interest between the maturity rate and previous position exchange rate
    // otherwise, calculate marginal exchange rate between current and previous exchange rate.
    uint256 yield;
    if (maturityRate > 0) { // Calculate marginal interest
      yield = ((maturityRate * 1e26) / vlt.exchangeRate) - 1e26;
    } else {
      // calculate marginal interest
      yield = ((exchangeRate * 1e26) / vlt.exchangeRate) - 1e26;
    }

    uint256 interest = (yield * vlt.notional) / 1e26;
    // remove amount from position, Add interest to position, reset cToken exchange rate
    vlt.redeemable += interest;
    vlt.notional -= a;
    vlt.exchangeRate = exchangeRate;

    vaults[o] = vlt;

    return true;
  }

  /// @notice Redeem's interest accrued by a given address
  /// @param o Address that owns a vault
  function redeemInterest(address o) external authorized(admin) returns (uint256) {

    Vault memory vlt = vaults[o];

    uint256 redeemable = vlt.redeemable;
    uint256 exchangeRate = Compounding.exchangeRate(protocol, cTokenAddr);

    // if market has matured, calculate marginal interest between the maturity rate and previous position exchange rate
    // otherwise, calculate marginal exchange rate between current and previous exchange rate.
    uint256 yield;
    if (maturityRate > 0) { // Calculate marginal interest
      yield = ((maturityRate * 1e26) / vlt.exchangeRate) - 1e26;
    } else {
      // calculate marginal interest
      yield = ((exchangeRate * 1e26) / vlt.exchangeRate) - 1e26;
    }

    uint256 interest = (yield * vlt.notional) / 1e26;

    vlt.exchangeRate = exchangeRate;
    vlt.redeemable = 0;

    vaults[o] = vlt;

    // return adds marginal interest to previously accrued redeemable interest
    return (redeemable + interest);
  }

  /// @notice Matures the vault
  /// @param c The current cToken exchange rate
  function matureVault(uint256 c) external authorized(admin) returns (bool) {
    maturityRate = c;
    return true;
  }

  /// @notice Transfers notional from one address to another
  /// @param f Owner of the amount
  /// @param t Recipient of the amount
  /// @param a Amount to transfer
  function transferNotionalFrom(address f, address t, uint256 a) external authorized(admin) returns (bool) {
    if (f == t) {
      revert Exception(32, 0, 0, f, t);
    }

    Vault memory from = vaults[f];
    Vault memory to = vaults[t];

    if (a > from.notional) {
      revert Exception(31, a, from.notional, address(0), address(0));
    }

    uint256 exchangeRate = Compounding.exchangeRate(protocol, cTokenAddr);

    // if market has matured, calculate marginal interest between the maturity rate and previous position exchange rate
    // otherwise, calculate marginal exchange rate between current and previous exchange rate.
    uint256 yield;
    if (maturityRate > 0) { 
      // calculate marginal interest
      yield = ((maturityRate * 1e26) / from.exchangeRate) - 1e26;
    } else {
      yield = ((exchangeRate * 1e26) / from.exchangeRate) - 1e26;
    }

    uint256 interest = (yield * from.notional) / 1e26;
    // remove amount from position, Add interest to position, reset cToken exchange rate
    from.redeemable += interest;
    from.notional -= a;
    from.exchangeRate = exchangeRate;

    vaults[f] = from;

    // transfer notional to address "t", calculate interest if necessary
    if (to.notional > 0) {
      // if market has matured, calculate marginal interest between the maturity rate and previous position exchange rate
      // otherwise, calculate marginal exchange rate between current and previous exchange rate.
      if (maturityRate > 0) { 
        // calculate marginal interest
        yield = ((maturityRate * 1e26) / to.exchangeRate) - 1e26;
      } else {
        yield = ((exchangeRate * 1e26) / to.exchangeRate) - 1e26;
      }

      uint256 newVaultInterest = (yield * to.notional) / 1e26;
      // add interest and amount to position, reset cToken exchange rate
      to.redeemable += newVaultInterest;
      to.notional += a;
    } else {
      to.notional = a;
    }

    to.exchangeRate = exchangeRate;
    vaults[t] = to;

    return true;
  }

  /// @notice Transfers, in notional, a fee payment to the Swivel contract without recalculating marginal interest for the owner
  /// @param f Owner of the amount
  /// @param a Amount to transfer
  function transferNotionalFee(address f, uint256 a) external authorized(admin) returns(bool) {
    Vault memory oVault = vaults[f];
    Vault memory sVault = vaults[swivel];

    // remove notional from its owner
    oVault.notional -= a;

    uint256 exchangeRate = Compounding.exchangeRate(protocol, cTokenAddr);

    // check if exchangeRate has been stored already this block. If not, calculate marginal interest + store exchangeRate
    uint256 yield;
    if (sVault.exchangeRate != exchangeRate) {
      // if market has matured, calculate marginal interest between the maturity rate and previous position exchange rate
      // otherwise, calculate marginal exchange rate between current and previous exchange rate.
      if (maturityRate > 0) { 
        // calculate marginal interest
          yield = ((maturityRate * 1e26) / sVault.exchangeRate) - 1e26;
      } else {
          yield = ((exchangeRate * 1e26) / sVault.exchangeRate) - 1e26;
      }
      uint256 interest = (yield * sVault.notional) / 1e26;
      // add interest and amount, reset cToken exchange rate
      sVault.redeemable += interest;
      sVault.exchangeRate = exchangeRate;
    }
    // add notional to swivel's vault
    sVault.notional += a;
    // store the adjusted vaults
    vaults[swivel] = sVault;
    vaults[f] = oVault;
    return true;
  }

  /// @notice Returns both relevant balances for a given user's vault
  /// @param o Address that owns a vault
  function balancesOf(address o) external view returns (uint256, uint256) {
    Vault memory vault = vaults[o];
    return (vault.notional, vault.redeemable);
  }

  modifier authorized(address a) {
    if(msg.sender != a) {
      revert Exception(0, 0, 0, msg.sender, a);
    }
    _;
  }
}
