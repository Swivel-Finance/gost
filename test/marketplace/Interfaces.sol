// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface IErc20 {
    function approve(address, uint256) external returns (bool);
}

interface IErc5095 {
    event Redeem(address, address, uint256);
    event Mature(uint256);

    function maturity() external returns (uint256);

    function underlying() external returns (address);

    function maturityRate() external returns (uint256);

    function convertToUnderlying(uint256) external returns (uint256);

    function convertToPrincipal(uint256) external returns (uint256);

    function maxRedeem(address) external returns (uint256);

    function previewRedeem(uint256) external returns (uint256);

    function maxWithdraw(address) external returns (uint256);

    function previewWithdraw(uint256) external returns (uint256);

    function withdraw(
        uint256,
        address,
        address
    ) external returns (uint256);

    function redeem(
        uint256,
        address,
        address
    ) external returns (uint256);
}

interface IPool {
    function principalToken() external returns (IErc5095);

    function underlying() external returns (address);

    function sellUnderlying(address, uint128) external returns (uint128);

    function buyUnderlying(
        address,
        uint128,
        uint128
    ) external returns (uint128);

    function sellPrincipalToken(address to, uint128 min) external returns (uint128);

    function buyPrincipalToken(
        address,
        uint128,
        uint128
    ) external returns (uint128);

    function sellUnderlyingPreview(uint128) external returns (uint128);

    function buyUnderlyingPreview(uint128) external returns (uint128);

    function sellPrincipalTokenPreview(uint128) external returns (uint128);

    function buyPrincipalTokenPreview(uint128) external returns (uint128);
}
