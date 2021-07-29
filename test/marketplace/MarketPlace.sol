// SPDX-License-Identifier: UNLICENSED

// TODO update to 0.8.4 (or whatever latest is...)

// NOTE the pattern [underlying, maturity*, cToken, ...]

pragma solidity 0.8.4;

import './Abstracts.sol';
import './ZcToken.sol';
import './VaultTracker.sol';

contract MarketPlace {
  struct Market {
    address cTokenAddr;
    address zcTokenAddr;
    address vaultAddr;
  }

  mapping (address => mapping (uint256 => Market)) public markets;
  mapping (address => mapping (uint256 => bool)) public mature;
  mapping (address => mapping (uint256 => uint256)) public maturityRate;

  address public immutable admin;
  address public swivel;

  event Create(address indexed underlying, uint256 indexed maturity, address cToken, address zcToken);
  event Mature(address indexed underlying, uint256 indexed maturity, uint256 maturityRate, uint256 matured);
  event RedeemZcToken(address indexed underlying, uint256 indexed maturity, address indexed sender, uint256 amount);
  event RedeemVaultInterest(address indexed underlying, uint256 indexed maturity, address indexed sender);
  event CustodialInitiate(address indexed underlying, uint256 indexed maturity, address zcTarget, address nTarget, uint256 amount);
  event CustodialExit(address indexed underlying, uint256 indexed maturity, address zcTarget, address nTarget, uint256 amount);
  event P2pZcTokenExchange(address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount);
  event P2pVaultExchange(address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount);
  event TransferVaultNotional(address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount);

  constructor() {
    admin = msg.sender;
  }

  /// @param s Address of the deployed swivel contract
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

    emit Create(u, m, c, zctAddr);

    return true;
  }

  /// @notice Can be called after maturity, allowing all of the zcTokens to earn floating interest on Compound until they release their funds
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  function matureMarket(address u, uint256 m) public returns (bool) {
    require(!mature[u][m], 'market already matured');
    require(block.timestamp >= ZcToken(markets[u][m].zcTokenAddr).maturity(), "maturity not reached");

    // Set the base maturity cToken exchange rate at maturity to the current cToken exchange rate
    uint256 currentExchangeRate = CErc20(markets[u][m].cTokenAddr).exchangeRateCurrent();
    maturityRate[u][m] = currentExchangeRate;
    // Set the maturity state to true (for zcb market)
    mature[u][m] = true;

    // Set Floating Market "matured" to true
    require(VaultTracker(markets[u][m].vaultAddr).matureVault(), 'maturity not reached');

    emit Mature(u, m, block.timestamp, currentExchangeRate);

    return true;
  }

  /// @notice Allows zcToken holders to redeem their tokens for underlying tokens after maturity has been reached.
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param a Amount of zcTokens being redeemed
  function redeemZcToken(address u, uint256 m, uint256 a) public returns (bool) {
    // If market hasn't matured, mature it and redeem exactly the amount

    Market memory mkt = markets[u][m];

    if (!mature[u][m]) {
      // Attempt to Mature it
      require(matureMarket(u, m), 'failed to mature the market');

      // Burn user's zcTokens
      require(ZcToken(mkt.zcTokenAddr).burn(msg.sender, a), 'could not burn');

      // Redeem principleReturned of underlying token to Swivel Contract from Compound
      require(CErc20(mkt.cTokenAddr).redeemUnderlying(a) == 0 ,'cToken redemption failed');

      // Transfer the principleReturned in underlying tokens to the user
      Erc20(u).transfer(msg.sender, a);
    } else { // If market has matured, redeem the amount + the marginal floating interest generated on Compound since maturity
      require(ZcToken(mkt.zcTokenAddr).burn(msg.sender, a), 'could not burn'); // Burn user's zcTokens

      // Call internal function to determine the amount of principle to return (including marginal interest since maturity)
      uint256 principleReturned = calculateReturn(u, m, a);

      // Redeem principleReturned of underlying token to Swivel Contract from Compound
      require(CErc20(mkt.cTokenAddr).redeemUnderlying(principleReturned) == 0 ,'cToken redemption failed');

      // Transfer the principleReturned in underlying tokens to the user
      Erc20(u).transfer(msg.sender, principleReturned);
    }

    emit RedeemZcToken(u, m, msg.sender, a);

    return true;
  }

  /// @notice Allows Vault owners to redeem any currently accrued interest within a given ____
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  function redeemVaultInterest(address u, uint256 m) public returns (bool) {
    // Call to the floating market contract to release the position and calculate the interest generated
    uint256 interestGenerated = VaultTracker(markets[u][m].vaultAddr).redeemInterest(msg.sender);

    // Redeem the interest generated by the position to Swivel Contract from Compound
    require(CErc20(markets[u][m].cTokenAddr).redeemUnderlying(interestGenerated) == 0, "redemption from Compound failed");

    // Transfer the interest generated in underlying tokens to the user
    Erc20(u).transfer(msg.sender, interestGenerated);

    emit RedeemVaultInterest(u, m, msg.sender);

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
    emit CustodialInitiate(u, m, z, n, a);
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
    emit CustodialExit(u, m, z, n, a);
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
    emit P2pZcTokenExchange(u, m, f, t, a);
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
    emit P2pVaultExchange(u, m, f, t, a);
    return true;
  }

  /// @notice External method giving access to this functionality within a given vault
  /// @dev Note that this method calculates yield and interest as well
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param t Target to be transferred to
  /// @param a Amount of notional to be transferred
  function transferVaultNotional(address u, uint256 m, address t, uint256 a) public returns (bool) {
    require(VaultTracker(markets[u][m].vaultAddr).transferNotionalFrom(msg.sender, t, a), 'vault transfer failed');
    emit TransferVaultNotional(u, m, msg.sender, t, a);
    return true;
  }

  modifier onlyAdmin(address a) {
    require(msg.sender == a, 'sender must be admin');
    _;
  }

  modifier onlySwivel(address s) {
    require(msg.sender == s, 'sender must be Swivel contract');
    _;
  }
}
