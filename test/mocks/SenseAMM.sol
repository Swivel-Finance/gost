// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract SenseAMM {
    uint256 private maturityReturn;

    function maturityReturns(uint256 a) public {
        maturityReturn = a;
    }

    function maturity() public view returns (uint256) {
        return maturityReturn;
    }
}
