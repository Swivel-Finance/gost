// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.8.4;

import "./IAToken.sol";
import "./IAavePool.sol";

contract AaveAdapter {

    function exchangeRateCurrent(address a) external view returns (uint256){
        IAToken aToken = IAToken(a);
        return(IAavePool(aToken.POOL()).getReserveNormalizedIncome(aToken.UNDERLYING_ASSET_ADDRESS()));
    }
}