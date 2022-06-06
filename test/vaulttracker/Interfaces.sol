// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface ICErc20 {
	function exchangeRateCurrent() external returns (uint256);
}
