// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract MarketPlace {
    struct MarketsArgs {
        uint256 maturity;
        uint8 precision;
    }

    struct PoolsArgs {
        uint256 maturity;
        uint8 precision;
    }

    address private princialTokenReturn;
    address private zcTokenReturn;
    address private zcTokenReturnSpoof;
    address private poolsReturn;
    bool private pausedReturn;

    mapping(address => MarketsArgs) public marketsCalled;
    mapping(address => PoolsArgs) public poolsCalled;
    uint8 public pausedCalled;

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

    function poolsReturns(address p) public {
        poolsReturn = p;
    }

    function pools(
        uint8 p,
        address u,
        uint256 m
    ) external returns (address) {
        poolsCalled[u] = PoolsArgs(m, p);
        return poolsReturn;
    }

    function pausedReturns(bool p) public {
        pausedReturn = p;
    }

    function paused(uint8 p) external returns (bool) {
        pausedCalled = p;
        return pausedReturn;
    }
}
