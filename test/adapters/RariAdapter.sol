// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity >= 0.8.10;

import "./LibFuse.sol";

contract RariAdapter {

    uint8 public protocol = 2;

    function exchangeRateCurrent(address a) external view returns (uint256){
        return(LibFuse.viewExchangeRate(CERC20(a)));
    }
}