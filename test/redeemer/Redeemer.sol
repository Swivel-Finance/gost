// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import "./Interfaces.sol";
import "./MarketPlace.sol";
import "./Safe.sol";

contract Redeemer {
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
  /// @param l: the lender contract
  /// @param s: the swivel contract
  /// @param p: the pendle contract
  /// @param t: the tempus contract
  /// @param a: the apwine contract
  constructor(address l, address s, address p, address t, address a) {
    admin = msg.sender;
    lender = l;
    swivelAddr = s;
    pendleAddr = p;
    tempusAddr = t;
    apwineAddr = a;
  }

  /// @notice Sets the address of the marketplace contract which contains the 
  /// addresses of all the fixed rate markets
  /// @param m the address of the marketplace contract
  /// @return bool true if the address was set, false otherwise
  function setMarketPlaceAddress(address m) authorized(admin) external returns (bool) {
    require(marketPlace == address(0));
    marketPlace = m;
    return true;
  }

  function setLenderAddress(address l) authorized(admin) external returns (bool) {
    require(lender == address(0));
    lender = l;
    return true;
  }

  /// @notice Redeems underlying token for illuminate, apwine and tempus 
  /// protocols
  /// @dev Illuminate burns its tokens prior to redemption, unlike APWine and
  /// Tempus, which withdraw their tokens after transferring the underlying to 
  /// the redeem contract
  /// @dev We can avoid a require check on the principal at the start of this
  /// redeem because there is no common business logic executed before the
  /// protocol specific code is executed.
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u the underlying token being redeemed
  /// @param m the maturity of the market being redeemed
  /// @param o address of the controller or contract that manages the underlying token
  /// @return bool true if the redemption was successful
  function redeem(uint8 p, address u, uint256 m, address o) public returns (bool) {
    // Get the address of the principal token being redeemed
    address principal = IMarketPlace(marketPlace).markets(u, m, p);

    // Get the amount of tokens to be redeemed from the principal token
    uint256 amount = IErc20(principal).balanceOf(o);

    // Transfer the underlying token to the redeem contract if it is not illuminate
    if (p != uint8(MarketPlace.Principals.Illuminate)) {
        Safe.transferFrom(IErc20(u), lender, address(this), amount);
    } 
    
    if (p == uint8(MarketPlace.Principals.Apwine)) {
        // Redeem the underlying token from APWine to illuminate
        IAPWine(apwineAddr).withdraw(o, amount);
    } else if (p == uint8(MarketPlace.Principals.Tempus)) {
        // Redeem the tokens from the tempus contract to illuminate
        ITempus(tempusAddr).redeemToBacking(o, amount, 0, address(this));
    } else if (p == uint8(MarketPlace.Principals.Illuminate)) {
        // Burn the prinicipal token from illuminate
        IZcToken(principal).burn(o, amount);
        // Transfer the original underlying token back to the user
        Safe.transferFrom(IErc20(u), lender, address(this), amount);
    } else {
        revert("Invalid principal");
    }

    emit Redeem(0, u, m, amount);

    return true;
  }

  /// @dev redeem method for swivel, yield, element and notional. This method redeems all
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u underlying token being redeemed
  /// @param m maturity of the market being redeemed
  /// @return bool true if the redemption was successful
  function redeem(uint8 p, address u, uint256 m) public returns (bool) {
    // Get the principal token that is being redeemed by the user
    address principal = IMarketPlace(marketPlace).markets(u, m, p);

    // Make sure we have the correct principal
    require(p == uint8(MarketPlace.Principals.Swivel) ||
        p == uint8(MarketPlace.Principals.Element) ||
        p == uint8(MarketPlace.Principals.Yield) ||
        p == uint8(MarketPlace.Principals.Notional), "Invalid principal");

    // The amount redeemed should be the balance of the principal token held by the illuminate contract
    uint256 amount = IErc20(principal).balanceOf(marketPlace);

    // Transfer the principal token from the marketplace contract to here
    Safe.transferFrom(IErc20(principal), lender, address(this), amount);

    if (p == uint8(MarketPlace.Principals.Swivel)) {
      // Redeems zc tokens to the sender's address
      require((ISwivel(swivelAddr).redeemZcToken(u, m, amount)));
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
  function redeem(uint8 p, address u, uint256 m, bytes32 i) public returns (bool) {
    // Check the principal is pendle
    require(p == uint8(MarketPlace.Principals.Pendle));

    // Get the principal token that is being redeemed by the user
    IErc20 token = IErc20(IMarketPlace(marketPlace).markets(u, m, p));

    // Get the balance of tokens to be redeemed by the user
    uint256 amount = token.balanceOf(marketPlace);

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
  function redeem(uint8 p, address u, uint256 m, address d, address o) public returns (bool) {
    // Check the principal is sense
    require(p == uint8(MarketPlace.Principals.Sense));
    
    // Get the principal token for the given market
    IErc20 token = IErc20(IMarketPlace(marketPlace).markets(u, m, p));

    // Set the redeemer contract address
    address self = address(this);

    // Get the balance of tokens to be redeemed by the user
    uint256 amount = token.balanceOf(self);

    // Transfer the user's tokens to the redeem contract
    Safe.transferFrom(token, lender, self, amount);

    // Redeem the tokens from the sense contract
    ISense(d).redeem(o, m, amount);

    emit Redeem(p, u, m, amount);

    return true;
  }


  /// @notice implements the redeem method for the contract to fulfil the 
  /// ERC-5095 interface
  /// @param u address of the underlying asset
  /// @param m maturity of the market
  /// @param f address from where the underlying asset will be burned
  /// @param t address to where the underlying asset will be transferred
  /// @param a amount of the underlying asset to be burned
  /// @return bool true if the underlying asset was burned successfully
  function authRedeem(address u, uint256 m, address f, address t, uint256 a) authorized(IMarketPlace(marketPlace).markets(u, m, 0)) public returns (bool) {
    // Get the principal token for the given market
    IZcToken pt = IZcToken(IMarketPlace(marketPlace).markets(u, m, 0));
    require(block.timestamp > pt.maturity(), 'maturity error');

    // Get the balance of tokens to be redeemed by the user
    uint256 amount = pt.balanceOf(f);

    // Burn the user's principal tokens
    pt.burn(f, a);

    // Transfer the original underlying token back to the user
    Safe.transfer(IErc20(u), t, a);
  }

  modifier authorized(address a) {
    require(msg.sender == a, 'sender must be authorized');
    _;
  }
}
