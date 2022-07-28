// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.13;

/// @dev Creator is a mock whose bindings are imported by unit tests in any pkg/*testing that needs it
contract Creator {
  struct CreateArgs {
    address underlying;
    uint256 maturity;
    address cToken; // is sender or maker depending on method
    address swivel; // same as above
    string name;
    string symbol;
    uint8 decimals;
  }

  mapping (uint8 => CreateArgs) public createCalled;

  address private zcTokenReturn;
  address private vaultTrackerReturn;

  /// @dev allow us to dictate the return of create()
  function createReturns(address z, address v) external {
    zcTokenReturn = z;
    vaultTrackerReturn = v;
  }

  function create(uint8 p, address u, uint256 m, address c, address sw, string memory n, string memory s, uint8 d) external returns (address, address) {
    CreateArgs memory args;
    args.underlying = u;
    args.maturity = m;
    args.cToken = c;
    args.swivel = sw;
    args.name = n;
    args.symbol = s;
    args.decimals = d;

    createCalled[p] = args;
    return (zcTokenReturn, vaultTrackerReturn);
  }
}
