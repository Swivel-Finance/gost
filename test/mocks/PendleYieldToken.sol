// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract IPendleYieldToken {
    address underlyingAssetReturn;
    address underlyingYieldTokenReturn;

    function underlyingAssetReturns(address a) external {
        underlyingAssetReturn = a;
    }

    function underlyingYieldTokenReturns(address a) external {
        underlyingYieldTokenReturn = a;
    }

    function underlyingAsset() external view returns (address) {
        return underlyingAssetReturn;
    }

    function underlyingYieldToken() external view returns (address) {
        return underlyingYieldTokenReturn;
    }
}