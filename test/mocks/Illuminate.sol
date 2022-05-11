// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract Illuminate {
    address[8] private marketsReturn;
    bool transferFromReturn;

    mapping (address => uint256) public marketsCalled;

    function marketsReturns(address[8] calldata m) external {
        marketsReturn = m;
    }

    function markets(address u, uint256 m) external returns (address[8] memory) {
        marketsCalled[u] = m;
        return marketsReturn;
    }

    function tranferFromReturns(bool t) external {
        transferFromReturn = t;
    }

    function transferFrom() external view returns (bool) {
        return transferFromReturn;
    }

}
