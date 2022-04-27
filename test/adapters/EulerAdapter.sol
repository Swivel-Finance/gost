// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.8.4;

import "./IEToken.sol";

contract EulerAdapter {

    uint8 public protocol = 5;

    function exchangeRateCurrent(address a) external view returns (uint256){
        return(IEToken(a).convertBalanceToUnderlying(1e27));
    }
}

