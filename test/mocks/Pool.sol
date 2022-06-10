// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface IErc5095 {
    event Redeem(address indexed from, address indexed to, uint256 amount);
    event Mature(uint256 timestamp);

    function maturity() external returns (uint256);

    function underlying() external returns (address);

    function maturityRate() external returns (uint256);

    function convertToUnderlying(uint256 principalAmount) external returns (uint256);

    function convertToPrincipal(uint256 underlyingAmount) external returns (uint256);

    function maxRedeem(address owner) external returns (uint256);

    function previewRedeem(uint256 principalAmount) external returns (uint256);

    function maxWithdraw(address owner) external returns (uint256);

    function previewWithdraw(uint256 principalAmount) external returns (uint256);

    function withdraw(
        uint256 underlyingAmount,
        address owner,
        address receiver
    ) external returns (uint256);

    function redeem(
        uint256 principalAmount,
        address owner,
        address receiver
    ) external returns (uint256);
}

contract Pool {
    struct BuyUnderlyingArgs {
        uint128 amount;
        uint128 min;
    }

    struct BuyPTArgs {
        uint128 ptOut;
        uint128 min;
    }

    IErc5095 private PTReturn;
    address private underlyingReturn;
    uint128 private sellUnderlyingReturn;
    uint128 private buyUnderlyingReturn;
    uint128 private sellPTReturn;
    uint128 private buyPTReturn;
    uint128 private sellUnderlyingPreviewReturn;
    uint128 private buyUnderlyingPreviewReturn;
    uint128 private sellPTPreviewReturn;
    uint128 private buyPTPreviewReturn;

    mapping(address => uint128) public sellUnderlyingCalled;
    mapping(address => BuyUnderlyingArgs) public buyUnderlyingCalled;
    mapping(address => uint128) public sellPTCalled;
    mapping(address => BuyPTArgs) public buyPTCalled;
    uint128 public sellUnderlyingPreviewCalled;
    uint128 public buyUnderlyingPreviewCalled;
    uint128 public sellPTPreviewCalled;
    uint128 public buyPTPreviewCalled;

    function PTReturns(address p) public {
        PTReturn = IErc5095(p);
    }

    function PT() external view returns (IErc5095) {
        return PTReturn;
    }

    function sellUnderlyingReturns(uint128 s) public {
        sellUnderlyingReturn = s;
    }

    function underlyingReturns(address u) public {
        underlyingReturn = u;
    }

    function underlying() public view returns (address) {
        return underlyingReturn;
    }

    function sellUnderlying(address t, uint128 m) external returns (uint128) {
        sellUnderlyingCalled[t] = m;
        return sellUnderlyingReturn;
    }

    function buyUnderlyingReturns(uint128 b) public {
        buyUnderlyingReturn = b;
    }

    function buyUnderlying(
        address t,
        uint128 u,
        uint128 m
    ) external returns (uint128) {
        buyUnderlyingCalled[t] = BuyUnderlyingArgs(u, m);
        return buyUnderlyingReturn;
    }

    function sellPTReturns(uint128 s) public {
        sellPTReturn = s;
    }

    function sellPT(address t, uint128 m) external returns (uint128) {
        sellPTCalled[t] = m;
        return sellPTReturn;
    }

    function buyPTReturns(uint128 b) public {
        buyPTReturn = b;
    }

    function buyPT(
        address t,
        uint128 p,
        uint128 m
    ) external returns (uint128) {
        buyPTCalled[t] = BuyPTArgs(p, m);
        return buyPTReturn;
    }

    function sellUnderlyingPreview(uint128 u) external returns (uint128) {
        sellUnderlyingPreviewCalled = u;
        return sellUnderlyingPreviewCalled;
    }

    function buyUnderlyingPreview(uint128 u) external returns (uint128) {
        buyUnderlyingPreviewCalled = u;
        return buyUnderlyingPreviewCalled;
    }

    function sellPTPReviewReturns(uint128 i) public {
        sellPTPreviewReturn = i;
    }

    function sellPTPreview(uint128 i) external returns (uint128) {
        sellPTPreviewCalled = i;
        return sellPTPreviewCalled;
    }

    function buyPTPReviewReturns(uint128 o) public {
        buyPTPreviewReturn = o;
    }

    function buyPTPreview(uint128 o) external returns (uint128) {
        buyPTPreviewCalled = o;
        return buyPTPreviewCalled;
    }
}
