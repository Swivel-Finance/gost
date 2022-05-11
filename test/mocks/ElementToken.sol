// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract ElementToken {
    // a struct to hold the arguments passed to transferFrom
    struct TransferFromArgs {
        address to;
        uint256 amount;
    }

    uint256 private unlockTimestampReturn;
    address private underlyingReturn;
    
    mapping (address => TransferFromArgs) public transferFromCalled;

    function unlockTimestampReturns(uint256 u) external {
        unlockTimestampReturn = u;
    }

    function underlyingReturns(address a) external {
        underlyingReturn = a;
    }

    function unlockTimestamp() external view returns (uint256) {
        return unlockTimestampReturn;
    }

    function underlying() external view returns (address) {
        return underlyingReturn;
    }

    function balanceOf(address _owner) external view returns (uint256) {
        return 0;
    }

    function transferFrom(address f, address t, uint256 a) public returns (bool) {
        TransferFromArgs memory args;
        args.to = t;
        args.amount = a;
        transferFromCalled[f] = args;
        return transferFromReturn;
    }

    function transferFromReturns(bool b) public {
        transferFromReturn = b;
    }

    function redeem(address o, uint256 a) external {
        tokenCalled = o;
        amountCalled = a;
    }

}