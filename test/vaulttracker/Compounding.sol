// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Protocols.sol'; // NOTE: if size restrictions become extreme we can use ints (implicit enum)

interface IErc4626 {
  /// @dev Converts the given 'assets' (uint256) to 'shares', returning that amount
  function convertToAssets(uint256) external view returns (uint256);
}

interface ICompoundToken {
  // TODO comment
  function exchangeRateCurrent() external view returns(uint256);
}

interface IYearnVault {
  // TODO comment
  function pricePerShare() external view returns (uint256);
}

interface IAavePool {
   /// @dev Returns the normalized income of the reserve given the address of the underlying asset of the reserve
  function getReserveNormalizedIncome(address) external view returns (uint256);
}

interface IAaveToken {
  // TODO comments
  function POOL() external view returns (address);
  function UNDERLYING_ASSET_ADDRESS() external view returns (address);
}

library Compounding {
  /// @param p Protocol Enum value
  /// @param c Compounding token address
  function exchangeRate(uint8 p, address c) internal view returns (uint256) {
    if (p == uint8(Protocols.Compound)) { // TODO is Rari a drop in here?
      return ICompoundToken(c).exchangeRateCurrent(); // TODO the alternative method to this
    } else if (p == uint8(Protocols.Yearn)) {
      return IYearnVault(c).pricePerShare();
    } else if (p == uint8(Protocols.Aave)) {
      IAaveToken aToken = IAaveToken(c);
      return IAavePool(aToken.POOL()).getReserveNormalizedIncome(aToken.UNDERLYING_ASSET_ADDRESS());
    } else if (p == uint8(Protocols.Euler)) {
      // TODO this
      return 0;
    } else {
      return IErc4626(c).convertToAssets(1e26);
    }
  }
}
