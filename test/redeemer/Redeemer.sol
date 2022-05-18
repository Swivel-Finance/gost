// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import "./Interfaces.sol";

import "./Illuminate.sol";

import "./Safe.sol";

contract Redeemer {
  address public admin;
  address public illuminate;

  address public apwineAddr;
  address public tempusAddr;
  address public pendleAddr;
  address public swivelAddr;

  event Redeem(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount);  

  /// @param m the deployed Illuminate contract
  constructor(address m, address a, address t, address p, address s) {
    admin = msg.sender;
    illuminate = m; // TODO add an authorized setter for this?
    apwineAddr = a;
    tempusAddr = t;
    pendleAddr = p;
    swivelAddr = s;
  }

  /// @notice Redeems underlying token for illuminate, apwine and tempus 
  /// protocols
  /// @notice Illuminate burns its tokens prior to redemption, unlike APWine and
  /// Tempus, which withdraw their tokens after transferring the underlying to 
  /// the redeem contract
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u the underlying token being redeemed
  /// @param m the maturity of the market being redeemed
  /// @param o address of the controller or contract that manages the underlying token
  function redeem(uint8 p, address u, uint256 m, address o) public returns (bool) {
    // Get the address of the principal token being redeemed
    address principal = IIlluminate(illuminate).markets(u, m)[p];

    // Get the amount of tokens to be redeemed from the principal token
    uint256 amount = IErc20(principal).balanceOf(o);

    // Transfer the underlying token to the redeem contract if it is not illuminate
    if (p != uint8(Illuminate.Principals.Illuminate)) {
        Safe.transferFrom(IErc20(u), illuminate, address(this), amount);
    }

    if (p == uint8(Illuminate.Principals.Apwine)) {
        // Redeem the underlying token from APWine to illuminate
        IAPWine(apwineAddr).withdraw(o, amount);
    } else if (p == uint8(Illuminate.Principals.Tempus)) {
        // Redeem the tokens from the tempus contract to illuminate
        ITempus(tempusAddr).redeemToBacking(o, amount, 0, address(this));
    } else if (p == uint8(Illuminate.Principals.Illuminate)) {
        // Burn the prinicipal token from illuminate
        IZcToken(principal).burn(o, amount);
        // Transfer the original underlying token back to the user
        Safe.transferFrom(IErc20(u), illuminate, address(this), amount);
    }

    emit Redeem(0, u, m, amount);

    return true;
  }

  /// @dev redeem method for swivel, yield, element. This method redeems all
  /// @param p prinicipal tokens to illuminate upon maturity
  /// @param p the protocol being used to redeem
  /// @param u underlying token being redeemed
  /// @param m maturity of the market being redeemed
  function redeem(uint8 p, address u, uint256 m) public returns (bool) {
    // Get the principal token that is being redeemed by the user
    address principal = IIlluminate(illuminate).markets(u, m)[p];

    // The amount redeemed should be the balance of the principal token held by the illuminate contract
    // TODO: Should we check if the principal token has matured?
    uint256 amount = IErc20(principal).balanceOf(illuminate);

    // Transfer the principal token to the illuminate contract from here
    Safe.transferFrom(IErc20(principal), illuminate, address(this), amount);

    if (p == uint8(Illuminate.Principals.Swivel)) {
      // Redeems zc tokens to the sender's address
      require((ISwivel(swivelAddr).redeemZcToken(u, m, amount)));
    } else if (p == uint8(Illuminate.Principals.Element)) {
      IElementToken(principal).withdrawPrincipal(amount, illuminate);
    } else if (p == uint8(Illuminate.Principals.Yield)) {
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
  function redeem(uint8 p, address u, uint256 m, bytes32 i) public returns (bool) {
    // Get the principal token that is being redeemed by the user
    IErc20 token = IErc20(IIlluminate(illuminate).markets(u, m)[p]);

    // Get the balance of tokens to be redeemed by the user
    uint256 amount = token.balanceOf(illuminate);

    // Transfer the user's tokens to the redeem contract
    Safe.transferFrom(token, illuminate, address(this), amount);

    // Redeem the tokens from the pendle contract
    IPendle(pendleAddr).redeemAfterExpiry(i, u, m);
    return true;
  }

  /// @dev redeem method signature for sense
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u underlying token being redeemed
  /// @param m maturity of the market being redeemed
  /// @param d sense contract that splits the loan's prinicpal and yield
  /// @param a sense contract that holds special logic for managing the underlying 
  /// token before and after it is sent to its respective Sense divider contract
  function redeem(uint8 p, address u, uint256 m, address d, address a) public returns (bool) {
    // Get the principal token for the given market
    IErc20 token = IErc20(IIlluminate(illuminate).markets(u, m)[p]);

    // Set the redeemer contract address
    address self = address(this);

    // Get the balance of tokens to be redeemed by the user
    uint256 amount = token.balanceOf(self);

    // Transfer the user's tokens to the redeem contract
    Safe.transferFrom(token, illuminate, self, amount);

    // Redeem the tokens from the sense contract
    ISense(d).redeem(a, m, amount);

    emit Redeem(p, u, m, amount);

    return true;
  }
}
