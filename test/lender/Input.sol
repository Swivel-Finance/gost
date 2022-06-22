// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Swivel.sol' as sw;

library Input {
    /// @notice used to pass data to Swivel's lend method
    /// @param principal value of a specific principal according to the Illuminate Principals Enum
    /// @param underlying address of an underlying asset
    /// @param maturity (timestamp) of the market
    /// @param amounts array of amounts of underlying tokens lent to each order in the orders array
    /// @param yield pool that leftover assets are sent to
    /// @param orders array of Swivel orders being filled
    /// @param swivel array of signatures for each order in the orders array
    struct Swivel {
        uint8 principal;
        address underlying;
        uint256 maturity;
        uint256[] amounts;
        address pool;
        sw.Swivel.Order[] orders;
        sw.Swivel.Components[] components;
    }

    /// @notice used to pass data to Element's lend function
    /// @param principal value of a specific principal according to the Illuminate Principals Enum
    /// @param underlying address of an underlying asset
    /// @param maturity (timestamp) of the market
    /// @param amount of principal tokens to lend
    /// @param returned minimum amount to return, this puts a cap on allowed slippage
    /// @param deadline is a timestamp by which the swap must be executed deadline is a timestamp by which the swap must be executed
    /// @param pool that is lent
    /// @param id of the pool
    struct Element {
        uint8 principal;
        address underlying;
        uint256 maturity;
        uint256 amount;
        uint256 returned;
        uint256 deadline;
        address pool;
        bytes32 id;
    }

    /// @param principal value of a specific principal according to the Illuminate Principals Enum
    /// @param underlying address of an underlying asset
    /// @param maturity (timestamp) of the market
    /// @param amount of principal tokens to lend
    /// @param returned minimum amount to return, this puts a cap on allowed slippage
    /// @param deadline is a timestamp by which the swap must be executed
    struct Pendle {
        uint8 principal;
        address underlying;
        uint256 maturity;
        uint256 amount;
        uint256 returned;
        uint256 deadline;
    }

    /// @param principal value of a specific principal according to the Illuminate Principals Enum
    /// @param underlying address of an underlying asset
    /// @param maturity (timestamp) of the market
    /// @param amount of underlying tokens to lend
    /// @param returned minimum number of tokens to lend (sets a limit on the order)
    /// @param amm that is used to conduct the swap
    /// @param divider address of contract that holds the principal token for this market
    struct Sense {
        uint8 principal;
        address underlying;
        uint256 maturity;
        uint128 amount;
        uint256 returned;
        address amm;
        address divider;
    }

    /// @param principal value of a specific principal according to the Illuminate Principals Enum
    /// @param underlying address of an underlying asset
    /// @param maturity (timestamp) of the market
    /// @param amount of principal tokens to lend
    /// @param returned minimum amount to return when executing the swap (sets a limit to slippage)
    /// @param deadline is a timestamp by which the swap must be executed
    /// @param pool that houses the underlying principal tokens
    /// @param amm that executes the swap
    struct Tempus {
        uint8 principal;
        address underlying;
        uint256 maturity;
        uint128 amount;
        uint256 returned;
        uint256 deadline;
        address pool;
        address amm;
    }

    /// @param principal value of a specific principal according to the Illuminate Principals Enum
    /// @param underlying address of an underlying asset
    /// @param maturity (timestamp) of the market
    /// @param amount of underlying tokens to lend
    /// @param returned the minimum amount of zero-coupon tokens to return accounting for slippage
    /// @param pool the address of a given APWine pool
    /// @param aave the aave that stores the lent capital
    /// @param id of the pool
    struct APWine {
        uint8 principal;
        address underlying;
        uint256 maturity;
        uint128 amount;
        uint256 returned;
        address pool;
        address aave;
        uint256 id;
    }
}
