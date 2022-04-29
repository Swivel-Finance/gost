// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract ZcToken {
    bool private mintReturn;
    mapping (address => uint256) public mintCalled;

    function mintReturns(bool r) external {
        mintReturn = r;
    }

    function mint(address a, uint256 u) external returns (bool) {
        mintCalled[a] = u;
        return mintReturn;
    }
}