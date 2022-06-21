// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

// TODO we'll likely separate these for adapters...
import './Interfaces.sol';

contract Compound is ICompounding, ICompound {
  // TODO this could be public with an authorized setter if we envision ourselves changing them?
  uint8 constant public PROTOCOL = 0;
  
  /// @param c Compounding token address
  function exchangeRate(address c) external override returns (uint256) {
    // in this case it _is_ a compound token...
    return ICompoundToken(c).exchangeRateCurrent();
  }

  /// @param c Compounding token address
  function underlying(address c) external override returns (address) {
    return ICompoundToken(c).underlying();
  }

  /// @param c Compounding token address
  /// @param a Amount to deposit (mint)
  function deposit(address c,  uint256 a) external override returns (uint256) {
    return ICompoundToken(c).mint(a);
  }

  /// @param c Compounding token address
  /// @param a Amount of underlying to redeem
  function redeemUnderlying(address c, uint256 a) external override returns (uint256) {
    return ICompoundToken(c).redeemUnderlying(a);
  }
}
