// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './IPErc20.sol';

interface IZcToken is IPErc20 {
    function burn(address, uint256) external returns (bool);

    function mint(address, uint256) external returns (bool);
}
