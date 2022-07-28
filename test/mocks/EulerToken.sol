// SPDX-License-Identifier: UNLICENSED

/**
  Euler is a mock IEuler Token token with stubs of the methods we need for testing.
*/

pragma solidity ^0.8.13;

contract EulerToken {
  /// @dev allows us to dictate return from underlyingAsset().
  address private underlyingAssetReturn;
  /// @dev allows us to dictate the return from convertBalanceToUnderlying()
  uint256 private convertBalanceToUnderlyingReturn;
  /// @dev the last amount deposit was called with
  mapping (uint256 => uint256) public depositCalled;
  /// @dev the last amount withdraw was called with
  mapping (uint256 => uint256) public withdrawCalled;

  function underlyingAsset() public view returns (address) {
    return underlyingAssetReturn;
  }

  function underlyingAssetReturns(address a) public {
    underlyingAssetReturn = a;
  }

  // NOTE using the solc feature to squelch the compiler nag about unused vars by not having a var name
  function convertBalanceToUnderlying(uint256) public view returns (uint256) {
    return convertBalanceToUnderlyingReturn;
  }

  function convertBalanceToUnderlyingReturns(uint256 n) public {
    convertBalanceToUnderlyingReturn = n;
  }

  function deposit(uint256 i, uint256 a) public {
    depositCalled[i] = a;
  }

  function withdraw(uint256 i, uint256 a) public {
    withdrawCalled[i] = a;
  }
}
