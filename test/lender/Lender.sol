// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Interfaces.sol';
import './MarketPlace.sol'; // library of market place specific constructs
import './Swivel.sol'; // library of swivel specific constructs
import './Element.sol'; // library of element specific constructs
import './Safe.sol';
import './Cast.sol';

contract Lender {
    address public admin;
    address public marketPlace;

    /// @dev addresses of the 3rd party protocol contracts
    address public swivelAddr;
    address public pendleAddr;
    address public tempusAddr;

    uint256 public feenominator;

    /// @notice maps underlying tokens to the amount of fees accumulated for that token
    mapping(address => uint256) public fees;

    event Lend(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 returned);
    event Mint(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount);

    /// @notice Initializes the Lender contract
    /// @dev the ctor sets a default value for the feenominator
    /// @param s: the swivel contract
    /// @param p: the pendle contract
    /// @param t: the tempus contract
    constructor(
        address s,
        address p,
        address t
    ) {
        admin = msg.sender;
        swivelAddr = s;
        pendleAddr = p;
        tempusAddr = t;
        feenominator = 1000;
    }

    /// @notice Allows the admin to bulk approve given compound addresses at the underlying token, saving marginal approvals
    /// @param u array of underlying token addresses
    /// @param c array of compound token addresses
    function approveUnderlying(address[] calldata u, address[] calldata c) external authorized(admin) returns (bool) {
        uint256 len = u.length;
        require(len == c.length, 'array length mismatch');

        uint256 max = 2**256 - 1;

        for (uint256 i; i < len; i++) {
            IErc20 uToken = IErc20(u[i]);
            Safe.approve(uToken, c[i], max);
        }

        return true;
    }

    /// @notice Sets the feenominator to the given value
    /// @param f: the new value of the feenominator, fees are not collected when the feenominator is 0
    /// @return bool true if successful
    function setFee(uint256 f) external authorized(admin) returns (bool) {
        feenominator = f;
        return true;
    }

    /// @notice Sets the address of the marketplace contract which contains the
    /// addresses of all the fixed rate markets
    /// @param m: the address of the marketplace contract
    /// @return bool true if the address was set, false otherwise
    function setMarketPlaceAddress(address m) external authorized(admin) returns (bool) {
        require(marketPlace == address(0));
        marketPlace = m;
        return true;
    }

    /// @notice mint swaps the sender's principal tokens for illuminate's zc tokens
    /// @notice in effect, this opens a new fixed rate position for the sender on illuminate
    /// @dev mint is uniform across all principals, thus there is no need for a 'minter'
    /// @param p value of a specific principal according to the MarketPlace Principals Enum
    /// @param u address of an underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a amount being minted
    /// @return bool true if the mint was successful, false otherwise
    function mint(
        uint8 p,
        address u,
        uint256 m,
        uint256 a
    ) public returns (bool) {
        //use market interface to fetch the market for the given market pair
        address principal = IMarketPlace(marketPlace).markets(u, m, p);
        //use safe transfer lib and ERC interface...
        Safe.transferFrom(IErc20(principal), msg.sender, address(this), a);
        //use zctoken interface...
        IZcToken(zcToken(u, m)).mint(msg.sender, a);

        emit Mint(p, u, m, a);

        return true;
    }

    /// @dev lend method signature for both illuminate and yield
    /// @param p value of a specific principal according to the MarketPlace Principals Enum
    /// @param u address of an underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a amount of underlying tokens to lend
    /// @param y yield pool that will execute the swap for the principal token
    /// @return uint256 the amount of principal tokens lent out
    function lend(
        uint8 p,
        address u,
        uint256 m,
        uint256 a,
        address y
    ) public returns (uint256) {
        // check the principal is illuminate or yield
        require(p == uint8(MarketPlace.Principals.Illuminate) || p == uint8(MarketPlace.Principals.Yield));

        // uses yield token interface...
        IYield pool = IYield(y);

        // the yield token must match the market pair
        require(address(pool.base()) == u, 'yield base != underlying');
        require(pool.maturity() == m, 'yield maturity != maturity');

        // transfer from user to illuminate
        Safe.transferFrom(IErc20(u), msg.sender, address(this), a);

        uint256 returned = yield(u, y, a - calculateFee(a));

        // this step is only needed when the lend is for yield
        if (p == uint8(MarketPlace.Principals.Yield)) {
            // TODO should we require on this?
            IZcToken(zcToken(u, m)).mint(msg.sender, returned);
        }

        emit Lend(p, u, m, returned);

        return returned;
    }

    /// @notice lends to yield pool. remaining balance is sent to the yield pool
    /// TODO: this will change when we implement a check on the gas market
    /// @dev lend method signature for swivel
    /// @param p value of a specific principal according to the Illuminate Principals Enum
    /// @param u address of an underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a array of amounts of underlying tokens lent to each order in the orders array
    /// @param y yield pool
    /// @param o array of swivel orders being filled
    /// @param s array of signatures for each order in the orders array
    /// @return uint256 the amount of principal tokens lent out
    function lend(
        uint8 p,
        address u,
        uint256 m,
        uint256[] calldata a,
        address y,
        Swivel.Order[] calldata o,
        Swivel.Components[] calldata s
    ) public returns (uint256) {
        // check the principal is swivel
        require(p == uint8(MarketPlace.Principals.Swivel));

        // lent represents the number of underlying tokens lent
        uint256 lent;
        // returned represents the number of underlying tokens to lend to yield
        uint256 returned;

        {
            uint256 totalFee;
            // iterate through each order a calculate the total lent and returned
            for (uint256 i = 0; i < o.length; i++) {
                Swivel.Order memory order = o[i];
                // Require the Swivel order provided matches the underlying and maturity market provided
                require(order.maturity == m, 'swivel maturity != maturity');
                require(order.underlying == u, 'swivel underlying != underlying');
                // Determine the fee
                uint256 fee = calculateFee(a[i]);
                // Track accumulated fees
                totalFee += fee;
                // Sum the total amount lent to Swivel (amount of zc tokens to mint) minus fees
                lent += a[i] - fee;
                // Sum the total amount of premium paid from Swivel (amount of underlying to lend to yield)
                returned += (a[i] - fee) * (order.premium / order.principal);
            }
            // Track accumulated fee
            fees[u] += totalFee;

            // transfer underlying tokens from user to illuminate
            Safe.transferFrom(IErc20(u), msg.sender, address(this), lent);
            // fill the orders on swivel protocol
            ISwivel(swivelAddr).initiate(o, a, s);

            yield(u, y, returned);
        }

        emit Lend(p, u, m, lent);
        return lent;
    }

    /// @dev lend method signature for element
    /// @param p value of a specific principal according to the Illuminate Principals Enum
    /// @param u address of an underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a amount of principal tokens to lend
    /// @param r minimum amount to return, this puts a cap on allowed slippage
    /// @param d deadline is a timestamp by which the swap must be executed deadline is a timestamp by which the swap must be executed
    /// @param e element pool that is lent to
    /// @param i the id of the pool
    function lend(
        uint8 p,
        address u,
        uint256 m,
        uint256 a,
        uint256 r,
        uint256 d,
        address e,
        bytes32 i
    ) public returns (uint256) {
        // check the principal is element
        require(p == uint8(MarketPlace.Principals.Element));
        // Get the principal token for this market for element
        address principal = IMarketPlace(marketPlace).markets(u, m, p);

        // the element token must match the market pair
        require(IElementToken(principal).underlying() == u, 'element underlying != underlying');
        require(IElementToken(principal).unlockTimestamp() == m, 'element maturity != maturity');

        // Transfer underlying token from user to illuminate
        Safe.transferFrom(IErc20(u), msg.sender, address(this), a);

        // Track the accumulated fees
        fees[u] += calculateFee(a);

        // Create the variables needed to execute an element swap
        Element.FundManagement memory fund = Element.FundManagement({
            sender: address(this),
            recipient: payable(address(this)),
            fromInternalBalance: false,
            toInternalBalance: false
        });

        Element.SingleSwap memory swap = Element.SingleSwap({
            userData: '0x00000000000000000000000000000000000000000000000000000000000000',
            poolId: i,
            amount: a - calculateFee(a),
            kind: Element.SwapKind.In,
            assetIn: Any(u),
            assetOut: Any(principal)
        });

        // Conduct the swap on element
        uint256 purchased = IElement(e).swap(swap, fund, r, d);

        emit Lend(p, u, m, purchased);
        return purchased;
    }

    /// @dev lend method signature for pendle
    /// @param p value of a specific principal according to the Illuminate Principals Enum
    /// @param u address of an underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a amount of principal tokens to lend
    /// @param r minimum amount to return, this puts a cap on allowed slippage
    /// @param d deadline is a timestamp by which the swap must be executed
    /// @return uint256 the amount of principal tokens lent out
    function lend(
        uint8 p,
        address u,
        uint256 m,
        uint256 a,
        uint256 r,
        uint256 d
    ) public returns (uint256) {
        // check the principal is pendle
        require(p == uint8(MarketPlace.Principals.Pendle));
        // Instantiate market and tokens
        address principal = IMarketPlace(marketPlace).markets(u, m, p);
        IPendleToken token = IPendleToken(principal);

        // confirm that we are in the correct market
        require(token.yieldToken() == u, 'pendle underlying != underlying');
        require(token.expiry() == m, 'pendle maturity != maturity');

        // Transfer funds from user to Illuminate
        Safe.transferFrom(IErc20(u), msg.sender, address(this), a);

        // Add the accumulated fees to the total
        fees[u] += calculateFee(a);

        address[] memory path = new address[](2);
        path[0] = u;
        path[1] = principal;

        // Swap on the Pendle Router using the provided market and params
        uint256 returned = IPendle(pendleAddr).swapExactTokensForTokens(a - calculateFee(a), r, path, address(this), d)[0];

        // Mint Illuminate zero coupons
        address illuminateToken = zcToken(u, m);
        IZcToken(illuminateToken).mint(msg.sender, returned);

        emit Lend(p, u, m, returned);

        return returned;
    }

    /// @dev lend method signature for tempus
    /// @notice Can be called before maturity to lend to Tempus while minting Illuminate tokens
    /// @param p value of a specific principal according to the Illuminate Principals Enum
    /// @param u address of an underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a amount of principal tokens to lend
    /// @param r minimum amount to return when executing the swap (sets a limit to slippage)
    /// @param d deadline is a timestamp by which the swap must be executed
    /// @param t tempus pool that houses the underlying principal tokens
    /// @param x tempus amm that executes the swap
    /// @return uint256 the amount of principal tokens lent out
    function lend(
        uint8 p,
        address u,
        uint256 m,
        uint256 a,
        uint256 r,
        uint256 d,
        address t,
        address x
    ) public returns (uint256) {
        // check the principal is tempus
        require(p == uint8(MarketPlace.Principals.Tempus));

        // Instantiate market and tokens
        address principal = IMarketPlace(marketPlace).markets(u, m, p);
        require(ITempus(principal).yieldBearingToken() == IErc20Metadata(u), 'tempus underlying != underlying');
        require(ITempus(principal).maturityTime() == m, 'tempus maturity != maturity');

        // Get the underlying token
        IErc20 underlyingToken = IErc20(u);

        // Transfer funds from user to Illuminate, Scope to avoid stack limit
        Safe.transferFrom(underlyingToken, msg.sender, address(this), a);

        // Add the accumulated fees to the total
        fees[u] += calculateFee(a);

        // Swap on the Tempus Router using the provided market and params
        IZcToken illuminateToken = IZcToken(zcToken(u, m));
        uint256 returned = ITempus(tempusAddr).depositAndFix(Any(x), Any(t), a - calculateFee(a), true, r, d) -
            illuminateToken.balanceOf(address(this));

        // Mint Illuminate zero coupons
        illuminateToken.mint(msg.sender, returned);

        emit Lend(p, u, m, returned);

        return returned;
    }

    /// @dev lend method signature for sense
    /// @notice Can be called before maturity to lend to Sense while minting Illuminate tokens
    /// @param p value of a specific principal according to the Illuminate Principals Enum
    /// @param u address of an underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a amount of underlying tokens to lend
    /// @param r minimum number of tokens to lend (sets a limit on the order)
    /// @param x amm that is used to conduct the swap
    /// @param s contract that holds the principal token for this market
    /// @return uint256 the amount of principal tokens lent out
    function lend(
        uint8 p,
        address u,
        uint256 m,
        uint128 a,
        uint256 r,
        address x,
        address s
    ) public returns (uint256) {
        // check the principal is sense
        require(p == uint8(MarketPlace.Principals.Sense));

        // Get the principal token for this market for this market
        ISenseToken token = ISenseToken(IMarketPlace(marketPlace).markets(u, m, p));

        // Verify that the underlying matches up
        require(token.underlying() == u, 'sense underlying != underlying');

        // Determine the fee
        uint256 fee = calculateFee(a);

        // Add the accumulated fees to the total
        fees[u] += fee;

        // Determine lent amount after fees
        uint256 lent = a - fee;

        // Transfer funds from user to Illuminate
        Safe.transferFrom(IErc20(u), msg.sender, address(this), a);

        // Swap those tokens for the principal tokens
        uint256 returned = ISense(x).swapUnderlyingForPTs(s, m, lent, r);

        // Get the address of the ZC token for this market
        IZcToken illuminateToken = IZcToken(zcToken(u, m));

        // Mint the illuminate tokens based on the returned amount
        illuminateToken.mint(msg.sender, returned);

        emit Lend(p, u, m, returned);

        return returned;
    }

    /// @notice Can be called before maturity to lend to APWine while minting Illuminate tokens
    /// @param p value of a specific principal according to the Illuminate Principals Enum
    /// @param u address of an underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a the amount of underlying tokens to lend
    /// @param r the minimum amount of zero-coupon tokens to return accounting for slippage
    /// @param pool the address of a given APWine pool
    /// @param i the id of the pool
    /// @return uint256 the amount of principal tokens lent out
    function lend(
        uint8 p,
        address u,
        uint256 m,
        uint256 a,
        uint256 r,
        address pool,
        address aave,
        uint256 i
    ) public returns (uint256) {
        // check the principal is apwine
        require(p == uint8(MarketPlace.Principals.Apwine));

        // Instantiate market and tokens
        address principal = IMarketPlace(marketPlace).markets(u, m, p);
        require(IAPWineToken(principal).getPTAddress() == u, 'apwine principle != principle');

        // Transfer funds from user to Illuminate
        Safe.transferFrom(IErc20(u), msg.sender, address(this), a);

        // Determine the fee
        uint256 fee = calculateFee(a);

        // Add the accumulated fees to the total
        fees[u] += fee;

        // Determine the amount lent after fees
        uint256 lent = a - fee;

        // Deposit into aave
        // TODO: Get a referral code and use it to deposit
        IAave(aave).deposit(u, lent, address(this), 0);
        require(1 == 2, 'alkjsdf');

        // Swap on the APWine Pool using the provided market and params
        uint256 returned = IAPWineRouter(pool).swapExactAmountIn(i, 1, lent, 0, r, address(this));

        // Mint Illuminate zero coupons
        IZcToken(zcToken(u, m)).mint(msg.sender, returned);

        emit Lend(p, u, m, returned);

        return returned;
    }

    /// @dev lend method signature for Notional
    /// @param p value of a specific principal according to the Illuminate Principals Enum
    /// @param u address of an underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a amount of principal tokens to lend
    /// @return uint256 the amount of principal tokens lent out
    function lend(
        uint8 p,
        address u,
        uint256 m,
        uint256 a
    ) public returns (uint256) {
        // check the principal is notional
        require(p == uint8(MarketPlace.Principals.Notional));

        // Instantiate market and tokens
        address principal = IMarketPlace(marketPlace).markets(u, m, p);

        INotional token = INotional(principal);

        // Verify that the maturity matches up
        require(token.getMaturity() == m, 'notional maturity != maturity');
        // Verify that the underlying matches up
        (IErc20 underlying, ) = token.getUnderlyingToken();
        require(address(underlying) == u, 'notional underlying != underlying');

        // Transfer funds from user to Illuminate
        Safe.transferFrom(IErc20(u), msg.sender, address(this), a);

        // Add the accumulated fees to the total
        uint256 fee = calculateFee(a);
        fees[u] += fee;

        // Swap on the Notional Token wrapper
        uint256 returned = token.deposit(a - fee, address(this));

        // Mint Illuminate zero coupons
        address illuminateToken = zcToken(u, m);
        IZcToken(illuminateToken).mint(msg.sender, returned);

        emit Lend(p, u, m, returned);

        return returned;
    }

    /// @notice transfers excess funds to yield pool after principal tokens have been lent out
    /// @dev this method is only used by the yield, illuminate and swivel protocols
    /// @param u address of an underlying asset
    /// @param y: the yield pool to lend to
    /// @param a: the amount of underlying tokens to lend
    /// @return uint256 the amount of tokens sent to the yield pool
    function yield(
        address u,
        address y,
        uint256 a
    ) internal returns (uint256) {
        // TODO: This is definitely a bug. The returned variable should be returned
        // by the method, but the original amount should still be sent to the pool.

        // preview exact swap slippage on yield
        uint128 returned = IYield(y).sellBasePreview(Cast.u128(a));

        // send the remaing amount to the given yield pool
        Safe.transfer(IErc20(u), y, returned);

        // lend out the remaining tokens in the yield pool
        IYield(y).sellBase(address(this), returned);

        return returned;
    }

    /// @notice Withdraws accumulated lending fees of the underlying token
    /// @param e Address of the underlying token to withdraw
    /// @return bool true if successful
    function withdraw(address e) external authorized(admin) returns (bool) {
        // Get the token to be withdrawn
        IErc20 token = IErc20(e);

        // Get the balance to be transferred
        uint256 balance = fees[e];

        // Reset accumulated fees of the token to 0
        fees[e] = 0;

        // Transfer the accumulated fees to the admin
        Safe.transfer(token, admin, balance);

        return true;
    }

    /// @notice This method returns the fee based on the amount passed to it. If
    /// feenominator is 0, then there is no fee.
    /// @param a Amount of underlying tokens to calculate the fee for
    /// @return uint256 The total for for the given amount
    function calculateFee(uint256 a) internal view returns (uint256) {
        return feenominator > 0 ? a / feenominator : 0;
    }

    /// @notice retrieves the zc token for the given market
    /// @param u address of the underlying
    /// @param m uint256 representing the maturity of the market
    /// @return address of the zc token for the market
    function zcToken(address u, uint256 m) internal returns (address) {
        return IMarketPlace(marketPlace).markets(u, m, 0);
    }

    modifier authorized(address a) {
        require(msg.sender == a, 'sender must be authorized');
        _;
    }
}
