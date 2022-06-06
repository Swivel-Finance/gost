// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract MarketPlace {
    address[9] private marketsReturn;

    mapping (address => uint256) public marketsCalled;

    function marketsReturns(address[9] calldata m) external {
        marketsReturn = m;
    }

    function principals(address, uint256) external view returns (address[9] memory) {
        return marketsReturn;
    }
}
