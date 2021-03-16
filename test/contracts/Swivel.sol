// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.0;

import './Sig.sol';
import './Hash.sol';
import './Abstracts.sol';

contract Swivel {
  // TODO visibility of these...
  string constant public NAME = "Swivel Finance";
  string constant public VERSION = "1.0.0";
  /// @dev DAI compound token, passed to constructor
  address public CTOKEN;
  /// @dev EIP712 domain separator.
  bytes32 public DOMAIN;

  struct Agreement {
    address maker;
    address taker;
    address underlying;
    bool floating;
    bool released;
    uint256 rate;
    uint256 principal;
    uint256 interest;
    uint256 duration;
    uint256 release;
  }

  /// @dev maps the key of an order to a boolean indicating if an order was cancelled
  mapping (bytes32 => bool) public cancelled;
  /// @dev maps the key of an order to an amount representing its taken volume
  mapping (bytes32 => uint256) public filled;
  /// @dev maps an order key to a mapping of agreement Key -> agreement
  mapping (bytes32 => mapping (bytes32 => Agreement)) public agreements;

  event Initiate (bytes32 indexed orderKey, bytes32 indexed agreementKey, uint256 filled, address taker);
  event Cancel (bytes32 indexed key);
  event Release (bytes32 indexed orderKey, bytes32 indexed agreementKey);
  
  /// @param i Chain ID of the network this contract is deployed to
  /// @param c Deployed address of the compound token to be used
  /// @param v [optional] Verifying contract if present. Defaults to this.
  constructor(uint256 i, address c, address v) {
    require(c != address(0), 'compound token address required');
    CTOKEN = c;
    address verifier = v == address(0) ? address(this) : v;
    DOMAIN = Hash.domain(NAME, VERSION, i, verifier);
  }

  /// @param o An offline Swivel.Order
  /// @param a order volume (interest) amount this agreement is filling
  /// @param k Key of this agreement
  /// @param c Components of a valid ECDSA signature
  function fillFixed(
    Hash.Order calldata o,
    uint256 a,
    bytes32 k,
    Sig.Components calldata c
  ) public valid(o, c) returns (bool) {
    require(o.floating == false, 'must be filled on fixed side');
    require(a <= (o.interest - filled[o.key]), 'taker amount > available volume');
    // .principal is principal * ratio / 1ETH were ratio is (a * 1ETH) / interest
    uint256 principal = o.principal * ((a * 1 ether) / o.interest) / 1 ether;


    // transfer tokens to this contract
    Erc20 uToken = Erc20(o.underlying);
    require(uToken.transferFrom(o.maker, address(this), principal), 'transfer from maker failed');
    require(uToken.transferFrom(msg.sender, address(this), a), 'transfer from taker failed');
    // interest is 'a' when fixed
    return fill(o, a, k, principal, a, false); 
  }

  /// @param o Array of offline Swivel.Orders
  /// @param a Array of order volume (interest) amounts relative to passed orders
  /// @param k Key for these agreements
  /// @param c Array of Components from valid ECDSA signatures
  function batchFillFixed(
    Hash.Order[] calldata o,
    uint256[] calldata a,
    bytes32 k,
    Sig.Components[] calldata c
  ) public returns (bool) {
    for (uint256 i=0; i < o.length; i++) {
      require(fillFixed(o[i], a[i], k, c[i]));     
    }

    return true;
  }

  /// @param o An offline Swivel.Order
  /// @param a order volume (principal) amount this agreement is filling
  /// @param k Key of this new agreement
  /// @param c Components of a valid ECDSA signature
  function fillFloating(
    Hash.Order calldata o,
    uint256 a,
    bytes32 k,
    Sig.Components calldata c
  ) public valid(o, c) returns (bool) {
    require(o.floating == true, 'must be filled on floating side');
    require(a <= (o.principal - filled[o.key]), 'taker amount > available volume');
    // .interest is interest * ratio / 1ETH where ratio is (a * 1ETH) / principal
    uint256 interest = o.interest * ((a * 1 ether) / o.principal) / 1 ether;

    // transfer tokens to this contract
    Erc20 uToken = Erc20(o.underlying);
    require(uToken.transferFrom(o.maker, address(this), interest), 'transfer from maker failed');
    require(uToken.transferFrom(msg.sender, address(this), a), 'transfer from taker failed');
    // .principal is 'a' when floating
    return fill(o, a, k, a, interest, true); 
  }

  /// @param o Array of offline Swivel.Order
  /// @param a Array of order volume (principal) amounts relative to passed orders
  /// @param k Key for these agreements
  /// @param c Array of Components from valid ECDSA signatures
  function batchFillFloating(
    Hash.Order[] calldata o,
    uint256[] calldata  a,
    bytes32 k,
    Sig.Components[] calldata c
  ) public returns (bool) {
     for (uint256 i=0; i < o.length; i++) {
      require(fillFloating(o[i], a[i], k, c[i]));     
    }

    return true;
  }

  /// @param o key of the order this agreement is associated with
  /// @param a key of the agreement being released
  function releaseFixed(bytes32 o, bytes32 a) public returns (bool) {
    require(agreements[o][a].floating == false, 'agreement must be fixed');
    return release(o, a, agreements[o][a].maker, agreements[o][a].taker);
  }

  /// @param o key of the order this agreement is associated with
  /// @param a key of the agreement being released
  function releaseFloating(bytes32 o, bytes32 a) public returns (bool) {
    require(agreements[o][a].floating, 'agreement must be floating');
    return release(o, a, agreements[o][a].taker, agreements[o][a].maker);
  }

  /// @param o Array of order keys the passed agreements are associated with
  /// @param a Array of agreement keys which are being released
  function batchRelease(bytes32[] calldata o, bytes32[] calldata a) public returns (bool) {
    for (uint256 i=0; i < o.length; i++) {
      if (agreements[o[i]][a[i]].floating) require(releaseFloating(o[i], a[i]), 'release of floating agreement failed');
      else require(releaseFixed(o[i], a[i]), 'release of fixed agreement failed');
    }

    return true; 
  }

  function cancel(Hash.Order calldata o, Sig.Components calldata c) public returns (bool) {
    require(o.maker == msg.sender || o.maker == tx.origin, 'must be authorized to cancel');
    require(o.maker == Sig.recover(Hash.message(DOMAIN, Hash.order(o)), c), 'invalid signature');

    cancelled[o.key] = true;

    emit Cancel(o.key);

    return true;
  }

  /// @param o An offline Swivel.Order
  /// @param a order volume (principal or interest) amount this agreement is filling
  /// @param k Key of this new agreement
  /// @param p Principal of the new agreement
  /// @param i Interest of the new agreement
  /// @param b Boolean which indicates a floating agreement or not
  function fill(
    Hash.Order calldata o,
    uint256 a,
    bytes32 k,
    uint256 p,
    uint256 i,
    bool b
  ) internal returns (bool) {
    Agreement memory newAgreement; 
    
    newAgreement.principal = p;
    newAgreement.interest = i;

    // mint compound
    require(mintCToken(o.underlying, newAgreement.interest + newAgreement.principal) == 0, 'CToken minting failed');

    // finish setting new agreement props and store
    CErc20 cToken = CErc20(CTOKEN);
    newAgreement.maker = o.maker;
    newAgreement.taker = msg.sender;
    newAgreement.underlying = o.underlying; 
    newAgreement.floating = b;
    newAgreement.rate = cToken.exchangeRateCurrent();
    newAgreement.duration = o.duration;
    newAgreement.release = newAgreement.duration + block.timestamp;

    agreements[o.key][k] = newAgreement;

    filled[o.key] += a;

    emit Initiate(o.key, k, a, msg.sender);

    return true;
  }

  /// @param o key of the order this agreement is associated with
  /// @param a key of the agreement being released
  /// @param t Recipient Address of the calculated total
  /// @param d Recipient Address of the calculated diff
  function release(bytes32 o, bytes32 a, address t, address d) internal returns (bool) {
    Agreement memory releasing = agreements[o][a]; // TODO we could just look it up each time...

    require(releasing.released == false, 'agreement is already released');
    require(block.timestamp >= releasing.release, 'agreement term is not complete');


    CErc20 cToken = CErc20(CTOKEN);

    // calc annualized interest
    uint256 total = releasing.principal + releasing.interest;
    // TODO dusting?
    uint256 yield = ((cToken.exchangeRateCurrent() * 1e26) / releasing.rate) - 1e26; 
    uint256 interest = ((yield * releasing.principal) + (yield * releasing.interest)) / 1e26;

    // update the total, protecting from underflow.
    uint256 diff;
    if (interest == releasing.interest) diff = interest;
    else {
      diff = interest > releasing.interest ? interest - releasing.interest :
        releasing.interest - interest;
    }

    // TODO is there a sprintf type interpolation available here, or do we just pass more args?
    require(redeemCToken(total + diff) == 0, 'redeem compound token failed');

    // return funds to appropriate parties
    Erc20 uToken = Erc20(releasing.underlying);

    require(uToken.transfer(t, total), 'transfer of total failed');
    require(uToken.transfer(d, diff), 'transfer of diff failed');

    agreements[o][a].released = true;

    emit Release(o, a);

    return true;
  }

  /// @param u address of the underlying token contract
  /// @param n number of token to be minted
  function mintCToken(address u, uint256 n) internal returns (uint256) {
    Erc20 uToken = Erc20(u); 
    // approve for n on uToken, facilitating the eventual transfer
    require(uToken.approve(CTOKEN, n), 'underlying approval failed');
    CErc20 cToken = CErc20(CTOKEN);
    return cToken.mint(n);
  }

  /// @param n Number of underlying token to be redeemed
  function redeemCToken(uint256 n) internal returns (uint256) {
    return CErc20(CTOKEN).redeemUnderlying(n);
  }

  /// @dev Agreements may only be Initiated if the Order is valid.
  /// @param o An offline Swivel.Order
  /// @param c Components of a valid ECDSA signature
  modifier valid(Hash.Order calldata o, Sig.Components calldata c) {
    require(cancelled[o.key] == false, 'order has been cancelled');
    require(o.expiry >= block.timestamp, 'order has expired');
    require(o.maker == Sig.recover(Hash.message(DOMAIN, Hash.order(o)), c), 'invalid signature');
    _;
  }
}
