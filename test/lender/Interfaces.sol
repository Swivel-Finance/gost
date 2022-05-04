// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Swivel.sol'; // library of swivel specific constructs
import './Element.sol'; // library of Element specific constructs

interface IErc20 {
	function approve(address, uint256) external returns (bool);
	function transfer(address, uint256) external returns (bool);
	function balanceOf(address) external returns (uint256);
	function transferFrom(address, address, uint256) external returns (bool);
}

// TODO move to ./MarketPlace.sol as a stand alone enum?
interface IMarketPlace {
  function markets(address, uint256) external returns (address[8] calldata);
}

interface ISwivel {
  function initiate(Swivel.Order[] calldata, uint256[] calldata, Swivel.Components[] calldata) external returns (bool);
}

// we'll use ...Token for interfaces that are Erc20s
interface IYieldToken { // TODO OG has `is ..Erc20` - is that necessary?
  function base() external returns (IErc20); // TODO can we use the wide Interface here?
  function maturity() external returns (uint32);
  function sellBase(address, uint128) external returns (uint128);
  function sellBasePreview(uint128) external returns (uint128);
}

interface IPendleRouter {
  function data() external view returns (IPendleData);
  function tokenizeYield(bytes32 f, address u, uint256 m, uint256 a, address t) external returns (address, address, uint256);
}

interface IPendleData {
  function isValidXYT(bytes32 f, address u, uint256 m) external view returns (bool);
  function isValidOT(bytes32 f, address u, uint256 m) external view returns (bool);
  function xytTokens(bytes32 f, address u, uint256 m) external view returns (IPendleYieldToken);
}

interface IPendleYieldToken {
  function underlyingAsset() external view returns (address);
  function underlyingYieldToken() external view returns (address);
}

interface IElement {
  function swap(Element.SingleSwap memory, Element.FundManagement memory, uint256, uint256) external returns (uint256);
}

interface IElementToken {
  function unlockTimestamp() external returns (uint256);
  function underlying() external returns (address);
}

interface ISense {
  function swapUnderlyingForPTs() external returns (uint256);
}

interface IZcToken {
  function mint(address, uint256) external returns (bool);
}

interface Any {}