// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract ZcToken {
    struct TransferFromArgs {
        address to;
        uint256 amount;
    }

    mapping (address => uint256) public balances;
    bool private mintReturn;
    bool private transferFromReturn;

    mapping (address => TransferFromArgs) public transferFromCalled;
    mapping (address => uint256) public mintCalled;
    mapping (address => uint256) public burnCalled;
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
        burnCalled[o] = a;
    }

    function transferFromReturns(bool b) public {
        transferFromReturn = b;
    }

    function transferFrom(address f, address t, uint256 a) public returns (bool) {
        TransferFromArgs memory args;
        args.to = t;
        args.amount = a;
        transferFromCalled[f] = args;
        return transferFromReturn;
    }
}