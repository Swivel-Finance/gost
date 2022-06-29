// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Protocols.sol'; // NOTE: if size restrictions become extreme we can use ints (implicit enum)

interface ICompoundToken {
  function exchangeRateCurrent() external view returns(uint256);
  function underlying() external view returns(address);
}

library Compounding {
  /// @param p Protocol Enum value
  /// @param c Compounding token address
  function underlying(uint8 p, address c) internal view returns (address) {
    if (p == uint8(Protocols.Compound)) { // TODO is Rari a drop in here?
      return ICompoundToken(c).underlying();
    } else if (p == uint8(Protocols.Yearn)) {
      // TODO this
      return c;
    } else if (p == uint8(Protocols.Aave)) {
      // TODO this
      return c;
    } else if (p == uint8(Protocols.Euler)) {
      // TODO this
      return c;
    } else {
      // TODO this
      return c;      
    }
  }

  /// @param p Protocol Enum value
  /// @param c Compounding token address
  function exchangeRate(uint8 p, address c) internal view returns (uint256) {
    if (p == uint8(Protocols.Compound)) { // TODO is Rari a drop in here?
      return ICompoundToken(c).exchangeRateCurrent(); // TODO the alternative method to this
    } else if (p == uint8(Protocols.Yearn)) {
      // TODO this
      return 0;
    } else if (p == uint8(Protocols.Aave)) {
      // TODO this
      return 0;
    } else if (p == uint8(Protocols.Euler)) {
      // TODO this
      return 0;
    } else {
      // TODO this
      return 0;      
    }
  }
}
