// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;


interface IErc20 {
	function approve(address, uint256) external returns (bool);
	function transfer(address, uint256) external returns (bool);
	function balanceOf(address) external returns (uint256);
	function transferFrom(address, address, uint256) external returns (bool);
}

contract Notional {
    struct TransferFromArgs {
        address to;
        uint256 amount;
    }

    IErc20 private getUnderlyingTokenReturn;
    uint256 private getMaturityReturn;
    bool private transferFromReturn;
    uint256 private balanceOfReturn;
    uint256 private depositReturn;

    mapping (address => TransferFromArgs) public transferFromCalled;
    mapping (address => uint256) public depositCalled;
    address public balanceOfCalled;

    function getUnderlyingTokenReturns(address a) external {
        getUnderlyingTokenReturn = IErc20(a);
    }

    function getMaturityReturns(uint256 m) external {
        getMaturityReturn = m;
    }

    function getUnderlyingToken() external view returns (IErc20, int256) {
        return (IErc20(getUnderlyingTokenReturn), 0);
    }

    function getMaturity() external view returns (uint256) {
        return getMaturityReturn;
    }

    function balanceOfReturns(uint256 b) public {
        balanceOfReturn = b;
    }

    function balanceOf(address b) public returns (uint256) {
        balanceOfCalled = b;
        return balanceOfReturn;
    }

    function transferFromReturns(bool b) public {
        transferFromReturn = b;
    }

    function transferFrom(address f, address t, uint256 a) public returns (bool) {
        TransferFromArgs memory args;
        args.to = t;
        args.amount = a;
        transferFromCalled[f] = args;
        return transferFromReturn;
    }

    function depositReturns(uint256 a) public {
        depositReturn = a;
    }

    function deposit(uint256 a, address r) public returns (uint256) {
        depositCalled[r] = a;
        return depositReturn;
    }
}