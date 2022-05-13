// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract PendleToken {
    struct TransferFromArgs {
        address to;
        uint256 amount;
    }

    address private yieldTokenReturn;
    uint256 private expiryReturn;
    bool private transferFromReturn;
    uint256 private balanceOfReturn;

    mapping (address => TransferFromArgs) public transferFromCalled;
    address public balanceOfCalled;

    function yieldTokenReturns(address a) external {
        yieldTokenReturn = a;
    }

    function expiryReturns(uint256 m) external {
        expiryReturn = m;
    }

    function yieldToken() external view returns (address) {
        return yieldTokenReturn;
    }

    function expiry() external view returns (uint256) {
        return expiryReturn;
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

    function transferFrom(address f, address t, uint256 a) public returns (bool) {
        TransferFromArgs memory args;
        args.to = t;
        args.amount = a;
        transferFromCalled[f] = args;
        return transferFromReturn;
    }

}