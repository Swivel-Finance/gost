// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Swivel.sol'; // library of swivel specific constructs
import './Element.sol'; // library of Element specific constructs

interface Any {}

interface IERC20 {
    function approve(address, uint256) external returns (bool);

    function transfer(address, uint256) external returns (bool);

    function balanceOf(address) external returns (uint256);

    function transferFrom(
        address,
        address,
        uint256
    ) external returns (bool);
}

interface IERC20Metadata is IERC20 {
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
}

interface IYield {
    function base() external returns (IERC20);

    function maturity() external returns (uint32);

    function sellBase(address, uint128) external returns (uint128);

    function sellBasePreview(uint128) external returns (uint128);
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
}

interface ISense {
    function swapUnderlyingForPTs(
        address,
        uint256,
        uint256,
        uint256
    ) external returns (uint256);

    function maturity() external returns (uint256);
}

interface ISenseToken {
    function underlying() external returns (address);
}

interface ISenseAMM {
    function maturity() external returns (uint256);
}

interface IZcToken {
    function mint(address, uint256) external returns (bool);

    function balanceOf(address) external returns (uint256);
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
}

interface ITempus {
    function maturityTime() external view returns (uint256);

    function yieldBearingToken() external view returns (IERC20Metadata);

    function depositAndFix(
        Any,
        Any,
        uint256,
        bool,
        uint256,
        uint256
    ) external returns (uint256);
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
    function getUnderlyingOfIBTAddress() external view returns (address);
}

interface INotional {
    function getUnderlyingToken() external view returns (IERC20, int256);

    function getMaturity() external view returns (uint40);

    function deposit(uint256 a, address r) external returns (uint256);
}

interface IAave {
    function deposit(
        address,
        uint256,
        address,
        uint16
    ) external;
}
