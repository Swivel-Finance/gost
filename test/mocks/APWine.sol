// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract APWine {
    address getPTAddressReturn;

    function getPTAddressReturns(address a) external {
        getPTAddressReturn = a;
    }

    function getPTAddress() external view returns (address) {
        return getPTAddressReturn;
    }
}