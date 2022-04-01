// SPDX-License-Identifier: GPL-3.0-or-later

pragma solidity 0.8.13;

contract Illuminate {
  address public admin;
  // a marketplace contracts has many markets
  // a market is a specific set of 3rd party protocol (token addresses)
  // a market is identified by a market-pair which is underlying:maturity (address:uint256)
  // unknown: volatility of the members of a market (same 3rd party members all the time and forever?)
  address public marketPlace;
  // Illuminite essenttially dictates the interface, Iilluminate, which is
  // * mint
  // * lend
  // * redeem
  // * ...
  // In order to adapt 3rd party protocols to this interface the adapter keeps interfaces for any/all
  // 3rd parties actual method names/signatures. Thus a call to Adapter().mint with some form of an
  // identifier it can delegate to the proper internal method for that intended 3rd party thus freeing
  // illuminate from being tied to those other parties. Illuminate can also possess admin-only methods to
  // change the internal address of the adapter so that contract may be changed without impacting illuminate itself.
  address public adapter;

  /// @param m address of a deployed market contract to use
  /// @param r address of a deployed relayer contract to use
  constructor(address m, address r) {
    admin = msg.sender;
  }

  /// @param u address of an underlying asset
  /// @param m maturity (timestamp) of the market
  /// @param a amount being minted
  function mint(address u, uint256 m, uint256 a) public returns (bool) {
    /*
      1. transfers from an ERC20[specific market protocol address] sender -> this contract for amount passed
      2. mints ZcToken[specific market illuminante ZcToken Address] (zctoken is unique to each market as "illuminate")

      get market from Market contract given the inputs.

      unknown: how to identify the intended market protocol (the 'from' in the transfer).
        * const style uint8s that represent third parties -> SWIVEL = 1, YIELD = 2
          * who makes that known? likely the market itself? it could provide getters for them.
        * other methods to identify which 3rd party
    */

    return true;
  }
