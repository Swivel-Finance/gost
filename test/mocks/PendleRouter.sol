// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface IPendleYieldToken {
    function underlyingAsset() external view returns (address);
    function underlyingYieldToken() external view returns (address);
}

interface IPendleData {
    function isValidXYT(bytes32 f, address u, uint256 m) external view returns (bool);
    function isValidOT(bytes32 f, address u, uint256 m) external view returns (bool);
    function xytTokens(bytes32 f, address u, uint256 m) external view returns (IPendleYieldToken);
}

contract PendleRouter {
    uint256 tokenizeYieldReturn;
    IPendleData dataReturn;

    bytes32 tokenizeYieldIdCalled;
    address tokenizeYieldUnderlyingCalled;
    uint256 tokenizeYieldMaturityCalled;
    uint256 tokenizeYieldAmountCalled;
    address tokenizeYieldTokenCalled;

    function tokenizeYieldReturns (uint256 x) external {
        tokenizeYieldReturn = x;
    }

    function dataReturns (IPendleData x) external {
        dataReturn = x;
    }

    function data() external view returns (IPendleData) {
        return dataReturn;
    }
    function tokenizeYield(bytes32 f, address u, uint256 m, uint256 a, address t) external returns (address, address, uint256) {
        tokenizeYieldIdCalled = f;
        tokenizeYieldUnderlyingCalled = u;
        tokenizeYieldMaturityCalled = m;
        tokenizeYieldAmountCalled = a;
        tokenizeYieldTokenCalled = t;
        return (address(0), address(0), tokenizeYieldReturn);
    }
}