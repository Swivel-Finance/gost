// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Interfaces.sol';
import './MarketPlace.sol'; // library of market place specific constructs
import './Element.sol'; // library of element specific constructs
import './Safe.sol';
import './Cast.sol';
import './Input.sol';

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
    uint256 public constant HOLD = 3 days;

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
    mapping(address => uint256) public withdrawals;

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
        } else {
            // Purchase illuminate PTs directly to msg.sender
            uint256 returned = yield(u, y, a - calculateFee(a), msg.sender);

            emit Lend(p, u, m, returned);

            return returned;
        }
    }

    /// @notice lend method signature for swivel
    /// @dev lends to yield pool. remaining balance is sent to the yield pool
    /// @param s Swivel input structure
    /// @return uint256 the amount of principal tokens lent out
    function lend(Input.Swivel memory s) public unpaused(s.principal) returns (uint256) {
        // lent represents the number of underlying tokens lent
        uint256 lent;
        // returned represents the number of underlying tokens to lend to yield
        uint256 returned;

        {
            uint256 totalFee;
            // iterate through each order a calculate the total lent and returned
            for (uint256 i = 0; i < s.orders.length; ) {
                Swivel.Order memory order = s.orders[i];
                // Require the Swivel order provided matches the underlying and maturity market provided
                if (order.underlying != s.underlying) {
                    revert NotEqual('underlying');
                } else if (order.maturity > s.maturity) {
                    revert NotEqual('maturity');
                }
                // Determine the fee
                uint256 fee = calculateFee(s.amounts[i]);
                // Track accumulated fees
                totalFee += fee;
                // Sum the total amount lent to Swivel (amount of ERC5095 tokens to mint) minus fees
                lent += s.amounts[i] - fee;
                // Sum the total amount of premium paid from Swivel (amount of underlying to lend to yield)
                returned += (s.amounts[i] - fee) * (order.premium / order.principal);

                unchecked {
                    i++;
                }
            }
            // Track accumulated fee
            fees[s.underlying] += totalFee;

            // transfer underlying tokens from user to illuminate
            Safe.transferFrom(IERC20(s.underlying), msg.sender, address(this), lent);
            // fill the orders on swivel protocol
            ISwivel(swivelAddr).initiate(s.orders, s.amounts, s.components);

            yield(s.underlying, s.pool, returned, address(this));
        }

        emit Lend(s.principal, s.underlying, s.maturity, lent);
        return lent;
    }

    /// @notice lend method signature for element
    /// @param e Element input struct
    /// @return uint256 the amount of principal tokens lent out
    function lend(Input.Element memory e) public unpaused(e.principal) returns (uint256) {
        // Get the principal token for this market for element
        address principal = IMarketPlace(marketPlace).markets(e.underlying, e.maturity, e.principal);

        // the element token must match the market pair
        if (IElementToken(principal).underlying() != e.underlying) {
            revert NotEqual('underlying');
        } else if (IElementToken(principal).unlockTimestamp() > e.maturity) {
            revert NotEqual('maturity');
        }
        // Transfer underlying token from user to illuminate
        Safe.transferFrom(IERC20(e.underlying), msg.sender, address(this), e.amount);

        // Track the accumulated fees
        fees[e.underlying] += calculateFee(e.amount);

        // Create the variables needed to execute an element swap
        Element.FundManagement memory fund = Element.FundManagement({
            sender: address(this),
            recipient: payable(address(this)),
            fromInternalBalance: false,
            toInternalBalance: false
        });

        Element.SingleSwap memory swap = Element.SingleSwap({
            userData: '0x00000000000000000000000000000000000000000000000000000000000000',
            poolId: e.id,
            amount: e.amount - calculateFee(e.amount),
            kind: Element.SwapKind.In,
            assetIn: Any(e.underlying),
            assetOut: Any(principal)
        });

        // Conduct the swap on element
        uint256 purchased = IElement(e.pool).swap(swap, fund, e.returned, e.deadline);

        emit Lend(e.principal, e.underlying, e.maturity, purchased);
        return purchased;
    }

    /// @notice lend method signature for pendle
    /// @param p pendle input structure
    /// @return uint256 the amount of principal tokens lent out
    function lend(Input.Pendle memory p) public unpaused(p.principal) returns (uint256) {
        // Instantiate market and tokens
        address principal = IMarketPlace(marketPlace).markets(p.underlying, p.maturity, p.principal);
        IPendleToken token = IPendleToken(principal);

        // confirm that we are in the correct market
        if (token.yieldToken() != p.underlying) {
            revert NotEqual('underlying');
        } else if (token.expiry() > p.maturity) {
            revert NotEqual('maturity');
        }

        // Transfer funds from user to Illuminate
        Safe.transferFrom(IERC20(p.underlying), msg.sender, address(this), p.amount);

        // Add the accumulated fees to the total
        uint256 fee = calculateFee(p.amount);
        fees[p.underlying] += fee;

        address[] memory path = new address[](2);
        path[0] = p.underlying;
        path[1] = principal;

        // Swap on the Pendle Router using the provided market and params
        uint256 returned = IPendle(pendleAddr).swapExactTokensForTokens(p.amount - fee, p.returned, path, address(this), p.deadline)[0];

        // Mint Illuminate zero coupons
        address illuminateToken = principalToken(p.underlying, p.maturity);
        IERC5095(illuminateToken).mint(msg.sender, returned);

        emit Lend(p.principal, p.underlying, p.maturity, returned);
        return returned;
    }

    /// @notice lend method signature for tempus
    /// @dev This method can be called before maturity to lend to Tempus while minting Illuminate tokens
    /// @param t Tempus input structure
    /// @return uint256 the amount of principal tokens lent out
    function lend(Input.Tempus memory t) public unpaused(t.principal) returns (uint256) {
        // Instantiate market and tokens
        address principal = IMarketPlace(marketPlace).markets(t.underlying, t.maturity, t.principal);
        if (ITempus(principal).yieldBearingToken() != IERC20Metadata(t.underlying)) {
            revert NotEqual('underlying');
        } else if (ITempus(principal).maturityTime() > t.maturity) {
            revert NotEqual('maturity');
        }

        // Get the underlying token
        IERC20 underlyingToken = IERC20(t.underlying);

        // Transfer funds from user to Illuminate, Scope to avoid stack limit
        Safe.transferFrom(underlyingToken, msg.sender, address(this), t.amount);

        // Add the accumulated fees to the total
        uint256 fee = calculateFee(t.amount);
        fees[t.underlying] += fee;

        // Swap on the Tempus Router using the provided market and params
        IERC5095 illuminateToken = IERC5095(principalToken(t.underlying, t.maturity));
        uint256 returned = ITempus(tempusAddr).depositAndFix(Any(t.amm), Any(t.pool), t.amount - fee, true, t.returned, t.deadline) -
            illuminateToken.balanceOf(address(this));

        // Mint Illuminate zero coupons
        illuminateToken.mint(msg.sender, returned);

        emit Lend(t.principal, t.underlying, t.maturity, returned);
        return returned;
    }

    /// @notice lend method signature for sense
    /// @dev this method can be called before maturity to lend to Sense while minting Illuminate tokens
    /// @dev sense provides a [divider] contract that splits [target] assets (underlying) into PTs and YTs. Each [target] asset has a [series] of contracts, each identifiable by their [maturity].
    /// @param s Sense input struct
    /// @return uint256 the amount of principal tokens lent out
    function lend(Input.Sense memory s) public unpaused(s.principal) returns (uint256) {
        // Get the principal token for this market for this market
        ISenseToken token = ISenseToken(IMarketPlace(marketPlace).markets(s.underlying, s.maturity, s.principal));

        // Verify that the underlying and maturity match up
        if (token.underlying() != s.underlying) {
            // gauruntee the input token is the right token
            revert NotEqual('underlying');
        } else if (ISense(s.divider).pt() != address(token)) {
            revert NotEqual('principal token');
        } else if (ISense(s.amm).maturity() > s.maturity) {
            // gauruntee the input amm has the correct maturity
            revert NotEqual('maturity');
        }

        // Determine the fee
        uint256 fee = calculateFee(s.amount);

        // Add the accumulated fees to the total
        fees[s.underlying] += fee;

        // Determine lent amount after fees
        uint256 lent = s.amount - fee;

        // Transfer funds from user to Illuminate
        Safe.transferFrom(IERC20(s.underlying), msg.sender, address(this), s.amount);

        // Swap those tokens for the principal tokens
        uint256 returned = ISense(s.amm).swapUnderlyingForPTs(s.divider, s.maturity, lent, s.returned);

        // Get the address of the ERC5095 token for this market
        IERC5095 illuminateToken = IERC5095(principalToken(s.underlying, s.maturity));

        // Mint the illuminate tokens based on the returned amount
        illuminateToken.mint(msg.sender, returned);

        emit Lend(s.principal, s.underlying, s.maturity, returned);
        return returned;
    }

    /// @notice this method can be called before maturity to lend to APWine while minting Illuminate tokens
    /// @param a APWine input structure
    /// @return uint256 the amount of principal tokens lent out
    function lend(Input.APWine memory a) public unpaused(a.principal) returns (uint256) {
        // Instantiate market and tokens
        address principal = IMarketPlace(marketPlace).markets(a.underlying, a.maturity, a.principal);
        if (IAPWineToken(principal).getUnderlyingOfIBTAddress() != a.underlying) {
            revert NotEqual('underlying');
        }
        // Dont necessarily need to validate APWINE maturity (They have 1 maturity per underlying)
        // Potentially add redundant implied maturity calculation

        // Transfer funds from user to Illuminate
        Safe.transferFrom(IERC20(a.underlying), msg.sender, address(this), a.amount);

        // Determine the fee
        uint256 fee = calculateFee(a.amount);

        // Add the accumulated fees to the total
        fees[a.underlying] += fee;

        // Determine the amount lent after fees
        uint256 lent = a.amount - fee;

        // Deposit into aave
        IAave(a.aave).deposit(a.underlying, lent, address(this), 0);

        // Swap on the APWine Pool using the provided market and params
        uint256 returned = IAPWineRouter(a.pool).swapExactAmountIn(a.id, 1, lent, 0, a.returned, address(this));

        // Mint Illuminate zero coupons
        IERC5095(principalToken(a.underlying, a.maturity)).mint(msg.sender, returned);

        emit Lend(a.principal, a.underlying, a.maturity, returned);
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
        address illuminateToken = principalToken(u, m);
        (u, m);
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
        require(when != 0, 'no withdrawal scheduled');

        require(block.timestamp >= when, 'withdrawal still on hold');

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
