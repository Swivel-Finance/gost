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

  event Initiate (bytes32 indexed orderKey, bytes32 indexed agreementKey);
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

  /// @dev Agreements may only be Initiated if the Order is valid.
  /// @param o An offline Swivel.Order
  /// @param c Components of a valid ECDSA signature
  modifier valid(Hash.Order calldata o, Sig.Components calldata c) {
    require(cancelled[o.key] == false, 'order has been cancelled');
    require(o.expiry >= block.timestamp, 'order has expired');
    require(o.maker == Sig.recover(Hash.message(DOMAIN, Hash.order(o)), c), 'invalid signature');
    _;
  }

  /// @param o An offline Swivel.Order
  /// @param a order volume amount this agreement is filling
  /// @param k Key of this new agreement
  /// @param c Components of a valid ECDSA signature
  function fillFixed(
    Hash.Order calldata o,
    uint256 a,
    bytes32 k,
    Sig.Components calldata c
  ) public valid(o, c) returns (bool) {
    require(a <= (o.interest - filled[o.key]), 'taker amount > available volume');

    Agreement memory newAgreement;
    // .principal is principal * ratio / 1ETH were ratio is (a * 1ETH) / interest
    newAgreement.principal = o.principal * ((a * 1 ether) / o.interest) / 1 ether;
    newAgreement.interest = a;

    // transfer tokens to this contract
    Erc20 uToken = Erc20(o.underlying);
    require(uToken.transferFrom(o.maker, address(this), newAgreement.principal), 'transfer from maker failed');
    require(uToken.transferFrom(msg.sender, address(this), newAgreement.interest), 'transfer from taker failed');
    // mint compound
    require(mintCToken(o.underlying, o.principal + o.interest) > 0, 'CToken minting failed');

    // finish setting new agreement props and store
    CErc20 cToken = CErc20(CTOKEN);
    newAgreement.maker = o.maker;
    newAgreement.taker = msg.sender;
    newAgreement.underlying = o.underlying; 
    newAgreement.floating = false;
    newAgreement.rate = cToken.exchangeRateCurrent();
    newAgreement.duration = o.duration;
    newAgreement.release = newAgreement.duration + block.timestamp;

    agreements[o.key][k] = newAgreement;

    emit Initiate(o.key, k);

    return true;
  }

  /// @param o An offline Swivel.Order
  /// @param a order volume amount this agreement is filling
  /// @param k Key of this new agreement
  /// @param c Components of a valid ECDSA signature
  function fillFloating(
    Hash.Order calldata o,
    uint256 a,
    bytes32 k,
    Sig.Components calldata c
  ) public valid(o, c) returns (bool) {
    require(a <= (o.interest - filled[o.key]), 'taker amount > available volume');

    Agreement memory newAgreement;
    // .interest is interest * ratio / 1ETH where ratio is (a * 1ETH) / principal
    newAgreement.interest = o.interest * ((a * 1 ether) / o.principal) / 1 ether;
    newAgreement.principal = a;

    // transfer tokens to this contract
    Erc20 uToken = Erc20(o.underlying);
    require(uToken.transferFrom(o.maker, address(this), newAgreement.interest), 'transfer from maker failed');
    require(uToken.transferFrom(msg.sender, address(this), newAgreement.principal), 'transfer from taker failed');
    // mint compound
    require(mintCToken(o.underlying, o.interest + o.principal) > 0, 'CToken minting failed');

    // finish setting new agreement props and store
    CErc20 cToken = CErc20(CTOKEN);
    newAgreement.maker = o.maker;
    newAgreement.taker = msg.sender;
    newAgreement.underlying = o.underlying; 
    newAgreement.floating = true;
    newAgreement.rate = cToken.exchangeRateCurrent();
    newAgreement.duration = o.duration;
    newAgreement.release = newAgreement.duration + block.timestamp;

    agreements[o.key][k] = newAgreement;

    emit Initiate(o.key, k);

    return true;
  }

  /// @param o key of the order this agreement is associated with
  /// @param a key of the agreement being released
  function releaseFixed(bytes32 o, bytes32 a) public returns (bool) {
    Agreement memory releasing = agreements[o][a];

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

    require(uToken.transfer(releasing.maker, total), 'transfer to maker failed');
    require(uToken.transfer(releasing.taker, diff), 'transfer to taker failed');

    releasing.released = true;

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

  function redeemCToken(uint256 n) internal returns (uint256) {
    return CErc20(CTOKEN).redeemUnderlying(n);
  }
}
