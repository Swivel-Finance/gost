// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract ZcToken {
    mapping (address => uint256) public balances;
    bool private mintReturn;

    mapping (address => uint256) public mintCalled;
    address public ownerCalled;
    uint256 public amountCalled;

    function mintReturns(bool r) external {
        mintReturn = r;
    }

    function balanceOfReturns(address a, uint256 b) external {
        balances[a] = b;
    }

    function mint(address a, uint256 u) external returns (bool) {
        balances[a] = u;
        mintCalled[a] = u;
        return mintReturn;
    }

    function balanceOf(address a) external view returns (uint256) {
        return balances[a];
    }

    function burn(address o, uint256 a) external {
        ownerCalled = o;
        amountCalled = a;
    }
}