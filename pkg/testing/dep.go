package testing

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/fakes"
	"github.com/swivel-finance/gost/test/mocks"
)

type Dep struct {
	SigFakeAddress       common.Address
	SigFake              *fakes.SigFake // fake sig lib test contract
	HashFakeAddress      common.Address
	HashFake             *fakes.HashFake // fake hash lib test contract
	Erc20Address         common.Address
	Erc20                *mocks.Erc20
	CompoundTokenAddress common.Address
	CompoundToken        *mocks.CompoundToken
	// FErc20Address   common.Address
	// FErc20          *mocks.FErc20
}

func Deploy(e *Env) (*Dep, error) {
	// deploying the lib testing contract "fakes"
	// NOTE these _could_ be moved into their own package as they are not needed
	// for swivel to operate. TODO
	sigAddress, _, sigContract, sigErr := fakes.DeploySigFake(e.Owner.Opts, e.Blockchain)

	if sigErr != nil {
		return nil, sigErr
	}

	e.Blockchain.Commit()

	hashAddress, _, hashContract, hashErr := fakes.DeployHashFake(e.Owner.Opts, e.Blockchain)

	if hashErr != nil {
		return nil, hashErr
	}

	e.Blockchain.Commit()

	// deploy the two mock tokens.
	ercAddress, _, ercContract, ercErr := mocks.DeployErc20(e.Owner.Opts, e.Blockchain)

	if ercErr != nil {
		return nil, ercErr
	}

	e.Blockchain.Commit()

	ctAddress, _, ctContract, ctErr := mocks.DeployCompoundToken(e.Owner.Opts, e.Blockchain)

	if ctErr != nil {
		return nil, ctErr
	}

	e.Blockchain.Commit()

	return &Dep{
		SigFakeAddress:       sigAddress,
		SigFake:              sigContract,
		HashFakeAddress:      hashAddress,
		HashFake:             hashContract,
		Erc20Address:         ercAddress,
		Erc20:                ercContract,
		CompoundTokenAddress: ctAddress,
		CompoundToken:        ctContract,
	}, nil
}
