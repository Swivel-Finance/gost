// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Interfaces.sol';
import './ZcToken.sol';
import './Safe.sol';

contract MarketPlace {
    /// @notice the available principals
    /// @dev the order of this enum is used to select principals from the markets
    /// mapping
    /// @dev e.g. Illuminate => 0, Swivel => 1, and so on
    enum Principals {
        Illuminate,
        Swivel,
        Yield,
        Element,
        Pendle,
        Tempus,
        Sense,
        Apwine,
        Notional
    }

    error Exists(string);
    error Unauthorized();
    error Invalid(string);

    /// markets are defined by a market pair which point to a fixed length array
    /// of principal token addresses. The principal tokens those addresses
    /// represent correspond to their Principals enum value, thus the array
    /// should be ordered in that way
    mapping(address => mapping(uint256 => address[9])) public markets;

    mapping(address => mapping(uint256 => address)) public pools;

    address public admin;
    /// @notice address of the deployed redeemer contract
    address public immutable redeemer;
    /// @notice flag that determines if a principal's pool is available
    bool[9] public paused = [false, false, false, false, false, false, false, false, false];

    event CreateMarket(address indexed underlying, uint256 indexed maturity);

    /// @param r address of the deployed redeemer contract
    constructor(address r) {
        admin = msg.sender;
        redeemer = r;
    }

    /// @notice allows the admin to create a market
    /// @param p enum value of the principal token
    /// @param u underlying token address
    /// @param m maturity timestamp for the market
    /// @param a address of the new market
    /// @return bool true if successful
    function setMarket(
        uint8 p,
        address u,
        uint256 m,
        address a
    ) external authorized(admin) returns (bool) {
        if (markets[u][m][p] != address(0)) {
            revert Exists('Market already exists');
        }
        markets[u][m][p] = a;
        return true;
    }

    /// @notice creates a new market for the given underlying token and maturity
    /// @notice all seven principals should be provided in the order of their enum value
    /// @param u address of an underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param t principal token addresses for this market minus the illuminate
    /// principal (which is added here)
    /// @param n name for the illuminate token
    /// @param s symbol for the illuminate token
    /// @param d decimals for the illuminate token
    /// @return bool true if successful
    function createMarket(
        address u,
        uint256 m,
        address[8] memory t,
        string calldata n,
        string calldata s,
        uint8 d
    ) external authorized(admin) returns (bool) {
        if (markets[u][m][uint256(Principals.Illuminate)] != address(0)) {
            revert Exists('market already exists');
        }

        // deploy an illuminate token with this new market
        // NOTE: ATM is using name as symbol args
        address iToken = address(new ZcToken(u, m, n, s, d));

        // the market will have the illuminate principal as its zeroth item, thus t should have Principals[1] as [0]
        // TODO we could choose to put illuminate last in
        address[9] memory market = [iToken, t[0], t[1], t[2], t[3], t[4], t[5], t[6], t[7]];

        // max is the maximum integer value for a 256 unsighed integer
        uint256 max = 2**256 - 1;

        // approve the underlying for max per given principal
        for (uint8 i; i < 8; ) {
            Safe.approve(IErc20(market[i]), redeemer, max);
            unchecked {
                i++;
            }
        }
        // set the market
        markets[u][m] = market;

        emit CreateMarket(u, m);

        return true;
    }

    /// @notice allows the admin to pause a principal's pools
    /// @param p principal's enum value
    /// @param s true if the pool should be paused, false otherwise
    /// @return bool true if successful
    function pause(uint8 p, bool s) external authorized(admin) returns (bool) {
        paused[p] = s;
        return true;
    }

    /// @notice sets the address for a pool
    /// @param u address of the underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a address of the pool
    /// @return bool true if successful
    function setPool(
        address u,
        uint256 m,
        address a
    ) external authorized(admin) returns (bool) {
        if (pools[u][m] != address(0)) {
            revert Exists('pool already exists');
        }
        pools[u][m] = a;
        return true;
    }

    /// @notice sells the PT for the PT via the pool
    /// @param p enum value of the principal token
    /// @param u address of the underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a amount of PT to swap
    /// @return uint128 amount of PT bought
    function sellPrincipalToken(
        uint8 p,
        address u,
        uint256 m,
        uint128 a
    ) external unpaused(p) returns (uint128) {
        IPool pool = IPool(pools[u][m]);
        Safe.transfer(IErc20(address(pool.principalToken())), address(pool), a);
        return pool.sellPrincipalToken(msg.sender, pool.sellPrincipalTokenPreview(a));
    }

    /// @notice buys the underlying for the PT via the pool
    /// @param p enum value of the principal token
    /// @param u address of the underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a amount of underlying tokens to sell
    /// @return uint128 amount of PT received
    function buyPrincipalToken(
        uint8 p,
        address u,
        uint256 m,
        uint128 a
    ) external unpaused(p) returns (uint128) {
        IPool pool = IPool(pools[u][m]);
        Safe.transfer(IErc20(address(pool.underlying())), address(pool), a);
        return pool.buyPrincipalToken(msg.sender, pool.buyPrincipalTokenPreview(a), a);
    }

    /// @notice sells the underlying for the PT via the pool
    /// @param p enum value of the principal token
    /// @param u address of the underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a amount of underlying to swap
    /// @return uint128 amount of underlying sold
    function sellUnderlying(
        uint8 p,
        address u,
        uint256 m,
        uint128 a
    ) external unpaused(p) returns (uint128) {
        IPool pool = IPool(pools[u][m]);
        Safe.transfer(IErc20(address(pool.underlying())), address(pool), a);
        return pool.sellUnderlying(msg.sender, pool.sellUnderlyingPreview(a));
    }

    /// @notice buys the underlying for the PT via the pool
    /// @param p enum value of the principal token
    /// @param u address of the underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a amount of PT to swap
    /// @return uint128 amount of underlying bought
    function buyUnderlying(
        uint8 p,
        address u,
        uint256 m,
        uint128 a
    ) external unpaused(p) returns (uint128) {
        IPool pool = IPool(pools[u][m]);
        Safe.transfer(IErc20(address(pool.principalToken())), address(pool), a);
        return pool.buyUnderlying(msg.sender, pool.buyUnderlyingPreview(a), a);
    }

    modifier authorized(address a) {
        if (msg.sender != a) {
            revert Unauthorized();
        }
        _;
    }

    modifier unpaused(uint8 p) {
        if (paused[p]) {
            revert Invalid('princpal paused');
        }
        _;
    }
}
