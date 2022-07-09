// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.13;

enum Protocols {
  Erc4626,
  Compound,
  Rari,
  Yearn,
  Aave,
  Euler
}

interface IErc4626 {
  /// @dev Converts the given 'assets' (uint256) to 'shares', returning that amount
  function convertToAssets(uint256) external view returns (uint256);
  /// @dev The address of the underlying asset
  function asset() external view returns (address);
}

interface ICompoundToken {
  function exchangeRateCurrent() external view returns(uint256);
  /// @dev The address of the underlying asset
  function underlying() external view returns(address);
}

interface IYearnVault {
  function pricePerShare() external view returns (uint256);
  /// @dev The address of the underlying asset
  function underlying() external view returns(address);
}

interface IAavePool {
   /// @dev Returns the normalized income of the reserve given the address of the underlying asset of the reserve
  function getReserveNormalizedIncome(address) external view returns (uint256);
}

interface IAaveToken {
  // @dev Deployed ddress of the associated Aave Pool
  function POOL() external view returns (address);
  /// @dev The address of the underlying asset
  function UNDERLYING_ASSET_ADDRESS() external view returns (address);
}

interface IEulerToken {
  /// @notice Convert an eToken balance to an underlying amount, taking into account current exchange rate
  function convertBalanceToUnderlying(uint256) external view returns(uint256);
  /// @dev The address of the underlying asset
  function underlyingAsset() external view returns(address);
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
      return IAaveToken(c).UNDERLYING_ASSET_ADDRESS();
    } else if (p == uint8(Protocols.Euler)) {
      return IEulerToken(c).underlyingAsset();
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
      IAaveToken aToken = IAaveToken(c);
      return IAavePool(aToken.POOL()).getReserveNormalizedIncome(aToken.UNDERLYING_ASSET_ADDRESS());
    } else if (p == uint8(Protocols.Euler)) {
      // NOTE: the 1e26 const is a degree of precision to enforce on the return
      return IEulerToken(c).convertBalanceToUnderlying(1e26);
    } else {
      // NOTE: the 1e26 const is a degree of precision to enforce on the return
      return IErc4626(c).convertToAssets(1e26);
    }
  }
}