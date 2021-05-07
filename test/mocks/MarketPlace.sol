// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.0;

/// @dev MarketPlace is a mock whose bindings are imported by unit tests in any pkg/*testing that needs it
contract MarketPlace {
  // the address that will be returned when marketCTokenAddress is called
  address private cTokenAddr;
  // the arguments that were passed to marketCTokenAddress when it was called
  // TODO: we should likely standardize on using the `Called` suffix for these mappings in the mocks.
  //       the token mocks use a differnt pattern, change them to match this...
  mapping (address => uint256) public marketCTokenAddressCalled;
  // the bool that will be returned from mintZcTokenAddingNotional
  bool private mintZcTokenAddingNotionalReturn;

  struct MintZcTokenAddingNotionalArgs {
    uint256 maturity;
    uint256 amount;
    address owner;
  }

  mapping (address => MintZcTokenAddingNotionalArgs) public mintZcTokenAddingNotionalCalled;

  function marketCTokenAddressReturns(address a) external {
    cTokenAddr = a;
  }

  function marketCTokenAddress(address u, uint256 m) external returns (address) {
    marketCTokenAddressCalled[u] = m;
    return cTokenAddr;
  }

  function mintZcTokenAddingNotionalReturns(bool b) external {
    mintZcTokenAddingNotionalReturn = b;
  }

  function mintZcTokenAddingNotional(address u, uint256 m, uint256 a, address o) external returns (bool) {
    MintZcTokenAddingNotionalArgs memory args; 
    args.maturity = m;
    args.amount = a;
    args.owner = o;
    mintZcTokenAddingNotionalCalled[u] = args;

    return mintZcTokenAddingNotionalReturn;
  }
}
