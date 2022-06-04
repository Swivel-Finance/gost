// SPDX-License-Identifier: MIT
pragma solidity >= 0.8.4;

import './IPool.sol'; 
import './SafeTransferLib.sol';

contract Router {

    address public immutable admin;
    bool public paused;

    constructor(address a) {
        admin = a;
    }

    mapping (address => mapping (uint256 => IPool)) public pools;

    // a = the amount of FYToken sold
    function sellFYToken(address u, uint256 m, uint128 a) external unpaused() returns (uint128 returned) {
        IPool pool = IPool(pools[u][m]);
        SafeTransferLib.safeTransfer(ERC20(address(pool.fyToken())), address(pool), a);
        return pool.sellFYToken(msg.sender, pool.sellFYTokenPreview(a));
    }

    // a = the amount of FYToken bought
    function buyFYToken(address u, uint256 m, uint128 a) external unpaused() returns (uint128 returned) {
        IPool pool = IPool(pools[u][m]);
        SafeTransferLib.safeTransfer(ERC20(address(pool.base())), address(pool), a);
        return pool.buyFYToken(msg.sender, pool.buyFYTokenPreview(a), a);
    }

    // a = the amount of underlying sold
    function sellBase(address u, uint256 m, uint128 a) external unpaused() returns (uint128 returned) {
        IPool pool = IPool(pools[u][m]);
        SafeTransferLib.safeTransfer(ERC20(address(pool.base())), address(pool), a);
        return pool.sellBase(msg.sender, pool.sellBasePreview(a));
    }

    // a = the amount of base bought
    function buyBase(address u, uint256 m, uint128 a) external unpaused() returns (uint128 returned) {
        IPool pool = IPool(pools[u][m]);
        SafeTransferLib.safeTransfer(ERC20(address(pool.fyToken())), address(pool), a);
        return pool.buyBase(msg.sender, pool.buyBasePreview(a), a);
    }   

    /// @notice Called by admin at any point to pause / unpause market transactions
    /// @param b Boolean which indicates the markets paused status
    function pause(bool b) external authorized(admin) returns (bool) {
        paused = b;
        return true;
    }

    modifier authorized(address a) {
        require(msg.sender == a, 'sender must be authorized');
        _;
    }

    modifier unpaused() {
        require(!paused, 'markets are paused');
        _;
    }
}