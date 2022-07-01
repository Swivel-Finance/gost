// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Protocols.sol'; // NOTE: if size restrictions become extreme we can use ints (implicit enum)

interface IErc4626 {
  /// @dev Converts the given 'assets' (uint256) to 'shares', returning that amount
  function convertToAssets(uint256) external view returns (uint256);
  /// @dev The address of the underlying ERC20 token
  function asset() external view returns (address);
}

interface ICompoundToken {
  function exchangeRateCurrent() external view returns(uint256);
  /// @dev The address of the underlying ERC20 token
  function underlying() external view returns(address);
}

interface IYearnVault {
  function pricePerShare() external view returns (uint256);
  /// @dev The address of the underlying ERC20 token
  function underlying() external view returns(address);
}

library Compounding {
  /// @param p Protocol Enum value
  /// @param c Compounding token address
  function underlying(uint8 p, address c) internal view returns (address) {
    if (p == uint8(Protocols.Compound)) { // TODO is Rari a drop in here?
      return ICompoundToken(c).underlying();
    } else if (p == uint8(Protocols.Yearn)) {
      return IYearnVault(c).underlying();
    } else if (p == uint8(Protocols.Aave)) {
      // TODO this
      return c;
    } else if (p == uint8(Protocols.Euler)) {
      // TODO this
      return c;
    } else {
      return IErc4626(c).asset();      
    }
  }

  /// @param p Protocol Enum value
  /// @param c Compounding token address
  function exchangeRate(uint8 p, address c) internal view returns (uint256) {
    if (p == uint8(Protocols.Compound)) { // TODO is Rari a drop in here?
      return ICompoundToken(c).exchangeRateCurrent(); // TODO the alternative method to this
    } else if (p == uint8(Protocols.Yearn)) {
      return IYearnVault(c).pricePerShare();
    } else if (p == uint8(Protocols.Aave)) {
      // TODO this
      return 0;
    } else if (p == uint8(Protocols.Euler)) {
      // TODO this
      return 0;
    } else {
      // NOTE: the 1e26 const is a degree of precision to enforce on the return
      return IErc4626(c).convertToAssets(1e26);
    }
  }
}
