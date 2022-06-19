// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Interfaces.sol';
import './ERC5095.sol';
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
    /// @notice address of the deployed lender contract
    address public immutable lender;
    /// @notice flag that determines if a principal's pool is available
    bool[9] public paused = [false, false, false, false, false, false, false, false, false];

    event CreateMarket(address indexed underlying, uint256 indexed maturity);

    /// @param r address of the deployed redeemer contract
    /// @param l address of the deployed lender contract
    constructor(address r, address l) {
        admin = msg.sender;
        redeemer = r;
        lender = l;
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
        address iToken = address(new ERC5095(u, m, redeemer, lender, n, s, d));

        // the market will have the illuminate principal as its zeroth item, thus t should have Principals[1] as [0]
        address[9] memory market = [iToken, t[0], t[1], t[2], t[3], t[4], t[5], t[6], t[7]];

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

    /// @param a Address of a new admin
    /// @return bool true if successful
    function setAdmin(address a) external authorized(admin) returns (bool) {
        admin = a;
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
        Safe.transfer(IERC20(address(pool.fyToken())), address(pool), a);
        return pool.sellFYToken(msg.sender, pool.sellFYTokenPreview(a));
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
        Safe.transfer(IERC20(address(pool.base())), address(pool), a);
        return pool.buyFYToken(msg.sender, pool.buyFYTokenPreview(a), a);
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
        Safe.transfer(IERC20(address(pool.base())), address(pool), a);
        return pool.sellBase(msg.sender, pool.sellBasePreview(a));
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
        Safe.transfer(IERC20(address(pool.fyToken())), address(pool), a);
        return pool.buyBase(msg.sender, pool.buyBasePreview(a), a);
    }

    //////////LIQUIDITY PROVIDER METHODS//////////////////
    /// @dev Mint liquidity tokens in exchange for adding underlying and PT
    /// The amount of liquidity tokens to mint is calculated from the amount of unaccounted for PT in this contract.
    /// A proportional amount of underlying tokens need to be present in this contract, also unaccounted for.
    /// @param u the address of the underlying token
    /// @param m the maturity of the principal token
    /// @param uA the underlying amount being sent
    /// @param ptA the principal token amount being sent
    /// @param minRatio Minimum ratio of underlying to PT in the pool.
    /// @param maxRatio Maximum ratio of underlying to PT in the pool.
    /// @return The amount of liquidity tokens minted.
    function mint(address u, uint256 m, uint256 uA, uint256 ptA, uint256 minRatio, uint256 maxRatio) external returns (uint256, uint256, uint256) {
        IPool pool = IPool(pools[u][m]);
        Safe.transferFrom(ERC20(address(pool.base())), msg.sender, address(pool), uA);  
        Safe.transferFrom(ERC20(address(pool.fyToken())), msg.sender, address(pool), ptA);
        return pool.mint(msg.sender, msg.sender, minRatio, maxRatio);
    }

    /// @dev Mint liquidity tokens in exchange for adding only underlying
    /// The amount of liquidity tokens is calculated from the amount of PT to buy from the pool,
    /// plus the amount of unaccounted for PT in this contract.
    /// The underlying tokens need to be present in this contract, unaccounted for.
    /// @param u the address of the underlying token
    /// @param m the maturity of the principal token
    /// @param a the underlying amount being sent
    /// @param ptBought Amount of `PT` being bought in the Pool, from this we calculate how much underlying it will be taken in.
    /// @param minRatio Minimum ratio of underlying to PT in the pool.
    /// @param maxRatio Maximum ratio of underlying to PT in the pool.
    /// @return The amount of liquidity tokens minted.
    function mintWithUnderlying(address u, uint256 m, uint256 a, uint256 ptBought, uint256 minRatio, uint256 maxRatio) external returns (uint256, uint256, uint256) {
        IPool pool = IPool(pools[u][m]);
        Safe.safeTransferFrom(ERC20(address(pool.underlying())), msg.sender, address(pool), a);
        return pool.mintWithUnderlying(msg.sender, msg.sender, ptBought, minRatio, maxRatio);
    }

    /// @dev Burn liquidity tokens in exchange for underlying and PT.
    /// The liquidity tokens need to be in this contract.
    /// @param minRatio Minimum ratio of underlying to PT in the pool.
    /// @param maxRatio Maximum ratio of underlying to PT in the pool.
    /// @return The amount of tokens burned and returned (tokensBurned, underlyings, PTs).
    function burn(address u, uint256 m, uint256 minRatio, uint256 maxRatio) external returns (uint256, uint256, uint256) {
        return IPool(pools[u][m]).burn(msg.sender, msg.sender, minRatio, maxRatio);
    }

    /// @dev Burn liquidity tokens in exchange for underlying.
    /// @param minRatio Minimum ratio of underlying to PT in the pool.
    /// @param maxRatio Minimum ratio of underlying to PT in the pool.
    /// @return tokensBurned The amount of lp tokens burned.
    /// @return underlyingOut The amount of underlying tokens returned.
    function burnForUnderlying(address u, uint256 m, uint256 minRatio, uint256 maxRatio) external override returns (uint256 tokensBurned, uint256 underlyingOut) {
        return IPool(pools[u][m]).burnForUnderlying(msg.sender, minRatio, maxRatio);
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
