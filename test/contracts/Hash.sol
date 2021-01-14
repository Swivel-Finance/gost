// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.0;

/**
  @notice Encapsulation of the logic to produce EIP712 hashed domain and messages.
  Also to produce / verify hashed and signed Orders.
  See https://github.com/ethereum/EIPs/blob/master/EIPS/eip-712.md
  See/attribute https://github.com/0xProject/0x-monorepo/blob/development/contracts/utils/contracts/src/LibEIP712.sol
*/

library Hash {
  enum Params { RATE, PRINCIPAL, INTEREST, DURATION, EXPIRY, NONCE }

  // EIP712 typeHash of an Order 
  // keccak256(abi.encodePacked(
  //     'Order(',
  //     'bytes32 key,',
  //     'address owner,',
  //     'address underlying,',
  //     'bool fixed',
  //     'uint256 rate,',
  //     'uint256 principal,',
  //     'uint256 interest,',
  //     'uint256 duration,',
  //     'uint256 expiry,',
  //     'uint256 nonce,',
  //     ')'
  // ));
  // TODO doube check
  bytes32 constant internal ORDER_TYPEHASH = 0xb64e83290881c04f7b95baa5dd00436e04f41d002805882789a249a4defea5dc;

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

  /// @param k Key of the Order
  /// @param o Owner of the order
  /// @param u Address of the underlying token contract
  /// @param f Boolean indicating if the Order is fixed or not
  /// @param p Array of integers ordered as per the Params enum
  function order(bytes32 k, address o, address u, bool f, uint256[6] calldata p) internal pure returns (bytes32) {
    // TODO assembly
    return keccak256(abi.encode(
      ORDER_TYPEHASH,
      k, o, u, f,
      p[uint(Params.RATE)],
      p[uint(Params.PRINCIPAL)],
      p[uint(Params.INTEREST)],
      p[uint(Params.DURATION)],
      p[uint(Params.EXPIRY)],
      p[uint(Params.NONCE)]
    ));
  }
}
