// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import "./Illuminate.sol";
import "./Interfaces.sol";
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
  /// @param p: the principal token
  /// @param u the underlying token being redeemed
  /// @param m the maturity of the market being redeemed
  /// @param o the owner of the underlying tokens being redeemed
  /// TODO: owner is a bad name for the 4th parameter. For illuminate, it makes
  /// sense, but for Tempus/APWine, it should be something like router. Ask
  /// Julian for clarification.
  function redeem(uint8 p, address u, uint256 m, address o) public returns (bool) {
    address principal = IIlluminate(illuminate).markets(u, m)[p];

    uint256 amount = IErc20(principal).balanceOf(o);

    if (p == uint8(Illuminate.Principals.Apwine)) {
        IAPWine(apwineAddr).withdraw(o, amount);
    } else if (p == uint8(Illuminate.Principals.Tempus)) {
        ITempus(tempusAddr).redeemToBacking(o, m, amount, u);
    } else if (p == uint8(Illuminate.Principals.Illuminate)) {
        IZcToken(principal).burn(o, amount);
    }

    Safe.transfer(IErc20(u), o, amount);

    emit Redeem(0, u, m, amount);

    return true;
  }

  /// @dev redeem method for swivel, yield, element. This method redeems all
  /// prinicipal tokens to illuminate upon maturity
  /// @param p value of a specific principal according to the Illuminate Principals Enum
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
  /// @param i forge id ?
  function redeem(uint8 p, address u, uint256 m, bytes32 i) public returns (bool) {    
    IErc20 token = IErc20(IIlluminate(illuminate).markets(u, m)[p]);

    uint256 amount = token.balanceOf(illuminate);

    Safe.transferFrom(token, illuminate, address(this), amount);

    IPendle(pendleAddr).redeemAfterExpiry(i, u, m);

    emit Redeem(p, u, m, amount);

    return true;
  }

  /// @dev redeem method signature for sense
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u underlying token being redeemed
  /// @param m maturity of the market being redeemed
  /// @param d sense wut (divider?) ?
  /// @param a sense wut (adapter?) ?
  function redeem(uint8 p, address u, uint256 m, address d, address a) public returns (bool) {
    IErc20 token = IErc20(IIlluminate(illuminate).markets(u, m)[p]);

    uint256 amount = token.balanceOf(address(this));

    Safe.transferFrom(token, illuminate, address(this), amount);

    ISense(d).redeem(a, m, amount);

    emit Redeem(p, u, m, amount);

    return true;
  }
}
