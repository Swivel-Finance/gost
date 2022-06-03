// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract MarketPlace {
    mapping (address => mapping (uint256 => address[9])) public markets;
}
