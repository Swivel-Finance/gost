// SPDX-License-Identifier: UNLICENSED

/**
  AaveToken is a mock aave token with stubs of the methods we need for testing.
*/

pragma solidity 0.8.16;

// TODO this could inherit from the ERC20 mock if needed
contract AaveToken {
  address private POOLReturn;
  // NOTE abbreviating UNDERLYING_ASSET_ADDRESS here
  address private UNDERLYINGReturn;

  function POOL() public view returns (address) {
    return POOLReturn;
  }

  function POOLReturns(address a) public {
    POOLReturn = a;
  }

  function UNDERLYING_ASSET_ADDRESS() public view returns (address) {
    return UNDERLYINGReturn;
  }

  function UNDERLYINGReturns(address a) public {
    UNDERLYINGReturn = a;
  }
}
