// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract Pendle {
    address private underlyingReturn;
    uint256 private maturityReturn;

    function underlyingReturns(address a) external {
        underlyingReturn = a;
    }

    function maturityReturns(uint256 m) external {
        maturityReturn = m;
    }

    function underlying() external view returns (address) {
        return underlyingReturn;
    }

    function maturity() external view returns (uint256) {
        return maturityReturn;
    }
}