package testing

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/swivel-finance/gost/test/contracts/token"
)

type Dep struct {
	TokenAddress        common.Address
	Token               *token.Token
	DeployedTransaction *types.Transaction // in case we want to view the deployment gas cost etc...
}

func Deploy(e *Env) (*Dep, error) {
	address, tx, contract, err := token.DeployToken(e.Owner, e.Blockchain, GOST_TOKEN_DECIMALS, GOST_TOKEN_SYMBOL)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	return &Dep{
		TokenAddress:        address,
		Token:               contract,
		DeployedTransaction: tx,
	}, nil
}
