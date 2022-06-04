// SPDX-License-Identifier: UNLICENSED

pragma solidity >= 0.8.13;

interface IErc20 {
	function approve(address, uint256) external returns (bool);
}

interface IRouter {
	createMarket(address u, uint256 m, address i) external returns (address);
}
