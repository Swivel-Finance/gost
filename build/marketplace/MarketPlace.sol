// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Interfaces.sol';
import './ZcToken.sol';
import './Safe.sol';

contract MarketPlace {
    /// @notice the available principals
    /// @dev the order of this enum is used to select protocols from the markets
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

    /// markets are defined by a market pair which point to a fixed length array
    /// of principal token addresses. The principal tokens those addresses
    /// represent correspond to their Principals enum value, thus the array
    /// should be ordered in that way
    mapping(address => mapping(uint256 => address[9])) public markets;

    mapping(address => mapping(uint256 => address[9])) public pools;

    address public admin;
    /// @notice address of the deployed redeemer contract
    address public immutable redeemer;

    event CreateMarket(address indexed underlying, uint256 indexed maturity);

    /// @param r address of the deployed redeemer contract
    constructor(address r) {
        admin = msg.sender;
        redeemer = r;
    }

    /// @notice creates a new market for the given underlying token and maturity
    /// @notice all seven protocols should be provided in the order of their enum value
    /// @param u address of an underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param t principal token addresses for this market minus the illuminate principal (which is added here)
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
        require(markets[u][m][uint256(Principals.Illuminate)] == address(0), 'market exists');

        // deploy an illuminate token with this new market
        // NOTE: ATM is using name as symbol args
        address iToken = address(new ZcToken(u, m, n, s, d));

        // the market will have the illuminate principal as its zeroth item, thus t should have Principals[1] as [0]
        // TODO we could choose to put illuminate last in
        address[9] memory market = [iToken, t[0], t[1], t[2], t[3], t[4], t[5], t[6], t[7]];

        // max is the maximum integer value for a 256 unsighed integer
        uint256 max = 2**256 - 1;

        // approve the underlying for max per given principal
        for (uint8 i; i < 8; i++) {
            Safe.approve(IErc20(market[i]), redeemer, max);
        }

        // set the market
        markets[u][m] = market;

        emit CreateMarket(u, m);

        return true;
    }

    /// @notice sets the address for a pool
    /// @param p enum value of the principal token
    /// @param u address of the underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a address of the pool
    /// @return bool true if successful
    function setPool(
        uint8 p,
        address u,
        uint256 m,
        address a
    ) external authorized(admin) returns (bool) {
        require(pools[u][m][p] == address(0), 'pool exists');
        pools[u][m][p] = a;
        return true;
    }

    /// @notice sells the PT for the PT via the pool
    /// @param p enum value of the principal token
    /// @param u address of the underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a amount of PT to swap
    /// @return uint128 amount of PT bought
    function sellPT(
        uint8 p,
        address u,
        uint256 m,
        uint128 a
    ) external returns (uint128) {
        IPool pool = IPool(pools[u][m][p]);
        Safe.transfer(IErc20(address(pool.PT())), address(pool), a);
        return pool.sellPT(msg.sender, pool.sellPTPreview(a));
    }

    /// @notice buys the underlying for the PT via the pool
    /// @param p enum value of the principal token
    /// @param u address of the underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a amount of underlying tokens to sell
    /// @return uint128 amount of PT received
    function buyPT(
        uint8 p,
        address u,
        uint256 m,
        uint128 a
    ) external returns (uint128) {
        IPool pool = IPool(pools[u][m][p]);
        Safe.transfer(IErc20(address(pool.underlying())), address(pool), a);
        return pool.buyPT(msg.sender, pool.buyPTPreview(a), a);
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
    ) external returns (uint128) {
        IPool pool = IPool(pools[u][m][p]);
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
    ) external returns (uint128) {
        IPool pool = IPool(pools[u][m][p]);
        Safe.transfer(IErc20(address(pool.PT())), address(pool), a);
        return pool.buyUnderlying(msg.sender, pool.buyUnderlyingPreview(a), a);
    }

    modifier authorized(address a) {
        require(msg.sender == a, 'sender must be authorized');
        _;
    }
}
