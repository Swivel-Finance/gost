// SPDX-License-Identifier: UNLICENSED

// TODO update to 0.8.4 (or whatever latest is...)

// NOTE the pattern [underlying, maturity*, cToken, ...]

pragma solidity 0.8.0;

import './Abstracts.sol';
import './ZcToken.sol';
import './VaultTracker.sol';

contract MarketPlace {
  address public admin = msg.sender;
  address public swivel;

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

  function setSwivelAddress(address s) external onlyAdmin(admin) returns (bool) {
    swivel = s;
    return true;
  }

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
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  function matureMarket(address u, uint256 m) public returns (bool) {
    require(mature[u][m] == false, 'market already matured');
    require(block.timestamp >= ZcToken(markets[u][m].zcTokenAddr).maturity(), "maturity not reached");

    // Set the base maturity cToken exchange rate at maturity to the current cToken exchange rate
    uint256 currentExchangeRate = CErc20(markets[u][m].cTokenAddr).exchangeRateCurrent();
    maturityRate[u][m] = currentExchangeRate;

    // Set Floating Market "matured" to true
    require(VaultTracker(markets[u][m].vaultAddr).matureVault() == true, 'maturity not reached');

    // Set the maturity state to true (for zcb market)
    mature[u][m] = true;

    // TODO review indexed args...
    emit Mature(u, m, block.timestamp, currentExchangeRate);

    return true;
  }

  /// @notice Allows zcToken holders to redeem their tokens for underlying tokens after maturity has been reached.
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param a Amount of zcTokens being redeemed
  function redeemZcToken(address u, uint256 m, uint256 a) public returns (bool) {
    // If market hasn't matured, mature it and redeem exactly the amount
    if (mature[u][m] == false) {
      // Attempt to Mature it
      require(matureMarket(u, m) == true, 'failed to mature the market');

      // Burn user's zcTokens
      require(ZcToken(markets[u][m].zcTokenAddr).burn(msg.sender, a), 'could not burn');

      // Redeem principleReturned of underlying token to Swivel Contract from Compound
      require(CErc20(markets[u][m].cTokenAddr).redeemUnderlying(a) == 0 ,'cToken redemption failed');

      // Transfer the principleReturned in underlying tokens to the user
      require(Erc20(u).transfer(msg.sender, a), 'transfer of redemption failed');
    }
    // If market has matured, redeem the amount + the marginal floating interest generated on Compound since maturity
    else {
      // Burn user's zcTokens
      require(ZcToken(markets[u][m].zcTokenAddr).burn(msg.sender, a), 'could not burn');

      // Call internal function to determine the amount of principle to return (including marginal interest since maturity)
      uint256 principleReturned = calculateReturn(u, m, a);

      // Redeem principleReturned of underlying token to Swivel Contract from Compound
      require(CErc20(markets[u][m].cTokenAddr).redeemUnderlying(principleReturned) == 0 ,'cToken redemption failed');

      // Transfer the principleReturned in underlying tokens to the user
      require(Erc20(u).transfer(msg.sender, principleReturned), 'transfer of redemption failed');

    }
    return true;
  }

  /// @notice Allows Vault owners to redeem any currently accrued interest within a given ____
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  function redeemVaultInterest(address u, uint256 m) public returns (bool) {
    // Call to the floating market contract to release the position and calculate the interest generated
    uint256 interestGenerated = VaultTracker(markets[u][m].vaultAddr).redeemInterest(msg.sender);

    // Redeem the interest generated by the position to Swivel Contract from Compound
    require(CErc20(markets[u][m].cTokenAddr).redeemUnderlying(interestGenerated) == 0, "redemption from Compound Failed");

    // Transfer the interest generated in underlying tokens to the user
    require(Erc20(u).transfer(msg.sender, interestGenerated), 'transfer of redeemable failed');

    return true;
  }

  /// @notice Calculates the total amount of underlying returned including interest generated since the `matureMarket` function has been called
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param a Amount of zcTokens being redeemed
  function calculateReturn(address u, uint256 m, uint256 a) internal returns (uint256) {
    // Calculate difference between the cToken exchange rate @ maturity and the current cToken exchange rate
    uint256 rateDifference = CErc20(markets[u][m].cTokenAddr).exchangeRateCurrent() - maturityRate[u][m];

    // Calculate the yield generated after maturity in %. Precise to 9 decimals (5.25% = .0525 = 52500000)
    uint256 residualYield = (((rateDifference * 1e26) / maturityRate[u][m]) / 1e17) + 1E9;

    // Calculate the total amount of underlying principle to return
    return (residualYield * a) / 1e9;
  }

  function cTokenAddress(address a, uint256 m) external view returns (address) {
    return markets[a][m].cTokenAddr;
  }

  /// @notice called by swivel IVFZI && IZFVI
  /// @dev call with underlying, maturity, mint-target, add-notional-target and an amount
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param z Recipient of the minted zcToken
  /// @param n Recipient of the added notional
  /// @param a Amount of zcToken minted and notional added
  function custodialInitiate(address u, uint256 m, address z, address n, uint256 a) external onlySwivel(swivel) returns (bool) {
    require(ZcToken(markets[u][m].zcTokenAddr).mint(z, a), 'mint failed');
    require(VaultTracker(markets[u][m].vaultAddr).addNotional(n, a), 'add notional failed');
    return true;
  }

  /// @notice called by swivel EVFZE FF EZFVE
  /// @dev call with underlying, maturity, burn-target, remove-notional-target and an amount
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param z Owner of the zcToken to be burned
  /// @param n Target to remove notional from
  /// @param a Amount of zcToken burned and notional removed
  function custodialExit(address u, uint256 m, address z, address n, uint256 a) external onlySwivel(swivel) returns (bool) {
    require(ZcToken(markets[u][m].zcTokenAddr).burn(z, a), 'burn failed');
    require(VaultTracker(markets[u][m].vaultAddr).removeNotional(n, a), 'remove notional failed');
    return true;
  }

  /// @notice called by swivel IZFZE, EZFZI
  /// @dev call with underlying, maturity, transfer-from, transfer-to, amount
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param f Owner of the zcToken to be transferred
  /// @param t Target to be transferred to
  /// @param a Amount of zcToken transfer
  function p2pZcTokenExchange(address u, uint256 m, address f, address t, uint256 a) external onlySwivel(swivel) returns (bool) {
    require(ZcToken(markets[u][m].zcTokenAddr).transferFrom(f, t, a), 'zcToken transfer failed');
    return true;
  }

  /// @notice called by swivel IVFVE, EVFVI
  /// @dev call with underlying, maturity, remove-from, add-to, amount
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param f Owner of the notional to be transferred
  /// @param t Target to be transferred to
  /// @param a Amount of notional transfer
  function p2pVaultExchange(address u, uint256 m, address f, address t, uint256 a) external onlySwivel(swivel) returns (bool) {
    require(VaultTracker(markets[u][m].vaultAddr).transferNotionalFrom(f, t, a), 'transfer notional failed');
    return true;
  }

  modifier onlyAdmin(address a) {
    require(msg.sender == a, 'sender must be admin');
    _;
  }

  modifier onlySwivel(address s) {
    require(msg.sender == s, 'sender must be swivel contract');
    _;
  }
}
