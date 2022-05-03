// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Interfaces.sol';
import './MarketPlace.sol'; // library of market place specific constructs
import './Swivel.sol'; // library of swivel specific constructs
import './Safe.sol';
import './Cast.sol';

contract Lender {
  address public admin;
  address public marketPlace; // TODO authorized setter for this

  // TODO the nature of these addresses?
  address public swivelAddr; // addresses of the 3rd party protocol contracts

  event Lend(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 returned);

  /// @param m the deployed MarketPlace contract
  constructor(address m, address s) {
    admin = msg.sender;
    marketPlace = m; // TODO add an authorized setter for this?
    swivelAddr = s;
  }

  // TODO since we are heavily using overriding here do we need any extra fallback function security?
  // likely not if we just don't define it but... worth asking

  /// @dev lend method signature for both illuminate and yield
  /// @notice Can be called before maturity to lend to Illuminate / Yield ?
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param y yield pool ?
  /// @param a amount of underlying tokens to lend ?
  function lend(uint8 p, address u, uint256 m, address y, uint256 a) public returns (uint256) {
    // uses yield token interface...
    IYieldToken yToken = IYieldToken(y);

    // the yield token must match the market pair
    // TODO this needs to be cast? the inteface says yToken.maturity() returns uint32
    // TODO Use the base to get the address and compare that to the underlying
    require(address(yToken.base()) == u, 'yield base != underlying');
    require(yToken.maturity() == m, 'yield maturity != maturity');

    IErc20 uToken = IErc20(u);
    address self = address(this);

    // transfer from user to illuminate
    Safe.transferFrom(uToken, msg.sender, self, a);
    // priview exact swap slippage on yield
    uint128 returned = yToken.sellBasePreview(Cast.u128(a));
    // tranfer to yield
    Safe.transfer(uToken, y, a);
    // 'sell base' meaning purchase the zero coupons from yield
    yToken.sellBase(self, returned);

    // this step is only needed when the lend is for yield
    if (p == uint8(MarketPlace.Principals.Yield)) {
      address[8] memory market = IMarketPlace(marketPlace).markets(u, m); 
      // TODO should we require on this?
      IZcToken(market[uint256(MarketPlace.Principals.Illuminate)]).mint(msg.sender, returned);
    }

    emit Lend(p, u, m, returned);

    return returned;
  }

  /// @dev lend method signature for swivel
  /// @notice can be called before maturity to lend to Swivel while minting Illuminate tokens
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param y yield pool
  /// @param o array of swivel orders being filled
  /// @param a array of amounts of underlying tokens lent to each order in the orders array
  /// @param s array of signatures for each order in the orders array
  function lend(uint8 p, address u, uint256 m, address y, Swivel.Order[] calldata o, uint256[] calldata a, Swivel.Components[] calldata s) public returns (uint256) {
    uint256 lent;
    uint256 returned;

    for (uint256 i = 0; i < o.length; i++) {
      Swivel.Order memory order = o[i];
      // Require the Swivel order provided matches the underlying and maturity market provided    
      require(order.maturity == m, 'swivel maturity != maturity');
      require(order.underlying == u, 'swivel underlying != underlying');
      // Sum the total amount lent to Swivel (amount of zcb to mint)
      lent += a[i];
      // Sum the total amount of premium paid from Swivel (amount of underlying to lend to yield)
      returned += a[i] * order.premium / order.principal; // TODO guard order of operation?
    }

    // transfer from user to illuminate
    Safe.transferFrom(IErc20(u), msg.sender, address(this), lent);

    // fill the orders on swivel protocol, TODO: require response?
    ISwivel(swivelAddr).initiate(o, a, s);

    // TODO: lend the remaining amount to yield? (this is prob wrong)
    Safe.transfer(IErc20(u), y, returned - lent);

    emit Lend(p, u, m, returned);
    return returned;
  }

  /// @dev lend method signature for element
  /// @notice can be called before maturity to lend to Element / Sense ?
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param e element pool ?
  /// @param i element pool id ?
  /// @param a amount ?
  /// @param r minimum amount to return ?
  /// @param d deadline ?
  // function lend(uint8 p, address u, uint256 m, address t, bytes32 i, uint256 a, uint256 r, uint256 d) public returns (uint256) {
  //   address self = address(this);

  //   // get the market...
  //   // uint8[8] market = MarketPlace(marketPlace).markets(u, m) etc...

  //   // the underlying...
  //   // Erc20 uToken = ...

  //   // safe transfer from uToken is uniform
  //   Safe.transferFrom(uToken, msg.Sender, self, a);

  //   address principal = market[market.Principals.Element];
  //   ElementToken eToken = ElementToken(principal);

  //   // the element token must match the market pair
  //   require(address(eToken.underlying()) == u, '');
  //   require(eToken.unlockTimestamp() == m, '');

  //   // safe transfer... self...
  //   Element.SingleSwap memory swap = Element.SingleSwap({
  //     userData: address(0),
  //     poolId: i, 
  //     amount: a,
  //     kind: Element.SwapKind.In, // TODO OG cantract has foo.Enum(0) ?
  //     assetIn: Element.Any(u),
  //     assetOut: Element.Any(principal)
  //   });

  //   Element.FundManagement memory fund = Element.FundManagement({
  //     sender: self,
  //     recipient: payable(self),
  //     fromInternalBalance: false,
  //     toInternalBalance: false
  //   });

  //   uint256 returned = Element(e).swap(swap, fund, r, d);

  //   emit Lend(p, u, m, returned);
  //   return returned;
  // }

  /// @dev lend method signature for pendle
  /// @notice Can be called before maturity to lend to Pendle while minting Illuminate tokens
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param i pendle id ?
  /// @param a amount ?
  /// @param r minimum amount to return ?
  // function lend(uint8 p, address u, uint256 m, bytes32 i, uint256 a, uint256 r) public returns (uint256) {
  //   // ...
  //   uint256 returned = 0;

  //   emit Lend(p, u, m, returned);
  //   return returned;
  // }

  /// @dev lend method signature for tempus
  /// @notice Can be called before maturity to lend to Tempus while minting Illuminate tokens
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param t tempus pool ?
  /// @param x tempus amm ?
  /// @param a amount ?
  /// @param r minimum amount to return ?
  /// @param d deadline ?
  // function lend(uint8 p, address u, uint256 m, address t, address x, uint256 a, uint256 r, uint256 d) public returns (uint256) {
  //   // ...
  //   uint256 returned = 0;

  //   emit Lend(p, u, m, returned);
  //   return returned;
  // }

  /// @dev lend method signature for sense
  /// @notice Can be called before maturity to lend to Sense while minting Illuminate tokens
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param s sense pool ?
  /// @param x sense wut ?
  /// @param a amount ?
  /// @param r minimum amount to return ?
  // function lend(uint8 p, address u, uint256 m, address s, address x, uint256 a, uint256 r) public returns (uint256) {
  //   // ...
  //   uint256 returned = 0;

  //   emit Lend(p, u, m, returned);
  //   return returned;
  // }

  /// @dev lend method signature for apwine
  /// @notice Can be called before maturity to lend to APWine while minting Illuminate tokens
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param w apwine pool ?
  /// @param i apwine pair id ?
  /// @param r minimum amount to return ?
  // function lend(uint8 p, address u, uint256 m, address w, uint256 i, uint256 r) public returns (uint256) {
  //   // ...
  //   uint256 returned = 0;

  //   emit Lend(p, u, m, returned);
  //   return returned;
  // }
}
