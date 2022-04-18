// SPDX-License-Identifier: GPL-3.0-or-later

pragma solidity 0.8.13;

contract Illuminate {
  address public admin;
  // a marketplace contracts has many markets
  // a market is a specific set of Principals (token addresses)
  // a market is identified by a market-pair which is underlying:maturity (address:uint256)
  address public marketPlace;
  // Illuminite essenttially dictates the interface, Iilluminate, which is
  // * mint
  // * lend
  // * redeem
  // * ...
  // in order to adapt Principals to this interface the lender keeps interfaces for any/all lending
  address public lender;
  // and the same goes for redeeming
  address public redeemer;

  event Mint(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount);

  /// @param m address of a deployed market contract to use
  /// @param r address of a deployed relayer contract to use
  constructor(address m, address r) {
    admin = msg.sender;
  }

  /// @dev mint is uniform across all principals, thus there is no need for a 'minter'
  /// @param p value of a specific principal according to the MarketPlace Principals Enum. NOTE (if using a local Principals this type can change)
  /// @param u address of an underlying asset
  /// @param m maturity (timestamp) of the market
  /// @param a amount being minted
  function mint(uint8 p, address u, uint256 m, uint256 a) public returns (bool) {
    mPlace = MarketPlace(marketPlace);
    // use market interface to fetch the market for the given market pair
    address[8] market = mPlace.markets(u, m);
    // use safe transfer lib and ERC interface...
    Safe.transferFrom(Erc20(market[p]), msg.sender, address(this), a);
    // use zctoken interface...
    ZcToken(market[mPlace.Principals.Illuminate]).mint(msg.sender, a);

    emit Mint(p, u, m, a);

    return true;
  }
}
