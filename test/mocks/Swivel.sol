// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.16;

/// @dev Swivel is a mock whose bindings are imported by unit tests in any pkg/*testing that needs it
contract Swivel {
  struct MethodArgs {
    address underlying;
    address one; // variable use addresses
    address two;
    uint256 amount;
  }

  mapping (uint8 => MethodArgs) public authRedeemCalled;
  
  bool private authRedeemReturn;

  function authRedeemReturns(bool b) external {
    authRedeemReturn = b;
  }

  function authRedeem(uint8 p, address u, address c, address t, uint256 a) external returns (bool) {
    MethodArgs memory args;
    args.underlying = u;
    args.one = c;
    args.two = t;
    args.amount = a;
    authRedeemCalled[p] = args;
    return authRedeemReturn;
  }
}
