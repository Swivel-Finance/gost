// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract ElementToken {
    uint256 private unlockTimestampReturn;
    address private underlyingReturn;

    function unlockTimestampReturns(uint256 u) external {
        unlockTimestampReturn = u;
    }

    function underlyingReturns(address a) external {
        underlyingReturn = a;
    }

    function unlockTimestamp() external view returns (uint256) {
        return unlockTimestampReturn;
    }
    
    function underlying() external view returns (address) {
        return underlyingReturn;
    }
}