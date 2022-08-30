// SPDX-License-Identifier: AGPL-3.0

/**
  @dev SigFake.sol is written specfically to test the functions which exist in our Sig.sol "embedded" library
*/

pragma solidity 0.8.16;

import './Sig.sol';

contract SigFake {
  function splitTest(bytes memory sig) public pure returns (uint8 v, bytes32 r, bytes32 s) {
    return Sig.split(sig);
  }

  function recoverTest(bytes32 h, Sig.Components calldata c) public pure returns (address) {
    return Sig.recover(h,c);
  }
}
