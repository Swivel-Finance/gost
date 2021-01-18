package testing

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/contracts/fakes"
)

type Dep struct {
	SigFakeAddress  common.Address
	SigFake         *fakes.SigFake
	HashFakeAddress common.Address
	HashFake        *fakes.HashFake
}

func Deploy(e *Env) (*Dep, error) {
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

	return &Dep{
		SigFakeAddress:  sigAddress,
		SigFake:         sigContract,
		HashFakeAddress: hashAddress,
		HashFake:        hashContract,
	}, nil
}
