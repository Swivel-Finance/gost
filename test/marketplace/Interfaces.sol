// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface IErc20 {
  event Transfer(address indexed from, address indexed to, uint256 value);
  event Approval(address indexed owner, address indexed spender, uint256 value);
  function totalSupply() external view returns (uint256);
  function balanceOf(address account) external view returns (uint256);
  function transfer(address to, uint256 amount) external returns (bool);
  function allowance(address owner, address spender) external view returns (uint256);
  function approve(address spender, uint256 amount) external returns (bool);
  function decimals() external returns (uint8);
  function transferFrom(
      address from,
      address to,
      uint256 amount
  ) external returns (bool);
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

interface InterestRateModel {
    function getBorrowRate(
        uint256,
        uint256,
        uint256
    ) external view returns (uint256);
    function getSupplyRate(
        uint256,
        uint256,
        uint256,
        uint256
    ) external view returns (uint256);
}

interface ICERC20 is IErc20 {
    function underlying() external view returns (IErc20);
    function totalBorrows() external view returns (uint256);
    function totalReserves() external view returns (uint256);
    function exchangeRateStored() external view returns (uint256);
    function reserveFactorMantissa() external view returns (uint256);
    function interestRateModel() external view returns (InterestRateModel);
    function initialExchangeRateMantissa() external view returns (uint256);
    function accrualBlockNumber() external view returns (uint256);
    function totalAdminFees() external view returns (uint256);
    function totalFuseFees() external view returns (uint256);
    function adminFeeMantissa() external view returns (uint256);
    function fuseFeeMantissa() external view returns (uint256);
}