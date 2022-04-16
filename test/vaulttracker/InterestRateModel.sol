pragma solidity >=0.8.4;

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