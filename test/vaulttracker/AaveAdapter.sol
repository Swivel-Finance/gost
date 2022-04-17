// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity >=0.8.4;

import "./IAToken.sol";
import "./IAavePool.sol";

contract AaveAdapter {

    function exchangeRateCurrent(address a) external view returns (uint256){
        IAToken aToken = IAToken(a);
        return(IAavePool(aToken.POOL()).getReserveNormalizedIncome(aToken.UNDERLYING_ASSET_ADDRESS()));
    }
}