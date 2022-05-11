// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract Sense {

    struct SwapUnderlyingForPTs {
        uint256 maturity;
        uint256 amount;
        uint256 minimumBought;
    }

    uint256 private swapUnderlyingForPTsReturn;
    mapping(address => SwapUnderlyingForPTs) public swapUnderlyingForPTsCalled;

    function swapUnderlyingForPTsReturns(uint256 s) external {
        swapUnderlyingForPTsReturn = s;
    }
    
    function swapUnderlyingForPTs(address sa, uint256 m, uint256 a, uint256 mb) external returns (uint256) {
        swapUnderlyingForPTsCalled[sa] = SwapUnderlyingForPTs(m, a, mb);
        return swapUnderlyingForPTsReturn;
    }
}