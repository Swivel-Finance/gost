package vaulttrackertesting

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/vaulttracker"
)

type Dep struct {
	Compounding        *mocks.Compound // deployed Compound Adapter, implements ICompounding
	CompoundingAddress common.Address

	CompoundingTokenAddress common.Address
	SwivelAddress           common.Address

	VaultTracker        *vaulttracker.VaultTracker
	VaultTrackerAddress common.Address

	Maturity *big.Int
}

func Deploy(e *Env) (*Dep, error) {
	maturity := big.NewInt(MATURITY)
	compAddress, _, compContract, compErr := mocks.DeployCompound(e.Owner.Opts, e.Blockchain)

	if compErr != nil {
		return nil, compErr
	}

	e.Blockchain.Commit()

	// compounding token address can be mocked (contract communicates with the adapter, not the token directly)
	ctAddress := common.HexToAddress("0x123Abc")
	// swivel address can be mocked
	swivelAddress := common.HexToAddress("0x234bCd")

	// deploy contract...
	trackerAddress, _, trackerContract, trackerErr := vaulttracker.DeployVaultTracker(
		e.Owner.Opts,
		e.Blockchain,
		maturity,
		ctAddress,
		compAddress,
		swivelAddress,
	)

	if trackerErr != nil {
		return nil, trackerErr
	}

	e.Blockchain.Commit()

	return &Dep{
		Compounding:             compContract,
		CompoundingAddress:      compAddress,
		CompoundingTokenAddress: ctAddress,
		Maturity:                maturity,
		SwivelAddress:           swivelAddress,
		VaultTrackerAddress:     trackerAddress,
		VaultTracker:            trackerContract,
	}, nil
}
