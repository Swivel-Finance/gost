// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface ICompoundToken {
  function mint(uint256) external returns(uint256);
  function redeemUnderlying(uint256) external returns(uint256);
}

library Compound {
  /// @param c Address of a deployed ICompoundToken
  /// @param a Amount to deposit / mint
  function deposit(address c, uint256 a) internal returns (uint256) {
    return ICompoundToken(c).mint(a);
  }

  /// @param c Address of a deployed ICompoundToken
  /// @param a Amount to withdraw / redeem
  function withdraw(address c, uint256 a) internal returns (uint256) {
    return ICompoundToken(c).redeemUnderlying(a);
  }
}
