// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

contract MarketPlace {
    struct MarketsArgs {
        uint256 maturity;
        uint8 precision;
    }

    address private princialTokenReturn;
    address private zcTokenReturn;
    address private spoof;

    bool public spoofing = false;

    mapping (address => MarketsArgs) public marketsCalled;

    function isSpoofing(bool a) public {
        spoofing = a;
    }

    function principalTokenReturns(address p) external {
        princialTokenReturn = p;
    }

    function zcTokenReturns(address z) external {
        zcTokenReturn = z;
    }

    function spoofingReturns(address s) external {
        spoof = s;
    }

    function markets(address u, uint256 m, uint8 p) external returns (address) {
        marketsCalled[u] = MarketsArgs(m, p);
        if (p == 0) {
          if (spoofing) {
            spoofing = false;
            return spoof;
          }
          return zcTokenReturn;
        }
        return princialTokenReturn;
    }
}
