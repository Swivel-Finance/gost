// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract NotionalToken {
    struct TransferFromArgs {
        address to;
        uint256 amount;
    }

    address private getUnderlyingTokenReturn;
    uint256 private getMaturityReturn;
    bool private transferFromReturn;
    uint256 private balanceOfReturn;

    mapping (address => TransferFromArgs) public transferFromCalled;
    mapping (address => uint256) public depositCalled;
    address public balanceOfCalled;

    function getUnderlyingTokenReturns(address a) external {
        getUnderlyingTokenReturn = a;
    }

    function getMaturityReturns(uint256 m) external {
        getMaturityReturn = m;
    }

    function getUnderlyingToken() external view returns (address) {
        return getUnderlyingTokenReturn;
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

    function deposit(address a, uint256 b) public {
        depositCalled[a] = b;
    }
}