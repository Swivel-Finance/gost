// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

library Element {
  // a blank interface for the assets, TODO this may be useful elsewhere (move if so...)
  interface Any {}

  // TODO are these established element names? kind? not type? etc...
  enum SwapKind { In, Out } // TODO In, Out vs GIVEN_IN, GIVEN_OUT. If those names are needed they should be GivenIn etc...

  // TODO audit structure / names / order-of-members etc...
  struct SingleSwap {
    bytes userData
    bytes32 poolId;
    uint256 amount;
    SwapKind kind;
    Any assetIn;
    Any assetOut;
  }

  struct FundManagement {
    address sender;
    address payable recipient;
    bool fromInternalBalance;
    bool toInternalBalance;
  }
}
