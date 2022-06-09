// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract MarketPlace {
    struct MarketsArgs {
        uint256 maturity;
        uint8 precision;
    }

    address private princialTokenReturn;
    address private zcTokenReturn;
    address private zcTokenReturnSpoof;

    mapping(address => MarketsArgs) public marketsCalled;

    function principalTokenReturns(address p) external {
        princialTokenReturn = p;
    }

    function zcTokenReturns(address z) external {
        zcTokenReturn = z;
    }

    function zcTokenSpoofReturns(address s) external {
        zcTokenReturnSpoof = s;
    }

    function markets(
        address u,
        uint256 m,
        uint8 p
    ) external returns (address) {
        marketsCalled[u] = MarketsArgs(m, p);
        address spoofed = zcTokenReturnSpoof;
        if (p == 0) {
            if (spoofed != address(0)) {
                zcTokenReturnSpoof = address(0);
                return spoofed;
            }
            return zcTokenReturn;
        }
        return princialTokenReturn;
    }
}
