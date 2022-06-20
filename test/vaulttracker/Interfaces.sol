// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

/// @notice tokenized representation of assets.
/// @dev This represents the 'port' exposed by the swivel protocol that various
/// deployed adapters implement in order to communicate with their given PROTOCOL
interface ICompounding {
  function exchangeRate(address) external returns (uint256);
}
