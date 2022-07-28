// SPDX-License-Identifier: UNLICENSED

/**
  AavePool is a mock with stubs of the methods we need for testing.
*/

pragma solidity ^0.8.13;

contract AavePool {
  struct TransactArgs {
    uint256 amount;
    address onBehalfOf;
    address to;
    uint16 code;
  }

  /// @dev allows us to dictate the return from getReserveNormalizedIncome()
  uint256 private getReserveNormalizedIncomeReturn;
  /// @dev the last args deposit was called with. NOTE that Aave deposit is void
  mapping (address => TransactArgs) public depositCalled;
  /// @dev allows us to dictate return from withdraw().
  uint256 private withdrawReturn;
  /// @dev the last args withdraw was called with
  mapping (address => TransactArgs) public withdrawCalled;

  // NOTE: using the .sol feature of omitting argument name to prevent compiler nagging
  function getReserveNormalizedIncome(address) public view returns (uint256) {
    return getReserveNormalizedIncomeReturn;
  }

  function getReserveNormalizedIncomeReturns(uint256 n) public {
    getReserveNormalizedIncomeReturn = n;
  }

  function deposit(address u, uint256 a, address o, uint16 c) public {
    TransactArgs memory args;
    args.amount = a;
    args.onBehalfOf = o;
    args.code = c;
    depositCalled[u] = args;
  }

  function withdraw(address u, uint256 a, address t) public returns (uint256) {
    TransactArgs memory args;
    args.amount = a;
    args.to = t;
    withdrawCalled[u] = args;
    return withdrawReturn;
  }

  function withdrawReturns(uint256 a) public {
    withdrawReturn = a;
  }
}
