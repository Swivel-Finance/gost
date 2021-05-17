package vaulttrackertesting

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/vaulttracker"
	"math/big"
)

type Dep struct {
	CErc20        *mocks.CErc20
	CErc20Address common.Address

	VaultTracker        *vaulttracker.VaultTracker
	VaultTrackerAddress common.Address

	Maturity      *big.Int
}

func Deploy(e *Env) (*Dep, error) {
	err := e.Blockchain.AdjustTime(0) // set bc timestamp to 0
	if err != nil {
		panic(err)
	}
	e.Blockchain.Commit()

	maturity := big.NewInt(MATURITY)
	cercAddress, _, cercContract, cercErr := mocks.DeployCErc20(e.Owner.Opts, e.Blockchain)

	if cercErr != nil {
		return nil, cercErr
	}

	e.Blockchain.Commit()

	// deploy contract...
	trackerAddress, _, trackerContract, trackerErr := vaulttracker.DeployVaultTracker(
		e.Owner.Opts,
		e.Blockchain,
		maturity,
		cercAddress,
	)

	if trackerErr != nil {
		return nil, trackerErr
	}

	e.Blockchain.Commit()

	return &Dep{
		VaultTrackerAddress: trackerAddress,
		VaultTracker:        trackerContract,
		Maturity:            maturity,
		CErc20:              cercContract,
		CErc20Address:       cercAddress,
	}, nil
}
