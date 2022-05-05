// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract ZcToken {
    mapping (address => uint256) public balances;
    bool private mintReturn;

    function mintReturns(bool r) external {
        mintReturn = r;
    }

    function balanceOfReturns(address a, uint256 b) external {
        balances[a] = b;
    }

    function mint(address a, uint256 u) external returns (bool) {
        balances[a] = u;
        return mintReturn;
    }

    function balanceOf(address a) external view returns (uint256) {
        return balances[a];
    }
}