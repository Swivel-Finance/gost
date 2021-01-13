// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.0;

library Sig {
  /// @param h Hashed data which was originally signed
  /// @param v Verifying bit taken from a valid ECDSA signature
  /// @param r First 32bytes split from an ECDSA signature
  /// @param s Next 32bytes split from an ECDSA signature
  /// @return The recovered address
  function recover(bytes32 h, uint8 v, bytes32 r, bytes32 s) internal pure returns (address) {
    // EIP-2 and malleable signatures...
    // see https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/cryptography/ECDSA.sol
    require(uint256(s) <= 0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0, 'invalid signature "s" value');
    require(v == 27 || v == 28, 'invalid signature "v" value');

    return ecrecover(h, v, r, s);
  }

  /// @param h Hashed data which was originally signed
  /// @param sig Valid ECDSA signature
  /// @dev splitAndRecover should only be used if it is known that the resulting 
  /// verifying bit (V) will be 27 || 28. Otherwise use recover, possibly calling split first.
  /// @return The recovered address
  function splitAndRecover(bytes32 h, bytes memory sig) internal pure returns (address) {
    (uint8 v, bytes32 r, bytes32 s) = split(sig);

    return ecrecover(h, v, r, s);
  }

  /// @param sig Valid ECDSA signature
  /// @return v The verification bit
  /// @return r First 32 bytes
  /// @return s Next 32 bytes
  function split(bytes memory sig) internal pure returns (uint8 v, bytes32 r, bytes32 s) {
    require(sig.length == 65, 'invalid signature length'); // TODO standardize error messages

    assembly {
      r := mload(add(sig, 32))
      s := mload(add(sig, 64))
      v := byte(0, mload(add(sig, 96)))
    }

    return (v, r, s);
  }
}
