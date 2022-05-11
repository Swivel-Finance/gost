// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import "./Interfaces.sol";
import "./Safe.sol";

contract Redeemer {

  // NOTE: the imported interfaces don't need to be named anything other than what they are:
  // ISwivel
  // IYield
  // etc...

  address public admin;
  address public illuminate;

  event Redeem(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount);  

  /// @param m the deployed Illuminate contract
  constructor(address m) {
    admin = msg.sender;
    illuminate = m; // TODO add an authorized setter for this?
  }

    /// @notice Can be called after maturity and after tokens have been redeemed to Illuminate to redeem underlying tokens 
    /// @param u the underlying token being redeemed
    /// @param m the maturity of the market being redeemed
    /// @param o the owner of the zcTokens being redeemed
    function redeem(address u, uint256 m, address o) public returns (bool) {
        IZcToken token = IZcToken(IIlluminate(illuminate).markets(u, m)[0]);

        uint256 amount = token.balanceOf(o);

        token.burn(o, amount);

        Safe.transfer(IErc20(u), o, amount);

        emit Redeem(0, u, m, amount);

        return true;
    }

  /// @dev redeem method signature for illuminate, tempus, apwine
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u underlying token being redeemed
  /// @param m maturity of the market being redeemed
  /// @param o owner / tempus pool / apiwine vault ?
  function redeem(uint8 p, address u, uint256 m, address o) public returns (bool) {
    // ...
    return true;
  }

  /// @dev redeem method signature for swivel, yield, element, 
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u underlying token being redeemed
  /// @param m maturity of the market being redeemed
  function redeem(uint8 p, address u, uint256 m) public returns (bool) {
    // ...
    return true;
  }

  /// @dev redeem method signature for pendle 
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u underlying token being redeemed
  /// @param m maturity of the market being redeemed
  /// @param i pendle id ?
  function redeem(uint8 p, address u, uint256 m, bytes32 i) public returns (bool) {
    // ...
    return true;
  }

  /// @dev redeem method signature for sense
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u underlying token being redeemed
  /// @param m maturity of the market being redeemed
  /// @param d sense wut (divider?) ?
  /// @param a sense wut (adapter?) ?
  function redeem(uint8 p, address u, uint256 m, address d, address a) public returns (bool) {
    // ...
    return true;
  }
}
