// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.4;

interface Erc20 {
	function approve(address, uint256) external returns (bool);
	function transfer(address, uint256) external returns (bool);
	function balanceOf(address) external view returns (uint256);
	function transferFrom(address, address, uint256) external returns (bool);
}

interface CErc20 is Erc20 {
	function mint(uint256) external returns (uint256);
	function redeem(uint256) external returns (uint256);
	function redeemUnderlying(uint256) external returns (uint256);
	function exchangeRateCurrent() external returns (uint256);
	function underlying() external view returns(address);
}
