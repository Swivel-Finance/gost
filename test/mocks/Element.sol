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
    struct SwapArgs {
        address recipient;
        uint256 swapAmount;
        uint256 limit;
        uint256 deadline;
    }

    uint256 private swapReturn;

    mapping (address => SwapArgs) public swapCalled;
    mapping (address => uint256) public withdrawPrincipalCalled;

    function swapReturns(uint256 s) external {
        swapReturn = s;
    }

    function swap(SingleSwap memory s, FundManagement memory f, uint256 l, uint256 d) external returns (uint256) {
        swapCalled[f.sender] = SwapArgs(f.recipient, s.amount, l, d);
        return swapReturn;
    }

    function withdrawPrincipal(uint256 a, address d) external {
        withdrawPrincipalCalled[d] = a;
    }
}