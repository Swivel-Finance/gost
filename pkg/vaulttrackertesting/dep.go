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
	Erc4626              *mocks.Erc4626
	Erc4626Address       common.Address
	YearnVault           *mocks.YearnVault
	YearnVaultAddress    common.Address
	AaveToken            *mocks.AaveToken
	AaveTokenAddress     common.Address
	AavePool             *mocks.AavePool
	AavePoolAddress      common.Address
	EulerToken           *mocks.EulerToken
	EulerTokenAddress    common.Address
	SwivelAddress        common.Address
	VaultTracker         *vaulttracker.VaultTracker
	VaultTrackerAddress  common.Address
	Maturity             *big.Int
}

func Deploy(e *Env) (*Dep, error) {
	maturity := big.NewInt(MATURITY)

	ctAddress, _, ctContract, ctErr := mocks.DeployCompoundToken(e.Owner.Opts, e.Blockchain)

	if ctErr != nil {
		return nil, ctErr
	}

	e.Blockchain.Commit()

	erc4626Address, _, erc4626Contract, erc4626Err := mocks.DeployErc4626(e.Owner.Opts, e.Blockchain)

	if erc4626Err != nil {
		return nil, erc4626Err
	}

	e.Blockchain.Commit()

	yvAddress, _, yvContract, yvErr := mocks.DeployYearnVault(e.Owner.Opts, e.Blockchain)

	if yvErr != nil {
		return nil, yvErr
	}

	e.Blockchain.Commit()

	atAddress, _, atContract, atErr := mocks.DeployAaveToken(e.Owner.Opts, e.Blockchain)

	if atErr != nil {
		return nil, atErr
	}

	e.Blockchain.Commit()

	apAddress, _, apContract, apErr := mocks.DeployAavePool(e.Owner.Opts, e.Blockchain)

	if apErr != nil {
		return nil, apErr
	}

	etAddress, _, etContract, etErr := mocks.DeployEulerToken(e.Owner.Opts, e.Blockchain)

	if etErr != nil {
		return nil, etErr
	}

	e.Blockchain.Commit()

	e.Blockchain.Commit()

	// swivel address can be mocked
	swivelAddress := common.HexToAddress("0x234bCd")

	trackerAddress, _, trackerContract, trackerErr := vaulttracker.DeployVaultTracker(
		e.Owner.Opts,
		e.Blockchain,
		uint8(1), // Compound here
		maturity,
		ctAddress,
		swivelAddress,
		e.Owner.Opts.From, // we use the bound session owner as "marketPlace" here as it will be the caller in these unit tests
	)

	if trackerErr != nil {
		return nil, trackerErr
	}

	e.Blockchain.Commit()

	return &Dep{
		CompoundToken:        ctContract,
		CompoundTokenAddress: ctAddress,
		Erc4626:              erc4626Contract,
		Erc4626Address:       erc4626Address,
		YearnVault:           yvContract,
		YearnVaultAddress:    yvAddress,
		AaveToken:            atContract,
		AaveTokenAddress:     atAddress,
		AavePool:             apContract,
		AavePoolAddress:      apAddress,
		EulerToken:           etContract,
		EulerTokenAddress:    etAddress,
		Maturity:             maturity,
		SwivelAddress:        swivelAddress,
		VaultTrackerAddress:  trackerAddress,
		VaultTracker:         trackerContract,
	}, nil
}
