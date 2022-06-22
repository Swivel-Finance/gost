// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Interfaces.sol';
import './MarketPlace.sol'; // library of market place specific constructs
import './Swivel.sol'; // library of swivel specific constructs
import './Element.sol'; // library of element specific constructs
import './Safe.sol';
import './Cast.sol';

/// @title Lender.sol
/// @author Sourabh Marathe, Julian Traversa, Rob Robbins
/// @notice The lender contract executes loans on behalf of users. 
/// @notice The contract holds the principal tokens for each market and mints an ERC-5095 position to users to represent their lent positions.
contract Lender {
    error Unauthorized();
    error NotEqual(string);
    error Exists(address);
    error Invalid(string);

    /// @notice minimum amount of time the admin must wait before executing a withdrawl
    uint256 constant public HOLD = 3 days;

    /// @notice address that is allowed to create markets, set fees, etc. It is commonly used in the authorized modifier.
    address public admin;
    /// @notice address of the MarketPlace.sol contract, used to access the markets mapping
    address public marketPlace;
    /// @notice mapping that determines if a principal may be used by a lender
    mapping(uint8 => bool) public paused;

    /// @notice third party contract needed to lend on Swivel
    address public swivelAddr;
    /// @notice third party contract needed to lend on Pendle
    address public immutable pendleAddr;
    /// @notice third party contract needed to lend on Tempus
    address public immutable tempusAddr;

    /// @notice this value determines the amount of fees paid on loans
    uint256 public feenominator;

    /// @notice maps underlying tokens to the amount of fees accumulated for that token
    mapping(address => uint256) public fees;
    /// @notice maps a token address to a point in time, a hold, after which a withdrawal can be made
    mapping (address => uint256) public withdrawals;

    /// @notice emitted upon executed lend
    event Lend(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 returned);
    /// @notice emitted upon minted ERC5095 to user
    event Mint(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount);
    /// @notice emitted on token withdrawal scheduling
    event ScheduleWithdrawal(address indexed token, uint256 hold);
    /// @notice emitted on token withdrawal blocking
    event BlockWithdrawal(address indexed token);
    /// @notice emitted on a change to the feenominators array

    /// @notice initializes the Lender contract
    /// @param s the swivel contract
    /// @param p the pendle contract
    /// @param t the tempus contract
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

    /// @notice approves the redeemer contract to spend the principal tokens held by the lender contract.
    /// @param u underlying token's address, used to define the market being approved
    /// @param m maturity of the underlying token, used to define the market being approved
    /// @param r the address being approved, in this case the redeemer contract
    /// @return bool true if the approval was successful, false otherwise
    function approve(
        address u,
        uint256 m,
        address r
    ) external authorized(admin) returns (bool) {
        // max is the maximum integer value for a 256 unsighed integer
        uint256 max = 2**256 - 1;

        // approve the underlying for max per given principal
        for (uint8 i; i < 9; ) {
            // get the principal token's address
            address token = IMarketPlace(marketPlace).markets(u, m, i);
            // check that the token is defined for this particular market
            if (token != address(0)) {
                // max approve the token
                Safe.approve(IERC20(token), r, max);
            }
            unchecked {
                i++;
            }
        }
        return true;
    }

    /// @notice bulk approves the usage of addresses at the given ERC20 addresses. 
    /// @dev the lengths of the inputs must match because the arrays are paired by index
    /// @param u array of ERC20 token addresses that will be approved on
    /// @param a array of addresses that will be approved
    /// @return true if successful
    function approve(address[] calldata u, address[] calldata a) external authorized(admin) returns (bool) {
        uint256 len = u.length;
        if (len != a.length) {
            revert NotEqual('array length');
        }
        uint256 max = 2**256 - 1;

        for (uint256 i; i < len; ) {
            IERC20 uToken = IERC20(u[i]);
            if (address(0) != (address(uToken))) {
                Safe.approve(uToken, a[i], max);
            }
            unchecked {
                i++;
            }
        }
        return true;
    }

    /// @notice sets the admin address
    /// @param a address of a new admin
    /// @return bool true if successful
    function setAdmin(address a) external authorized(admin) returns (bool) {
        admin = a;
        return true;
    }

    /// @notice sets the feenominator to the given value
    /// @param f the new value of the feenominator, fees are not collected when the feenominator is 0
    /// @return bool true if successful
    function setFee(uint256 f) external authorized(admin) returns (bool) {
        feenominator = f;
        return true;
    }

    /// @notice sets the address of the marketplace contract which contains the addresses of all the fixed rate markets
    /// @param m the address of the marketplace contract
    /// @return bool true if the address was set, false otherwise
    function setMarketPlace(address m) external authorized(admin) returns (bool) {
        if (marketPlace != address(0)) {
            revert Exists(marketPlace);
        }
        marketPlace = m;
        return true;
    }

    /// @notice sets the feenominator to the given value
    /// @param s the address of the Swivel.sol Router
    /// @return bool true if successful
    function setSwivel(address s) external authorized(admin) returns (bool) {
        swivelAddr = s;
        return true;
    }

    /// @notice mint swaps the sender's principal tokens for illuminate's ERC5095 tokens in effect, this opens a new fixed rate position for the sender on illuminate
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
        Safe.transferFrom(IERC20(principal), msg.sender, address(this), a);
        //use ERC5095 interface...
        IERC5095(principalToken(u, m)).mint(msg.sender, a);

        emit Mint(p, u, m, a);

        return true;
    }

    /// @notice lend method signature for both illuminate and yield
    /// @param p value of a specific principal according to the MarketPlace Principals Enum
    /// @param u address of an underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a amount of underlying tokens to lend
    /// @param y yieldspace pool that will execute the swap for the principal token
    /// @return uint256 the amount of principal tokens lent out
    function lend(
        uint8 p,
        address u,
        uint256 m,
        uint256 a,
        address y
    ) public unpaused(p) returns (uint256) {
        // check the principal is illuminate or yield
        if (p != uint8(MarketPlace.Principals.Illuminate) && p != uint8(MarketPlace.Principals.Yield)) {
            revert Invalid('principal');
        }

        // uses yield token interface...
        IYield pool = IYield(y);

        // the yield token must match the market pair
        if (address(pool.base()) != u) {
            revert NotEqual('underlying');
        } else if (pool.maturity() > m) {
            revert NotEqual('maturity');
        }

        // transfer from user to illuminate
        Safe.transferFrom(IERC20(u), msg.sender, address(this), a);

        if (p == uint8(MarketPlace.Principals.Yield)) {
            // Purchase yield PTs to lender.sol (address(this))
            uint256 returned = yield(u, y, a - calculateFee(a), address(this));
            // Mint and distribute equivalent illuminate PTs
            IERC5095(principalToken(u, m)).mint(msg.sender, returned);
            
            emit Lend(p, u, m, returned);

            return returned;
        }
        else {
            // Purchase illuminate PTs directly to msg.sender
            uint256 returned = yield(u, y, a - calculateFee(a), msg.sender);

            emit Lend(p, u, m, returned);

            return returned;
        }
    }

    /// @notice lend method signature for swivel
    /// @dev lends to yield pool. remaining balance is sent to the yield pool
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
    ) public unpaused(p) returns (uint256) {

        // lent represents the number of underlying tokens lent
        uint256 lent;
        // returned represents the number of underlying tokens to lend to yield
        uint256 returned;

        {
            uint256 totalFee;
            // iterate through each order a calculate the total lent and returned
            for (uint256 i = 0; i < o.length; ) {
                Swivel.Order memory order = o[i];
                // Require the Swivel order provided matches the underlying and maturity market provided
                if (order.underlying != u) {
                    revert NotEqual('underlying');
                } else if (order.maturity > m) {
                    revert NotEqual('maturity');
                }
                // Determine the fee
                uint256 fee = calculateFee(a[i]);
                // Track accumulated fees
                totalFee += fee;
                // Sum the total amount lent to Swivel (amount of ERC5095 tokens to mint) minus fees
                lent += a[i] - fee;
                // Sum the total amount of premium paid from Swivel (amount of underlying to lend to yield)
                returned += (a[i] - fee) * (order.premium / order.principal);

                unchecked {
                    i++;
                }
            }
            // Track accumulated fee
            fees[u] += totalFee;

            // transfer underlying tokens from user to illuminate
            Safe.transferFrom(IERC20(u), msg.sender, address(this), lent);
            // fill the orders on swivel protocol
            ISwivel(swivelAddr).initiate(o, a, s);

            yield(u, y, returned, address(this));
        }

        emit Lend(p, u, m, lent);
        return lent;
    }

    /// @notice lend method signature for element
    /// @param p value of a specific principal according to the Illuminate Principals Enum
    /// @param u address of an underlying asset
    /// @param m maturity (timestamp) of the market
    /// @param a amount of principal tokens to lend
    /// @param r minimum amount to return, this puts a cap on allowed slippage
    /// @param d deadline is a timestamp by which the swap must be executed deadline is a timestamp by which the swap must be executed
    /// @param e element pool that is lent to
    /// @param i the id of the pool
    /// @return uint256 the amount of principal tokens lent out
    // function lend(
    //     uint8 p,
    //     address u,
    //     uint256 m,
    //     uint256 a,
    //     uint256 r,
    //     uint256 d,
    //     address e,
    //     bytes32 i
    // ) public unpaused(p) returns (uint256) {
    //     // Get the principal token for this market for element
    //     address principal = IMarketPlace(marketPlace).markets(u, m, p);

    //     // the element token must match the market pair
    //     if (IElementToken(principal).underlying() != u) {
    //         revert NotEqual('underlying');
    //     } else if (IElementToken(principal).unlockTimestamp() > m) {
    //         revert NotEqual('maturity');
    //     }
    //     // Transfer underlying token from user to illuminate
    //     Safe.transferFrom(IERC20(u), msg.sender, address(this), a);

    //     // Track the accumulated fees
    //     fees[u] += calculateFee(a);

    //     // Create the variables needed to execute an element swap
    //     Element.FundManagement memory fund = Element.FundManagement({
    //         sender: address(this),
    //         recipient: payable(address(this)),
    //         fromInternalBalance: false,
    //         toInternalBalance: false
    //     });

    //     Element.SingleSwap memory swap = Element.SingleSwap({
    //         userData: '0x00000000000000000000000000000000000000000000000000000000000000',
    //         poolId: i,
    //         amount: a - calculateFee(a),
    //         kind: Element.SwapKind.In,
    //         assetIn: Any(u),
    //         assetOut: Any(principal)
    //     });

    //     // Conduct the swap on element
    //     uint256 purchased = IElement(e).swap(swap, fund, r, d);

    //     emit Lend(p, u, m, purchased);
    //     return purchased;
    // }

    /// @notice lend method signature for pendle
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
    ) public unpaused(p) returns (uint256) {

        // Instantiate market and tokens
        address principal = IMarketPlace(marketPlace).markets(u, m, p);
        IPendleToken token = IPendleToken(principal);

        // confirm that we are in the correct market
        if (token.yieldToken() != u) {
            revert NotEqual('underlying');
        } else if (token.expiry() > m) {
            revert NotEqual('maturity');
        }

        // Transfer funds from user to Illuminate
        Safe.transferFrom(IERC20(u), msg.sender, address(this), a);

        uint256 returned;
        {
            // Add the accumulated fees to the total
            uint256 fee = calculateFee(a);
            fees[u] += fee;

            address[] memory path = new address[](2);
            path[0] = u;
            path[1] = principal;

            // Swap on the Pendle Router using the provided market and params
            returned = IPendle(pendleAddr).swapExactTokensForTokens(a - fee, r, path, address(this), d)[0];

        }
        // Mint Illuminate zero coupons
        address illuminateToken = principalToken(u, m);
        IERC5095(illuminateToken).mint(msg.sender, returned);

        emit Lend(p, u, m, returned);
        return returned;
    }

    /// @notice lend method signature for tempus
    /// @dev This method can be called before maturity to lend to Tempus while minting Illuminate tokens
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
    ) public unpaused(p) returns (uint256) {

        // Instantiate market and tokens
        address principal = IMarketPlace(marketPlace).markets(u, m, p);
        if (ITempus(principal).yieldBearingToken() != IERC20Metadata(u)) {
            revert NotEqual('underlying');
        } else if (ITempus(principal).maturityTime() > m) {
            revert NotEqual('maturity');
        }

        // Get the underlying token
        IERC20 underlyingToken = IERC20(u);

        // Transfer funds from user to Illuminate, Scope to avoid stack limit
        Safe.transferFrom(underlyingToken, msg.sender, address(this), a);

        uint256 returned;
        {
        // Add the accumulated fees to the total
        uint256 fee = calculateFee(a);
        fees[u] += fee;

        // Swap on the Tempus Router using the provided market and params
        IERC5095 illuminateToken = IERC5095(principalToken(u, m));
        uint256 returned = ITempus(tempusAddr).depositAndFix(Any(x), Any(t), a - fee, true, r, d) -
            illuminateToken.balanceOf(address(this));

        // Mint Illuminate zero coupons
        illuminateToken.mint(msg.sender, returned);

        emit Lend(p, u, m, returned);
        return returned;
    }

    /// @notice lend method signature for sense
    /// @dev this method can be called before maturity to lend to Sense while minting Illuminate tokens
    /// @dev sense provides a [divider] contract that splits [target] assets (underlying) into PTs and YTs. Each [target] asset has a [series] of contracts, each identifiable by their [maturity].
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
    ) public unpaused(p) returns (uint256) {

        // Get the principal token for this market for this market
        ISenseToken token = ISenseToken(IMarketPlace(marketPlace).markets(u, m, p));

        // Verify that the underlying and maturity match up
        if (token.underlying() != u) { // gauruntee the input token is the right token
            revert NotEqual('underlying'); 
        } else if (ISense(s).pt() != address(token)) {
            revert NotEqual('principal token'); 
        } else if (ISense(x).maturity() > m) { // gauruntee the input amm has the correct maturity
            revert NotEqual('maturity');
        }

        uint256 lent;
        {
            // Determine the fee
            uint256 fee = calculateFee(a);

            // Add the accumulated fees to the total
            fees[u] += fee;

            // Determine lent amount after fees
            lent = a - fee;
        }

        // Transfer funds from user to Illuminate
        Safe.transferFrom(IERC20(u), msg.sender, address(this), a);

        // Swap those tokens for the principal tokens
        uint256 returned = ISense(x).swapUnderlyingForPTs(s, m, lent, r);

        // Get the address of the ERC5095 token for this market
        IERC5095 illuminateToken = IERC5095(principalToken(u, m));

        // Mint the illuminate tokens based on the returned amount
        illuminateToken.mint(msg.sender, returned);

        emit Lend(p, u, m, returned);
        return returned;
    }

    /// @notice this method can be called before maturity to lend to APWine while minting Illuminate tokens
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
    ) public unpaused(p) returns (uint256) {
        // Instantiate market and tokens
        address principal = IMarketPlace(marketPlace).markets(u, m, p);
        if (IAPWineToken(principal).getUnderlyingOfIBTAddress() != u) {
            revert NotEqual('underlying');
        }
        // Dont necessarily need to validate APWINE maturity (They have 1 maturity per underlying)
        // Potentially add redundant implied maturity calculation

        // Transfer funds from user to Illuminate
        Safe.transferFrom(IERC20(u), msg.sender, address(this), a);

        // Determine the fee
        uint256 fee = calculateFee(a);

        // Add the accumulated fees to the total
        fees[u] += fee;

        // Determine the amount lent after fees
        uint256 lent = a - fee;

        // Deposit into aave
        IAave(aave).deposit(u, lent, address(this), 0);

        // Swap on the APWine Pool using the provided market and params
        uint256 returned = IAPWineRouter(pool).swapExactAmountIn(i, 1, lent, 0, r, address(this));

        // Mint Illuminate zero coupons
        IERC5095(principalToken(u, m)).mint(msg.sender, returned);

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
    ) public unpaused(p) returns (uint256) {
        // Instantiate market and tokens
        address principal = IMarketPlace(marketPlace).markets(u, m, p);

        INotional token = INotional(principal);

        // Verify that the underlying and maturity match up
        (IERC20 underlying, ) = token.getUnderlyingToken();
        if (address(underlying) != u) {
            revert NotEqual('underlying');
        } else if (token.getMaturity() > m) {
            revert NotEqual('maturity');
        }

        // Transfer funds from user to Illuminate
        Safe.transferFrom(IERC20(u), msg.sender, address(this), a);

        // Add the accumulated fees to the total
        uint256 fee = calculateFee(a);
        fees[u] += fee;

        // Swap on the Notional Token wrapper
        uint256 returned = token.deposit(a - fee, address(this));

        // Mint Illuminate zero coupons
        address illuminateToken = principalToken(u, m);(u, m);
        IERC5095(illuminateToken).mint(msg.sender, returned);

        emit Lend(p, u, m, returned);
        return returned;
    }

    /// @notice transfers excess funds to yield pool after principal tokens have been lent out
    /// @dev this method is only used by the yield, illuminate and swivel protocols
    /// @param u address of an underlying asset
    /// @param y the yield pool to lend to
    /// @param a the amount of underlying tokens to lend
    /// @param r the receiving address for PTs
    /// @return uint256 the amount of tokens sent to the yield pool
    function yield(
        address u,
        address y,
        uint256 a,
        address r
    ) internal returns (uint256) {
        // preview exact swap slippage on yield
        uint128 returned = IYield(y).sellBasePreview(Cast.u128(a));

        // send the remaing amount to the given yield pool
        Safe.transfer(IERC20(u), y, a);

        // lend out the remaining tokens in the yield pool
        IYield(y).sellBase(r, returned);

        return returned;
    }

    /// @notice withdraws accumulated lending fees of the underlying token
    /// @param e address of the underlying token to withdraw
    /// @return bool true if successful
    function withdrawFee(address e) external authorized(admin) returns (bool) {
        // Get the token to be withdrawn
        IERC20 token = IERC20(e);

        // Get the balance to be transferred
        uint256 balance = fees[e];

        // Reset accumulated fees of the token to 0
        fees[e] = 0;

        // Transfer the accumulated fees to the admin
        Safe.transfer(token, admin, balance);
        return true;
    }

    /// @notice this method returns the fee based on the amount passed to it. If the feenominator is 0, then there is no fee.
    /// @param a amount of underlying tokens to calculate the fee for
    /// @return uint256 The total for for the given amount
    function calculateFee(uint256 a) internal view returns (uint256) {
        return feenominator > 0 ? a / feenominator : 0;
    }

    /// @notice allows the admin to schedule the withdrawal of tokens
    /// @param e address of (erc20) token to withdraw
    /// @return bool true if successful
    function scheduleWithdrawal(address e) external authorized(admin) returns (bool) {
        uint256 when = block.timestamp + HOLD;
        withdrawals[e] = when;

        emit ScheduleWithdrawal(e, when);
        return true;
    }

    /// @notice emergency function to block unplanned withdrawals
    /// @param e address of token withdrawal to block
    /// @return bool true if successful
    function blockWithdrawal(address e) external authorized(admin) returns (bool) {
        withdrawals[e] = 0;

        emit BlockWithdrawal(e);
        return true;
    }

    /// @notice allows the admin to withdraw the given token, provided the holding period has been observed
    /// @param e Address of token to withdraw
    /// @return bool true if successful
    function withdraw(address e) external authorized(admin) returns (bool) {
        uint256 when = withdrawals[e];
        require (when != 0, 'no withdrawal scheduled');
  
        require (block.timestamp >= when, 'withdrawal still on hold');
  
        withdrawals[e] = 0;
  
        IERC20 token = IERC20(e);
        Safe.transfer(token, admin, token.balanceOf(address(this)));
  
        return true;
    }

    /// @notice retrieves the ERC5095 token for the given market
    /// @param u address of the underlying
    /// @param m uint256 representing the maturity of the market
    /// @return address of the ERC5095 token for the market
    function principalToken(address u, uint256 m) internal returns (address) {
        return IMarketPlace(marketPlace).markets(u, m, 0);
    }

    /// @notice pauses a market and prevents execution of all lending for that market
    /// @param p principal enum value
    /// @param b bool representing whether to pause or unpause
    /// @return bool true if successful
    function pause(uint8 p, bool b) external authorized(admin) returns (bool) {
        paused[p] = b;
        return true;
    }

    /// @notice ensures that only a certain address can call the function
    /// @param a address that msg.sender must be to be authorized
    modifier authorized(address a) {
        if (msg.sender != a) {
            revert Unauthorized();
        }
        _;
    }

    /// @notice reverts on all markets where the paused mapping returns true
    /// @param p principal enum value
    modifier unpaused(uint8 p) {
        if (paused[p]) {
            revert Invalid('paused');
        }
        _;
    }
}
