// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract SenseToken {
    address private underlyingReturn;

    function underlyingReturns(address a) external {
        underlyingReturn = a;
    }

    function underlying() external view returns (address) {
        return underlyingReturn;
    }
}