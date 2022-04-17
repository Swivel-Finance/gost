// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity >=0.8.4;

import "./LibFuse.sol";

contract RariAdapter {

    function exchangeRateCurrent(address a) external view returns (uint256){
        return(LibFuse.viewExchangeRate(CERC20(a)));
    }
}