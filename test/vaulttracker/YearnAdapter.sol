// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.8.4;

import "./IYearnVault.sol";

contract YearnAdapter {

    function exchangeRateCurrent(address a) external view returns (uint256){
        return(IYearnVault(a).pricePerShare());
    }
}