// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.13;

// the behavioral VaultTracker Interface, Implemented by VaultTracker.sol
interface IVaultTracker {
  function addNotional(address, uint256) external returns (bool);
  function removeNotional(address, uint256) external returns (bool);
  function redeemInterest(address) external returns (uint256);
  function matureVault(uint256) external returns (bool);
  function transferNotionalFrom(address, address, uint256) external returns (bool);
  function transferNotionalFee(address, uint256) external returns (bool);
  function rates() external view returns (uint256, uint256);
  function balancesOf(address) external view returns (uint256, uint256);
}
