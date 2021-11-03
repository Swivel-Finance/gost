// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.4;

interface Erc20 {
	function approve(address, uint256) external returns (bool);
	function transfer(address, uint256) external returns (bool);
	function balanceOf(address) external view returns (uint256);
	function transferFrom(address, address, uint256) external returns (bool);
}

interface CErc20 is Erc20 {
	function mint(uint256) external returns (uint256);
	function redeem(uint256) external returns (uint256);
	function redeemUnderlying(uint256) external returns (uint256);
	function exchangeRateCurrent() external returns (uint256);
	function underlying() external view returns(address);
}

interface MarketPlace {
  // adds notional and mints zctokens
  function mintZcTokenAddingNotional(address, uint256, address, uint256) external returns (bool);
  // removes notional and burns zctokens
  function burnZcTokenRemovingNotional(address, uint256, address, uint256) external returns (bool);
  // returns the amount of underlying interest to send
  function redeemVaultInterest(address u, uint256 m, address a) external returns (uint256);
  // returns the amount of underlying principal to send
  function redeemZcToken(address u, uint256 m, address t, uint256 a) external returns (uint256);
  // returns the cToken address for a given market
  function cTokenAddress(address, uint256) external returns (address);
  // EVFZE FF EZFVE call this which would then burn zctoken and remove notional
  function custodialExit(address, uint256, address, address, uint256) external returns (bool);
  // IVFZI && IZFVI call this which would then mint zctoken and add notional
  function custodialInitiate(address, uint256, address, address, uint256) external returns (bool);
  // IZFZE && EZFZI call this, tranferring zctoken from one party to another
  function p2pZcTokenExchange(address, uint256, address, address, uint256) external returns (bool);
  // IVFVE && EVFVI call this, removing notional from one party and adding to the other
  function p2pVaultExchange(address, uint256, address, address, uint256) external returns (bool);
  // IVFZI && IVFVE call this which then transfers notional from msg.sender (taker) to swivel
  function transferVaultNotionalFee(address, uint256, address, uint256) external returns (bool);
}