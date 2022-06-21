// SPDX-License-Identifier: UNLICENSED

/**
  Compound is the mock for a deployed adapter for Compound.
  NOTE: that it implements ICompounding, thus fulfulling the ICompound is ICompounding agreement
*/

pragma solidity 0.8.13;

contract Compound {
  uint8 public PROTOCOL = 0;
  /// @dev allows us to dictate return from deposit().
  uint256 private depositReturn;
  /// @dev the last amount deposit was called with
  mapping (address => uint256) public depositCalled;
  /// @dev allows us to dictate return from redeemUnderlying().
  uint256 private redeemUnderlyingReturn;
  /// @dev the last amount redeemUnderlying was called with
  mapping (address => uint256) public redeemUnderlyingCalled;

  address public underlyingCalled;
  address private underlyingReturn;

  address public exchangeRateCalled;
  /// @dev allows us to dictate return from exchangeRate().
  uint256 private exchangeRateReturn;

  function deposit(address c, uint256 n) public returns (uint256) {
    depositCalled[c] = n;
    return depositReturn;
  }

  function depositReturns(uint256 n) public {
    depositReturn = n;
  }

  function redeemUnderlying(address c, uint256 n) public returns (uint256) {
    redeemUnderlyingCalled[c] = n;
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
