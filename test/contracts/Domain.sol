// SPDX-License-Identifier: UNLICENSED

// NOTE this contract is simply a placeholder for logic that does into the Swivel.sol constructor

pragma solidity 0.8.0;

/**
  @dev Domain holds the computed (EIP712) hashed domain of the Swivel Finance Protocol.
*/

import './Hash.sol';

contract Domain {
  string constant internal NAME = "Swivel Finance";
  string constant internal VERSION = "1.0.0";
  // while technically not a const, DOMAIN_HASH is treated as one.
  // TODO solhint
  bytes32 public HASH;

  /// @param c Chain ID of the network this contract is deployed to
  /// @param v [optional] Verifying contract if present. Defaults to this.
  constructor(uint256 c, address v) public {
    address verifier = v == address(0) ? address(this) : v;
    HASH = Hash.domain(NAME, VERSION, c, v);
  }
}
