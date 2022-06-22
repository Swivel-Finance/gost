// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Swivel.sol' as sw;

library Input {
    struct Swivel {
        uint8 principal;
        address underlying;
        uint256 maturity;
        uint256[] amounts;
        address pool;
        sw.Swivel.Order[] orders;
        sw.Swivel.Components[] components;
    }

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

    struct Pendle {
        uint8 principal;
        address underlying;
        uint256 maturity;
        uint256 amount;
        uint256 returned;
        uint256 deadline;
    }

    struct Sense {
        uint8 principal;
        address underlying;
        uint256 maturity;
        uint128 amount;
        uint256 returned;
        address amm;
        address divider;
    }

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
