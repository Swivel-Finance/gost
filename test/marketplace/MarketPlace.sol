// SPDX-License-Identifier: UNLICENSED

// TODO update to 0.8.4 (or whatever latest is...)
// TODO add swivel address
// TODO add setter for swivel address (only admin, once only if desired...)
// TODO add onlySwivel
// TODO add cTokenAddress
// TODO add p2p methods
// TODO add custodial methods

// NOTE the pattern [underlying, maturity*, cToken, ...]

pragma solidity 0.8.0;

import './Abstracts.sol';
import './ZcToken.sol';
import './VaultTracker.sol';

contract MarketPlace {
  address public admin = msg.sender;

  struct Market {
    address cTokenAddr;
    address zcTokenAddr;
    address vaultAddr;
  }

  mapping (address => mapping (uint256 => Market)) public markets;
  mapping (address => mapping (uint256 => bool)) public mature;
  mapping (address => mapping (uint256 => uint256)) public maturityRate;

  // TODO what should be indexed here?
  event Create(address indexed underlying, uint256 indexed maturity, address cToken, address zcToken); 
  event Mature(address indexed underlying, uint256 indexed maturity, uint256 maturityRate, uint256 matured);

  /// @notice Allows the owner to create new markets
  /// @param u Underlying token address associated with the new market
  /// @param m Maturity timestamp of the new market
  /// @param c cToken address associated with underlying for the new market 
  /// @param n Name of the new zcToken market
  /// @param s Symbol of the new zcToken market
  function createMarket(
    address u,
    uint256 m,
    address c,
    string memory n,
    string memory s
  ) public onlyAdmin(admin) returns (bool) {
    // TODO can we live with the factory pattern here both bytecode size wise and CREATE opcode cost wise?
    address zctAddr = address(new ZcToken(u, m, n, s));
    address vAddr = address(new VaultTracker(m, c));
    markets[u][m] = Market(c, zctAddr, vAddr);

    // TODO review indexed args...
    emit Create(u, m, c, zctAddr);

    return true;
  }

  /// @notice Can be called after maturity, allowing all of the zcTokens to earn floating interest on Compound until they release their funds
  /// @param u Underlying token address associated with the new market
  /// @param m Maturity timestamp of the new market
  function matureMarket(address u, uint256 m) public returns (bool) {
    require(mature[u][m] == false, 'Market already matured');
    require(block.timestamp >= ZcToken(markets[u][m].zcTokenAddr).maturity(), "Maturity not reached");

    // Set the base maturity cToken exchange rate at maturity to the current cToken exchange rate
    uint256 currentExchangeRate = CErc20(markets[u][m].cTokenAddr).exchangeRateCurrent();
    maturityRate[u][m] = currentExchangeRate;

    // Set Floating Market "matured" to true
    require(VaultTracker(markets[u][m].vaultAddr).matureVault() == true, 'Maturity not reached');

    // Set the maturity state to true (for zcb market)
    mature[u][m] = true;

    // TODO review indexed args...
    emit Mature(u, m, block.timestamp, currentExchangeRate);

    return true;
  }

  function redeemZcToken(address u, uint256 m, uint256 a) public returns (bool) {
    // ...  

    return true;
  }

  function reddemVaultInterest(address u, uint256 m) public returns (bool) {

    return true;
  }

  function calculateReturn(address u, uint256 m, uint256 a) internal returns (uint256) {
    uint256 total;

    return total;
  }

  modifier onlyAdmin(address a) {
    require(msg.sender == a, 'sender must be admin');
    _;
  }
}
