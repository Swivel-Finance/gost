// SPDX-License-Identifier: AGPL-3.0

/**
  YearnVault is a mock yearn vault with stubs of the methods we need for testing.
*/

pragma solidity 0.8.16;

contract YearnVault {
  /// @dev allows us to dictate return from underlying().
  address private tokenReturn;
  /// @dev allows us to dictate return from pricePerShare().
  uint256 private pricePerShareReturn;
  /// @dev allows us to dictate return from deposit().
  uint256 private depositReturn;
  /// @dev the last amount deposit was called with, msg.sender -> amount
  mapping (address => uint256) public depositCalled;
  /// @dev allows us to dictate return from withdraw().
  uint256 private withdrawReturn;
  /// @dev the last amount withdraw was called with, msg.sender -> amount
  mapping (address => uint256) public withdrawCalled;

  function token() public view returns (address) {
    return tokenReturn;
  }

  function tokenReturns(address a) public {
    tokenReturn = a;
  }

  function pricePerShare() public view returns (uint256) {
    return pricePerShareReturn;
  }

  function pricePerShareReturns(uint256 p) public {
    pricePerShareReturn = p;
  }

  function deposit(uint256 a) public returns (uint256) {
    depositCalled[msg.sender] = a;
    return depositReturn;
  }

  function depositReturns(uint256 a) public {
    depositReturn = a;
  }

  function withdraw(uint256 a) public returns (uint256) {
    withdrawCalled[msg.sender] = a;
    return withdrawReturn;
  }

  function withdrawReturns(uint256 a) public {
    withdrawReturn = a;
  }
}
