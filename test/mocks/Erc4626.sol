// SPDX-License-Identifier: UNLICENSED

/**
  Erc4626 is a mock tokenized vault with stubs of the methods we need for testing.
*/

pragma solidity ^0.8.13;

contract Erc4626 {
  struct WithdrawArgs {
    uint256 assets;
    address owner;
  }

  /// @dev allows us to dictate return from asset().
  address private assetReturn;
  /// @dev allows us to dictate return from convertToAssets().
  uint256 private convertToAssetsReturn;
  /// @dev allows us to dictate return from deposit().
  uint256 private depositReturn;
  /// @dev the last amount deposit was called with, "reciever" -> assets
  mapping (address => uint256) public depositCalled;
  /// @dev allows us to dictate return from withdraw().
  uint256 private withdrawReturn;
  /// @dev the last amount withdraw was called with, reciever -> WithdrawArgs
  mapping (address => WithdrawArgs) public withdrawCalled;

  function asset() public view returns (address) {
    return assetReturn;
  }

  function assetReturns(address a) public {
    assetReturn = a;
  }

  // NOTE using the solc feature to squelch the compiler nag about unused vars by not having a var name
  function convertToAssets(uint256) public view returns (uint256) {
    return convertToAssetsReturn;
  }

  // NOTE we refrain from a convertToAssetsCalled here as it's a hard coded 1e26
  // also we would lose our ability to have it interfaced as `view`

  function convertToAssetsReturns(uint256 a) public {
    convertToAssetsReturn = a;
  }

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
