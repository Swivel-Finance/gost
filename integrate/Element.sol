// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Interfaces.sol';

library Element {
    // TODO are these established element names? kind? not type? etc...
    // TODO In, Out vs GIVEN_IN, GIVEN_OUT. If those names are needed they should be GivenIn etc...
    enum SwapKind {
        In,
        Out
    }

    // TODO audit structure / names / order-of-members etc...
    struct SingleSwap {
        bytes userData;
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
