// SPDX-License-Identifier: UNLICENSED

/**
  @dev SigTest.sol is written specfically to test the functions which exist in our Sig.sol  "embedded" library
*/

pragma solidity 0.8.0;

import './Sig.sol';

contract SigTest {
  function splitTest(bytes memory sig) public pure returns (uint8 v, bytes32 r, bytes32 s) {
    return Sig.split(sig);
  }

  function recoverTest(bytes32 h, uint8 v, bytes32 r, bytes32 s) public pure returns (address) {
    return Sig.recover(h,v,r,s);
  }
}
