// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

/// @notice this mocks the actual marketplace contract by providing access to
/// the public markets mapping.
contract MarketPlace 
{
  mapping (address => mapping (uint256 => address[9])) public markets;

  /// @dev this enum must match the Principals enum in the MarketPlace's contract
  enum Principals {
    Illuminate,
    Swivel,
    Yield,
    Element,
    Pendle,
    Tempus,
    Sense,
    Apwine,
    Notional
  }
}
