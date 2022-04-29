// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

interface IErc20 {
	function approve(address, uint256) external returns (bool);
	function transfer(address, uint256) external returns (bool);
	function balanceOf(address) external returns (uint256);
	function transferFrom(address, address, uint256) external returns (bool);
}

contract YieldToken {
    uint32 private maturityReturn;
    IErc20 private baseReturn;
    uint128 private sellBaseReturn;
    uint128 private sellBasePreviewReturn;

    // mapping of arguments sent to sellBase. key is the passed in address.
    mapping (address => uint256) public sellBaseCalled;
    uint128 public sellBasePreviewCalled;

    function base() external view returns (IErc20) {
        return baseReturn;
    }

    function baseReturns(IErc20 i) external {
        baseReturn = i;
    }

    function maturity() external view returns (uint32) {
        return maturityReturn;
    }

    function maturityReturns(uint32 m) external {
        maturityReturn = m;
    } 

    function sellBaseReturns(uint128 r) external {
        sellBaseReturn = r;
    }

    function sellBasePreviewReturns(uint128 r) external {
        sellBasePreviewReturn = r;
    }

    function sellBase(address a, uint128 u) external returns (uint128) {
        sellBaseCalled[a] = u;
        return sellBaseReturn;
    }

    function sellBasePreview(uint128 u) external returns (uint128) {
        sellBasePreviewCalled = u;
        return sellBasePreviewReturn;
    }
}
