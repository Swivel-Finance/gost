// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract ElementToken {
    struct TransferFromArgs {
        address to;
        uint256 amount;
    }

    uint256 private unlockTimestampReturn;
    uint256 private balanceOfReturn;
    address private underlyingReturn;
    bool private transferFromReturn;

    address public balanceOfCalled;
    mapping(address => TransferFromArgs) public transferFromCalled;
    mapping(address => uint256) public withdrawPrincipalCalled;

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

    function withdrawPrincipal(uint256 a, address d) external {
        withdrawPrincipalCalled[d] = a;
    }
}
