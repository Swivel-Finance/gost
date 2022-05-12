// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract ZcToken {
    struct TransferFromArgs {
        address to;
        uint256 amount;
    }

    bool private mintReturn;
    bool private transferFromReturn;
    uint256 private balanceOfReturn;

    mapping (address => TransferFromArgs) public transferFromCalled;
    mapping (address => uint256) public mintCalled;
    mapping (address => uint256) public burnCalled;
    uint256 public amountCalled;
    address public balanceOfCalled;

    function mintReturns(bool r) external {
        mintReturn = r;
    }

    function mint(address a, uint256 u) external returns (bool) {
        mintCalled[a] = u;
        return mintReturn;
    }

    function balanceOfReturns(uint256 b) public {
        balanceOfReturn = b;
    }

    function balanceOf(address b) public returns (uint256) {
        balanceOfCalled = b;
        return balanceOfReturn;
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