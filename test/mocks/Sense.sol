// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract Sense {
    uint256 swapUnderlyingForPTsReturn;

    address public sensePoolCalled;
    uint256 public maturityCalled;
    uint256 public amountCalled;
    uint256 public minimumBoughtCalled;

    function swapUnderlyingForPTsReturns(uint256 s) external {
        swapUnderlyingForPTsReturn = s;
    }
    
    function swapUnderlyingForPTs(address sp, uint256 m, uint256 a, uint256 mb) external returns (uint256) {
        sensePoolCalled = sp;
        maturityCalled = m;
        amountCalled = a;
        minimumBoughtCalled = mb;
        
        return swapUnderlyingForPTsReturn;
    }
}