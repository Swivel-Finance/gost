// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract Aave {
    struct DepositArgs {
        address onBehalfOf;
        uint256 amount;
        uint16 referralCode;
    }

    mapping(address => DepositArgs) public depositCalled;

    function deposit(
        address a,
        uint256 amount,
        address b,
        uint16 r
    ) external {
        depositCalled[a] = DepositArgs(b, amount, r);
    }
}
