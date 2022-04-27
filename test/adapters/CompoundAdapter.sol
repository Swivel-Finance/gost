// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity >= 0.8.10;

import "./LibCompound.sol";

contract CompoundAdapter {

    uint8 public protocol = 1;

    function exchangeRateCurrent(address a) external view returns (uint256){
        return(LibCompound.viewExchangeRate(CERC20(a)));
    }
}