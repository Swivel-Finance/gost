// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract Pendle {
    struct SwapExactTokensForTokensArgs {
        uint256 amount;
        uint256 minimumBought;
        address[] path;
        uint256 deadline;
    }

    uint[] private swapExactTokensForTokensReturn;

    mapping (address => SwapExactTokensForTokensArgs) public swapExactTokensForTokensCalled;

    function swapExactTokensForTokensReturns(uint[] memory r) external {
        swapExactTokensForTokensReturn = r;
    }
    
    function swapExactTokensForTokens(uint256 a, uint256 m, address[] calldata p, address t, uint256 d) external returns (uint[] memory) {
        swapExactTokensForTokensCalled[t] = SwapExactTokensForTokensArgs(a, m, p, d);
        return swapExactTokensForTokensReturn;
    }
}