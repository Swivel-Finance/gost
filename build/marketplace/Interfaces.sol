// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface IErc20 {
  function decimals() external returns (uint8);
}

/// @notice The deployed Creator contract creates and deploys instances of the
/// ZcToken and VaultTracker contracts, returning their addresses
interface ICreator {
  function create(
    uint8 p,
    address u,
    uint256 m,
    address c,
    address sw,
    string memory n,
    string memory s,
    uint8 d
  ) external returns (address, address);
}

interface IZcToken {
  function mint(address t, uint256 a) external returns (bool);
  function burn(address f, uint256 a) external returns (bool);
}

interface IVaultTracker {
  function addNotional(address o, uint256 a) external returns (bool);
  function removeNotional(address o, uint256 a) external returns (bool);
  function redeemInterest(address o) external returns (uint256);
  function matureVault(uint256 c) external returns (bool);
  function transferNotionalFrom(address f, address t, uint256 a) external returns (bool);
  function transferNotionalFee(address f, uint256 a) external returns (bool);
  function balancesOf(address o) external view returns (uint256, uint256);
}

interface ISwivel {
  function authRedeem(uint8 p, address u, address c, address t, uint256 a) external returns (bool);
}
