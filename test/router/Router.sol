// SPDX-License-Identifier: BUSL-1.1
pragma solidity >= 0.8.4;

import './IPool.sol'; 
import './SafeTransferLib.sol';
import '../pool/Pool.sol';

contract Router {

    address public immutable admin;
    address public immutable marketplace;
    bool public paused;

    constructor(address a, address m) {
        admin = a;
        marketplace = m;
    }

    mapping (address => mapping (uint256 => IPool)) public pools;

    // a = the amount of PT sold
    function sellPT(address u, uint256 m, uint128 a) external unpaused() returns (uint128 returned) {
        IPool pool = IPool(pools[u][m]);
        SafeTransferLib.safeTransfer(ERC20(address(pool.PT())), address(pool), a);
        return pool.sellPT(msg.sender, pool.sellPTPreview(a));
    }

    // a = the amount of PT bought
    function buyPT(address u, uint256 m, uint128 a) external unpaused() returns (uint128 returned) {
        IPool pool = IPool(pools[u][m]);
        SafeTransferLib.safeTransfer(ERC20(address(pool.underlying())), address(pool), a);
        return pool.buyPT(msg.sender, pool.buyPTPreview(a), a);
    }

    // a = the amount of underlying sold
    function sellUnderlying(address u, uint256 m, uint128 a) external unpaused() returns (uint128 returned) {
        IPool pool = IPool(pools[u][m]);
        SafeTransferLib.safeTransfer(ERC20(address(pool.underlying())), address(pool), a);
        return pool.sellunderlying(msg.sender, pool.sellunderlyingPreview(a));
    }

    // a = the amount of underlying bought
    function buyUnderlying(address u, uint256 m, uint128 a) external unpaused() returns (uint128 returned) {
        IPool pool = IPool(pools[u][m]);
        SafeTransferLib.safeTransfer(ERC20(address(pool.PT())), address(pool), a);
        return pool.buyunderlying(msg.sender, pool.buyunderlyingPreview(a), a);
    }   

    function createPool(address u, uint256 m, address i) external authorized(marketplace) returns (address p) {
        ////////////////////////////////////////////Secs in 10 yr//sell underlying fee co-eff//sell PT co-eff
        Pool pool = new Pool(IERC20(u), IPT(i), 23381681843, 13835058055282163712, 24595658764946068821);
        pools[u][m] = pool;
        return (address(pool));
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