// SPDX-License-Identifier: UNLICENSED

pragma solidity >= 0.8.4;

interface Erc20 {
  function decimals() external returns (uint8);
}

interface CErc20 {
	function exchangeRateCurrent() external returns (uint256);
  function underlying() external returns (address);
}
