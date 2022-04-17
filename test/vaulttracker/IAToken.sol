// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.8.4;

interface IAToken {
    function POOL() external view returns (address);
    function UNDERLYING_ASSET_ADDRESS() external view returns (address);
}