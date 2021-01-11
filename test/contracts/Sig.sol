// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.0;

library Sig {
  /**
    @dev recover is passed v,r and s separately. This accounts for the need some environments will have
    to inspect and correct the version bit ahead of time
  */
  function recover(bytes32 h, uint8 v, bytes32 r, bytes32 s) public pure returns (address) {
    // EIP-2 and malleable signatures...
    // see https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/cryptography/ECDSA.sol
    require(uint256(s) <= 0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0, 'invalid signature "s" value');
    require(v == 27 || v == 28, 'invalid signature "v" value');

    return ecrecover(h, v, r, s);
  }

  /**
    @dev splitAndRecover is present for use if the developer is certain the version bit
    resulting from Sig.split will be 27 or 28, thus saving the need for 2 calls.
    NOTE: while it is true solidity provides for function overloading, we are intentionally avoiding it.
  */
  function splitAndRecover(bytes32 h, bytes memory sig) public pure returns (address) {
    (uint8 v, bytes32 r, bytes32 s) = split(sig);

    return ecrecover(h, v, r, s);
  }

  function split(bytes memory sig) public pure returns (uint8 v, bytes32 r, bytes32 s) {
    require(sig.length == 65, 'invalid signature length'); // TODO standardize error messages

    assembly {
      r := mload(add(sig, 32))
      s := mload(add(sig, 64))
      v := byte(0, mload(add(sig, 96)))
    }


    return (v, r, s);
  }
}
