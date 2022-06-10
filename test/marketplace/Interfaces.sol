// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface IErc20 {
    function approve(address, uint256) external returns (bool);
}

interface IErc5095 {
    event Redeem(address indexed from, address indexed to, uint256 amount);
    event Mature(uint256 timestamp);

    function maturity() external returns (uint256);

    function underlying() external returns (address);

    function maturityRate() external returns (uint256);

    function convertToUnderlying(uint256 principalAmount) external returns (uint256);

    function convertToPrincipal(uint256 underlyingAmount) external returns (uint256);

    function maxRedeem(address owner) external view returns (uint256);

    function previewRedeem(uint256 principalAmount) external returns (uint256);

    function maxWithdraw(address owner) external returns (uint256);

    function previewWithdraw(uint256 principalAmount) external returns (uint256);

    function withdraw(
        uint256 underlyingAmount,
        address owner,
        address receiver
    ) external returns (uint256);

    function redeem(
        uint256 principalAmount,
        address owner,
        address receiver
    ) external returns (uint256);
}

interface IPool {
    function PT() external returns (IErc5095);

    function sellUnderlying(address to, uint128 min) external returns (uint128);

    function buyUnderlying(
        address to,
        uint128 underlyingOut,
        uint128 max
    ) external returns (uint128);

    function sellPT(address to, uint128 min) external returns (uint128);

    function buyPT(
        address to,
        uint128 PTOut,
        uint128 max
    ) external returns (uint128);

    function sellUnderlyingPreview(uint128 underlyingIn) external returns (uint128);

    function buyUnderlyingPreview(uint128 underlyingOut) external returns (uint128);

    function sellPTPreview(uint128 PTIn) external returns (uint128);

    function buyPTPreview(uint128 PTOut) external returns (uint128);
}
