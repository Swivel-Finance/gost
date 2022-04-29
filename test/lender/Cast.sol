// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;


library Cast {
    /// @dev Safely cast an uint256 to an uint128
    /// @param n the u256 to cast to u128
    function u128(uint256 n) internal pure returns (uint128) {
        require (n <= type(uint128).max, 'failed to cast');
        return uint128(n);
    }
}
