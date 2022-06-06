// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract MarketPlace {
    address[9] private marketsReturn;

    mapping (address => uint256) public marketsCalled;

    function marketsReturns(address[9] calldata m) external {
        marketsReturn = m;
    }

    function principals(address u, uint256 m) external returns (address[9] memory) {
        marketsCalled[u] = m;
        return marketsReturn;
    }
}
