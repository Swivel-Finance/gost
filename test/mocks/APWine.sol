// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract APWine {
    struct SwapExactAmountIn {
        uint256 id;
        uint256 tokenIn;
        uint256 tokenOut;
        uint256 amount;
        uint256 minimumAmount;
    }
    uint256 private swapExactAmountInReturn;

    mapping (address => uint256) public withdrawCalled;
    mapping (address => SwapExactAmountIn) public swapExactAmountInCalled;

    function swapExactAmountInReturns(uint256 s) external {
        swapExactAmountInReturn = s;
    }

    function swapExactAmountIn(uint256 i, uint256 tokenIn, uint256 a, uint256 tokenOut, uint256 m, address to) external returns (uint256) {
        swapExactAmountInCalled[to] = SwapExactAmountIn(i, tokenIn, tokenOut, a, m);
        return swapExactAmountInReturn;
    }

    function withdraw(address o, uint256 a) external {
        withdrawCalled[o] = a;
    }
}