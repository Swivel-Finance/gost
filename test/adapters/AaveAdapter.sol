// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity >=0.8.4;

import "./IAToken.sol";
import "./IAavePool.sol";

contract AaveAdapter {

    uint8 public protocol = 4;

    function exchangeRateCurrent(address a) external view returns (uint256){
        return(IAavePool(IAToken(a).POOL()).getReserveNormalizedIncome(aToken.UNDERLYING_ASSET_ADDRESS()));
    }
}