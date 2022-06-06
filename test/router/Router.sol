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

    function createPool(address u, uint256 m, address i) external authorized(marketplace) returns (address p) {
        ////////////////////////////////////////////Secs in 10 yr//sell base fee co-eff//sell fyToken co-eff
        Pool pool = new Pool(IERC20(u), IFYToken(i), 23381681843, 13835058055282163712, 24595658764946068821);
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