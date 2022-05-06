// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Interfaces.sol';
import './Illuminate.sol'; // library of market place specific constructs
import './Swivel.sol'; // library of swivel specific constructs
import './Safe.sol';
import './Cast.sol';
import './Element.sol';

contract Lender {
  address public admin;
  address public illuminate; // TODO authorized setter for this

  // TODO the nature of these addresses?
  address public swivelAddr; // addresses of the 3rd party protocol contracts
  address public sushiRouter;

  event Lend(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 returned);
  event Mint(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount);

  /// @param i the deployed Illuminate contract
  constructor(address i, address s, address su) {
    admin = msg.sender;
    illuminate = i; // TODO add an authorized setter for this?
    swivelAddr = s;
    sushiRouter = su;
  }

  /// @dev mint is uniform across all principals, thus there is no need for a 'minter'
  // @param p value of a specific principal according to the Illuminate Principals Enum.
  // @param u address of an underlying asset
  // @param m maturity (timestamp) of the market
  // @param a amount being minted
  function mint(uint8 p, address u, uint256 m, uint256 a) public returns (bool) {
    //use market interface to fetch the market for the given market pair
    address[8] memory market = IIlluminate(illuminate).markets(u, m);
    //use safe transfer lib and ERC interface...
    Safe.transferFrom(IErc20(market[p]), msg.sender, address(this), a);
    //use zctoken interface...
    IZcToken(market[uint256(Illuminate.Principals.Illuminate)]).mint(msg.sender, a);

    emit Mint(p, u, m, a);

    return true;
  }

  // TODO since we are heavily using overriding here do we need any extra fallback function security?
  // likely not if we just don't define it but... worth asking

  /// @dev lend method signature for both illuminate and yield
  /// @notice Can be called before maturity to lend to Illuminate / Yield ?
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param y yield pool ?
  /// @param a amount of underlying tokens to lend ?
  function lend(uint8 p, address u, uint256 m, address y, uint256 a) public returns (uint256) {
    // uses yield token interface...
    IYield yToken = IYield(y);

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
    if (p == uint8(Illuminate.Principals.Yield)) {
      address[8] memory market = IIlluminate(illuminate).markets(u, m); 
      // TODO should we require on this?
      IZcToken(market[uint256(Illuminate.Principals.Illuminate)]).mint(msg.sender, returned);
    }

    emit Lend(p, u, m, returned);

    return returned;
  }

  /// @dev lend method signature for swivel
  /// @notice can be called before maturity to lend to Swivel while minting Illuminate tokens
  /// @param p value of a specific principal according to the Illuminate Principals Enum
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
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param e element pool ?
  /// @param i element pool id ?
  /// @param a amount ?
  /// @param r minimum amount to return ?
  /// @param d deadline ?
  function lend(uint8 p, address u, uint256 m, address e, bytes32 i, uint256 a, uint256 r, uint256 d) public returns (uint256) {
    // safe transfer from uToken is uniform
    Safe.transferFrom(IErc20(u), msg.sender, address(this), a);


    IElementToken eToken = IElementToken(IIlluminate(illuminate).markets(u, m)[p]);

    // the element token must match the market pair
    require(eToken.underlying() == u, '');
    require(eToken.unlockTimestamp() == m, '');


    // safe transfer... self...
    
    // TODO: userData should be set to address(0) in bytes form
    Element.SingleSwap memory swap = Element.SingleSwap({
      userData: "",
      poolId: i, 
      amount: a,
      kind: Element.SwapKind.In, // TODO OG cantract has foo.Enum(0) ?
      assetIn: Any(u),
      assetOut: Any(e)
    });

    Element.FundManagement memory fund = Element.FundManagement({
      sender: address(this),
      recipient: payable(address(this)),
      fromInternalBalance: false,
      toInternalBalance: false
    });

    emit Lend(p, u, m, IElement(e).swap(swap, fund, r, d));
    return IElement(e).swap(swap, fund, r, d);
  }

  /// @dev lend method signature for pendle
  /// @notice Can be called before maturity to lend to Pendle while minting Illuminate tokens
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u the underlying token being redeemed
  /// @param m the maturity of the market being redeemed
  /// @param a the amount of underlying tokens to lend
  /// @param mb the minimum amount of zero-coupon tokens to return accounting for slippage
  /// @param d the maximum timestamp at which the transaction can be executed
  function lend(uint8 p, address u, uint256 m, uint256 a, uint256 mb, uint256 d) public returns (uint256) {
      // Instantiate market and tokens
      address market = IIlluminate(illuminate).markets(u, m)[p];
      IPendle pendle = IPendle(market);

      // confirm that we are in the correct market
      require(pendle.underlying() == u, 'pendle underlying != underlying');
      require(pendle.maturity() == m, 'pendle maturity != maturity');

      // Transfer funds from user to Illuminate
      Safe.transferFrom(IErc20(u), msg.sender, address(this), a);


      address[] memory path = new address[](2);
      path[0] = u;
      path[1] = market;

      // Swap on the Pendle Router using the provided market and params
      uint256 returned = ISushi(sushiRouter).swapExactTokensForTokens(a, mb, path, address(this), d)[0];

      // Mint Illuminate zero coupons
      address[8] memory markets = IIlluminate(illuminate).markets(u, m); 
      IZcToken(markets[uint256(Illuminate.Principals.Illuminate)]).mint(msg.sender, returned);

      emit Lend(p, u, m, returned);

      return returned;
  }

  /// @dev lend method signature for tempus
  /// @notice Can be called before maturity to lend to Tempus while minting Illuminate tokens
  /// @param p value of a specific principal according to the Illuminate Principals Enum
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
  /// @param p value of a specific principal according to the Illuminate Principals Enum
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
  /// @param p value of a specific principal according to the Illuminate Principals Enum
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
