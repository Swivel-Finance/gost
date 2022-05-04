// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface IPendleYieldToken {
  function underlyingAsset() external view returns (address);
  function underlyingYieldToken() external view returns (address);
}

contract PendleData {
    bool isValidXYTReturn;
    bool isValidOTReturn;
    address xytTokensReturn;

    bytes32 isValidOTId;
    address isValidOTUnderlying;
    uint256 isValidOTMaturity;

    bytes32 isValidXYTId;
    address isValidXYTUnderlying;
    uint256 isValidXYTMaturity;

    bytes32 xytTokensCalledId;
    address xytTokensCalledUnderlying;
    uint256 xytTokensCalledMaturity;


    function isValidXYTReturns(bool x) external {
        isValidXYTReturn = x;
    }

    function isValidOTReturns(bool o) external {
        isValidOTReturn = o;
    }

    function xytTokensReturns(address a) external {
        xytTokensReturn = a;
    }

    function isValidXYT(bytes32 f, address u, uint256 m) external returns (bool) {
        isValidXYTId = f;
        isValidXYTUnderlying = u;
        isValidXYTMaturity = m;
        return isValidXYTReturn;
    }

    function isValidOT(bytes32 f, address u, uint256 m) external returns (bool) {
        isValidOTId = f;
        isValidOTUnderlying = u;
        isValidOTMaturity = m;
        return isValidOTReturn;
    }

    function xytTokens(bytes32 f, address u, uint256 m) external returns (IPendleYieldToken) {
        xytTokensCalledId = f;
        xytTokensCalledUnderlying = u;
        xytTokensCalledMaturity = m;
        return IPendleYieldToken(xytTokensReturn);
    }
}