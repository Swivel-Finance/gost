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

  struct agreement {
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
  mapping (bytes32 => mapping (bytes32 => agreement)) public agreements;

  event Initiate (bytes32 orderKey, bytes32 agreementKey);
  event Cancel (bytes32 key);
  event Release (bytes32 orderKey, bytes32 agreementKey);

  /// @param c Chain ID of the network this contract is deployed to
  /// @param v [optional] Verifying contract if present. Defaults to this.
  constructor(uint256 c, address v) {
    address verifier = v == address(0) ? address(this) : v;
    DOMAIN = Hash.domain(NAME, VERSION, c, verifier);
  }

  /// @param o Key of the order being filled by this new agreement
  /// @param m address of the order's maker
  /// @param p Params of the order.
  /// @param a order volume this agreement is filling
  /// @param k Key of this new agreement
  /// @param u address of the underlying token contract
  /// @param v Signature verification bit
  /// @param r First 32 bytes of a valid ECDSA Signature
  /// @param s Next 32 bytes of a valid ECDSA Signature
  function fillFixed(
    bytes32 o,
    address m,
    uint256[6] calldata p,
    uint256 a,
    bytes32 k,
    address u,
    uint8 v,
    bytes32 r,
    bytes32 s
  ) public returns (bool) {
    require(cancelled[o] == false, 'order has been cancelled');
    require(p[uint256(Hash.Params.EXPIRY)] >= block.timestamp, 'order has expired');
    require(m == Sig.recover(Hash.message(DOMAIN, Hash.order(o, m, u, false, p)), v, r, s), 'invalid signature'); 
    require(a <= (p[uint256(Hash.Params.INTEREST)] - filled[o]), 'taker amount > available volume');

    agreement memory newAgreement;
    // .principal is principal * ratio / 1ETH were ratio is (a) * 1ETH / interest
    newAgreement.principal = p[uint256(Hash.Params.PRINCIPAL)] * ((a * (1 ether)) / p[uint256(Hash.Params.INTEREST)]) / (1 ether);
    newAgreement.interest = a;

    // transfer tokens to this contract
    Erc20 uToken = Erc20(u);
    require(uToken.transferFrom(m, address(this), newAgreement.principal), 'transfer from maker failed');
    require(uToken.transferFrom(msg.sender, address(this), newAgreement.interest), 'transfer from taker failed');
    // mint compound
    require(mintCToken(u, p[uint256(Hash.Params.PRINCIPAL)] + p[uint256(Hash.Params.INTEREST)]) > 0, 'CToken minting failed');

    // finish setting new agreement props and store
    CErc20 cToken = CErc20(CTOKEN);
    newAgreement.maker = m;
    newAgreement.taker = msg.sender;
    newAgreement.underlying = u; 
    newAgreement.floating = false;
    newAgreement.initialRate = cToken.exchangeRateCurrent();
    newAgreement.rate = p[uint256(Hash.Params.RATE)];
    newAgreement.duration = p[uint256(Hash.Params.DURATION)];
    newAgreement.release = newAgreement.duration + block.timestamp;

    agreements[o][k] = newAgreement;

    emit Initiate(o, k);

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
