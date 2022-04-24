// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.8.4;

import "./ILidoETH.sol";

contract LidoAdapter {

    function exchangeRateCurrent(address a) external view returns (uint256){
        return(ILidoETH(a).getPooledEthByShares(1e26));
    }
}