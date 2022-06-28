// SPDX-License-Identifier: UNLICENSED

/**
  Euler is a mock IEuler Token token with stubs of the methods we need for testing.
*/

pragma solidity 0.8.13;

contract EulerToken {
  /// @dev the last amount deposit was called with
  mapping (uint256 => uint256) public depositCalled;
  /// @dev the last amount withdraw was called with
  mapping (uint256 => uint256) public withdrawCalled;

  function deposit(uint256 i, uint256 a) public {
    depositCalled[i] = a;
  }

  function withdraw(uint256 i, uint256 a) public {
    withdrawCalled[i] = a;
  }
}
