// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Interfaces.sol';
import './Illuminate.sol'; // library of market place specific constructs
import './Swivel.sol'; // library of swivel specific constructs
import './Element.sol'; // library of element specific constructs
import './Safe.sol';
import './Cast.sol';


contract Lender {
  address public admin;
  address public illuminate; // TODO authorized setter for this

  address public swivelAddr; // addresses of the 3rd party protocol contracts
  address public pendleAddr;
  address public tempusAddr;

  uint256 public feenominator;

  event Lend(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 returned);
  event Mint(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount);

  /// @notice Initializes the Lender contract
  /// @dev the ctor sets a default value for the feenominator
  /// @param s: the swivel contract
  /// @param p: the pendle contract
  /// @param t: the tempus contract
  constructor(address s, address p, address t) {
    admin = msg.sender;
    swivelAddr = s;
    pendleAddr = p;
    tempusAddr = t;
    feenominator = 1000;
  }

  /// @notice Sets the feenominator to the given value
  /// @param f: the new value of the feenominator
  /// @return true if successful
  function setFee(uint256 f) external authorized(admin) returns (bool) {
    feenominator = f;
    return true;
  }
  
  /// Sets the address of the illuminate contract, contains the addresses of all
  /// the aggregated markets
  /// @param i: the address of the illumninate contract
  /// @return bool: true if the address was set, false otherwise
  function setIlluminateAddress(address i) authorized(admin) external returns (bool) {
    require(illuminate == address(0));
    illuminate = i;
    return true;
  }

  /// @dev mint is uniform across all principals, thus there is no need for a 'minter'
  /// @param p value of a specific principal according to the Illuminate Principals Enum.
  /// @param u address of an underlying asset
  /// @param m maturity (timestamp) of the market
  /// @param a amount being minted
  function mint(uint8 p, address u, uint256 m, uint256 a) public returns (bool) {
    //use market interface to fetch the market for the given market pair
    address[8] memory market = IIlluminate(illuminate).markets(u, m);
    // determine fee based on the amount
    uint256 fee = a / feenominator;
    // determine the fill amount
    uint256 amount = a - fee;
    //use safe transfer lib and ERC interface...
    Safe.transferFrom(IErc20(market[p]), msg.sender, address(this), amount);
    // extract the fee
    Safe.transferFrom(IErc20(market[p]), msg.sender, illuminate, fee);
    //use zctoken interface...
    IZcToken(market[uint256(Illuminate.Principals.Illuminate)]).mint(msg.sender, amount);

    emit Mint(p, u, m, a);

    return true;
  }

  // TODO since we are heavily using overriding here do we need any extra fallback function security?
  // likely not if we just don't define it but... worth asking

  /// @dev lend method signature for both illuminate and yield
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u address of an underlying asset
  /// @param m maturity (timestamp) of the market
  /// @param y yield pool that will execute the swap for the principal token
  /// @param a amount of underlying tokens to lend
  function lend(uint8 p, address u, uint256 m, address y, uint256 a) public returns (uint256) {
    // uses yield token interface...
    IYield yToken = IYield(y);

    // the yield token must match the market pair
    require(address(yToken.base()) == u, 'yield base != underlying');
    require(yToken.maturity() == m, 'yield maturity != maturity');

    // transfer from user to illuminate
    Safe.transferFrom(IErc20(u), msg.sender, address(this), a);
    
    uint256 returned = yield(u, y, a);

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
  /// @param u address of an underlying asset
  /// @param m maturity (timestamp) of the market
  /// @param y yield pool
  /// @param o array of swivel orders being filled
  /// @param a array of amounts of underlying tokens lent to each order in the orders array
  /// @param s array of signatures for each order in the orders array
  function lend(uint8 p, address u, uint256 m, address y, Swivel.Order[] calldata o, uint256[] calldata a, Swivel.Components[] calldata s) public returns (uint256) {
    // lent represents the number of underlying tokens lent
    uint256 lent;
    // returned represents the number of underlying tokens to lend to yield
    uint256 returned;

    // iterate through each order a calculate the total lent and returned
    for (uint256 i = 0; i < o.length; i++) {
      Swivel.Order memory order = o[i];
      // Require the Swivel order provided matches the underlying and maturity market provided    
      require(order.maturity == m, 'swivel maturity != maturity');
      require(order.underlying == u, 'swivel underlying != underlying');
      // Sum the total amount lent to Swivel (amount of zc tokens to mint)
      lent += a[i];
      // Sum the total amount of premium paid from Swivel (amount of underlying to lend to yield)
      returned += a[i] * (order.premium / order.principal);
    }

    // transfer underlying tokens from user to illuminate
    Safe.transferFrom(IErc20(u), msg.sender, address(this), lent);

    // fill the orders on swivel protocol
    ISwivel(swivelAddr).initiate(o, a, s);

    yield(u, y, returned);

    emit Lend(p, u, m, lent);
    return lent;
  }

  /// @dev lend method signature for element
  /// @notice fees are calculated in line as the amount divided by the 
  /// feenominator (a / feenominator)
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u address of an underlying asset
  /// @param m maturity (timestamp) of the market
  /// @param e element pool that is lent to
  /// @param i element pool id 
  /// @param a amount of principal tokens to lend
  /// @param r minimum amount to return 
  /// @param d deadline ?
  function lend(uint8 p, address u, uint256 m, address e, bytes32 i, uint256 a, uint256 r, uint256 d) public returns (uint256) {
    // Get the principal token for this market for element
    IElementToken token = IElementToken(IIlluminate(illuminate).markets(u, m)[p]);

    // the element token must match the market pair
    require(token.underlying() == u, '');
    require(token.unlockTimestamp() == m, '');

    // Transfer fee to illuminate
    Safe.transferFrom(IErc20(u), msg.sender, illuminate, a/feenominator);

    // Transfer underlying token from user to illuminate
    Safe.transferFrom(IErc20(u), msg.sender, address(this), a - (a / feenominator));
    
    // Create the variables needed to execute an element swap
    Element.FundManagement memory fund = Element.FundManagement({
      sender: address(this),
      recipient: payable(address(this)),
      fromInternalBalance: false,
      toInternalBalance: false
    });

    Element.SingleSwap memory swap = Element.SingleSwap({
      userData: "0x00000000000000000000000000000000000000000000000000000000000000",
      poolId: i, 
      amount: a - (a / feenominator),
      kind: Element.SwapKind.In,
      assetIn: Any(u),
      assetOut: Any(IIlluminate(illuminate).markets(u, m)[p])
    });


    // Conduct the swap on element
    uint256 purchased = IElement(e).swap(swap, fund, r, d);

    emit Lend(p, u, m, purchased);
    return purchased;
  }

  /// @dev lend method signature for pendle
  /// @notice Can be called before maturity to lend to Pendle while minting Illuminate tokens
  /// @param p value of a specific principal according to the MarketPlace Principals Enum
  /// @param u address of an underlying asset
  /// @param m maturity (timestamp) of the market
  /// @param a the amount of underlying tokens to lend
  /// @param r the minimum amount of zero-coupon tokens to return accounting for slippage
  /// @param d the maximum timestamp at which the transaction can be executed
  function lend(uint8 p, address u, uint256 m, uint256 a, uint256 r, uint256 d) public returns (uint256) {
      // Instantiate market and tokens
      address[8] memory markets = IIlluminate(illuminate).markets(u, m); 
      address principal = markets[p];
      IPendleToken token = IPendleToken(principal); // rename to pendletoken

      // confirm that we are in the correct market
      require(token.yieldToken() == u, 'pendle underlying != underlying');
      require(token.expiry() == m, 'pendle maturity != maturity');

      // Transfer funds from user to Illuminate
      Safe.transferFrom(IErc20(u), msg.sender, address(this), a);


      address[] memory path = new address[](2);
      path[0] = u;
      path[1] = principal;

      // Swap on the Pendle Router using the provided market and params
      uint256 returned = IPendle(pendleAddr).swapExactTokensForTokens(a, r, path, address(this), d)[0];

      // Mint Illuminate zero coupons
      address illuminateToken = markets[uint8(Illuminate.Principals.Illuminate)];
      IZcToken(illuminateToken).mint(msg.sender, returned);

      emit Lend(p, u, m, returned);

      return returned;
  }

  /// @dev lend method signature for tempus
  /// @notice Can be called before maturity to lend to Tempus while minting Illuminate tokens
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u address of an underlying asset
  /// @param m maturity (timestamp) of the market
  /// @param a amount of underlying tokens to swap for
  /// @param r minimum amount to return when executing the swap (sets a limit to slippage)
  /// @param x tempus amm that executes the swap
  /// @param t tempus pool that houses the underlying principal tokens
  /// @param d timestamp by which the swap must occur otherwise the lend is reverted
  /// @return returned number amount of underlying tokens that were lent
  function lend(uint8 p, address u, uint256 m, uint256 a, uint256 r, address x, address t, uint256 d) public returns (uint256) {
      // Instantiate market and tokens
      address principal = IIlluminate(illuminate).markets(u, m)[p];
      require(ITempus(principal).yieldBearingToken() == IErc20Metadata(u), 'tempus underlying != underlying');
      require(ITempus(principal).maturityTime() == m, 'tempus maturity != maturity');

      // Transfer funds from user to Illuminate, Scope to avoid stack limit
      IErc20 underlyingToken = IErc20(u);
      Safe.transferFrom(underlyingToken, msg.sender, address(this), a);

      // Swap on the Tempus Router using the provided market and params
      IZcToken illuminateToken = IZcToken(IIlluminate(illuminate).markets(u, m)[uint256(Illuminate.Principals.Illuminate)]);
      uint256 returned = ITempus(tempusAddr).depositAndFix(Any(x), Any(t), a, true, r, d) - illuminateToken.balanceOf(address(this));

      // Mint Illuminate zero coupons
      illuminateToken.mint(msg.sender, returned);

      emit Lend(p, u, m, returned);

      return returned;
  }

  /// @dev lend method signature for sense
  /// @notice Can be called before maturity to lend to Sense while minting Illuminate tokens
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u address of an underlying asset
  /// @param m maturity (timestamp) of the market
  /// @param x amm that is used to conduct the swap
  /// @param s contract that holds the principal token for this market
  /// @param a amount of underlying tokens to lend
  /// @param r minimum number of tokens to lend (sets a limit on the order)
  function lend(uint8 p, address u, uint256 m, address x, address s, uint128 a, uint256 r) public returns (uint256){
    // Get the principal token for this market for this market
    ISenseToken token = ISenseToken(IIlluminate(illuminate).markets(u, m)[p]);

    // Verify that the underlying matches up
    require(token.underlying() == u, "sense underlying != underlying");

    // Determine the fee
    uint256 fee = a / feenominator;

    // Determine lent amount after fees
    uint256 lent = a - fee;

    // Transfer fee to illuminate
    Safe.transferFrom(IErc20(u), msg.sender, illuminate, fee);
    
    // Transfer funds from user to Illuminate
    Safe.transferFrom(IErc20(u), msg.sender, address(this), lent);
    
    // Swap those tokens for the principal tokens
    uint256 returned = ISense(x).swapUnderlyingForPTs(s, m, lent, r);

    // Get the address of the ZC token for this market
    IZcToken illuminateToken = IZcToken(IIlluminate(illuminate).markets(u, m)[uint256(Illuminate.Principals.Illuminate)]);
    
    // Mint the illuminate tokens based on the returned amount
    illuminateToken.mint(msg.sender, returned);

    emit Lend(p, u, m, returned);

    return returned;
  }

  /// @notice Can be called before maturity to lend to APWine while minting Illuminate tokens
  /// @param p value of a specific principal according to the Illuminate Principals Enum
  /// @param u the underlying token being redeemed
  /// @param m the maturity of the market being redeemed
  /// @param a the amount of underlying tokens to lend
  /// @param r the minimum amount of zero-coupon tokens to return accounting for slippage
  /// @param pool the address of a given APWine pool
  /// @param i the id of the pool
  /// @return returned the amount of underlying tokens that were lent
  function lend(uint8 p, address u, uint256 m, uint256 a, uint256 r, address pool, uint256 i) public returns (uint256) {
      // Instantiate market and tokens
      address[8] memory markets = IIlluminate(illuminate).markets(u, m);
      require(IAPWineToken(markets[p]).getPTAddress() == u, "apwine principle != principle");

      // Determine the fee
      uint256 fee = a / feenominator;

      // Determine the amount lent after fees
      uint256 lent = a - fee;
      
      // Transfer fee from user to Illuminate
      Safe.transferFrom(IErc20(u), msg.sender, illuminate, fee);

      // Transfer funds from user to Illuminate    
      Safe.transferFrom(IErc20(u), msg.sender, address(this), lent);   

      // Swap on the APWine Pool using the provided market and params
      uint256 returned = IAPWineRouter(pool).swapExactAmountIn(i, 1, lent, 0, r, address(this));

      // Mint Illuminate zero coupons
      IZcToken(markets[uint256(Illuminate.Principals.Illuminate)]).mint(msg.sender, returned);

      emit Lend(p, u, m, returned);

      return returned;
  }

  /// @notice transfers funds to yield pool and executes a lend
  /// @param u: the underlying asset
  /// @param y: the yield pool to lend to
  /// @param a: the amount of underlying tokens to lend
  function yield(address u, address y, uint256 a) internal returns (uint256) {
    // preview exact swap slippage on yield
    uint128 returned = IYield(y).sellBasePreview(Cast.u128(a));

    // send the remaing amount to the given yield pool
    Safe.transfer(IErc20(u), y, returned);
    
    // lend out the remaining tokens in the yield pool
    IYield(y).sellBase(address(this), returned);

    return returned;
  }

  modifier authorized(address a) {
    require(msg.sender == a, 'sender must be authorized');
    _;
  }
}
