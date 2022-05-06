// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract APWine {
    uint256 private swapExactAmountInReturn;

    uint256 public idCalled;
    uint256 public tokenInCalled;
    uint256 public amountCalled;
    uint256 public tokenOutCalled;
    uint256 public minimumAmountCalled;
    address public toCalled;

    function swapExactAmountInReturns(uint256 s) external {
        swapExactAmountInReturn = s;
    }

    function swapExactAmountIn(uint256 id, uint256 tokenIn, uint256 amount, uint256 tokenOut, uint256 minimumAmount, address to) external returns (uint256) {
        idCalled = id;
        tokenInCalled = tokenIn;
        amountCalled = amount;
        tokenOutCalled = tokenOut;
        minimumAmountCalled = minimumAmount;
        toCalled = to;
        
        return swapExactAmountInReturn;
    }
}