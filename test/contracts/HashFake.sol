// SPDX-License-Identifier: UNLICENSED

/**
  @dev HashFake.sol is written specfically to test the functions which exist in our Hash.sol "embedded" library
*/

pragma solidity 0.8.0;

import './Hash.sol';

contract HashFake {
  function domainTest(string memory n, string memory version, uint256 c, address verifier) public pure returns (bytes32) {
    return Hash.domain(n, version, c, verifier);  
  }

  function messageTest(bytes32 d, bytes32 h) public pure returns (bytes32) {
    return Hash.message(d, h);
  }

  function orderTest(Hash.Order calldata o) external pure returns (bytes32) {
    return Hash.order(o);
  }
}
