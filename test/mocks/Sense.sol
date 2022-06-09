// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract Sense {
    struct SwapUnderlyingForPTsArgs {
        uint256 maturity;
        uint256 amount;
        uint256 minimumBought;
    }

    struct RedeemArgs {
        uint256 maturity;
        uint256 amount;
    }

    uint256 private swapUnderlyingForPTsReturn;
    mapping(address => SwapUnderlyingForPTsArgs) public swapUnderlyingForPTsCalled;
    mapping(address => RedeemArgs) public redeemCalled;

    function swapUnderlyingForPTsReturns(uint256 s) external {
        swapUnderlyingForPTsReturn = s;
    }

    function swapUnderlyingForPTs(
        address sa,
        uint256 m,
        uint256 a,
        uint256 mb
    ) external returns (uint256) {
        swapUnderlyingForPTsCalled[sa] = SwapUnderlyingForPTsArgs(m, a, mb);
        return swapUnderlyingForPTsReturn;
    }

    function redeem(
        address a,
        uint256 m,
        uint256 amount
    ) external {
        redeemCalled[a] = RedeemArgs(m, amount);
    }
}
