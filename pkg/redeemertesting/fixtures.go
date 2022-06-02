package redeemertesting

import "github.com/ethereum/go-ethereum/common"

/// @notice Convenience method that returns a list of repeated addresses
/// @param a: the address used to represent the zc token
/// @param b: the address that fills the rest of the list
/// @return: list of repeated addresses
func marketsList(a common.Address, b common.Address) [9]common.Address {
	return [9]common.Address{a, b, b, b, b, b, b, b, b}
}
