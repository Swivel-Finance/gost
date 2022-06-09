// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract Pendle {
    struct SwapExactTokensForTokensArgs {
        uint256 amount;
        uint256 minimumBought;
        address[] path;
        uint256 deadline;
    }

    struct RredeemAfterExpiryArgs {
        bytes32 forgeId;
        uint256 maturity;
    }

    uint256[] private swapExactTokensForTokensReturn;

    mapping(address => SwapExactTokensForTokensArgs) public swapExactTokensForTokensCalled;
    mapping(address => RredeemAfterExpiryArgs) public redeemAfterExpiryCalled;

    function swapExactTokensForTokensReturns(uint256[] memory r) external {
        swapExactTokensForTokensReturn = r;
    }

    function swapExactTokensForTokens(
        uint256 a,
        uint256 m,
        address[] calldata p,
        address t,
        uint256 d
    ) external returns (uint256[] memory) {
        swapExactTokensForTokensCalled[t] = SwapExactTokensForTokensArgs(a, m, p, d);
        return swapExactTokensForTokensReturn;
    }

    function redeemAfterExpiry(
        bytes32 i,
        address u,
        uint256 m
    ) external {
        redeemAfterExpiryCalled[u] = RredeemAfterExpiryArgs(i, m);
    }
}
