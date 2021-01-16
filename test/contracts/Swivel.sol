// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.0;

import './Sig.sol';
import './Hash.sol';
import './Abstracts.sol';

contract Swivel {
  string constant internal NAME = "Swivel Finance";
  string constant internal VERSION = "1.0.0";

  /// @dev DAI compound token
  address constant internal CTOKEN = 0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb;

  /// @dev EIP712 domain separator.
  bytes32 public DOMAIN;

  struct Agreement {
    address maker;
    address taker;
    address underlying;
    bool floating;
    bool released;
    uint256 initialRate;
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

  event Initiate (bytes32 orderKey, bytes32 agreementKey);
  event Cancel (bytes32 key);
  event Release (bytes32 orderKey, bytes32 agreementKey);

  /// @param c Chain ID of the network this contract is deployed to
  /// @param v [optional] Verifying contract if present. Defaults to this.
  constructor(uint256 c, address v) {
    address verifier = v == address(0) ? address(this) : v;
    DOMAIN = Hash.domain(NAME, VERSION, c, verifier);
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
  ) public returns (bool) {
    require(cancelled[o.key] == false, 'order has been cancelled');
    require(o.expiry >= block.timestamp, 'order has expired');
    require(o.maker == Sig.recover(Hash.message(DOMAIN, Hash.order(o)), c), 'invalid signature'); 
    require(a <= (o.interest - filled[o.key]), 'taker amount > available volume');

    Agreement memory newAgreement;
    // .principal is principal * ratio / 1ETH were ratio is (a) * 1ETH / interest
    newAgreement.principal = o.principal * ((a * (1 ether)) / o.interest) / (1 ether);
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
    newAgreement.initialRate = cToken.exchangeRateCurrent();
    newAgreement.rate = o.rate;
    newAgreement.duration = o.duration;
    newAgreement.release = newAgreement.duration + block.timestamp;

    agreements[o.key][k] = newAgreement;

    emit Initiate(o.key, k);

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
}
