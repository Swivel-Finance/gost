// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import "./Interfaces.sol";
import "./ZcToken.sol";
import "./Safe.sol";

contract Illuminate {
  /// @notice for any given illuminate deployment the available principles are contained here
  enum Principals {
    Illuminate,
    Swivel,
    Yield,
    Element,
    Pendle,
    Tempus,
    Sense,
    Apwine
  }

  /// markets are defined by a market pair which point to a fixed length array of principal token addresses.
  /// the principal tokens those addresses represent correspond to their Principals enum value, thus the
  /// array should be ordered in that way
  mapping (address => mapping (uint256 => address[8])) public markets;

  address public admin;
  /// @notice address of the deployed redeemer contract
  address public immutable redeemer;

  event CreateMarket(address indexed underlying, uint256 indexed maturity);

  /// @param r address of the deployed redeemer contract 
  constructor(address r) {
    admin = msg.sender;
    redeemer = r;
  }

  /// @param u underlying
  /// @param m maturity
  /// @param t principal token addresses for this market minus the illuminate principal (which is added here)
  /// @param n name for the illuminate token
  /// @param s symbol for the illuminate token
  /// @param d decimals for the illuminate token
  function createMarket(address u, uint256 m, address[7] memory t, string calldata n, string calldata s, uint8 d) external authorized(admin) returns (bool) {
    require(markets[u][m][uint256(Principals.Illuminate)] == address(0), 'market exists'); 

    // deploy an illuminate token with this new market NOTE: ATM is using name as symbol args
    address iToken = address(new ZcToken(u, m, n, s, d));

    // the market will have the illuminate principal as its zeroth item, thus t should have Principals[1] as [0]
    // TODO we could choose to put illuminate last in
    address[8] memory market = [iToken, t[0], t[1], t[2], t[3], t[4], t[5], t[6]];

    uint256 max = 2**256 - 1;
    // approve the underlying for max per given principal
    for (uint8 i; i < 8; i++) {
      Safe.approve(IErc20(market[i]), redeemer, max);
    }

    markets[u][m] = market;

    // TODO "args" for this event?
    emit CreateMarket(u, m);
    // ...
    return true;
  }

  modifier authorized(address a) {
    require(msg.sender == a, 'sender must be authorized');
    _;
  }
}