// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Interfaces.sol';
import './ZcToken.sol';
import './VaultTracker.sol';

contract MarketPlace {
  struct Market {
    address cTokenAddr;
    address adapterAddr;
    address zcToken;
    address vaultTracker;
    uint256 maturityRate;
  }

  mapping (uint8 => mapping (address => mapping (uint256 => Market))) public markets;

  address public admin;
  address public swivel;
  bool public paused;

  // NOTE I don't believe the address of the adapter needs to be published as it can be fetched at any point in relation to the compounding token
  event Create(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address cToken, address zcToken, address vaultTracker, address adapter);
  event Mature(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, uint256 maturityRate, uint256 matured);
  event RedeemZcToken(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address sender, uint256 amount);
  event RedeemVaultInterest(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address sender);
  event CustodialInitiate(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address zcTarget, address nTarget, uint256 amount);
  event CustodialExit(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address zcTarget, address nTarget, uint256 amount);
  event P2pZcTokenExchange(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount);
  event P2pVaultExchange(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount);
  event TransferVaultNotional(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount);

  constructor() {
    admin = msg.sender;
  }

  // TODO becomes `setSwivel`
  /// @param s Address of the deployed swivel contract
  /// @notice We only allow this to be set once
  function setSwivelAddress(address s) external authorized(admin) returns (bool) {
    require(swivel == address(0), 'swivel contract address already set');
    swivel = s;
    return true;
  }

  // TODO becomes `setAdmin`
  /// @param a Address of a new admin
  function transferAdmin(address a) external authorized(admin) returns (bool) {
    admin = a;
    return true;
  }

  /// @notice Allows the owner to create new markets
  /// @param m Maturity timestamp of the new market
  /// @param c Compounding Token address associated with the new market
  /// @param a Deployed adapter address to use for this new market
  /// @param n Name of the new market zcToken
  /// @param s Symbol of the new market zcToken
  function createMarket(
    uint256 m,
    address c,
    address a,
    string memory n,
    string memory s
  ) external authorized(admin) unpaused() returns (bool) {
    require(swivel != address(0), 'swivel contract address not set');

    uint8 protocol;
    address underAddr;

    {
      ICompounding adapter = ICompounding(a);
      protocol = adapter.PROTOCOL();
      underAddr = adapter.underlying(c);
    }

    require(markets[protocol][underAddr][m].vaultTracker == address(0), 'market already exists');

    address zct;
    address tracker;

    {
      uint8 decimals = IErc20(underAddr).decimals();
      zct = address(new ZcToken(underAddr, m, n, s, decimals));
      tracker = address(new VaultTracker(m, c, a, swivel));
    }

    markets[protocol][underAddr][m] = Market(c, a, zct, tracker, 0);

    emit Create(protocol, underAddr, m, c, zct, tracker, a);

    return true;
  }

  /// @notice Can be called after maturity, allowing all of the zcTokens to earn floating interest on Compound until they release their funds
  /// @param p Protocol Enum value associated with the market being matured
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  function matureMarket(uint8 p, address u, uint256 m) public unpaused() returns (bool) {
    Market memory mkt = markets[p][u][m];

    require(mkt.maturityRate == 0, 'market already matured');
    require(block.timestamp >= m, "maturity not reached");

    // set the base maturity cToken exchange rate at maturity to the current cToken exchange rate
    uint256 exchangeRate = ICompounding(mkt.adapterAddr).exchangeRate(mkt.cTokenAddr);
    markets[p][u][m].maturityRate = exchangeRate;

    // TODO remove the require here? why?
    require(VaultTracker(mkt.vaultTracker).matureVault(exchangeRate), 'mature vault failed');

    emit Mature(p, u, m, exchangeRate, block.timestamp);

    return true;
  }

  /// @notice Allows Swivel caller to deposit their underlying, in the process splitting it - minting both zcTokens and vault notional.
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param t Address of the depositing user
  /// @param a Amount of notional being added
  function mintZcTokenAddingNotional(uint8 p, address u, uint256 m, address t, uint256 a) external authorized(swivel) unpaused() returns (bool) {
    Market memory mkt = markets[p][u][m];
    require(ZcToken(mkt.zcToken).mint(t, a), 'mint zcToken failed');
    require(VaultTracker(mkt.vaultTracker).addNotional(t, a), 'add notional failed');
    
    return true;
  }

  /// @notice Allows Swivel caller to deposit/burn both zcTokens + vault notional. This process is "combining" the two and redeeming underlying.
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param t Address of the combining/redeeming user
  /// @param a Amount of zcTokens being burned
  function burnZcTokenRemovingNotional(uint8 p, address u, uint256 m, address t, uint256 a) external authorized(swivel) unpaused() returns(bool) {
    Market memory mkt = markets[p][u][m];
    require(ZcToken(mkt.zcToken).burn(t, a), 'burn failed');
    require(VaultTracker(mkt.vaultTracker).removeNotional(t, a), 'remove notional failed');
    
    return true;
  }

  /// @notice Allows (via swivel) zcToken holders to redeem their tokens for underlying tokens after maturity has been reached.
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param t Address of the redeeming user
  /// @param a Amount of zcTokens being redeemed
  function redeemZcToken(uint8 p, address u, uint256 m, address t, uint256 a) external authorized(swivel) unpaused() returns (uint256) {
    Market memory mkt = markets[p][u][m];

    // if the market has not matured, mature it and redeem exactly the amount
    if (mkt.maturityRate == 0) {
      require(matureMarket(p, u, m), 'failed to mature the market');
    }

    // burn user's zcTokens
    require(ZcToken(mkt.zcToken).burn(t, a), 'could not burn');

    emit RedeemZcToken(p, u, m, t, a);

    if (mkt.maturityRate == 0) {
      return a;
    } else { 
      // if the market was already mature the return should include the amount + marginal floating interest generated on Compound since maturity
      return calculateReturn(p, u, m, a);
    }
  }

  /// @notice Allows Vault owners (via Swivel) to redeem any currently accrued interest
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param t Address of the redeeming user
  function redeemVaultInterest(uint8 p, address u, uint256 m, address t) external authorized(swivel) unpaused() returns (uint256) {
    // call to the floating market contract to release the position and calculate the interest generated
    uint256 interest = VaultTracker(markets[p][u][m].vaultTracker).redeemInterest(t);

    emit RedeemVaultInterest(p, u, m, t);

    return interest;
  }

  /// @notice Calculates the total amount of underlying returned including interest generated since the `matureMarket` function has been called
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param a Amount of zcTokens being redeemed
  function calculateReturn(uint8 p, address u, uint256 m, uint256 a) internal returns (uint256) {
    Market memory mkt = markets[p][u][m];

    uint256 exchangeRate = ICompounding(mkt.adapterAddr).exchangeRate(mkt.cTokenAddr);

    return (a * exchangeRate) / mkt.maturityRate;
  }

  /// @notice Return the compounding token address for a given market
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  function cTokenAndAdapterAddress(uint8 p, address u, uint256 m) external view returns (address, address) {
    Market memory mkt = markets[p][u][m];
    return (mkt.cTokenAddr, mkt.adapterAddr);
  }

  /// @notice Called by swivel IVFZI && IZFVI
  /// @dev Call with protocol, underlying, maturity, mint-target, add-notional-target and an amount
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param z Recipient of the minted zcToken
  /// @param n Recipient of the added notional
  /// @param a Amount of zcToken minted and notional added
  function custodialInitiate(uint8 p, address u, uint256 m, address z, address n, uint256 a) external authorized(swivel) unpaused() returns (bool) {
    Market memory mkt = markets[p][u][m];
    require(ZcToken(mkt.zcToken).mint(z, a), 'mint failed');
    require(VaultTracker(mkt.vaultTracker).addNotional(n, a), 'add notional failed');
    emit CustodialInitiate(p, u, m, z, n, a);
    return true;
  }

  /// @notice Called by swivel EVFZE FF EZFVE
  /// @dev Call with protocol, underlying, maturity, burn-target, remove-notional-target and an amount
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param z Owner of the zcToken to be burned
  /// @param n Target to remove notional from
  /// @param a Amount of zcToken burned and notional removed
  function custodialExit(uint8 p, address u, uint256 m, address z, address n, uint256 a) external authorized(swivel) unpaused() returns (bool) {
    Market memory mkt = markets[p][u][m];
    require(ZcToken(mkt.zcToken).burn(z, a), 'burn failed');
    require(VaultTracker(mkt.vaultTracker).removeNotional(n, a), 'remove notional failed');
    emit CustodialExit(p, u, m, z, n, a);
    return true;
  }

  /// @notice Called by swivel IZFZE, EZFZI
  /// @dev Call with underlying, maturity, transfer-from, transfer-to, amount
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param f Owner of the zcToken to be burned
  /// @param t Target to be minted to
  /// @param a Amount of zcToken transfer
  function p2pZcTokenExchange(uint8 p, address u, uint256 m, address f, address t, uint256 a) external authorized(swivel) unpaused() returns (bool) {
    Market memory mkt = markets[p][u][m];
    require(ZcToken(mkt.zcToken).burn(f, a), 'zcToken burn failed');
    require(ZcToken(mkt.zcToken).mint(t, a), 'zcToken mint failed');
    emit P2pZcTokenExchange(p, u, m, f, t, a);
    return true;
  }

  /// @notice Called by swivel IVFVE, EVFVI
  /// @dev Call with protocol, underlying, maturity, remove-from, add-to, amount
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param f Owner of the notional to be transferred
  /// @param t Target to be transferred to
  /// @param a Amount of notional transfer
  function p2pVaultExchange(uint8 p, address u, uint256 m, address f, address t, uint256 a) external authorized(swivel) unpaused() returns (bool) {
    require(VaultTracker(markets[p][u][m].vaultTracker).transferNotionalFrom(f, t, a), 'transfer notional failed');
    emit P2pVaultExchange(p, u, m, f, t, a);
    return true;
  }

  /// @notice External method giving access to this functionality within a given vault
  /// @dev Note that this method calculates yield and interest as well
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param t Target to be transferred to
  /// @param a Amount of notional to be transferred
  function transferVaultNotional(uint8 p, address u, uint256 m, address t, uint256 a) external unpaused() returns (bool) {
    require(VaultTracker(markets[p][u][m].vaultTracker).transferNotionalFrom(msg.sender, t, a), 'vault transfer failed');
    emit TransferVaultNotional(p, u, m, msg.sender, t, a);
    return true;
  }

  /// @notice Transfers notional fee to the Swivel contract without recalculating marginal interest for from
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param f Owner of the amount
  /// @param a Amount to transfer
  function transferVaultNotionalFee(uint8 p, address u, uint256 m, address f, uint256 a) external authorized(swivel) returns (bool) {
    VaultTracker(markets[p][u][m].vaultTracker).transferNotionalFee(f, a);
    return true;
  }

  // TODO this becomes per-protocol
  /// @notice Called by admin at any point to pause / unpause market transactions
  /// @param b Boolean which indicates the markets paused status
  function pause(bool b) external authorized(admin) returns (bool) {
    paused = b;
    return true;
  }

  modifier authorized(address a) {
    require(msg.sender == a, 'sender must be authorized');
    _;
  }

  modifier unpaused() {
    require(!paused, 'markets are paused');
    _;
  }
}
