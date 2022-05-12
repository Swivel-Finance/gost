// SPDX-License-Identifier: UNLICENSED

/**
  ZcToken is a mock which records arguments passed to its methods as well as
  provides setters allowing us to dictate method return values
*/

pragma solidity 0.8.13;

contract ZcToken {
  // a struct to hold the arguments passed to transferFrom
  // struct TransferFromArgs {
    // address to;
    // uint256 amount;
  // }

  // mapping for arguments passed to approve
  mapping (address => uint256) public approveCalled;
  // mapping of arguments sent to burn. key is the passed in address.
  // mapping (address => uint256) public burnCalled;
  // mapping of arguments sent to mint. key is the passed in address.
  // mapping (address => uint256) public mintCalled;
  // mapping of arguments sent to transferFrom. key is passed from address.
  // mapping (address => TransferFromArgs) public transferFromCalled;

  address private underlyingReturn;
  uint256 private maturityReturn;
  bool private approveReturn;
  string public name;
  string public symbol;
  uint8 public decimals;
  // a boolean flag which allows us to dictate the return of burn().
  // bool private burnReturn;
  // a boolean flag which allows us to dictate the return of mint().
  // bool private mintReturn;
  // a boolean flag which allows us to dictate the return of transferFrom().
  // bool private transferFromReturn;

  /// @param u Underlying
  /// @param m Maturity
  /// @param n Name
  /// @param s Symbol
  /// @param d Decimals
  constructor(address u, uint256 m, string memory n, string memory s, uint8 d) {
    // we can set the privates in the constructor as well...
    underlyingReturn = u;
    maturityReturn = m;

    name = n;
    symbol = s;
    decimals = d;

    // set approve to return true by default so that createMarket won't revert
    approveReturn = true;
  }

  function approve(address s, uint256 a) public returns (bool) {
    approveCalled[s] = a;
    return approveReturn;
  }

  function approveReturns(bool b) public {
    approveReturn = b;
  }
}
