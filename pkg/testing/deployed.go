package testing

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/swivel-finance/gost/test/contracts/swivel"
)

type Dep struct {
	SigTestAddress      common.Address
	SigTest             *swivel.SigTest
	SigTestTransaction  *types.Transaction // in case we want to view the deployment gas cost etc...
	HashTestAddress     common.Address
	HashTest            *swivel.HashTest
	HashTestTransaction *types.Transaction // in case we want to view the deployment gas cost etc...
}

func Deploy(e *Env) (*Dep, error) {
	sigAddress, sigTx, sigContract, sigErr := swivel.DeploySigTest(e.Owner.Opts, e.Blockchain)

	if sigErr != nil {
		return nil, sigErr
	}

	e.Blockchain.Commit()

	hashAddress, hashTx, hashContract, hashErr := swivel.DeployHashTest(e.Owner.Opts, e.Blockchain)

	if hashErr != nil {
		return nil, hashErr
	}

	e.Blockchain.Commit()

	return &Dep{
		SigTestAddress:      sigAddress,
		SigTest:             sigContract,
		SigTestTransaction:  sigTx,
		HashTestAddress:     hashAddress,
		HashTest:            hashContract,
		HashTestTransaction: hashTx,
	}, nil
}
