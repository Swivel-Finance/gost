// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract APWineToken {
    // a struct to hold the arguments passed to transferFrom
    struct TransferFromArgs {
        address to;
        uint256 amount;
    }

    mapping (address => uint256) public balances;
    address private getPTAddressReturn;
    bool private transferFromReturn;

    mapping (address => TransferFromArgs) public transferFromCalled;

    function getPTAddressReturns(address a) external {
        getPTAddressReturn = a;
    }

    function getPTAddress() external view returns (address) {
        return getPTAddressReturn;
    }

    function balanceOfReturns(address a, uint256 b) external {
        balances[a] = b;
    }

    function balanceOf(address a) external view returns (uint256) {
        return balances[a];
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

}