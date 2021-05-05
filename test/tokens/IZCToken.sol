// SPDX-License-Identifier: MIT

pragma solidity 0.8.0;

import "./IERC20Metadata.sol";
import "./IERC20.sol";

/**
 * @dev Mint and burn interface for the ZCToken
 *
 */
interface IZCToken is IERC20, IERC20Metadata {
    /**
     * @dev Mints...
     */
    function mint(address, uint256) external returns(bool);

    /**
     * @dev Burns...
     */
    function burn(address, uint256) external returns(bool);
}
