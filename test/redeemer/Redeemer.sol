// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import "./Illuminate.sol";
import "./Interfaces.sol";

contract Redeemer {

  // NOTE: the imported interfaces don't need to be named anything other than what they are:
  // ISwivel
  // IYield
  // etc...

  address public admin;
  address public illuminate;
  address public pendleRouter;
  address public tempusRouter;
  address public apRouter;
  address public illuminateRouter;

  event Redeem(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount);  

  /// @param m the deployed Illuminate contract
  constructor(address m, address p, address t, address a, address i) {
    admin = msg.sender;
    illuminate = m; // TODO add an authorized setter for this?
    pendleRouter = p;
    tempusRouter = t;
    apRouter = a;
    illuminateRouter = i;
  }
  /// @dev redeem method signature for illuminate, tempus, apwine
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u underlying token being redeemed
  /// @param m maturity of the market being redeemed
  /// @param o owner / tempus pool / apiwine vault ?
  function redeem(uint8 p, address u, uint256 m, address o) public returns (bool) {
    IErc20 token = IErc20(IIlluminate(illuminate).markets(u, m)[p]);

    IRouter router;
    if (p == uint8(Illuminate.Principals.Apwine)) {
      router = IRouter(apRouter);
    } else if (p == uint8(Illuminate.Principals.Tempus)) {
      router = IRouter(tempusRouter);
    } else {
      router = IRouter(pendleRouter);
    }

    uint256 amount = token.balanceOf(illuminate);

    token.transferFrom(illuminate, address(this), amount);
    require(1 == 2, "made it here");

    router.redeemToBacking(o, amount, 0, address(this));

    emit Redeem(p, u, m, amount);

    return (true);
  }

  /// @dev redeem method signature for swivel, yield, element, 
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u underlying token being redeemed
  /// @param m maturity of the market being redeemed
  function redeem(uint8 p, address u, uint256 m) public returns (bool) {
    
    IRedeemable token = IRedeemable(IIlluminate(illuminate).markets(u, m)[p]);

    uint256 amount = token.balanceOf(illuminate);

    token.transferFrom(illuminate, address(this), amount);

    token.redeem(address(this), address(this), amount);

    emit Redeem(p, u, m, amount);

    return true;
  }

  /// @dev redeem method signature for pendle 
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u underlying token being redeemed
  /// @param m maturity of the market being redeemed
  /// @param i pendle id ?
  function redeem(uint8 p, address u, uint256 m, bytes32 i) public returns (bool) {
    IErc20 pendleToken = IErc20(IIlluminate(illuminate).markets(u, m)[p]);
    IPendleRouter Router = IPendleRouter(pendleRouter);

    uint256 amount = pendleToken.balanceOf(illuminate);

    pendleToken.transferFrom(illuminate, address(this), amount);

    Router.redeemAfterExpiry(i, u, m);

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
    IErc20 senseToken = IErc20(IIlluminate(illuminate).markets(u, m)[p]);

    uint256 amount = senseToken.balanceOf(address(this));

    senseToken.transferFrom(illuminate, address(this), amount);

    ISenseDivider(d).redeem(a, m, amount);

    emit Redeem(p, u, m, amount);

    return true;
  }
}
