// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract SenseToken {
    struct TransferFromArgs {
        address to;
        uint256 amount;
    }

    address private underlyingReturn;
    bool private transferFromReturn;
    uint256 private balanceOfReturn;

    mapping(address => TransferFromArgs) public transferFromCalled;
    address public balanceOfCalled;

    function underlyingReturns(address a) external {
        underlyingReturn = a;
    }

    function underlying() external view returns (address) {
        return underlyingReturn;
    }

    function balanceOfReturns(uint256 b) public {
        balanceOfReturn = b;
    }

    function balanceOf(address b) public returns (uint256) {
        balanceOfCalled = b;
        return balanceOfReturn;
    }

    function transferFromReturns(bool b) public {
        transferFromReturn = b;
    }

    function transferFrom(
        address f,
        address t,
        uint256 a
    ) public returns (bool) {
        TransferFromArgs memory args;
        args.to = t;
        args.amount = a;
        transferFromCalled[f] = args;
        return transferFromReturn;
    }
}
