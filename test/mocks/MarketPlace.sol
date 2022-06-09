// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract MarketPlace {
    struct MarketsArgs {
        uint256 maturity;
        uint8 precision;
    }

    address private princialTokenReturn;
    address private zcTokenReturn;
    address private authZcTokenReturn;

    bool public isAuthorized;

    mapping (address => MarketsArgs) public marketsCalled;

    function principalTokenReturns(address p) external {
        princialTokenReturn = p;
    }

    function zcTokenReturns(address z, bool authorized) external {
        if (authorized) {
            authZcTokenReturn = z;
        } else {
            zcTokenReturn = z;
        }
    }

    function markets(address u, uint256 m, uint8 p) external returns (address) {
        marketsCalled[u] = MarketsArgs(m, p);
        if (p == 0) {
          if (isAuthorized) {
            isAuthorized = false;
            return authZcTokenReturn;
          }
          return zcTokenReturn;
        }
        return princialTokenReturn;
    }
}
