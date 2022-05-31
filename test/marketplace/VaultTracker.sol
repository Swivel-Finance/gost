// SPDX-License-Identifier: UNLICENSED

/**
  VaultTracker is a mock which records arguments passed to its methods as well as
  provides setters allowing us to dictate method return values
*/

pragma solidity 0.8.13;

contract VaultTracker {
  struct TransferNotionalFromArgs {
    address to;
    uint256 amount;
  }

  // mapping of arguments sent to addNotional. key is the passed in address.
  mapping (address => uint256) public addNotionalCalled;
  // mapping of arguments sent to removeNotional. key is the passed in address.
  mapping (address => uint256) public removeNotionalCalled;
  // mapping of arguments sent to transferNotionalFrom. key is the passed in address.
  mapping (address => TransferNotionalFromArgs) public transferNotionalFromCalled;
  // mapping of args sent to transferFee, key is the given payer's address
  mapping (address => uint256) public transferNotionalFeeCalled;

  address public cTokenAddr;
  address public swivel;
  address public redeemInterestCalled;
  uint256 public matureVaultCalled;

  uint256 private maturityReturn;
  uint256 private redeemInterestReturn;
  bool private matureVaultReturn;
  // a boolean flag which allows us to dictate the return of addNotional().
  bool private addNotionalReturn;
  // a boolean flag which allows us to dictate the return of removeNotional().
  bool private removeNotionalReturn;
  // a boolean flag which allows us to dictate the return of transferNotionalFrom().
  bool private transferNotionalFromReturn;
  bool private transferNotionalFeeReturn;

  /// @param m maturity
  /// @param c cToken address
  /// @param s deployed swivel contract address
  constructor(uint256 m, address c, address s) {
    maturityReturn = m;
    cTokenAddr = c;
    swivel = s;
  }

  function redeemInterestReturns(uint256 a) public {
    redeemInterestReturn = a;
  }

  function redeemInterest(address o) public returns (uint256) {
    redeemInterestCalled = o;
    return redeemInterestReturn;
  }

  function maturityReturns(uint256 n) public {
    maturityReturn = n;
  }

  // override what would be the autogenerated getter...
  function maturity() public view returns (uint256) {
    return maturityReturn;
  }

  function matureVault(uint256 c) public returns (bool) {
    matureVaultCalled = c;
    return matureVaultReturn;
  }

  function matureVaultReturns(bool b) public {
    matureVaultReturn = b;
  }

  function addNotionalReturns(bool b) public {
    addNotionalReturn = b;
  }

  function addNotional(address o, uint256 a) public returns (bool) {
    addNotionalCalled[o] = a;
    return addNotionalReturn;
  }

  function removeNotionalReturns(bool b) public {
    removeNotionalReturn = b;
  }

  function removeNotional(address o, uint256 a) public returns (bool) {
    removeNotionalCalled[o] = a;
    return removeNotionalReturn;
  }

  function transferNotionalFromReturns(bool b) public {
    transferNotionalFromReturn = b;
  }

  function transferNotionalFrom(address f, address t, uint256 a) public returns (bool) {
    TransferNotionalFromArgs memory args;
    args.to = t;
    args.amount = a;
    transferNotionalFromCalled[f] = args;
    return transferNotionalFromReturn;
  }

  function transferNotionalFeeReturns(bool b) public {
    transferNotionalFeeReturn = b;
  }

  function transferNotionalFee(address f, uint256 a) public returns (bool) {
    transferNotionalFeeCalled[f] = a;
    return transferNotionalFeeReturn;
  }
}
