// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.0;

// NOTE: the agreement hydrating order parameters array is
// uint256[6] in this order [duration, rate, interest, principal, nonce, expiry]

contract Swivel {
  bytes32 SEPARATOR;
  bytes32 constant DOMAIN_HASH = keccak256("");
  bytes32 constant ORDER_HASH = keccak256("");

  // NOTE seems to be a style to upcase enums
  enum Sides { Fixed, Floating };

  // TODO watch the order of the uint256 re: `params` calldata array...
  struct agreement {
    address owner;
    bytes key;
    bytes orderKey;
    address underlying;
    uint256 duration;
    uint256 rate;
    uint256 interest;
    uint256 principal;
    uint256 release;
    uint256 initialRate;
    Sides side;
  }

  // only an Order can be cancelled
  mapping (bytes => bool) public cancelled
  // only an Order can be filled
  mapping (bytes => uint256) public filled

  Event Initiate(bytes orderKey, bytes agreementKey);
  Event Cancel(bytes key);
  Event Release(bytes orderKey, bytes agreementKey);

  // function fillFixed(bytes orderKey, uint256[6] params, bytes agreementKey, bytes memory signature) public returns (bool) {}

  // function fillFloating(bytes orderKey, uint256[6] params, bytes agreementKey, bytes memory signature) public returns (bool) {}
}
