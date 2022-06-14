// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Swivel.sol';
import './Element.sol';

import './Redeemer.sol';

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
    function mint(address, uint256) external returns (bool);

    function balanceOf(address) external returns (uint256);

    function burn(address, uint256) external;

    function maturity() external returns (uint256);
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

    function transfer(address, uint256) external returns (bool);

    function balanceOf(address) external returns (uint256);

    function transferFrom(
        address,
        address,
        uint256
    ) external returns (bool);
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

interface Any {}

interface IErc20Metadata is IErc20 {
    function name() external view returns (string memory);

    function symbol() external view returns (string memory);

    function decimals() external view returns (uint8);
}

interface IMarketPlace {
    function markets(
        address,
        uint256,
        uint8
    ) external returns (address);
}

interface ISwivel {
    function initiate(
        Swivel.Order[] calldata,
        uint256[] calldata,
        Swivel.Components[] calldata
    ) external returns (bool);

    function redeemZcToken(
        address u,
        uint256 m,
        uint256 a
    ) external returns (bool);
}

// we'll use ...Token for interfaces that are Erc20s
interface IYield {
    // TODO OG has `is ..Erc20` - is that necessary?
    function base() external returns (IErc20); // TODO can we use the wide Interface here?

    function maturity() external returns (uint32);

    function sellBase(address, uint128) external returns (uint128);

    function sellBasePreview(uint128) external returns (uint128);

    function redeem(address to, uint256 amount) external returns (uint256);
}

interface IElement {
    function swap(
        Element.SingleSwap memory,
        Element.FundManagement memory,
        uint256,
        uint256
    ) external returns (uint256);
}

interface IElementToken {
    function unlockTimestamp() external returns (uint256);

    function underlying() external returns (address);

    function withdrawPrincipal(uint256 amount, address destination) external;
}

interface ISense {
    function swapUnderlyingForPTs(
        address,
        uint256,
        uint256,
        uint256
    ) external returns (uint256);

    function redeem(
        address,
        uint256,
        uint256
    ) external;
}

interface ISenseToken {
    function underlying() external returns (address);
}

interface IPendleToken {
    function yieldToken() external returns (address);

    function expiry() external returns (uint256);
}

interface IPendle {
    function swapExactTokensForTokens(
        uint256,
        uint256,
        address[] calldata,
        address,
        uint256
    ) external returns (uint256[] memory amounts);

    function redeemAfterExpiry(
        bytes32,
        address,
        uint256
    ) external;
}

interface ITempus {
    function maturityTime() external view returns (uint256);

    function yieldBearingToken() external view returns (IErc20Metadata);

    function depositAndFix(
        Any,
        Any,
        uint256,
        bool,
        uint256,
        uint256
    ) external returns (uint256);

    function redeemToBacking(
        address,
        uint256,
        uint256,
        address
    ) external;
}

interface IAPWineRouter {
    function swapExactAmountIn(
        uint256,
        uint256,
        uint256,
        uint256,
        uint256,
        address
    ) external returns (uint256);
}

interface IAPWineToken {
    function getPTAddress() external view returns (address);
}

interface INotional {
    function getUnderlyingToken() external view returns (IErc20, int256);

    function getMaturity() external view returns (uint40);

    function deposit(uint256 a, address r) external returns (uint256);

    function maxRedeem(address) external returns (uint256);
}

interface IAave {
    function deposit(
        address,
        uint256,
        address,
        uint16
    ) external;
}

interface IAPWine {
    function withdraw(address, uint256) external;
}

interface IYieldToken {
    function redeem(
        address,
        address,
        uint256
    ) external;
}
