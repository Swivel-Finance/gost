// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.13;

// the behavioral MarketPlace Interface, Implemented by MarketPlace.sol
interface IMarketPlace {
  function setSwivel(address) external returns (bool);
  function setAdmin(address) external returns (bool);
  function createMarket(uint8, uint256, address, string memory, string memory) external returns (bool);
  function matureMarket(uint8, address, uint256) external returns (bool);
  function mintZcTokenAddingNotional(uint8, address, uint256, address, uint256) external returns (bool);
  function burnZcTokenRemovingNotional(uint8, address, uint256, address, uint256) external returns (bool);
  function authRedeem(uint8, address, uint256, address, address, uint256) external returns(uint256);
  function redeemZcToken(uint8, address, uint256, address, uint256) external returns (uint256);
  function redeemVaultInterest(uint8, address, uint256, address) external returns (uint256);
  function cTokenAddress(uint8, address, uint256) external returns (address);
  function exchangeRate(uint8, address) external returns (uint256);
  function rates(uint8, address, uint256) external returns (uint256, uint256);
  function custodialExit(uint8, address, uint256, address, address, uint256) external returns (bool);
  function custodialInitiate(uint8, address, uint256, address, address, uint256) external returns (bool);
  function p2pZcTokenExchange(uint8, address, uint256, address, address, uint256) external returns (bool);
  function p2pVaultExchange(uint8, address, uint256, address, address, uint256) external returns (bool);
  function transferVaultNotional(uint8, address, uint256, address, uint256) external returns (bool);
  function transferVaultNotionalFee(uint8, address, uint256, address, uint256) external returns (bool);
}

// methods requried on other contracts which are expected to, at least, implement the following:
interface IErc20 {
  function decimals() external returns (uint8);
}

interface ICreator {
  function create(
    uint8,
    address,
    uint256,
    address,
    address,
    string calldata,
    string calldata,
    uint8
  ) external returns (address, address);
}

interface IZcToken {
  function mint(address, uint256) external returns (bool);
  function burn(address, uint256) external returns (bool);
}

interface IVaultTracker {
  function addNotional(address, uint256) external returns (bool);
  function removeNotional(address, uint256) external returns (bool);
  function redeemInterest(address) external returns (uint256);
  function matureVault(uint256) external returns (bool);
  function transferNotionalFrom(address, address, uint256) external returns (bool);
  function transferNotionalFee(address, uint256) external returns (bool);
  function rates() external view returns (uint256, uint256);
  function balancesOf(address) external view returns (uint256, uint256);
}

interface ISwivel {
  function authRedeem(uint8, address, address, address, uint256) external returns (bool);
}
