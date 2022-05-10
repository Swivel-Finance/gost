// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Swivel.sol'; // library of swivel specific constructs
import './Element.sol'; // library of Element specific constructs

interface Any {}

interface IErc20 {
	function approve(address, uint256) external returns (bool);
	function transfer(address, uint256) external returns (bool);
	function balanceOf(address) external returns (uint256);
	function transferFrom(address, address, uint256) external returns (bool);
}

interface IErc20Metadata is IErc20 {
    function name() external view returns (string memory);
    function symbol() external view returns (string memory);
    function decimals() external view returns (uint8);
}

interface IIlluminate {
  function markets(address, uint256) external returns (address[8] calldata);
}

interface ISwivel {
  function initiate(Swivel.Order[] calldata, uint256[] calldata, Swivel.Components[] calldata) external returns (bool);
}

// we'll use ...Token for interfaces that are Erc20s
interface IYield { // TODO OG has `is ..Erc20` - is that necessary?
  function base() external returns (IErc20); // TODO can we use the wide Interface here?
  function maturity() external returns (uint32);
  function sellBase(address, uint128) external returns (uint128);
  function sellBasePreview(uint128) external returns (uint128);
}

interface IElement {
  function swap(Element.SingleSwap memory, Element.FundManagement memory, uint256, uint256) external returns (uint256);
}

interface IElementToken {
  function unlockTimestamp() external returns (uint256);
  function underlying() external returns (address);
}

interface ISense {
  function swapUnderlyingForPTs(address, uint256, uint256, uint256) external returns (uint256);
}

interface ISenseToken {
  function underlying() external returns (address);
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
  function swapExactTokensForTokens(uint256, uint256, address[] calldata, address, uint256) external returns (uint[] memory amounts);
}

interface ITempus {
  function maturityTime() external view returns (uint256);
  function yieldBearingToken() external view returns (IErc20Metadata);
  function depositAndFix(Any, Any, uint256, bool, uint256, uint256) external returns (uint256);
}

interface IAPWineRouter { 
  function swapExactAmountIn(uint256, uint256, uint256, uint256, uint256, address) external returns (uint256);
}

interface IAPWineToken {
  function getPTAddress() external view returns (address);
}
