pragma solidity >=0.8.4;

import {PErc20} from "../tokens/PErc20.sol";

import {InterestRateModel} from "./InterestRateModel.sol";

abstract contract CERC20 is Erc20 {
    function mint(uint256) external virtual returns (uint256);

    function borrow(uint256) external virtual returns (uint256);

    function underlying() external view virtual returns (Erc20);

    function totalBorrows() external view virtual returns (uint256);

    function totalFuseFees() external view virtual returns (uint256);

    function repayBorrow(uint256) external virtual returns (uint256);

    function totalReserves() external view virtual returns (uint256);

    function exchangeRateCurrent() external virtual returns (uint256);

    function totalAdminFees() external view virtual returns (uint256);

    function fuseFeeMantissa() external view virtual returns (uint256);

    function adminFeeMantissa() external view virtual returns (uint256);

    function exchangeRateStored() external view virtual returns (uint256);

    function accrualBlockNumber() external view virtual returns (uint256);

    function redeemUnderlying(uint256) external virtual returns (uint256);

    function balanceOfUnderlying(address) external virtual returns (uint256);

    function reserveFactorMantissa() external view virtual returns (uint256);

    function borrowBalanceCurrent(address) external virtual returns (uint256);

    function interestRateModel() external view virtual returns (InterestRateModel);

    function initialExchangeRateMantissa() external view virtual returns (uint256);

    function repayBorrowBehalf(address, uint256) external virtual returns (uint256);
}