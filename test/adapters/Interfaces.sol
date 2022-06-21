// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

/// @notice tokenized representation of assets.
/// @dev This represents the 'port' exposed by the swivel protocol that various
/// deployed adapters implement in order to communicate with their given PROTOCOL
interface ICompounding {
  function PROTOCOL() external returns (uint8);
  function exchangeRate(address) external returns (uint256);
  function underlying(address) external returns (address);
}

interface ICompound {
	function deposit(address, uint256) external returns (uint256);
	function redeemUnderlying(address, uint256) external returns (uint256);
}

interface ICompoundToken {
	function exchangeRateCurrent() external returns (uint256);
  function underlying() external returns (address);
  function mint(uint256) external returns(uint256);
  function redeemUnderlying(uint256) external returns(uint256);
}
