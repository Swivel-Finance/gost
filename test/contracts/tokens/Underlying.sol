// SPDX-License-Identifier: UNLICENSED

/**
  Underlying is a (for now) ERC20 compatible token interface
*/

pragma solidity 0.8.0;

contract Erc20 {
  mapping (address => uint256) public balances;

  // For v1 we only need to read the balance of underlying
  function balanceOf(address o) public view returns (uint256) {
    return balances[o]; // this should not matter to the abi. i think...
  }
}
