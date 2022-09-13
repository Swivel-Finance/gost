// SPDX-License-Identifier: AGPL-3.0

pragma solidity ^0.8.13;

import './Protocols.sol';
import './LibCompound.sol';

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
  function token() external view returns(address);
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

interface ILidoToken {
  /// @dev The address of the stETH underlying asset
  function stETH() external view returns (address);
  /// @notice Returns amount of stETH for one wstETH
  function stEthPerToken() external view returns (uint256);
}

library Compounding {
  /// @param p Protocol Enum value
  /// @param c Compounding token address
  function underlying(uint8 p, address c) internal view returns (address) {
    if (p == uint8(Protocols.Compound) || p == uint8(Protocols.Rari)) {
      return ICompoundToken(c).underlying();
    } else if (p == uint8(Protocols.Yearn)) {
      return IYearnVault(c).token();
    } else if (p == uint8(Protocols.Aave)) {
      return IAaveToken(c).UNDERLYING_ASSET_ADDRESS();
    } else if (p == uint8(Protocols.Euler)) {
      return IEulerToken(c).underlyingAsset();
    } else if (p == uint8(Protocols.Lido)) {
      return ILidoToken(c).stETH();
    } else {
      return IErc4626(c).asset();      
    }
  }

  /// @param p Protocol Enum value
  /// @param c Compounding token address
  function exchangeRate(uint8 p, address c) internal view returns (uint256) {
    // in contrast to the below, LibCompound provides a lower gas alternative to exchangeRateCurrent()
    if (p == uint8(Protocols.Compound)) {
      return LibCompound.viewExchangeRate(ICERC20(c));
      // with the removal of LibFuse we will direct Rari to the exposed Compound CToken methodology
    } else if (p == uint8(Protocols.Rari)) { 
      return ICompoundToken(c).exchangeRateCurrent();
    } else if (p == uint8(Protocols.Yearn)) {
      return IYearnVault(c).pricePerShare();
    } else if (p == uint8(Protocols.Aave)) {
      IAaveToken aToken = IAaveToken(c);
      return IAavePool(aToken.POOL()).getReserveNormalizedIncome(aToken.UNDERLYING_ASSET_ADDRESS());
    } else if (p == uint8(Protocols.Euler)) {
      // NOTE: the 1e26 const is a degree of precision to enforce on the return
      return IEulerToken(c).convertBalanceToUnderlying(1e26);
    } else if (p == uint8(Protocols.Lido)) {
      return ILidoToken(c).stEthPerToken();
    } else {
      // NOTE: the 1e26 const is a degree of precision to enforce on the return
      return IErc4626(c).convertToAssets(1e26);
    }
  }
}
