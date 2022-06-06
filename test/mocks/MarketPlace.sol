// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract MarketPlace {
    struct MarketsCalledArgs {
        uint256 maturity;
        uint8 precision;
    }

    address private marketsReturn;
    address private zcTokenMarketReturn;

    mapping (address => MarketsCalledArgs) public marketsCalled;

    function marketsReturns(address m) external {
        marketsReturn = m;
    }

    function marketZcTokenReturns(address z) external {
        zcTokenMarketReturn = z;
    }

    function markets(address u, uint256 m, uint8 p) external returns (address) {
        marketsCalled[u] = MarketsCalledArgs(m, p);
        if (p == 0) {
          return zcTokenMarketReturn;
        }
        return marketsReturn;
    }
}
