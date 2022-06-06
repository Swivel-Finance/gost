// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Redeemer.sol';

interface IMarketPlace {
  function markets(address, uint256, uint8) external returns (address[9] calldata);
}

interface IErc20 {
	function approve(address, uint256) external returns (bool);
	function transfer(address, uint256) external returns (bool);
	function balanceOf(address) external returns (uint256);
	function transferFrom(address, address, uint256) external returns (bool);
}

interface IZcToken {
	function mint(address, uint256) external returns (bool);
	function balanceOf(address) external returns (uint256);
	function burn(address, uint256) external;
}

interface IAPWine {
	function withdraw(address, uint256) external;
}

interface ITempus {
	function redeemToBacking(address, uint256, uint256, address) external;
}

interface ISwivel {
	function redeemZcToken(address u, uint256 m, uint256 a) external returns (bool);
}

interface IYield {
	function redeem(address to, uint256 amount) external returns (uint256);
}

interface IPendle {
	function redeemAfterExpiry(bytes32 forgeId, address underlying, uint256 maturity) external;
}

interface ISense {
	function redeem(address adapter, uint256 maturity, uint256 amount) external;
}

interface IElementToken {
	function unlockTimestamp() external returns (uint256);
	function underlying() external returns (address);
	function withdrawPrincipal(uint256 amount, address destination) external;
}

interface IYieldToken {
	function redeem(address from, address to, uint256 amount) external;
}

interface INotional { 
  function maxRedeem(address) external returns (uint256);
}