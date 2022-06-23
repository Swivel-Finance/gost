// SPDX-License-Identifier: UNLICENSED

/**
  Erc4626 is a mock tokenized vault with stubs of the methods we need for testing.
*/

pragma solidity 0.8.13;

// TODO this could inherit from the ERC20 mock if needed
contract Erc4626 {
  struct WithdrawArgs {
    uint256 assets;
    address owner;
  }

  /// @dev allows us to dictate return from deposit().
  uint256 private depositReturn;
  /// @dev the last amount deposit was called with, "reciever" -> assets
  mapping (address => uint256) public depositCalled;
  /// @dev allows us to dictate return from withdraw().
  uint256 private withdrawReturn;
  /// @dev the last amount withdraw was called with, reciever -> WithdrawArgs
  mapping (address => WithdrawArgs) public withdrawCalled;

  function deposit(uint256 a, address r) public returns (uint256) {
    depositCalled[r] = a;
    return depositReturn;
  }

  function depositReturns(uint256 s) public {
    depositReturn = s;
  }

  function withdraw(uint256 a, address r, address o) public returns (uint256) {
    WithdrawArgs memory args;
    args.assets = a;
    args.owner = o;
    withdrawCalled[r] = args;
    return withdrawReturn;
  }

  function withdrawReturns(uint256 s) public {
    withdrawReturn = s;
  }
}
