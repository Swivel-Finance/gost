// SPDX-License-Identifier: UNLICENSED

/**
  Underlying is a (for now) ERC20 compatible token interface

  NOTE: with the addition of the ZCToken, we may just be able to use its bindings vs maintaining this... TODO
*/

pragma solidity 0.8.0;

contract Underlying {
  mapping (address => uint256) public balances;

  // For v1 we only need to read the balance of underlying
  function balanceOf(address o) public view returns (uint256) {
    return balances[o]; // this should not matter to the abi. i think...
  }
}
