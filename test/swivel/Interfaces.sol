// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.13;

import './Hash.sol';
import './Sig.sol';

// the behavioral Swivel Interface, Implemented by Swivel.sol
interface ISwivel {
  function initiate(Hash.Order[] calldata, uint256[] calldata, Sig.Components[] calldata) external returns (bool);
  function exit(Hash.Order[] calldata, uint256[] calldata, Sig.Components[] calldata) external returns (bool);
  function cancel(Hash.Order calldata, Sig.Components calldata) external returns (bool);
  function setAdmin(address) external returns (bool);
  function scheduleWithdrawal(address) external returns (bool);
  function scheduleFeeChange() external returns (bool);
  function blockWithdrawal(address) external returns (bool);
  function blockFeeChange() external returns (bool);
  function withdraw(address) external returns (bool);
  function changeFee(uint16[4] calldata) external returns (bool);
  function approveUnderlying(address[] calldata, address[] calldata) external returns (bool);
  function splitUnderlying(uint8, address, uint256, uint256) external returns (bool);
  function combineTokens(uint8, address, uint256, uint256) external returns (bool);
  function authRedeem(uint8, address, address, address, uint256) external returns (bool);
  function redeemZcToken(uint8, address, uint256, uint256) external returns (bool);
  function redeemVaultInterest(uint8, address, uint256) external returns (bool);
  function redeemSwivelVaultInterest(uint8, address, uint256) external returns (bool);
}

// methods requried on other contracts which are expected to, at least, implement the following:
interface IErc20 {
	function approve(address, uint256) external returns (bool);
	function transfer(address, uint256) external returns (bool);
	function balanceOf(address) external returns (uint256);
	function transferFrom(address, address, uint256) external returns (bool);
}

interface IErc4626 {
  function deposit(uint256, address) external returns (uint256);
  function withdraw(uint256, address, address) external returns (uint256);
}

interface ICompound {
  function mint(uint256) external returns (uint256);
  function redeemUnderlying(uint256) external returns (uint256);
}

interface IYearn {
  function deposit(uint256) external returns (uint256);
  function withdraw(uint256) external returns (uint256);
}

interface IAave {
  function deposit(address, uint256, address, uint16) external; // void
  function withdraw(address, uint256, address) external returns (uint256);
}

interface IEuler {
  function deposit(uint256, uint256) external; // void
  function withdraw(uint256, uint256) external; // void
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
  function cTokenAddress(uint8, address, uint256) external returns (address);
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
