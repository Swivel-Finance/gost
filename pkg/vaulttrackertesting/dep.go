package vaulttrackertesting

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/vaulttracker"
)

type Dep struct {
	VaultTrackerAddress common.Address
	VaultTracker        *vaulttracker.VaultTracker
}

func Deploy(e *Env) (*Dep, error) {
	// TODO deploy mock tokens?

	bsUnderlying := common.HexToAddress("0xnotunderlying")
	bsCToken := common.HexToAddress("0xnotctoken")

	// deploy contract...
	trackerAddress, _, trackerContract, trackerErr := vaulttracker.DeployVaultTracker(
		e.Owner.Opts,
		e.Blockchain,
		bsUnderlying,
		big.NewInt(123456789), // bs maturity
		bsCToken,
	)

	if trackerErr != nil {
		return nil, trackerErr
	}

	e.Blockchain.Commit()

	return &Dep{
		VaultTrackerAddress: trackerAddress,
		VaultTracker:        trackerContract,
	}, nil
}
