// SPDX-License-Identifier: UNLICENSED

/**
  Compound is the mock for a deployed adapter for Compound.
  NOTE: that it implements ICompounding, thus fulfulling the ICompound is ICompounding agreement
*/

pragma solidity 0.8.13;

contract Compound {
  /// @dev allows us to dictate return from mint().
  uint256 private mintReturn;
  /// @dev the last amount mint was called with
  uint256 public mintCalled;
  /// @dev allows us to dictate return from redeemUnderlying().
  uint256 private redeemUnderlyingReturn;
  /// @dev the last amount redeemUnderlying was called with
  uint256 public redeemUnderlyingCalled;

  address public underlyingCalled;
  address private underlyingReturn;

  address public exchangeRateCalled;
  /// @dev allows us to dictate return from exchangeRate().
  uint256 private exchangeRateReturn;

  function mint(uint256 n) public returns (uint256) {
    mintCalled = n;
    return mintReturn;
  }

  function mintReturns(uint256 n) public {
    mintReturn = n;
  }

  function redeemUnderlying(uint256 n) public returns (uint256) {
    redeemUnderlyingCalled = n;
    return redeemUnderlyingReturn;
  }

  function redeemUnderlyingReturns(uint256 n) public {
    redeemUnderlyingReturn = n;
  }

  function underlying(address a) public returns (address) {
    underlyingCalled = a;
    return underlyingReturn;
  }

  function underlyingReturns(address a) public {
    underlyingReturn = a;
  }

  function exchangeRate(address a) public returns (uint256) {
    exchangeRateCalled = a;
    return exchangeRateReturn;
  }

  function exchangeRateReturns(uint256 n) public {
    exchangeRateReturn = n;
  }
}
