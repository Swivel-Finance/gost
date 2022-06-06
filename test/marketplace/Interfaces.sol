// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface IErc20 {
  function decimals() external returns (uint8);
}

interface ICErc20 {
	function exchangeRateCurrent() external returns (uint256);
  function underlying() external returns (address);
}
