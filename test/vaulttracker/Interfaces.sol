// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.4;

interface CErc20 {
	function exchangeRateCurrent() external returns (uint256);
}
