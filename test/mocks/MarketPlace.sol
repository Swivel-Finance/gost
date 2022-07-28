// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.13;

/// @dev MarketPlace is a mock whose bindings are imported by unit tests in any pkg/*testing that needs it
contract MarketPlace {
  struct Market {
    address cTokenAddr;
    address zcToken;
    address vaultTracker;
    uint256 maturityRate;
  }

  struct MethodArgs {
    address underlying;
    uint256 maturity;
    address one; // is sender or maker depending on method
    address two; // same as above
    uint256 amount;
  }

  mapping (uint8 => mapping (address => mapping (uint256 => Market))) public markets;
  mapping (uint8 => MethodArgs) public cTokenAddressCalled;
  mapping (uint8 => MethodArgs) public custodialInitiateCalled;
  mapping (uint8 => MethodArgs) public custodialExitCalled;
  mapping (uint8 => MethodArgs) public p2pZcTokenExchangeCalled;
  mapping (uint8 => MethodArgs) public p2pVaultExchangeCalled;
  mapping (uint8 => MethodArgs) public mintZcTokenAddingNotionalCalled;
  mapping (uint8 => MethodArgs) public burnZcTokenRemovingNotionalCalled;
  mapping (uint8 => MethodArgs) public transferVaultNotionalFeeCalled;
  mapping (uint8 => MethodArgs) public redeemZcTokenCalled;
  mapping (uint8 => MethodArgs) public redeemVaultInterestCalled;
  mapping (uint8 => address) public exchangeRateCalled;

  address private cTokenAddr;
  bool private custodialInitiateReturn;
  bool private custodialExitReturn;
  bool private p2pZcTokenExchangeReturn;
  bool private p2pVaultExchangeReturn;
  bool private mintZcTokenAddingNotionalReturn;
  bool private burnZcTokenRemovingNotionalReturn;
  bool private transferVaultNotionalFeeReturn;
  uint256 private redeemZcTokenReturn;
  uint256 private redeemVaultInterestReturn;
  uint256 private exchangeRateReturn;

  /// @param p Protocol enum value
  /// @param u Underlying asset address
  /// @param m Maturity
  /// @param c Compounding token address
  /// @param z ZcToken address for the market
  /// @param v VaultTracker address for the market
  function createMarket(uint8 p, address u, uint256 m, address c, address z, address v) external {
    markets[p][u][m] = Market(c, z, v, 0);
  }

  function cTokenAddressReturns(address c) external {
    cTokenAddr = c;
  }

  function cTokenAddress(uint8 p, address u, uint256 m) external returns (address) {
    MethodArgs memory args; 
    args.underlying = u;
    args.maturity = m;
    cTokenAddressCalled[p] = args;
    return cTokenAddr;
  }

  function custodialInitiateReturns(bool b) external {
    custodialInitiateReturn = b;
  }

  // called by swivel IVFZI && IZFVI 
  // call with underlying, maturity, mint-target, add-notional-target and an amount
  function custodialInitiate(uint8 p, address u, uint256 m, address o, address t, uint256 a) external returns (bool) {
    MethodArgs memory args; 
    args.underlying = u;
    args.maturity = m;
    args.one = o; // will be the recipient of minted zctoken
    args.two = t; // will be the recipient of added notional
    args.amount = a; // the amount of minted zctoken and notional added
    custodialInitiateCalled[p] = args;

    return custodialInitiateReturn;
  }

  function custodialExitReturns(bool b) external {
    custodialExitReturn = b;
  }

  // called by swivel EVFZE FF EZFVE
  // call with underlying, maturity, burn-target, remove-notional-target and an amount
  function custodialExit(uint8 p, address u, uint256 m, address o, address t, uint256 a) external returns (bool) {
    MethodArgs memory args; 
    args.underlying = u;
    args.maturity = m;
    args.one = o; // will be the burn-from target
    args.two = t; // will be the remove-notional target
    args.amount = a; // zctoken burned, notional removed
    custodialExitCalled[p] = args;

    return custodialExitReturn;
  }

  function p2pZcTokenExchangeReturns(bool b) external {
    p2pZcTokenExchangeReturn = b;
  }

  // called by swivel IZFZE, EZFZI
  // call with underlying, maturity, transfer-from, transfer-to, amount
  function p2pZcTokenExchange(uint8 p, address u, uint256 m, address o, address t, uint256 a) external returns (bool) {
    MethodArgs memory args;
    args.underlying = u;
    args.maturity = m;
    args.one = o;
    args.two = t;
    args.amount = a;
    p2pZcTokenExchangeCalled[p] = args;

    return p2pZcTokenExchangeReturn;
  }

  function p2pVaultExchangeReturns(bool b) external {
    p2pVaultExchangeReturn = b;
  }

  // called by swivel IVFVE, EVFVI
  // call with underlying, maturity, remove-from, add-to, amount
  function p2pVaultExchange(uint8 p, address u, uint256 m, address o, address t, uint256 a) external returns (bool) {
    MethodArgs memory args;
    args.underlying = u;
    args.maturity = m;
    args.one = o;
    args.two = t;
    args.amount = a;
    p2pVaultExchangeCalled[p] = args;

    return p2pVaultExchangeReturn;
  }

  function mintZcTokenAddingNotionalReturns(bool b) external {
    mintZcTokenAddingNotionalReturn = b;
  }

  // call with underlying, maturity, mint-to, amount
  function mintZcTokenAddingNotional(uint8 p, address u, uint256 m, address t, uint256 a) external returns (bool) {
    MethodArgs memory args;
    args.underlying = u;
    args.maturity = m;
    args.one = t;
    args.amount = a;
    mintZcTokenAddingNotionalCalled[p] = args;

    return mintZcTokenAddingNotionalReturn;
  }

  function burnZcTokenRemovingNotionalReturns(bool b) external {
    burnZcTokenRemovingNotionalReturn = b;
  }

  // call with underlying, maturity, mint-to, amount
  function burnZcTokenRemovingNotional(uint8 p, address u, uint256 m, address t, uint256 a) external returns (bool) {
    MethodArgs memory args;
    args.underlying = u;
    args.maturity = m;
    args.one = t;
    args.amount = a;
    burnZcTokenRemovingNotionalCalled[p] = args;

    return burnZcTokenRemovingNotionalReturn;
  }

  function transferVaultNotionalFeeReturns(bool b) external {
    transferVaultNotionalFeeReturn = b;
  }

  function transferVaultNotionalFee(uint8 p, address u, uint256 m, address f, uint256 a) external returns (bool) {
    MethodArgs memory args;
    args.underlying = u;
    args.maturity = m;
    args.one = f;
    args.amount = a;
    transferVaultNotionalFeeCalled[p] = args;

    return transferVaultNotionalFeeReturn;
  }

  function redeemZcTokenReturns(uint256 a) external {
    redeemZcTokenReturn = a;
  }

  function redeemZcToken(uint8 p, address u, uint256 m, address t, uint256 a) external returns (uint256) {
    MethodArgs memory args;
    args.underlying = u;
    args.maturity = m;
    args.one = t;
    args.amount = a;
    redeemZcTokenCalled[p] = args;

    return redeemZcTokenReturn;
  }

  function redeemVaultInterestReturns(uint256 a) external {
    redeemVaultInterestReturn = a;
  }

  function redeemVaultInterest(uint8 p, address u, uint256 m, address t) external returns (uint256) {
    MethodArgs memory args;
    args.underlying = u;
    args.maturity = m;
    args.one = t;
    redeemVaultInterestCalled[p] = args;

    return redeemVaultInterestReturn;
  }

  function exchangeRateReturns(uint256 a) external {
    exchangeRateReturn = a;
  }

  function exchangeRate(uint8 p, address c) external returns (uint256) {
    exchangeRateCalled[p] = c;
    return exchangeRateReturn;
  }
}
