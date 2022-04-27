// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Swivel.sol'; // library of swivel specific constructs
import './Element.sol'; // library of Element specific constructs

interface Erc20 {
	function approve(address, uint256) external returns (bool);
	function transfer(address, uint256) external returns (bool);
	function balanceOf(address) external returns (uint256);
	function transferFrom(address, address, uint256) external returns (bool);
}

// TODO move to ./MarketPlace.sol as a stand alone enum?
interface MarketPlace {
  enum Principals { Illuminate, Swivel, Yield, Element, Pendle, Tempus, Sense, Apwine }
}

interface Swivel {
  function initiate(Swivel.Order[], uint256[], Swivel.Components[]) external returns (bool);
  function redeemZcToken(address, uint256, uint256) external returns (bool);
}

// we'll use ...Token for interfaces that are Erc20s
interface YieldToken { // TODO OG has `is ..Erc20` - is that necessary?
  function base() external returns (Erc20); // TODO can we use the wide Interface here?
  function maturity() external returns (uint32);
  function sellBase(address, uint128) external returns (uint128);
  function sellBasePreview(uint128) external returns (uint128);
}

interface Element {
  // TODO man those structs names are verbose... can it be ELement.Swap, Element.Fund etc... ?
  function swap(Element.SingleSwap, Element.FundManagement, uint256, uint256) external returns (uint256);
}

interface ElementToken {
  function unlockTimestamp() external returns (uint256);
  function underlying() external returns (address);
}

interface Sense {
  function swapUnderlyingForPTs() external returns (uint256);
}
