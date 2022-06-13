// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract Aave {
    struct DepositArgs {
        address onBehalfOf;
        uint256 amount;
        uint256 minimumAmount;
    }

    bool private depositReturn;

    mapping(address => DepositArgs) public depositCalled;

    function depositReturns(bool b) external {
        depositReturn = b;
    }

    function depsoit(
        address a,
        address b,
        uint256 amount,
        uint16 r
    ) external returns (bool) {
        depositCalled[a] = DepositArgs(b, amount, r);
        return depositReturn;
    }
}
