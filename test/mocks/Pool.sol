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

    struct BuyPrincipalTokenArgs {
        uint128 PrincipalTokenOut;
        uint128 min;
    }

    IErc5095 private PrincipalTokenReturn;
    address private underlyingReturn;
    uint128 private sellUnderlyingReturn;
    uint128 private buyUnderlyingReturn;
    uint128 private sellPrincipalTokenReturn;
    uint128 private buyPrincipalTokenReturn;
    uint128 private sellUnderlyingPreviewReturn;
    uint128 private buyUnderlyingPreviewReturn;
    uint128 private sellPrincipalTokenPreviewReturn;
    uint128 private buyPrincipalTokenPreviewReturn;

    mapping(address => uint128) public sellUnderlyingCalled;
    mapping(address => BuyUnderlyingArgs) public buyUnderlyingCalled;
    mapping(address => uint128) public sellPrincipalTokenCalled;
    mapping(address => BuyPrincipalTokenArgs) public buyPrincipalTokenCalled;
    uint128 public sellUnderlyingPreviewCalled;
    uint128 public buyUnderlyingPreviewCalled;
    uint128 public sellPrincipalTokenPreviewCalled;
    uint128 public buyPrincipalTokenPreviewCalled;

    function principalTokenReturns(address p) public {
        PrincipalTokenReturn = IErc5095(p);
    }

    function principalToken() external view returns (IErc5095) {
        return PrincipalTokenReturn;
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

    function sellPrincipalTokenReturns(uint128 s) public {
        sellPrincipalTokenReturn = s;
    }

    function sellPrincipalToken(address t, uint128 m) external returns (uint128) {
        sellPrincipalTokenCalled[t] = m;
        return sellPrincipalTokenReturn;
    }

    function buyPrincipalTokenReturns(uint128 b) public {
        buyPrincipalTokenReturn = b;
    }

    function buyPrincipalToken(
        address t,
        uint128 p,
        uint128 m
    ) external returns (uint128) {
        buyPrincipalTokenCalled[t] = BuyPrincipalTokenArgs(p, m);
        return buyPrincipalTokenReturn;
    }

    function sellUnderlyingPreviewReturns(uint128 s) public {
        sellUnderlyingPreviewReturn = s;
    }

    function sellUnderlyingPreview(uint128 u) external returns (uint128) {
        sellUnderlyingPreviewCalled = u;
        return sellUnderlyingPreviewReturn;
    }

    function buyUnderlyingPreviewReturns(uint128 p) public {
        buyUnderlyingPreviewReturn = p;
    }

    function buyUnderlyingPreview(uint128 u) external returns (uint128) {
        buyUnderlyingPreviewCalled = u;
        return buyUnderlyingPreviewReturn;
    }

    function sellPrincipalTokenPreviewReturns(uint128 i) public {
        sellPrincipalTokenPreviewReturn = i;
    }

    function sellPrincipalTokenPreview(uint128 i) external returns (uint128) {
        sellPrincipalTokenPreviewCalled = i;
        return sellPrincipalTokenPreviewReturn;
    }

    function buyPrincipalTokenPreviewReturns(uint128 o) public {
        buyPrincipalTokenPreviewReturn = o;
    }

    function buyPrincipalTokenPreview(uint128 o) external returns (uint128) {
        buyPrincipalTokenPreviewCalled = o;
        return buyPrincipalTokenPreviewReturn;
    }
}
