// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import "./Interfaces.sol";
import "./MarketPlace.sol";
import "./Safe.sol";

contract Redeemer {
  address public admin;
  address public marketplace;

  /// @dev addresses of the 3rd party protocol contracts
  address public swivelAddr;
  address public pendleAddr;
  address public tempusAddr;
  address public apwineAddr;

  event Redeem(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount);  

  /// @param s: the swivel contract
  /// @param p: the pendle contract
  /// @param t: the tempus contract
  /// @param a: the apwine contract
  constructor(address s, address p, address t, address a) {
    admin = msg.sender;
    swivelAddr = s;
    pendleAddr = p;
    tempusAddr = t;
    apwineAddr = a;
  }

  /// @notice sets the address of the marketplace contract which contains the addresses of the markets
  /// @param i the address of the illumninate contract
  /// @return bool true if the address was set, false otherwise
  function setMarketplaceAddress(address i) authorized(admin) external returns (bool) {
    require(marketplace == address(0));
    marketplace = i;
    return true;
  }

  /// @notice Redeems underlying token for illuminate, apwine and tempus 
  /// protocols
  /// @dev Illuminate burns its tokens prior to redemption, unlike APWine and
  /// Tempus, which withdraw their tokens after transferring the underlying to 
  /// the redeem contract
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u the underlying token being redeemed
  /// @param m the maturity of the market being redeemed
  /// @param o address of the controller or contract that manages the underlying token
  /// @return bool true if the redemption was successful
  function redeem(uint8 p, address u, uint256 m, address o) public returns (bool) {
    // Get the address of the principal token being redeemed
    address principal = IMarketPlace(marketplace).markets(u, m)[p];

    // Get the amount of tokens to be redeemed from the principal token
    uint256 amount = IErc20(principal).balanceOf(o);

    // Transfer the underlying token to the redeem contract if it is not illuminate
    if (p != uint8(MarketPlace.Principals.Illuminate)) {
        Safe.transferFrom(IErc20(u), marketplace, address(this), amount);
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
        Safe.transferFrom(IErc20(u), marketplace, address(this), amount);
    }

    emit Redeem(0, u, m, amount);

    return true;
  }

  /// @dev redeem method for swivel, yield, element. This method redeems all
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u underlying token being redeemed
  /// @param m maturity of the market being redeemed
  /// @return bool true if the redemption was successful
  function redeem(uint8 p, address u, uint256 m) public returns (bool) {
    // Get the principal token that is being redeemed by the user
    address principal = IMarketPlace(marketplace).markets(u, m)[p];

    // The amount redeemed should be the balance of the principal token held by the illuminate contract
    // TODO: Should we check if the principal token has matured?
    uint256 amount = IErc20(principal).balanceOf(marketplace);

    // Transfer the principal token to the marketplace contract from here
    Safe.transferFrom(IErc20(principal), marketplace, address(this), amount);

    if (p == uint8(MarketPlace.Principals.Swivel)) {
      // Redeems zc tokens to the sender's address
      require((ISwivel(swivelAddr).redeemZcToken(u, m, amount)));
    } else if (p == uint8(MarketPlace.Principals.Element)) {
      // Redeems principal tokens from element
      IElementToken(principal).withdrawPrincipal(amount, marketplace);
    } else if (p == uint8(MarketPlace.Principals.Yield)) {
      // Redeems prinicipal tokens from yield
      IYieldToken(principal).redeem(address(this), address(this), amount);
    }

    emit Redeem(p, u, m, amount); 

    return true;
  }

  /// @dev redeem method signature for pendle 
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u underlying token being redeemed
  /// @param m maturity of the market being redeemed
  /// @param i forge id used by pendle to redeem the underlying token
  /// @return bool true if the redemption was successful
  function redeem(uint8 p, address u, uint256 m, bytes32 i) public returns (bool) {
    // Get the principal token that is being redeemed by the user
    IErc20 token = IErc20(IMarketPlace(marketplace).markets(u, m)[p]);

    // Get the balance of tokens to be redeemed by the user
    uint256 amount = token.balanceOf(marketplace);

    // Transfer the user's tokens to the redeem contract
    Safe.transferFrom(token, marketplace, address(this), amount);

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
    // Get the principal token for the given market
    IErc20 token = IErc20(IMarketPlace(marketplace).markets(u, m)[p]);

    // Set the redeemer contract address
    address self = address(this);

    // Get the balance of tokens to be redeemed by the user
    uint256 amount = token.balanceOf(self);

    // Transfer the user's tokens to the redeem contract
    Safe.transferFrom(token, marketplace, self, amount);

    // Redeem the tokens from the sense contract
    ISense(d).redeem(o, m, amount);

    emit Redeem(p, u, m, amount);

    return true;
  }

  modifier authorized(address a) {
    require(msg.sender == a, 'sender must be authorized');
    _;
  }
}
