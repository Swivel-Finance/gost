// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.0;

contract Erc20 {
  // interation with these should be thru the methods, why should they be public?
  mapping (address => mapping (address => uint256)) allowances; // vs allowed
  mapping (address => uint256) balances;
  uint256 supply; // vs totalSupply

  // vanities. TODO keep or remove?
  uint8 public decimals;
  string public symbol;

  event Transfer(address indexed f, address indexed t, uint256 a);


  constructor(uint8 d, string memory s) {
    decimals = d;
    symbol = s;
    // this will be a simple "fixed" initial balance token that mints 1 ETH for the maker...
    mint(msg.sender, 1 ether);
  }

  function balanceOf(address a) public view returns (uint256) {
    return balances[a];
  }

  // TODO could be virtual a la OZ
  function mint(address r, uint256 a) internal {
    require(r != address(0), 'cannot mint for zero address');
    supply += a;
    balances[r] += a;
    emit Transfer(address(0), r, a);
  }

  function totalSupply() public view returns (uint256) {
    return supply;
  }

  function transfer(address t, uint256 a) public returns (bool) {
    require(msg.sender != address(0), 'cannot transfer from zero address');
    require(t != address(0), 'cannot transfer to zero address');

    balances[msg.sender] -= a; // this will revert if underflow. TODO could put a require in order to have a message
    balances[t] += a;
    emit Transfer(msg.sender, t, a);
    return true;
  }

  
}
