// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract Sense {
    uint256 swapUnderlyingForPTsReturn;

    address public senseTokenCalled;
    uint256 public maturityCalled;
    uint256 public amountCalled;
    uint256 public minimumBoughtCalled;

    function swapUnderlyingForPTsReturns(uint256 s) external {
        swapUnderlyingForPTsReturn = s;
    }
    
    function swapUnderlyingForPTs(address sa, uint256 m, uint256 a, uint256 mb) external returns (uint256) {
        senseTokenCalled = sa;
        maturityCalled = m;
        amountCalled = a;
        minimumBoughtCalled = mb;
        
        return swapUnderlyingForPTsReturn;
    }
}