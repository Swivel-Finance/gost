// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.0;

/// @dev MarketPlace is a mock whose bindings are imported by unit tests in any pkg/*testing that needs it
contract MarketPlace {
  // the address that will be returned when marketCTokenAddress is called
  address private cTokenAddr;
  // the arguments that were passed to marketCTokenAddress when it was called
  // TODO: we should likely standardize on using the `Called` suffix for these mappings in the mocks.
  //       the token mocks use a differnt pattern, change them to match this...
  mapping (address => uint256) public cTokenAddressCalled;
  bool private mintZcTokenAddingNotionalReturn;
  bool private burnZcTokenRemovingNotionalReturn;
  bool private transferFromZcTokenReturn;
  bool private transferFromNotionalReturn;

  struct MintZcTokenAddingNotionalArgs {
    uint256 maturity;
    address owner;
    address sender;
    uint256 amount;
  }

  struct BurnZcTokenRemovingNotionalArgs {
    uint256 maturity;
    address owner;
    address sender;
    uint256 amount;
  }

  struct TransferFromZcTokenArgs {
    uint256 maturity;
    address owner; // should become to
    address sender; // should become from
    uint256 amount;
  }

  struct TransferFromNotionalArgs {
    uint256 maturity;
    address owner; // should become from
    address sender; // should become to
    uint256 amount;
  }

  mapping (address => MintZcTokenAddingNotionalArgs) public mintZcTokenAddingNotionalCalled;
  mapping (address => BurnZcTokenRemovingNotionalArgs) public burnZcTokenRemovingNotionalCalled;
  mapping (address => TransferFromZcTokenArgs) public transferFromZcTokenCalled;
  mapping (address => TransferFromNotionalArgs) public transferFromNotionalCalled;

  function cTokenAddressReturns(address a) external {
    cTokenAddr = a;
  }

  function cTokenAddress(address u, uint256 m) external returns (address) {
    cTokenAddressCalled[u] = m;
    return cTokenAddr;
  }

  function mintZcTokenAddingNotionalReturns(bool b) external {
    mintZcTokenAddingNotionalReturn = b;
  }

  function mintZcTokenAddingNotional(address u, uint256 m, address o, address s, uint256 a) external returns (bool) {
    MintZcTokenAddingNotionalArgs memory args; 
    args.maturity = m;
    args.owner = o;
    args.sender = s;
    args.amount = a;
    mintZcTokenAddingNotionalCalled[u] = args;

    return mintZcTokenAddingNotionalReturn;
  }

  function burnZcTokenRemovingNotionalReturns(bool b) external {
    burnZcTokenRemovingNotionalReturn = b;
  }

  function burnZcTokenRemovingNotional(address u, uint256 m, address o, address s, uint256 a) external returns (bool) {
    BurnZcTokenRemovingNotionalArgs memory args; 
    args.maturity = m;
    args.owner = o;
    args.sender = s;
    args.amount = a;
    burnZcTokenRemovingNotionalCalled[u] = args;

    return burnZcTokenRemovingNotionalReturn;
  }

  function transferFromZcTokenReturns(bool b) external {
    transferFromZcTokenReturn = b;
  }

  function transferFromZcToken(address u, uint256 m, address o, address s, uint256 a) external returns (bool) {
    TransferFromZcTokenArgs memory args;
    args.maturity = m;
    args.owner = o;
    args.sender = s;
    args.amount = a;
    transferFromZcTokenCalled[u] = args;

    return transferFromZcTokenReturn;
  }

  function transferFromNotionalReturns(bool b) external {
    transferFromNotionalReturn = b;
  }

  function transferFromNotional(address u, uint256 m, address o, address s, uint256 a) external returns (bool) {
    TransferFromNotionalArgs memory args;
    args.maturity = m;
    args.owner = o;
    args.sender = s;
    args.amount = a;
    transferFromNotionalCalled[u] = args;

    return transferFromNotionalReturn;
  }
}
