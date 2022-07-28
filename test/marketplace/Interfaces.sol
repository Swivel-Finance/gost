// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.13;

interface IErc20 {
  function decimals() external returns (uint8);
}

/// @notice The deployed Creator contract creates and deploys instances of the
/// ZcToken and VaultTracker contracts, returning their addresses
interface ICreator {
  function create(
    uint8,
    address,
    uint256,
    address,
    address,
    string calldata,
    string calldata,
    uint8
  ) external returns (address, address);
}

interface IZcToken {
  function mint(address, uint256) external returns (bool);
  function burn(address, uint256) external returns (bool);
}

interface IVaultTracker {
  function addNotional(address, uint256) external returns (bool);
  function removeNotional(address, uint256) external returns (bool);
  function redeemInterest(address) external returns (uint256);
  function matureVault(uint256) external returns (bool);
  function transferNotionalFrom(address, address, uint256) external returns (bool);
  function transferNotionalFee(address, uint256) external returns (bool);
  function balancesOf(address) external view returns (uint256, uint256);
}

interface ISwivel {
  function authRedeem(uint8, address, address, address, uint256) external returns (bool);
}
