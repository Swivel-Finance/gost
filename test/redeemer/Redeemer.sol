// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Interfaces.sol';
import './MarketPlace.sol';
import './Safe.sol';

contract Redeemer {
    error Invalid(string);
    error Unauthorized();
    error Exists(string);

    address public admin;
    address public marketPlace;
    address public lender;

    /// @dev addresses of the 3rd party protocol contracts
    address public swivelAddr;
    address public pendleAddr;
    address public tempusAddr;
    address public apwineAddr;

    event Redeem(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount);

    /// @notice Initializes the Redeemer contract
    /// @param l the lender contract
    /// @param s the swivel contract
    /// @param p the pendle contract
    /// @param t the tempus contract
    /// @param a the apwine contract
    constructor(
        address l,
        address s,
        address p,
        address t,
        address a
    ) {
        admin = msg.sender;
        lender = l;
        swivelAddr = s;
        pendleAddr = p;
        tempusAddr = t;
        apwineAddr = a;
    }

    /// @param a Address of a new admin
    function setAdmin(address a) external authorized(admin) returns (bool) {
        admin = a;

        return true;
    }

    /// @notice Sets the address of the marketplace contract which contains the
    /// addresses of all the fixed rate markets
    /// @param m the address of the marketplace contract
    /// @return bool true if the address was set, false otherwise
    function setMarketPlaceAddress(address m) external authorized(admin) returns (bool) {
        if (marketPlace != address(0)) {
            revert Exists('marketplace');
        }
        marketPlace = m;
        return true;
    }

    /// @notice Sets the address of the lender contract which contains the
    /// addresses of all the fixed rate markets
    /// @param l the address of the lender contract
    /// @return bool true if the address was set, false otherwise
    function setLenderAddress(address l) external authorized(admin) returns (bool) {
        if (lender != address(0)) {
            revert Exists('lender');
        }
        lender = l;
        return true;
    }

    /// @notice Sets the feenominator to the given value
    /// @param s the address of the Swivel.sol Router
    /// @return bool true if successful
    function setSwivel(address s) external authorized(admin) returns (bool) {
        swivelAddr = s;
        return true;
    }

    /// @notice Redeems underlying token for illuminate, apwine and tempus
    /// protocols
    /// @dev Illuminate burns its tokens prior to redemption, unlike APWine and
    /// Tempus, which redeem PTs to the redeemer, transferring the underlying to
    /// this redeemer contract. Consequently, only Illuminate's redeem returns funds
    /// to the user.
    /// @dev We can avoid a require check on the principal at the start of this
    /// redeem because there is no common business logic executed before the
    /// protocol specific code is executed.
    /// @param p value of a specific principal according to the Illuminate Principals Enum
    /// @param u the underlying token being redeemed
    /// @param m the maturity of the market being redeemed
    /// @param o address of the controller or contract that manages the underlying token
    /// @return bool true if the redemption was successful
    function redeem(
        uint8 p,
        address u,
        uint256 m,
        address o
    ) public returns (bool) {
        // Get the address of the principal token being redeemed
        address principal = IMarketPlace(marketPlace).markets(u, m, p);

        if (p == uint8(MarketPlace.Principals.Illuminate)) {
            // Get Illuminate's principal token
            IERC5095 token = IERC5095(principal);
            // Get the amount of tokens to be redeemed from the sender
            uint256 amount = token.balanceOf(msg.sender);
            // Make sure the market has matured
            if (block.timestamp < token.maturity()) {
                revert Invalid('not matured');
            }
            // Burn the prinicipal token from illuminate
            token.burn(o, amount);
            // Transfer the original underlying token back to the user
            Safe.transferFrom(IERC20(u), lender, address(this), amount);

            emit Redeem(0, u, m, amount);
        }
        else {
            // Get the amount of tokens to be redeemed from the principal token
            uint256 amount = IERC20(principal).balanceOf(lender);
            // Transfer the principal token from the lender contract to here
            Safe.transferFrom(IERC20(u), lender, address(this), amount);

            if (p == uint8(MarketPlace.Principals.Apwine)) {
                // Redeem the underlying token from APWine to illuminate
                IAPWine(apwineAddr).withdraw(o, amount);
            } else if (p == uint8(MarketPlace.Principals.Tempus)) {
                // Redeem the tokens from the tempus contract to illuminate
                ITempus(tempusAddr).redeemToBacking(o, amount, 0, address(this));
            } else {
                revert Invalid('principal');
            }
            emit Redeem(0, u, m, amount);
        }

        return true;
    }

    /// @dev redeem method for swivel, yield, element and notional. This method redeems all
    /// @param p value of a specific principal according to the Illuminate Principals Enum
    /// @param u underlying token being redeemed
    /// @param m maturity of the market being redeemed
    /// @return bool true if the redemption was successful
    function redeem(
        uint8 p,
        address u,
        uint256 m
    ) public returns (bool) {
        // Get the principal token that is being redeemed by the user
        address principal = IMarketPlace(marketPlace).markets(u, m, p);

        // Make sure we have the correct principal
        if (
            p != uint8(MarketPlace.Principals.Swivel) &&
            p != uint8(MarketPlace.Principals.Element) &&
            p != uint8(MarketPlace.Principals.Yield) &&
            p != uint8(MarketPlace.Principals.Notional)
        ) {
            revert Invalid('principal');
        }

        // The amount redeemed should be the balance of the principal token held by the illuminate contract
        uint256 amount = IERC20(principal).balanceOf(lender);

        // Transfer the principal token from the lender contract to here
        Safe.transferFrom(IERC20(principal), lender, address(this), amount);

        if (p == uint8(MarketPlace.Principals.Swivel)) {
            // Redeems zc tokens to the sender's address
            ISwivel(swivelAddr).redeemZcToken(u, m, amount);
        } else if (p == uint8(MarketPlace.Principals.Element)) {
            // Redeems principal tokens from element
            IElementToken(principal).withdrawPrincipal(amount, marketPlace);
        } else if (p == uint8(MarketPlace.Principals.Yield)) {
            // Redeems prinicipal tokens from yield
            IYieldToken(principal).redeem(address(this), address(this), amount);
        } else if (p == uint8(MarketPlace.Principals.Notional)) {
            // Redeems the principal token from notional
            amount = INotional(principal).maxRedeem(address(this));
        }

        emit Redeem(p, u, m, amount);

        return true;
    }

    /// @notice redeem method signature for pendle
    /// @param p value of a specific principal according to the Illuminate Principals Enum
    /// @param u underlying token being redeemed
    /// @param m maturity of the market being redeemed
    /// @param i forge id used by pendle to redeem the underlying token
    /// @return bool true if the redemption was successful
    function redeem(
        uint8 p,
        address u,
        uint256 m,
        bytes32 i
    ) public returns (bool) {
        // Check the principal is pendle
        if (p != uint8(MarketPlace.Principals.Pendle)) {
            revert Invalid('principal');
        }

        // Get the principal token that is being redeemed by the user
        IERC20 token = IERC20(IMarketPlace(marketPlace).markets(u, m, p));

        // Get the balance of tokens to be redeemed by the user
        uint256 amount = token.balanceOf(lender);

        // Transfer the user's tokens to the redeem contract
        Safe.transferFrom(token, lender, address(this), amount);

        // Redeem the tokens from the pendle contract
        IPendle(pendleAddr).redeemAfterExpiry(i, u, m);
        return true;
    }

    /// @dev redeem method signature for sense
    /// @param p value of a specific principal according to the Illuminate Principals Enum
    /// @param u underlying token being redeemed
    /// @param m maturity of the market being redeemed
    /// @param d sense contract that splits the loan's prinicpal and yield
    /// @param o sense contract that [d] calls into to adapt the underlying to sense
    /// @return bool true if the redemption was successful
    function redeem(
        uint8 p,
        address u,
        uint256 m,
        address d,
        address o
    ) public returns (bool) {
        // Check the principal is sense
        if (p != uint8(MarketPlace.Principals.Sense)) {
            revert Invalid('principal');
        }

        // Get the principal token for the given market
        IERC20 token = IERC20(IMarketPlace(marketPlace).markets(u, m, p));

        // Get the balance of tokens to be redeemed by the user
        uint256 amount = token.balanceOf(lender);

        // Transfer the user's tokens to the redeem contract
        Safe.transferFrom(token, lender, address(this), amount);

        // Redeem the tokens from the sense contract
        ISense(d).redeem(o, m, amount);

        emit Redeem(p, u, m, amount);

        return true;
    }

    /// @notice implements the redeem method for the contract to fulfill the
    /// ERC-5095 interface
    /// @param u address of the underlying asset
    /// @param m maturity of the market
    /// @param f address from where the underlying asset will be burned
    /// @param t address to where the underlying asset will be transferred
    /// @param a amount of the underlying asset to be burned and sent to the to
    /// @return bool true if the underlying asset was burned successfully
    function authRedeem(
        address u,
        uint256 m,
        address f,
        address t,
        uint256 a
    ) public authorized(IMarketPlace(marketPlace).markets(u, m, 0)) returns (bool) {
        // Get the principal token for the given market
        IERC5095 pt = IERC5095(IMarketPlace(marketPlace).markets(u, m, 0));

        // Make sure the market has matured
        if (block.timestamp < pt.maturity()) {
            revert Invalid('not matured');
        }

        // Burn the user's principal tokens
        pt.burn(f, a);

        // Transfer the original underlying token back to the user
        Safe.transfer(IERC20(u), t, a);
        return true;
    }

    modifier authorized(address a) {
        if (msg.sender != a) {
            revert Unauthorized();
        }
        _;
    }
}
