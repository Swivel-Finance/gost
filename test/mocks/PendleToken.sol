// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract PendleToken {
    address private yieldTokenReturn;
    uint256 private expiryReturn;

    function yieldTokenReturns(address a) external {
        yieldTokenReturn = a;
    }

    function expiryReturns(uint256 m) external {
        expiryReturn = m;
    }

    function yieldToken() external view returns (address) {
        return yieldTokenReturn;
    }

    function expiry() external view returns (uint256) {
        return expiryReturn;
    }
}