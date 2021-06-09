// SPDX-License-Identifier: MIT

pragma solidity 0.8.4;

import "./IERC20Metadata.sol";
import "./IERC20.sol";

/**
 * @dev Mint and burn interface for the ZCToken
 * NOTE: The OZStyle naming conventions (vs our OzStlye) are kept 
 * in the OZ '20 and '2612 members.
 *
 */
interface IZcToken is IERC20, IERC20Metadata {
    /**
     * @dev Mints...
     */
    function mint(address, uint256) external returns(bool);

    /**
     * @dev Burns...
     */
    function burn(address, uint256) external returns(bool);
}
