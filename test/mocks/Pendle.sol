// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract Pendle {
    struct SwapExactTokensForTokens {
        uint256 id;
        uint256 tokenIn;
        address[] path;
        uint256 amount;
    }

    uint[] private swapExactTokensForTokensReturn;

    mapping (address => SwapExactTokensForTokens) public swapExactTokensForTokensCalled;

    function swapExactTokensForTokensReturns(uint[] memory r) external {
        swapExactTokensForTokensReturn = r;
    }
    
    function swapExactTokensForTokens(uint256 i, uint256 o, address[] calldata p, address t, uint256 d) external returns (uint[] memory) {
        swapExactTokensForTokensCalled[t] = SwapExactTokensForTokens(i, o, p, d);
        return swapExactTokensForTokensReturn;
    }
}