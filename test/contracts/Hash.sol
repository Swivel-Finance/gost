// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.0;

/**
  @notice Encapsulation of the logic to produce EIP712 hashed domain and messages.
  Also to produce / verify hashed and signed Orders.
  See https://github.com/ethereum/EIPs/blob/master/EIPS/eip-712.md
  See/attribute https://github.com/0xProject/0x-monorepo/blob/development/contracts/utils/contracts/src/LibEIP712.sol
*/

library Hash {
  // EIP712 Domain Separator typeHash
  // keccak256(abi.encodePacked(
  //     'EIP712Domain(',
  //     'string name,',
  //     'string version,',
  //     'uint256 chainId,',
  //     'address verifyingContract',
  //     ')'
  // ));
  // TODO double check
  bytes32 constant internal DOMAIN_TYPEHASH = 0x8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f;

  // EIP712 typeHash of an Order 
  // keccak256(abi.encodePacked(
  //     'Order(',
  //     'bytes32 key,',
  //     'address maker,',
  //     'address underlying,',
  //     'bool floating',
  //     'uint256 rate,',
  //     'uint256 principal,',
  //     'uint256 interest,',
  //     'uint256 duration,',
  //     'uint256 expiry,',
  //     'uint256 nonce,',
  //     ')'
  // ));
  // TODO doube check
  bytes32 constant internal ORDER_TYPEHASH = 0x982d366ee870e6c64d27e7b149dff6bf737fd1c909c2392b1b6dda92d31a24e2;

  /// @dev struct represents the attributes of an offchain Swivel.Order
  struct Order {
    bytes32 key;
    address maker;
    address underlying;
    bool floating;
    uint256 rate;
    uint256 principal;
    uint256 interest;
    uint256 duration;
    uint256 expiry;
    uint256 nonce;
  }

  /// @param n EIP712 domain name
  /// @param version EIP712 semantic version string
  /// @param c Chain ID
  /// @param verifier address of the verifying contract
  function domain(string memory n, string memory version, uint256 c, address verifier) internal pure returns (bytes32 hash) {
    assembly {
      let nameHash := keccak256(add(n, 32), mload(n))
      let versionHash := keccak256(add(version, 32), mload(version))
      let pointer := mload(64)
      mstore(pointer, DOMAIN_TYPEHASH)
      mstore(add(pointer, 32), nameHash)
      mstore(add(pointer, 64), versionHash)
      mstore(add(pointer, 96), c)
      mstore(add(pointer, 128), verifier)
      hash := keccak256(pointer, 160)
    }
    return hash;
  }

  /// @param d Type hash of the domain separator (see Hash.domain)
  /// @param h EIP712 hash struct (order for example)
  function message(bytes32 d, bytes32 h) internal pure returns (bytes32 hash) {
    assembly {
      let pointer := mload(64)
      mstore(pointer, 0x1901000000000000000000000000000000000000000000000000000000000000)
      mstore(add(pointer, 2), d)
      mstore(add(pointer, 34), h)
      hash := keccak256(pointer, 66)
    }
    return hash;
  }

  /// @param o A Swivel Order
  function order(Order calldata o) internal pure returns (bytes32) {
    // TODO assembly
    return keccak256(abi.encode(
      ORDER_TYPEHASH,
      o.key,
      o.maker,
      o.underlying,
      o.floating,
      o.rate,
      o.principal,
      o.interest,
      o.duration,
      o.expiry,
      o.nonce
    ));
  }
}
