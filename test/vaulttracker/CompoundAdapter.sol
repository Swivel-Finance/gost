// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.8.4;

import "./LibCompound.sol";

contract CompoundAdapter {

    function exchangeRateCurrent(address a) external view returns (uint256){
        return(LibCompound.viewExchangeRate(CERC20(a)));
    }
}