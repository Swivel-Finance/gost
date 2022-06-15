// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract APWineToken {
    // a struct to hold the arguments passed to transferFrom
    struct TransferFromArgs {
        address to;
        uint256 amount;
    }
    // mapping of arguments sent to transfer. key is the passed in address.
    mapping(address => uint256) public transferCalled;
    // mapping of arguments sent to transferFrom. key is passed from address.
    mapping(address => TransferFromArgs) public transferFromCalled;

    // balanceOf does not require a mapping.
    address public balanceOfCalled;

    address private getUnderlyingOfIBTAddressReturn;

    // a uint to return for balanceOf calls
    uint256 private balanceOfReturn;
    // a boolean flag which allows us to dictate the return of transfer().
    bool private transferReturn;
    // a boolean flag which allows us to dictate the return of transferFrom().
    bool private transferFromReturn;

    function getUnderlyingOfIBTAddressReturns(address a) external {
        getUnderlyingOfIBTAddressReturn = a;
    }

    function getUnderlyingOfIBTAddress() external view returns (address) {
        return getUnderlyingOfIBTAddressReturn;
    }

    function balanceOfReturns(uint256 b) public {
        balanceOfReturn = b;
    }

    function balanceOf(address t) public returns (uint256) {
        balanceOfCalled = t;
        return balanceOfReturn;
    }

    function transfer(address t, uint256 a) public returns (bool) {
        transferCalled[t] = a;
        return transferReturn;
    }

    function transferReturns(bool b) public {
        transferReturn = b;
    }

    function transferFrom(
        address f,
        address t,
        uint256 a
    ) public returns (bool) {
        TransferFromArgs memory args;
        args.to = t;
        args.amount = a;
        transferFromCalled[f] = args;
        return transferFromReturn;
    }

    function transferFromReturns(bool b) public {
        transferFromReturn = b;
    }
}
