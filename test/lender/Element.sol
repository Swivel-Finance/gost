// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

library Element {
  // TODO audit structure / order
  struct SingleSwap {
    bytes userData
    bytes32 poolId;
    uint256 amount;
    SwapKind kind;
    IAsset assetIn;
    IAsset assetOut;
  }

  struct FundManagement {
    bytes32 key;
    address maker;
    address underlying;
    bool vault;
    bool exit;
    uint256 principal;
    uint256 premium;
    uint256 maturity;
    uint256 expiry;
  }
}
