// SPDX-License-Identifier: MIT

pragma solidity 0.8.4;

import "./IErc20Metadata.sol";
import "./IErc20.sol";

/**
 * @dev Mint and burn interface for the ZCToken
 *
 */
interface IZcToken is IErc20, IErc20Metadata {
    /**
     * @dev Mints...
     */
    function mint(address, uint256) external returns(bool);

    /**
     * @dev Burns...
     */
    function burn(address, uint256) external returns(bool);
}
