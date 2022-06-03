// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract MarketPlace {
    mapping (address => mapping (uint256 => address[9])) public markets;

    function setMarket(address u, uint256 m, uint8 p, address t) public {
        markets[u][m][p] = t;
    }
}
