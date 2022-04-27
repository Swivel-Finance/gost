// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.8.4;

/// @notice Tokenised representation of assets
interface IAdapter {
    function exchangeRateCurrent(address cToken) external view returns(uint256);

    function protocol() external view returns(uint8);
}