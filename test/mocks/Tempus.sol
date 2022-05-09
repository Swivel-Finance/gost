// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface Any {}

interface IErc20 {
	function approve(address, uint256) external returns (bool);
	function transfer(address, uint256) external returns (bool);
	function balanceOf(address) external returns (uint256);
	function transferFrom(address, address, uint256) external returns (bool);
}

interface IErc20Metadata is IErc20 {
    function name() external view returns (string memory);
    function symbol() external view returns (string memory);
    function decimals() external view returns (uint8);
}

contract Tempus {
    uint256 private depositAndFixReturn;
    uint256 private maturityTimeReturn;
    IErc20Metadata private yieldBearingTokenReturn;

    Any public tempusAMMCalled;
    Any public tempusPoolCalled;
    uint256 public amountCalled;
    bool public isBackingTokenCalled;
    uint256 public minimumReturnCalled;
    uint256 public deadlineCalled;

    function depositAndFixReturns(uint256 r) external {
        depositAndFixReturn = r;
    }

    function maturityTimeReturns(uint256 m) external {
        maturityTimeReturn = m;
    }

    function yieldBearingTokenReturns(IErc20Metadata t) external {
        yieldBearingTokenReturn = t;
    }

    function maturityTime() external view returns (uint256) {
        return maturityTimeReturn;
    }
    
    function yieldBearingToken() external view returns (IErc20Metadata) {
        return yieldBearingTokenReturn;
    }

    function depositAndFix(Any x, Any p, uint256 a, bool bt, uint256 mr, uint256 d) external returns (uint256) {
        tempusAMMCalled = x;
        tempusPoolCalled = p;
        amountCalled = a;
        isBackingTokenCalled = bt;
        minimumReturnCalled = mr;
        deadlineCalled = d;

        return depositAndFixReturn;
    }
}