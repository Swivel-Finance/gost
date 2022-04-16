// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity >=0.8.4;

import "./LibCompound.sol";

contract CompoundAdapter {

    function exchangeRateCurrent(CERC20 cToken) external view returns (uint256){
        return(LibCompound.viewExchangeRate(cToken));
    }
}