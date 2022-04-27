// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.8.4;

import "./IYearnVault.sol";

contract YearnAdapter {

    uint8 public protocol = 3;

    function exchangeRateCurrent(address a) external view returns (uint256){
        return(IYearnVault(a).pricePerShare());
    }
}