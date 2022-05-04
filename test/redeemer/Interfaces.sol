// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Redeemer.sol';

interface IErc20 {
	function approve(address, uint256) external returns (bool);
	function transfer(address, uint256) external returns (bool);
	function balanceOf(address) external returns (uint256);
	function transferFrom(address, address, uint256) external returns (bool);
}

interface ISwivel {
    function redeemZcToken(address u, uint256 m, uint256 a) external returns (bool);
}

interface IYield {
    function redeem(address to, uint256 amount) external returns (uint256);
}

