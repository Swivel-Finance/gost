package vaulttrackertesting

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/vaulttracker"
)

type Dep struct {
	CompoundToken        *mocks.CompoundToken
	CompoundTokenAddress common.Address

	SwivelAddress common.Address

	VaultTracker        *vaulttracker.VaultTracker
	VaultTrackerAddress common.Address

	Maturity *big.Int
}

func Deploy(e *Env) (*Dep, error) {
	maturity := big.NewInt(MATURITY)

	ctAddress, _, ctContract, ctErr := mocks.DeployCompoundToken(e.Owner.Opts, e.Blockchain)

	if ctErr != nil {
		return nil, ctErr
	}

	e.Blockchain.Commit()

	// swivel address can be mocked
	swivelAddress := common.HexToAddress("0x234bCd")

	// deploy contract...NOTE that the vaulttracker has to be deployed with a particular protocol. deploy them separately to mitigate
	trackerAddress, _, trackerContract, trackerErr := vaulttracker.DeployVaultTracker(
		e.Owner.Opts,
		e.Blockchain,
		uint8(1), // Compound for now TODO change to 0 when Erc4626 is in place
		maturity,
		ctAddress,
		swivelAddress,
	)

	if trackerErr != nil {
		return nil, trackerErr
	}

	e.Blockchain.Commit()

	return &Dep{
		CompoundToken:        ctContract,
		CompoundTokenAddress: ctAddress,
		Maturity:             maturity,
		SwivelAddress:        swivelAddress,
		VaultTrackerAddress:  trackerAddress,
		VaultTracker:         trackerContract,
	}, nil
}
