// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.8.4;

import "./IEToken.sol";

contract EulerAdapter {

    function exchangeRateCurrent(address a) external view returns (uint256){
        return(IEToken(a).convertBalanceToUnderlying(1e18));
    }
}