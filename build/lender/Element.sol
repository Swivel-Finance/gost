// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Interfaces.sol';

library Element {

    enum SwapKind {
        In,
        Out
    }

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
