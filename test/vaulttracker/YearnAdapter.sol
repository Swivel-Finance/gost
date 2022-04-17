// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity >=0.8.4;

import "./IYearnVault.sol";

contract CompoundAdapter {

    function exchangeRateCurrent(address yToken) external view returns (uint256){
        return(IYearnVault(yToken).pricePerShare());
    }
}