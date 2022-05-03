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
    uint256 amountCalled;
    uint256 returnCalled;
    address fundManagementSenderCalled;
    uint256 singleSwapAmountCalled;

    function amount() external view returns (uint256) {
        return amountCalled;
    }

    function return_() external view returns (uint256) {
        return returnCalled;
    }

    function fundManagementSender() external view returns (address) {
        return fundManagementSenderCalled;
    }
    
    function singleSwapAmount() external view returns (uint256) {
        return singleSwapAmountCalled;
    }
    
    function swap(SingleSwap memory s, FundManagement memory f, uint256 a, uint256 r) external returns (uint256) {
        amountCalled = a;
        returnCalled = r;
        fundManagementSenderCalled = f.sender;
        singleSwapAmountCalled = s.amount;
        return 0;
    }
}