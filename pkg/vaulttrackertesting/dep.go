package vaulttrackertesting

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/vaulttracker"
)

type Dep struct {
	CErc20              *mocks.CErc20
	CErc20Address       common.Address
	VaultTracker        *vaulttracker.VaultTracker
	VaultTrackerAddress common.Address
	Maturity            *big.Int
}

func Deploy(e *Env) (*Dep, error) {
	bsMaturity := big.NewInt(17)

	cercAddress, _, cercContract, cercErr := mocks.DeployCErc20(e.Owner.Opts, e.Blockchain)

	if cercErr != nil {
		return nil, cercErr
	}

	e.Blockchain.Commit()

	// deploy contract...
	trackerAddress, _, trackerContract, trackerErr := vaulttracker.DeployVaultTracker(
		e.Owner.Opts,
		e.Blockchain,
		bsMaturity,
		cercAddress,
	)

	if trackerErr != nil {
		return nil, trackerErr
	}

	e.Blockchain.Commit()

	return &Dep{
		VaultTrackerAddress: trackerAddress,
		VaultTracker:        trackerContract,
		Maturity:            bsMaturity,
		CErc20:              cercContract,
		CErc20Address:       cercAddress,
	}, nil
}
