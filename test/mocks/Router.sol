// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract Router {
    address public ownerCalled;
    uint256 public amountCalled;
    uint256 public yieldAmountCalled;
    address public recipientCalled;

    function redeemToBacking(address o, uint256 a, uint256 y, address r) public {
        ownerCalled = o;
        amountCalled = a;
        yieldAmountCalled = y;
        recipientCalled = r;
    }
}