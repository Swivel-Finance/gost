// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract MarketPlace {
    address[8] private marketsReturn;

    mapping (address => uint256) public marketsCalled;

    function marketsReturns(address[8] calldata m) external {
        marketsReturn = m;
    }

    function markets(address u, uint256 m) external returns (address[8] memory) {
        marketsCalled[u] = m;
        return marketsReturn;
    }
}
