// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.14;

import "./IERC20.sol";
import "./IERC20Metadata.sol";
import "./MinimalTransferHelper.sol";
import "./CastU256U128.sol";
import "./CastU256U112.sol";
import "./CastU256I256.sol";
import "./CastU128U112.sol";
import "./CastU128I128.sol";
import "./IPool.sol";
import "../tokens/IERC5095.sol";
import "./YieldMath.sol";
import "./ERC20Permit.sol";


/// @dev The Pool contract exchanges underlying for PT at a price defined by a specific formula.
contract Pool is IPool, ERC20Permit {
    using CastU256U128 for uint256;
    using CastU256U112 for uint256;
    using CastU256I256 for uint256;
    using CastU128U112 for uint128;
    using CastU128I128 for uint128;
    using MinimalTransferHelper for IERC20;

    event Trade(uint32 maturity, address indexed from, address indexed to, int256 underlyings, int256 PTs);
    event Liquidity(uint32 maturity, address indexed from, address indexed to, address indexed PTTo, int256 underlyings, int256 PTs, int256 poolTokens);
    event Sync(uint112 underlyingCached, uint112 PTCached, uint256 cumulativeBalancesRatio);

    int128 public immutable override ts;              // 1 / Seconds in 10 years, in 64.64
    int128 public immutable override g1;             // To be used when selling underlying to the pool
    int128 public immutable override g2;             // To be used when selling PT to the pool
    uint32 public immutable override maturity;
    uint96 public immutable override scaleFactor;    // Scale up to 18 low decimal tokens to get the right precision in YieldMath

    IERC20 public immutable override underlying;
    IERC5095 public immutable override PT;

    uint112 private underlyingCached;              // uses single storage slot, accessible via getCache
    uint112 private PTCached;           // uses single storage slot, accessible via getCache
    uint32  private blockTimestampLast;      // uses single storage slot, accessible via getCache

    uint256 public override cumulativeBalancesRatio;  // Fixed point factor with 27 decimals (ray)

    /// @dev Deploy a Pool.
    /// Make sure that the PT follows ERC20 standards with regards to name, symbol and decimals
    constructor(IERC20 underlying_, IERC4095 PT_, int128 ts_, int128 g1_, int128 g2_)
        ERC20Permit(
            string(abi.encodePacked(IERC20Metadata(address(PT_)).name(), " LP")),
            string(abi.encodePacked(IERC20Metadata(address(PT_)).symbol(), "LP")),
            IERC20Metadata(address(PT_)).decimals()
        )
    {
        PT = PT_;
        underlying = underlying_;

        uint256 maturity_ = PT_.maturity();
        require (maturity_ <= type(uint32).max, "Pool: Maturity too far in the future");
        maturity = uint32(maturity_);

        ts = ts_;
        g1 = g1_;
        g2 = g2_;

        scaleFactor = uint96(10 ** (18 - uint96(decimals)));
    }

    /// @dev Trading can only be done before maturity
    modifier beforeMaturity() {
        require(
            block.timestamp < maturity,
            "Pool: Too late"
        );
        _;
    }

    // ---- Balances management ----

    /// @dev Updates the cache to match the actual balances.
    function sync() external {
        _update(_getUnderlyingBalance(), _getPTBalance(), underlyingCached, PTCached);
    }

    /// @dev Returns the cached balances & last updated timestamp.
    /// @return Cached underlying token balance.
    /// @return Cached virtual FY token balance.
    /// @return Timestamp that balances were last cached.
    function getCache()
        external view override
        returns (uint112, uint112, uint32)
    {
        return (underlyingCached, PTCached, blockTimestampLast);
    }

    /// @dev Returns the "virtual" PT balance, which is the real balance plus the pool token supply.
    function getPTBalance()
        public view override
        returns(uint112)
    {
        return _getPTBalance();
    }

    /// @dev Returns the underlying balance
    function getUnderlyingBalance()
        public view override
        returns(uint112)
    {
        return _getUnderlyingBalance();
    }

    /// @dev Returns the "virtual" PT balance, which is the real balance plus the pool token supply.
    function _getPTBalance()
        internal view
        returns(uint112)
    {
        return (PT.balanceOf(address(this)) + _totalSupply).u112();
    }

    /// @dev Returns the underlying balance
    function _getUnderlyingBalance()
        internal view
        returns(uint112)
    {
        return underlying.balanceOf(address(this)).u112();
    }

    /// @dev Retrieve any underlying tokens not accounted for in the cache
    function retrieveUnderlying(address to)
        external override
        returns(uint128 retrieved)
    {
        retrieved = _getUnderlyingBalance() - underlyingCached; // Cache can never be above balances
        underlying.safeTransfer(to, retrieved);
        // Now the current balances match the cache, so no need to update the TWAR
    }

    /// @dev Retrieve any PTs not accounted for in the cache
    function retrievePT(address to)
        external override
        returns(uint128 retrieved)
    {
        retrieved = _getPTBalance() - PTCached; // Cache can never be above balances
        IERC20(address(PT)).safeTransfer(to, retrieved);
        // Now the balances match the cache, so no need to update the TWAR
    }

    /// @dev Update cache and, on the first call per block, ratio accumulators
    function _update(uint128 underlyingBalance, uint128 fyBalance, uint112 _underlyingCached, uint112 _PTCached) private {
        uint32 blockTimestamp = uint32(block.timestamp);
        uint32 timeElapsed = blockTimestamp - blockTimestampLast; // overflow is desired
        if (timeElapsed > 0 && _underlyingCached != 0 && _PTCached != 0) {
            // We multiply by 1e27 here so that r = t * y/x is a fixed point factor with 27 decimals 
            uint256 scaledPTCached = uint256(_PTCached) * 1e27;
            cumulativeBalancesRatio += scaledPTCached  * timeElapsed / _underlyingCached;
        }
        underlyingCached = underlyingBalance.u112();
        PTCached = fyBalance.u112();
        blockTimestampLast = blockTimestamp;
        emit Sync(underlyingCached, PTCached, cumulativeBalancesRatio);
    }

    // ---- Liquidity ----

    /// @dev Mint liquidity tokens in exchange for adding underlying and PT
    /// The amount of liquidity tokens to mint is calculated from the amount of unaccounted for PT in this contract.
    /// A proportional amount of underlying tokens need to be present in this contract, also unaccounted for.
    /// @param to Wallet receiving the minted liquidity tokens.
    /// @param remainder Wallet receiving any surplus underlying.
    /// @param minRatio Minimum ratio of underlying to PT in the pool.
    /// @param maxRatio Maximum ratio of underlying to PT in the pool.
    /// @return The amount of liquidity tokens minted.
    function mint(address to, address remainder, uint256 minRatio, uint256 maxRatio)
        external override
        returns (uint256, uint256, uint256)
    {
        return _mintInternal(to, remainder, 0, minRatio, maxRatio);
    }

    /// @dev Mint liquidity tokens in exchange for adding only underlying
    /// The amount of liquidity tokens is calculated from the amount of PT to buy from the pool,
    /// plus the amount of unaccounted for PT in this contract.
    /// The underlying tokens need to be present in this contract, unaccounted for.
    /// @param to Wallet receiving the minted liquidity tokens.
    /// @param remainder Wallet receiving any surplus underlying.
    /// @param PTToBuy Amount of `PT` being bought in the Pool, from this we calculate how much underlying it will be taken in.
    /// @param minRatio Minimum ratio of underlying to PT in the pool.
    /// @param maxRatio Maximum ratio of underlying to PT in the pool.
    /// @return The amount of liquidity tokens minted.
    function mintWithUnderlying(address to, address remainder, uint256 PTToBuy, uint256 minRatio, uint256 maxRatio)
        external override
        returns (uint256, uint256, uint256)
    {
        return _mintInternal(to, remainder, PTToBuy, minRatio, maxRatio);
    }

    /// @dev Mint liquidity tokens, with an optional internal trade to buy PT beforehand.
    /// The amount of liquidity tokens is calculated from the amount of PT to buy from the pool,
    /// plus the amount of unaccounted for PT in this contract.
    /// The underlying tokens need to be present in this contract, unaccounted for.
    /// @param to Wallet receiving the minted liquidity tokens.
    /// @param remainder Wallet receiving any surplus underlying.
    /// @param PTToBuy Amount of `PT` being bought in the Pool, from this we calculate how much underlying it will be taken in.
    /// @param minRatio Minimum ratio of underlying to PT in the pool.
    /// @param maxRatio Minimum ratio of underlying to PT in the pool.
    function _mintInternal(address to, address remainder, uint256 PTToBuy, uint256 minRatio, uint256 maxRatio)
        internal
        returns (uint256 underlyingIn, uint256 PTIn, uint256 tokensMinted)
    {
        // Gather data
        uint256 supply = _totalSupply;
        (uint112 _underlyingCached, uint112 _PTCached) =
            (underlyingCached, PTCached);
        uint256 _realPTCached = _PTCached - supply;    // The PT cache includes the virtual PT, equal to the supply
        uint256 underlyingBalance = underlying.balanceOf(address(this));
        uint256 PTBalance = PT.balanceOf(address(this));
        uint256 underlyingAvailable = underlyingBalance - _underlyingCached;

        // Check the burn wasn't sandwiched
        require (
            _realPTCached == 0 || (
                uint256(_underlyingCached) * 1e18 / _realPTCached >= minRatio &&
                uint256(_underlyingCached) * 1e18 / _realPTCached <= maxRatio
            ),
            "Pool: Reserves ratio changed"
        );

        // Calculate token amounts
        if (supply == 0) { // Initialize at 1 pool token minted per underlying token supplied
            underlyingIn = underlyingAvailable;
            tokensMinted = underlyingIn;
        } else if (_realPTCached == 0) { // Edge case, no PT in the Pool after initialization
            underlyingIn = underlyingAvailable;
            tokensMinted = supply * underlyingIn / _underlyingCached;
        } else {
            // There is an optional virtual trade before the mint
            uint256 underlyingToSell;
            if (PTToBuy > 0) {
                underlyingToSell = _buyPTPreview(
                    PTToBuy.u128(),
                    _underlyingCached,
                    _PTCached
                ); 
            }

            // We use all the available PTs, plus a virtual trade if it happened, surplus is in underlying tokens
            PTIn = PTBalance - _realPTCached;
            tokensMinted = (supply * (PTToBuy + PTIn)) / (_realPTCached - PTToBuy);
            underlyingIn = underlyingToSell + ((_underlyingCached + underlyingToSell) * tokensMinted) / supply;
            require(underlyingAvailable >= underlyingIn, "Pool: Not enough underlying token in");
        }

        // Update TWAR
        _update(
            (_underlyingCached + underlyingIn).u128(),
            (_PTCached + PTIn + tokensMinted).u128(), // Account for the "virtual" PT from the new minted LP tokens
            _underlyingCached,
            _PTCached
        );

        // Execute mint
        _mint(to, tokensMinted);

        // Return any unused underlying
        if (underlyingAvailable - underlyingIn > 0) underlying.safeTransfer(remainder, underlyingAvailable - underlyingIn);

        emit Liquidity(maturity, msg.sender, to, address(0), -(underlyingIn.i256()), -(PTIn.i256()), tokensMinted.i256());
    }

    /// @dev Burn liquidity tokens in exchange for underlying and PT.
    /// The liquidity tokens need to be in this contract.
    /// @param underlyingTo Wallet receiving the underlying.
    /// @param PTTo Wallet receiving the PT.
    /// @param minRatio Minimum ratio of underlying to PT in the pool.
    /// @param maxRatio Maximum ratio of underlying to PT in the pool.
    /// @return The amount of tokens burned and returned (tokensBurned, underlyings, PTs).
    function burn(address underlyingTo, address PTTo, uint256 minRatio, uint256 maxRatio)
        external override
        returns (uint256, uint256, uint256)
    {
        return _burnInternal(underlyingTo, PTTo, false, minRatio, maxRatio);
    }

    /// @dev Burn liquidity tokens in exchange for underlying.
    /// The liquidity provider needs to have called `pool.approve`.
    /// @param to Wallet receiving the underlying and PT.
    /// @param minRatio Minimum ratio of underlying to PT in the pool.
    /// @param maxRatio Minimum ratio of underlying to PT in the pool.
    /// @return tokensBurned The amount of lp tokens burned.
    /// @return underlyingOut The amount of underlying tokens returned.
    function burnForUnderlying(address to, uint256 minRatio, uint256 maxRatio)
        external override
        returns (uint256 tokensBurned, uint256 underlyingOut)
    {
        (tokensBurned, underlyingOut, ) = _burnInternal(to, address(0), true, minRatio, maxRatio);
    }


    /// @dev Burn liquidity tokens in exchange for underlying.
    /// The liquidity provider needs to have called `pool.approve`.
    /// @param underlyingTo Wallet receiving the underlying.
    /// @param PTTo Wallet receiving the PT.
    /// @param tradeTounderlying Whether the resulting PT should be traded for underlying tokens.
    /// @param minRatio Minimum ratio of underlying to PT in the pool.
    /// @param maxRatio Minimum ratio of underlying to PT in the pool.
    /// @return tokensBurned The amount of pool tokens burned.
    /// @return tokenOut The amount of underlying tokens returned.
    /// @return PTOut The amount of PTs returned.
    function _burnInternal(address underlyingTo, address PTTo, bool tradeTounderlying, uint256 minRatio, uint256 maxRatio)
        internal
        returns (uint256 tokensBurned, uint256 tokenOut, uint256 PTOut)
    {
        
        // Gather data
        tokensBurned = _balanceOf[address(this)];
        uint256 supply = _totalSupply;
        (uint112 _underlyingCached, uint112 _PTCached) =
            (underlyingCached, PTCached);
        uint256 _realPTCached = _PTCached - supply;    // The PT cache includes the virtual PT, equal to the supply

        // Check the burn wasn't sandwiched
        require (
            _realPTCached == 0 || (
                uint256(_underlyingCached) * 1e18 / _realPTCached >= minRatio &&
                uint256(_underlyingCached) * 1e18 / _realPTCached <= maxRatio
            ),
            "Pool: Reserves ratio changed"
        );

        // Calculate trade
        tokenOut = (tokensBurned * _underlyingCached) / supply;
        PTOut = (tokensBurned * _realPTCached) / supply;

        if (tradeTounderlying) {
            tokenOut += YieldMath.underlyingOutForPTIn(                      // This is a virtual sell
                (_underlyingCached - tokenOut.u128()) * scaleFactor,              // Cache, minus virtual burn
                (_PTCached - PTOut.u128()) * scaleFactor,         // Cache, minus virtual burn
                PTOut.u128() * scaleFactor,                            // Sell the virtual PT obtained
                maturity - uint32(block.timestamp),                         // This can't be called after maturity
                ts,
                g2
            ) / scaleFactor;
            PTOut = 0;
        }

        // Update TWAR
        _update(
            (_underlyingCached - tokenOut).u128(),
            (_PTCached - PTOut - tokensBurned).u128(),
            _underlyingCached,
            _PTCached
        );

        // Transfer assets
        _burn(address(this), tokensBurned);
        underlying.safeTransfer(underlyingTo, tokenOut);
        if (PTOut > 0) IERC20(address(PT)).safeTransfer(PTTo, PTOut);

        emit Liquidity(maturity, msg.sender, underlyingTo, PTTo, tokenOut.i256(), PTOut.i256(), -(tokensBurned.i256()));
    }

    // ---- Trading ----

    /// @dev Sell underlying for PT.
    /// The trader needs to have transferred the amount of underlying to sell to the pool before in the same transaction.
    /// @param to Wallet receiving the PT being bought
    /// @param min Minimm accepted amount of PT
    /// @return Amount of PT that will be deposited on `to` wallet
    function sellUnderlying(address to, uint128 min)
        external override
        returns(uint128)
    {
        // Calculate trade
        (uint112 _underlyingCached, uint112 _PTCached) =
            (underlyingCached, PTCached);
        uint112 _underlyingBalance = _getUnderlyingBalance();
        uint112 _PTBalance = _getPTBalance();
        uint128 underlyingIn = _underlyingBalance - _underlyingCached;
        uint128 PTOut = _sellunderlyingPreview(
            underlyingIn,
            _underlyingCached,
            _PTBalance
        );

        // Slippage check
        require(
            PTOut >= min,
            "Pool: Not enough PT obtained"
        );

        // Update TWAR
        _update(
            _underlyingBalance,
            _PTBalance - PTOut,
            _underlyingCached,
            _PTCached
        );

        // Transfer assets
        IERC20(address(PT)).safeTransfer(to, PTOut);

        emit Trade(maturity, msg.sender, to, -(underlyingIn.i128()), PTOut.i128());
        return PTOut;
    }

    /// @dev Returns how much PT would be obtained by selling `underlyingIn` underlying
    /// @param underlyingIn Amount of underlying hypothetically sold.
    /// @return Amount of PT hypothetically bought.
    function sellUnderlyingPreview(uint128 underlyingIn)
        external view override
        returns(uint128)
    {
        (uint112 _underlyingCached, uint112 _PTCached) =
            (underlyingCached, PTCached);
        return _sellunderlyingPreview(underlyingIn, _underlyingCached, _PTCached);
    }

    /// @dev Returns how much PT would be obtained by selling `underlyingIn` underlying
    function _sellunderlyingPreview(
        uint128 underlyingIn,
        uint112 underlyingBalance,
        uint112 PTBalance
    )
        private view
        beforeMaturity
        returns(uint128)
    {
        uint128 PTOut = YieldMath.PTOutForunderlyingIn(
            underlyingBalance * scaleFactor,
            PTBalance * scaleFactor,
            underlyingIn * scaleFactor,
            maturity - uint32(block.timestamp),             // This can't be called after maturity
            ts,
            g1
        ) / scaleFactor;

        require(
            PTBalance - PTOut >= underlyingBalance + underlyingIn,
            "Pool: PT balance too low"
        );

        return PTOut;
    }

    /// @dev Buy underlying for PT
    /// The trader needs to have called `PT.approve`
    /// @param to Wallet receiving the underlying being bought
    /// @param tokenOut Amount of underlying being bought that will be deposited in `to` wallet
    /// @param max Maximum amount of PT that will be paid for the trade
    /// @return Amount of PT that will be taken from caller
    function buyUnderlying(address to, uint128 tokenOut, uint128 max)
        external override
        returns(uint128)
    {
        // Calculate trade
        uint128 PTBalance = _getPTBalance();
        (uint112 _underlyingCached, uint112 _PTCached) =
            (underlyingCached, PTCached);
        uint128 PTIn = _buyunderlyingPreview(
            tokenOut,
            _underlyingCached,
            _PTCached
        );
        require(
            PTBalance - _PTCached >= PTIn,
            "Pool: Not enough PT in"
        );

        // Slippage check
        require(
            PTIn <= max,
            "Pool: Too much PT in"
        );

        // Update TWAR
        _update(
            _underlyingCached - tokenOut,
            _PTCached + PTIn,
            _underlyingCached,
            _PTCached
        );

        // Transfer assets
        underlying.safeTransfer(to, tokenOut);

        emit Trade(maturity, msg.sender, to, tokenOut.i128(), -(PTIn.i128()));
        return PTIn;
    }

    /// @dev Returns how much PT would be required to buy `tokenOut` underlying.
    /// @param tokenOut Amount of underlying hypothetically desired.
    /// @return Amount of PT hypothetically required.
    function buyUnderlyingPreview(uint128 tokenOut)
        external view override
        returns(uint128)
    {
        (uint112 _underlyingCached, uint112 _PTCached) =
            (underlyingCached, PTCached);
        return _buyunderlyingPreview(tokenOut, _underlyingCached, _PTCached);
    }

    /// @dev Returns how much PT would be required to buy `tokenOut` underlying.
    function _buyunderlyingPreview(
        uint128 tokenOut,
        uint112 underlyingBalance,
        uint112 PTBalance
    )
        private view
        beforeMaturity
        returns(uint128)
    {
        return YieldMath.PTInForunderlyingOut(
            underlyingBalance * scaleFactor,
            PTBalance * scaleFactor,
            tokenOut * scaleFactor,
            maturity - uint32(block.timestamp),             // This can't be called after maturity
            ts,
            g2
        ) / scaleFactor;
    }

    /// @dev Sell PT for underlying
    /// The trader needs to have transferred the amount of PT to sell to the pool before in the same transaction.
    /// @param to Wallet receiving the underlying being bought
    /// @param min Minimm accepted amount of underlying
    /// @return Amount of underlying that will be deposited on `to` wallet
    function sellPT(address to, uint128 min)
        external override
        returns(uint128)
    {
        // Calculate trade
        (uint112 _underlyingCached, uint112 _PTCached) =
            (underlyingCached, PTCached);
        uint112 _PTBalance = _getPTBalance();
        uint112 _underlyingBalance = _getUnderlyingBalance();
        uint128 PTIn = _PTBalance - _PTCached;
        uint128 underlyingOut = _sellPTPreview(
            PTIn,
            _underlyingCached,
            _PTCached
        );

        // Slippage check
        require(
            underlyingOut >= min,
            "Pool: Not enough underlying obtained"
        );

        // Update TWAR
        _update(
            _underlyingBalance - underlyingOut,
            _PTBalance,
            _underlyingCached,
            _PTCached
        );

        // Transfer assets
        underlying.safeTransfer(to, underlyingOut);

        emit Trade(maturity, msg.sender, to, underlyingOut.i128(), -(PTIn.i128()));
        return underlyingOut;
    }

    /// @dev Returns how much underlying would be obtained by selling `PTIn` PT.
    /// @param PTIn Amount of PT hypothetically sold.
    /// @return Amount of underlying hypothetically bought.
    function sellPTPreview(uint128 PTIn)
        external view override
        returns(uint128)
    {
        (uint112 _underlyingCached, uint112 _PTCached) =
            (underlyingCached, PTCached);
        return _sellPTPreview(PTIn, _underlyingCached, _PTCached);
    }

    /// @dev Returns how much underlying would be obtained by selling `PTIn` PT.
    function _sellPTPreview(
        uint128 PTIn,
        uint112 underlyingBalance,
        uint112 PTBalance
    )
        private view
        beforeMaturity
        returns(uint128)
    {
        return YieldMath.underlyingOutForPTIn(
            underlyingBalance * scaleFactor,
            PTBalance * scaleFactor,
            PTIn * scaleFactor,
            maturity - uint32(block.timestamp),             // This can't be called after maturity
            ts,
            g2
        ) / scaleFactor;
    }

    /// @dev Buy PT for underlying
    /// The trader needs to have called `underlying.approve`
    /// @param to Wallet receiving the PT being bought
    /// @param PTOut Amount of PT being bought that will be deposited in `to` wallet
    /// @param max Maximum amount of underlying token that will be paid for the trade
    /// @return Amount of underlying that will be taken from caller's wallet
    function buyPT(address to, uint128 PTOut, uint128 max)
        external override
        returns(uint128)
    {
        // Calculate trade
        uint128 underlyingBalance = _getUnderlyingBalance();
        (uint112 _underlyingCached, uint112 _PTCached) =
            (underlyingCached, PTCached);
        uint128 underlyingIn = _buyPTPreview(
            PTOut,
            _underlyingCached,
            _PTCached
        );
        require(
            underlyingBalance - _underlyingCached >= underlyingIn,
            "Pool: Not enough underlying token in"
        );

        // Slippage check
        require(
            underlyingIn <= max,
            "Pool: Too much underlying token in"
        );

        // Update TWAR
        _update(
            _underlyingCached + underlyingIn,
            _PTCached - PTOut,
            _underlyingCached,
            _PTCached
        );

        // Transfer assets
        IERC20(address(PT)).safeTransfer(to, PTOut);

        emit Trade(maturity, msg.sender, to, -(underlyingIn.i128()), PTOut.i128());
        return underlyingIn;
    }

    /// @dev Returns how much underlying would be required to buy `PTOut` PT.
    /// @param PTOut Amount of PT hypothetically desired.
    /// @return Amount of underlying hypothetically required.
    function buyPTPreview(uint128 PTOut)
        external view override
        returns(uint128)
    {
        (uint112 _underlyingCached, uint112 _PTCached) =
            (underlyingCached, PTCached);
        return _buyPTPreview(PTOut, _underlyingCached, _PTCached);
    }

    /// @dev Returns how much underlying would be required to buy `PTOut` PT.
    function _buyPTPreview(
        uint128 PTOut,
        uint128 underlyingBalance,
        uint128 PTBalance
    )
        private view
        beforeMaturity
        returns(uint128)
    {
        uint128 underlyingIn = YieldMath.underlyingInForPTOut(
            underlyingBalance * scaleFactor,
            PTBalance * scaleFactor,
            PTOut * scaleFactor,
            maturity - uint32(block.timestamp),             // This can't be called after maturity
            ts,
            g1
        ) / scaleFactor;

        require(
            PTBalance - PTOut >= underlyingBalance + underlyingIn,
            "Pool: PT balance too low"
        );

        return underlyingIn;
    }
}