// SPDX-License-Identifier: UNLICENSED

/**
  Compound is the mock for a deployed adapter for Compound.
  NOTE: Adapters are read-only. Implements ICompounding
*/

pragma solidity 0.8.13;

contract Compound {
  uint8 public PROTOCOL = 0;
  address public underlyingCalled;
  address private underlyingReturn;

  address public exchangeRateCalled;
  /// @dev allows us to dictate return from exchangeRate().
  uint256 private exchangeRateReturn;

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
