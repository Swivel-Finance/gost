// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.0;

abstract contract Erc20 {
	function approve(address, uint256) virtual external returns (bool);
	function transfer(address, uint256) virtual external returns (bool);
	function balanceOf(address) virtual external returns (uint256);
	function transferFrom(address, address, uint256) virtual public returns (bool);
}

abstract contract CErc20 is Erc20 {
	function mint(uint256) virtual external returns (uint256);
	function redeem(uint256) virtual external returns (uint256);
	function redeemUnderlying(uint256) virtual external returns (uint256);
	function exchangeRateCurrent() virtual external returns (uint256);
}

abstract contract MarketPlace {
  function cTokenAddress(address, uint256) virtual external returns (address);
  // EVFZE FF EZFVE call this which would then burn zctoken and remove notional
  function custodialExit(address, uint256, address, address, uint256) virtual external returns (bool);
  // IVFZI && IZFVI call this which would then mint zctoken and add notional
  function custodialInitiate(address, uint256, address, address, uint256) virtual external returns (bool);
  function transferFromZcToken(address, uint256, address, address, uint256) virtual external returns (bool);
  function transferFromNotional(address, uint256, address, address, uint256) virtual external returns (bool);
}
