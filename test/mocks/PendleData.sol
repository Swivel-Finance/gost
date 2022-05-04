// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface IPendleYieldToken {
  function underlyingAsset() external view returns (address);
  function underlyingYieldToken() external view returns (address);
}

contract PendleData {
    bool private isValidXYTReturn;
    bool private isValidOTReturn;
    address private xytTokensReturn;

    bytes32 public isValidOTId;
    address public isValidOTUnderlying;
    uint256 public isValidOTMaturity;

    bytes32 public isValidXYTId;
    address public isValidXYTUnderlying;
    uint256 public isValidXYTMaturity;

    bytes32 public xytTokensCalledId;
    address public xytTokensCalledUnderlying;
    uint256 public xytTokensCalledMaturity;


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