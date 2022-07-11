// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

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

interface ICERC20 is IErc20 {
    function mint(uint256) external returns (uint256);
    function borrow(uint256) external returns (uint256);
    function underlying() external view returns (IErc20);
    function totalBorrows() external view returns (uint256);
    function totalFuseFees() external view returns (uint256);
    function repayBorrow(uint256) external returns (uint256);
    function totalReserves() external view returns (uint256);
    function exchangeRateCurrent() external returns (uint256);
    function totalAdminFees() external view returns (uint256);
    function fuseFeeMantissa() external view returns (uint256);
    function adminFeeMantissa() external view returns (uint256);
    function exchangeRateStored() external view returns (uint256);
    function accrualBlockNumber() external view returns (uint256);
    function redeemUnderlying(uint256) external returns (uint256);
    function balanceOfUnderlying(address) external returns (uint256);
    function reserveFactorMantissa() external view returns (uint256);
    function borrowBalanceCurrent(address) external returns (uint256);
    function interestRateModel() external view returns (InterestRateModel);
    function initialExchangeRateMantissa() external view returns (uint256);
    function repayBorrowBehalf(address, uint256) external returns (uint256);
}