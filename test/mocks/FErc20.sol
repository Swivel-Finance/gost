// SPDX-License-Identifier: UNLICENSED

/**
  FErc20 is a mock rari token with stubs of the methods we need for testing.
*/

pragma solidity 0.8.13;

import './Erc20.sol';
import './CErc20.sol';

// TODO this could inherit from the ERC20 mock if needed
contract FErc20 is Erc20, CErc20 {
  mapping (address => uint256) public allocateToCalled;
  /// @dev allows us to dictate the return from allocateTo calls
  bool private allocateToReturn;

  function allocateTo(address o, uint256 a) public returns (bool) {
    allocateToCalled[o] = a;
    return allocateToReturn;
  }

  function allocateToReturns(bool b) public {
    allocateToReturn = b;
  }
}
