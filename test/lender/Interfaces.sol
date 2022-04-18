// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './IPErc20.sol';
import './Swivel.sol'; // library of swivel specific constructs
import './Element.sol'; // library of Element specific constructs

interface ISwivel {
  function initiate(Swivel.Order[], uint256[], Swivel.Components[]) external returns (bool);
  function redeemZcToken(address, uint256, uint256) external returns (bool);
}

// we'll use ...Token for interfaces that are Erc20s
interface IYieldToken is IPErc20 {
  function base() external returns (IPErc20);
  function maturity() external returns (uint32);
  function sellBase(address, uint128) external returns (uint128);
  function sellBasePreview(uint128) external returns (uint128);
}

interface IElement {
  function swap() external returns (uint256);
}
