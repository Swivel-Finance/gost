// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface IErc20 {
	function approve(address, uint256) external returns (bool);
	function transfer(address, uint256) external returns (bool);
	function balanceOf(address) external returns (uint256);
	function transferFrom(address, address, uint256) external returns (bool);
}

interface IMarketPlace {
  // adds notional and mints zctokens
  function mintZcTokenAddingNotional(uint8, address, uint256, address, uint256) external returns (bool);
  // removes notional and burns zctokens
  function burnZcTokenRemovingNotional(uint8, address, uint256, address, uint256) external returns (bool);
  // returns the amount of underlying principal to send
  function redeemZcToken(uint8, address, uint256, address, uint256) external returns (uint256);
  // returns the amount of underlying interest to send
  function redeemVaultInterest(uint8, address, uint256, address) external returns (uint256);
  // returns the cToken address for a given market
  function cTokenAndAdapterAddress(uint8, address, uint256) external returns (address, address);
  // EVFZE FF EZFVE call this which would then burn zctoken and remove notional
  function custodialExit(uint8, address, uint256, address, address, uint256) external returns (bool);
  // IVFZI && IZFVI call this which would then mint zctoken and add notional
  function custodialInitiate(uint8, address, uint256, address, address, uint256) external returns (bool);
  // IZFZE && EZFZI call this, tranferring zctoken from one party to another
  function p2pZcTokenExchange(uint8, address, uint256, address, address, uint256) external returns (bool);
  // IVFVE && EVFVI call this, removing notional from one party and adding to the other
  function p2pVaultExchange(uint8, address, uint256, address, address, uint256) external returns (bool);
  // IVFZI && IVFVE call this which then transfers notional from msg.sender (taker) to swivel
  function transferVaultNotionalFee(uint8, address, uint256, address, uint256) external returns (bool);
}
