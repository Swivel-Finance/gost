// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Redeemer.sol';

interface IErc20 {
	function approve(address, uint256) external returns (bool);
	function transfer(address, uint256) external returns (bool);
	function balanceOf(address) external returns (uint256);
	function transferFrom(address, address, uint256) external returns (bool);
}

interface IRedeemable {
	function approve(address, uint256) external returns (bool);
	function transfer(address, uint256) external returns (bool);
	function balanceOf(address) external returns (uint256);
	function transferFrom(address, address, uint256) external returns (bool);
	function redeem(address, address, uint256) external;
}

interface IIlluminate {
  function markets(address, uint256) external returns (address[8] calldata);
}

interface IPendleRedeemable {
	function redeem(address, address, uint256, bytes32) external;
}

interface IPendleRouter {
	function redeemAfterExpiry(bytes32, address, uint256) external;
}

interface IRouter {
	function redeemToBacking(address, uint256, uint256, address) external;
}

interface ISwivel {
    function redeemZcToken(address u, uint256 m, uint256 a) external returns (bool);
}

interface IYield {
    function redeem(address to, uint256 amount) external returns (uint256);
}

interface ISenseDivider {
	function redeem(address, uint256, uint256) external;
}
