// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract Sushi {
    uint[] private swapExactTokensForTokensReturn;

    uint256 public inCalled;
    uint256 public outMinimumCalled;
    address[] public pathCalled;
    address public toCalled;
    uint256 public deadlineCalled;

    function swapExactTokensForTokensReturns(uint[] memory r) external {
        swapExactTokensForTokensReturn = r;
    }
    
    function swapExactTokensForTokens(uint256 i, uint256 o, address[] calldata p, address t, uint256 d) external returns (uint[] memory) {
        inCalled = i;
        outMinimumCalled = o;
        pathCalled = p;
        toCalled = t;
        deadlineCalled = d;

        return swapExactTokensForTokensReturn;
    }
}