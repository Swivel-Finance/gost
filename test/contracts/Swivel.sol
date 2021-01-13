// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.0;

// NOTE: the agreement hydrating order parameters array is
// uint256[6] in this order [duration, rate, interest, principal, nonce, expiry]

contract Swivel {
  struct agreement {
    bytes32 key;
    bytes32 orderKey;
    address owner;
    address underlying;
    bool fixed;
    uint256 initialRate;
    uint256 rate;
    uint256 principal;
    uint256 interest;
    uint256 duration;
    uint256 release;
  }

  // only an Order can be cancelled
  mapping (bytes => bool) public cancelled
  // only an Order can be filled
  mapping (bytes => uint256) public filled

  Event Initiate(bytes orderKey, bytes agreementKey);
  Event Cancel(bytes key);
  Event Release(bytes orderKey, bytes agreementKey);

  // function fillFixed(bytes32 orderKey, uint256[6] params, bytes32 agreementKey, bytes memory signature) public returns (bool) {}

  // function fillFloating(bytes32 orderKey, uint256[6] params, bytes32 agreementKey, bytes memory signature) public returns (bool) {}
}
