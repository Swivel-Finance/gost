// SPDX-License-Identifier: UNLICENSED

/**
  @dev HashTest.sol is written specfically to test the functions which exist in our Hash.sol "embedded" library
*/

pragma solidity 0.8.0;

import './Hash.sol';

contract HashTest {
  function domainTest(string memory n, string memory version, uint256 c, address verifier) public pure returns (bytes32) {
    return Hash.domain(n, version, c, verifier);  
  }

  function messageTest(bytes32 d, bytes32 h) public pure returns (bytes32) {
    return Hash.message(d, h);
  }


  function orderTest(bytes32 k, address o, address u, bool f, uint256[6] calldata p) external pure returns (bytes32) {
    return Hash.order(k,o,u,f,p);
  }
}
