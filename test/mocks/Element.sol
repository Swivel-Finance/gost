// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

enum SwapKind { In, Out }

interface Any {}

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

contract Element {
    struct Swap {
        FundManagement fundManagement;
        SingleSwap swap;
        uint256 _return;
        uint256 deadline;
    }

    uint256 private swapReturn;
    uint256 private deadlineReturn;
    uint256 private return_Return;
    address private fundManagementSenderReturn;
    uint256 private singleSwapAmountReturn;

    mapping (address => Swap) public swapCalled;

    function deadline() external view returns (uint256) {
        return deadlineReturn;
    }

    function deadlineReturns(uint256 d) external {
        deadlineReturn = d;
    }

    function return_() external view returns (uint256) {
        return return_Return;
    }

    function returnReturns(uint256 r) external {
        return_Return = r;
    }

    function fundManagementSender() external view returns (address) {
        return fundManagementSenderReturn;
    }

    function fundManagementSenderReturns(address s) external {
        fundManagementSenderReturn = s;
    }
    
    function singleSwapAmount() external view returns (uint256) {
        return singleSwapAmountReturn;
    }

    function singleSwapAmountReturns(uint256 a) external {
        singleSwapAmountReturn = a;
    }

    function swap(SingleSwap memory s, FundManagement memory f, uint256 r, uint256 d) external returns (uint256) {
        swapCalled[f.sender] = Swap(f, s, r, d);
        return swapReturn;
    }
}