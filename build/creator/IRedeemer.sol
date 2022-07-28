// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IRedeemer {
    struct Market {
        address cTokenAddr;
        address zcToken;
        address vaultTracker;
        uint256 maturityRate;
    }

    function authRedeem(uint8, address, uint256, address, address, uint256) external returns (uint256);
    function markets(uint8, address, uint256) external view returns (Market memory market);
    function exchangeRate(uint8, address) external view returns (uint256);
}
