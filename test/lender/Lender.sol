// SPDX-License-Identifier: GPL-3.0-or-later

pragma solidity 0.8.13;

contract Lender {
  import './Interfaces.sol';
  import './Swivel.sol'; // library of swivel specific constructs

  // TODO any reason to allow an admin to change these? (not use constant or immutable)
  address constant SWIVEL = '0xswivel'; // TODO the actual deployed addr (also pass to ctor and be immutable? why over just constant?)

  address public admin;
  address public marketPlace;

  event Lend(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 returned);

  constructor() {
    admin = msg.sender;
  }

  /// @dev lend method signature for both illuminate and yield
  /// @notice Can be called before maturity to lend to Illuminate / Yield ?
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param y yield pool ?
  /// @param a amount of underlying tokens to lend ?
  function lend(uint8 p, address u, uint256 m, address y, uint128 a) public returns (uint256) {
    address self = address(this);

    // Erc20 uToken ..
    // ZcToken iToken ... TODO this can likely be inlined

    // uses iyield token interface...
    IYieldToken yToken = IYieldToken(y);

    // the yield token must match the market pair
    // TODO this needs to be cast? the inteface says yToken.maturity() returns uint32
    require(address(yToken.base()) == u, ''); // TODO err msgs
    require(yToken.maturity() == m, '');

    // transfer from user to illuminate
    Safe.TransferFrom(uToken, msg.sender, self, a);
    // priview exact swap slippage on yield
    uint128 returned yToken.sellBasePreview(a);
    // tranfer to yield
    Safe.Transfer(u, y, a);
    // 'sell base' meaning purchase the zero coupons from yield
    yToken.sellBase(self, returned);

    // this step is only needed when the lend is for yield
    if (p == MarketPlace(marketPlace).Principals.Yield) {
      // mint illuminate zct TODO returned must be cast? TODO its only used once so inline the instantiation and invocation
      iToken.mint(msg.sender, returned);
    }

    emit Lend(p, u, m, returned);

    // TODO this is going to need te be cast?
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
    // ...

    // uses iswivel interface
    ISwivel(SWIVEL).initiate(o, a, s);
    
    // calls lend method for yield?
    // ... = lend(MarketPlace(marketPlace).Principals.Yield, u, m, y, ...)

    uint256 returned = 0;

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
  function lend(uint8 p, address u, uint256 m, address t, bytes32 i, uint256 a, uint256 r, uint256 d) public returns (uint256) {
    address self = address(this);

    // get the market...
    // uint8[8] market = MarketPlace(marketPlace).markets(u, m) etc...

    // the underlying...
    // Erc20 uToken = ...

    // safe transfer from uToken is uniform
    Safe.transferFrom(uToken, msg.Sender, self, a);

    address principal = market[market.Principals.Element];
    IElementToken eToken = IElementToken(principal);

    // the element token must match the market pair
    require(address(eToken.underlying()) == u, '');
    require(eToken.unlockTimestamp() == m, '');

    // safe transfer... self...
    IElement.SingleSwap memory swap = IElement.SingleSwap({
      userData: '0x00000000000000000000000000000000000000000000000000000000000000',
      poolId: i, 
      amount: a,
      kind: IElement.SwapKind.In, // TODO OG cantract has foo.Enum(0) ?
      assetIn: IElement.Any(u),
      assetOut: IElement.Any(principal)
    });

    IElement.FundManagement memory fund = IElement.FundManagement({
      sender: self,
      recipient: payable(self),
      fromInternalBalance: false,
      toInternalBalance: false
    });

    uint256 returned = IElement(e).swap(swap, fund, r, d);

    emit Lend(p, u, m, returned);
    return returned;
  }

  /// @dev lend method signature for pendle
  /// @notice Can be called before maturity to lend to Pendle while minting Illuminate tokens
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param i pendle id ?
  /// @param a amount ?
  /// @param r minimum amount to return ?
  function lend(uint8 p, address u, uint256 m, bytes32 i, uint256 a, uint256 r) public returns (uint256) {
    // ...
    uint256 returned = 0;

    emit Lend(p, u, m, returned);
    return returned;
  }

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
  function lend(uint8 p, address u, uint256 m, address t, address x, uint256 a, uint256 r, uint256 d) public returns (uint256) {
    // ...
    uint256 returned = 0;

    emit Lend(p, u, m, returned);
    return returned;
  }

  /// @dev lend method signature for sense
  /// @notice Can be called before maturity to lend to Sense while minting Illuminate tokens
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param s sense pool ?
  /// @param x sense wut ?
  /// @param a amount ?
  /// @param r minimum amount to return ?
  function lend(uint8 p, address u, uint256 m, address s, address x, uint256 a, uint256 r) public returns (uint256) {
    // ...
    uint256 returned = 0;

    emit Lend(p, u, m, returned);
    return returned;
  }

  /// @dev lend method signature for apwine
  /// @notice Can be called before maturity to lend to APWine while minting Illuminate tokens
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u underlying token being ?
  /// @param m maturity of the market being ?
  /// @param w apwine pool ?
  /// @param i apwine pair id ?
  /// @param r minimum amount to return ?
  function lend(uint8 p, address u, uint256 m, address w, uint256 i, uint256 r) public returns (uint256) {
    // ...
    uint256 returned = 0;

    emit Lend(p, u, m, returned);
    return returned;
  }
}
