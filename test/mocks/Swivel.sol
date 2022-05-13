// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

struct Components {
    uint8 v;
    bytes32 r;
    bytes32 s;
}

struct Order {
    bytes32 key;
    address maker;
    address underlying;
    bool vault;
    bool exit;
    uint256 principal;
    uint256 premium;
    uint256 maturity;
    uint256 expiry;
}

contract Swivel {
    struct RedeemZcTokenArgs {
        uint256 underlying;
        uint256 maturity;
    }

    bool initateReturn;

    mapping (address => uint256) public initiateCalledAmount;
    mapping (address => uint8) public initiateCalledSignature;
    mapping (address => RedeemZcTokenArgs) public redeemZcTokenCalled;

    function initiateReturns(bool i) external {
        initateReturn = i;
    }

    function initiate(
        Order[] calldata o, 
        uint256[] calldata a, 
        Components[] calldata s) 
    external returns (bool) {
        for (uint256 i = 0; i < o.length; i++) {
            initiateCalledAmount[o[i].maker] = a[i];
            initiateCalledSignature[o[i].maker] = s[i].v;
        }
        return initateReturn;
    }

    function redeemZcToken(address u, uint256 m, uint256 a) external {
        redeemZcTokenCalled[u] = RedeemZcTokenArgs(m, a);
    }
}