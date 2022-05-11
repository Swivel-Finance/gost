// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract Pendle {
    struct SwapExactTokensForTokensArgs {
        uint256 tokenIn;
        uint256 tokenOut;
        address[] path;
        uint256 deadline;
    }

    uint[] private swapExactTokensForTokensReturn;

    mapping (address => SwapExactTokensForTokensArgs) public swapExactTokensForTokensCalled;

    function swapExactTokensForTokensReturns(uint[] memory r) external {
        swapExactTokensForTokensReturn = r;
    }
    
    function swapExactTokensForTokens(uint256 i, uint256 o, address[] calldata p, address t, uint256 d) external returns (uint[] memory) {
        swapExactTokensForTokensCalled[t] = SwapExactTokensForTokensArgs(i, o, p, d);
        return swapExactTokensForTokensReturn;
    }
}