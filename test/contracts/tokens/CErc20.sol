// SPDX-License-Identifier: UNLICENSED

/**
  CErc20 is a mock compound token with stubs of the methods we need for testing.
*/

pragma solidity 0.8.0;

// TODO this could inherit from the ERC20 mock if needed
contract CErc20 {
  uint256 private currentExchangeRate;
  /// @dev allows us to dictate return from mint().
  uint256 private mintReturn;
  /// @dev the last amount mint was called with
  uint256 public mintedArgs;

  function exchangeRateCurrentReturns(uint256 n) public {
    currentExchangeRate = n;
  }
  
  function exchangeRateCurrent() public view returns (uint256) {
    return currentExchangeRate;
  }

  function mintReturns(uint256 n) public {
    mintReturn = n;
  }

  function mint(uint256 n) public returns (uint256) {
    mintedArgs = n;
    return mintReturn;
  }
}
