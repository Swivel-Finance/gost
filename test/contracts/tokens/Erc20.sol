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

  /// @dev mapping of arguments sent to transferFrom. key is passed from address.
  mapping (address => TransferFromArgs) public transferFrom;

  /// @dev a boolean flag which allows us to dictate the return. 
  bool private transferFromReturn;

  constructor() {
    transferFromReturnTrue();
  }

  function transferFromReturnTrue() public {
    transferFromReturn = true;
  }

  function transferFromReturnFalse() public {
    transferFromReturn = false;
  }

  function transferFrom(address f, address t, uint256 a) public returns (bool) {
    TransferFromArgs memory args;
    args.to = t;
    args.amount = a;
    transferFrom[f] = args;
    return transferFromReturn;
  }
}
