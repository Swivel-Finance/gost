// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.8.4;

import "./IERC4626.sol";

contract ERC4626Adapter {

    uint8 public protocol = 6;

    function exchangeRateCurrent(address a) public view returns (uint256){
        return (IERC4626(a).convertToAssets(1e26));
    }
}