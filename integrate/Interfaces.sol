// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface IPErc20 {
    function balanceOf(address) external view returns (uint256);

    function transfer(address, uint256) external returns (bool);

    function allowance(address, address) external view returns (uint256);

    function approve(address, uint256) external returns (bool);

    function transferFrom(
        address,
        address,
        uint256
    ) external returns (bool);

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

interface IZcToken is IPErc20 {
    function burn(address, uint256) external returns (bool);

    function mint(address, uint256) external returns (bool);
}

interface IErc2612 {
    function permit(
        address,
        address,
        uint256,
        uint256,
        uint8,
        bytes32,
        bytes32
    ) external;

    function nonces(address) external view returns (uint256);
}

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

    function sellPrincipalToken(address, uint128) external returns (uint128);

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
