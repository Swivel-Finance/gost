// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface IErc20 {
  function decimals() external returns (uint8);
}

interface IZcToken{
  function mint(address t, uint256 a) external returns (bool);
  function burn(address f, uint256 a) external returns (bool);
}

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
  ) external returns (address zc, address vt);
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