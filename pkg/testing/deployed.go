package testing

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/swivel-finance/gost/test/contracts/swivel"
)

type Dep struct {
	SigAddress     common.Address
	Sig            *swivel.Sig
	SigTransaction *types.Transaction // in case we want to view the deployment gas cost etc...
}

func Deploy(e *Env) (*Dep, error) {
	// deploy with the owner auth object
	address, tx, contract, err := swivel.DeploySig(e.Owner.Opts, e.Blockchain)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	return &Dep{
		SigAddress:     address,
		Sig:            contract,
		SigTransaction: tx,
	}, nil
}
