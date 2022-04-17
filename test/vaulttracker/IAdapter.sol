// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.8.4;

interface Adapter {
  function exchangeRateCurrent(address cToken) external view returns(uint256);
}