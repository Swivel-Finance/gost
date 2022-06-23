// SPDX-License-Identifier: UNLICENSED

/**
  CompoundToken is a mock compound token with stubs of the methods we need for testing.
*/

pragma solidity 0.8.13;

// TODO this could inherit from the ERC20 mock if needed
contract CompoundToken {
  /// @dev allows us to dictate return from mint().
  uint256 private mintReturn;
  /// @dev the last amount mint was called with, msg.sender -> amount
  mapping (address => uint256) public mintCalled;
  /// @dev allows us to dictate return from redeemUnderlying().
  uint256 private redeemUnderlyingReturn;
  /// @dev the last amount redeemUnderlying was called with, msg.sender -> amount
  mapping (address => uint256) public redeemUnderlyingCalled;

  address private underlyingReturn;
  /// @dev allows us to dictate return from exchangeRateCurrent().
  uint256 private exchangeRateCurrentReturn;
  uint256 private supplyRatePerBlockReturn;

  function mint(uint256 n) public returns (uint256) {
    mintCalled[msg.sender] = n;
    return mintReturn;
  }

  function mintReturns(uint256 n) public {
    mintReturn = n;
  }

  function redeemUnderlying(uint256 n) public returns (uint256) {
    redeemUnderlyingCalled[msg.sender] = n;
    return redeemUnderlyingReturn;
  }

  function redeemUnderlyingReturns(uint256 n) public {
    redeemUnderlyingReturn = n;
  }

  function underlying() public view returns (address) {
    return underlyingReturn;
  }

  function underlyingReturns(address a) public {
    underlyingReturn = a;
  }

  function exchangeRateCurrent() public view returns (uint256) {
    return exchangeRateCurrentReturn;
  }

  function exchangeRateCurrentReturns(uint256 n) public {
    exchangeRateCurrentReturn = n;
  }

  function supplyRatePerBlock() public view returns (uint256) {
    return supplyRatePerBlockReturn;
  }

  function supplyRatePerBlockReturns(uint256 n) public {
    supplyRatePerBlockReturn = n;
  }
}
