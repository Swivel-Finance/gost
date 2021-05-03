// SPDX-License-Identifier: UNLICENSED

/**
  Erc20 is a mock which records arguments passed to its methods as well as
  provides setters allowing us to dictate method return values
*/

pragma solidity 0.8.0;

contract Erc20 {
  /// @dev a struct to hold the arguments passed to transferFrom
  struct TransferFromArgs {
    address to;
    uint256 amount;
  }

  /// @dev mapping for arguments passed to approve
  mapping (address => uint256) public approvedArgs;
  /// @dev a boolean flag which allows us to dictate the return of approve(). 
  bool private approveReturn;
  /// @dev mapping of arguments sent to transfer. key is the passed in address.
  mapping (address => uint256) public transferredArgs;
  bool private transferReturn;
  /// @dev mapping of arguments sent to transferFrom. key is passed from address.
  mapping (address => TransferFromArgs) public transferredFromArgs;
  bool private transferFromReturn;

  function approveReturns(bool b) public {
    approveReturn = b;
  }

  function approve(address s, uint256 a) public returns (bool) {
    approvedArgs[s] = a;  
    return approveReturn;
  }

  function transferReturns(bool b) public {
    transferReturn = b;
  }

  function transfer(address t, uint256 a) public returns (bool) {
    transferredArgs[t] = a;
    return transferReturn;
  }

  function transferFromReturns(bool b) public {
    transferFromReturn = b;
  }

  function transferFrom(address f, address t, uint256 a) public returns (bool) {
    TransferFromArgs memory args;
    args.to = t;
    args.amount = a;
    transferredFromArgs[f] = args;
    return transferFromReturn;
  }
}
