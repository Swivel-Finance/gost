// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface Any {}

contract Tempus {
    uint256 depositAndFixReturn;

    Any public tempusAMMCalled;
    Any public tempusPoolCalled;
    uint256 public amountCalled;
    bool public isBackingTokenCalled;
    uint256 public minimumReturnCalled;
    uint256 public deadlineCalled;

    function depositAndFixReturns(uint256 r) external {
        depositAndFixReturn = r;
    }

    function depositAndFix(Any x, Any p, uint256 a, bool bt, uint256 mr, uint256 d) external returns (uint256) {
        tempusAMMCalled = x;
        tempusPoolCalled = p;
        amountCalled = a;
        isBackingTokenCalled = bt;
        minimumReturnCalled = mr;
        deadlineCalled = d;

        return depositAndFixReturn;
    }
}