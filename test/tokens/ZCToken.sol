// SPDX-License-Identifier: MIT

// NOTE: whomever decided methods which are the implementation of a stated interface
//       need to be labeled 'override' should to be publicly flogged.

pragma solidity 0.8.0;

import "./ERC20.sol";
import "./IZCToken.sol";

contract ZCToken is ERC20, IZCToken {
  address public admin;
  address public underlying;
  uint256 public maturity;

  /// @param u Underlying
  /// @param m Maturity
  /// @param n Name
  /// @param s Symbol
  constructor(address u, uint256 m, string memory n, string memory s) ERC20(n, s) {
      underlying = u;
      maturity = m;
      admin = msg.sender;
  }
  
  /// @param f From
  /// @param a Amount
  function burn(address f, uint256 a) external restricted(admin) override returns(bool) {
      _burn(f, a);
      return true;
  }

  /// @param t To
  /// @param a Amount
  function mint(address t, uint256 a) external restricted(admin) override returns(bool) {
      _mint(t, a);
      return true;
  }

  /// @param a Admin address
  modifier restricted(address a) {
    require(msg.sender == a, 'sender must be admin');
    _;
  }
}
