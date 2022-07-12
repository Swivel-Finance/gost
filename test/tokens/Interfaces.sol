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
  function totalSupply() external view returns (uint256);
  function balanceOf(address account) external view returns (uint256);  
  function decimals() external returns (uint8);
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