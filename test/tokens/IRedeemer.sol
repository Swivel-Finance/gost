// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IRedeemer {
    function authRedeem(uint8, address, uint256, address, address, uint256) external returns (uint256);
}
