// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IRedeemer {
    function authRedeem(address underlying, uint256 maturity, address from, address to, uint256 amount) external returns (uint256);
}