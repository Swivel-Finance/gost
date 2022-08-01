package zctokentesting

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/mocks"
)

type Dep struct {
	MarketPlace        *mocks.MarketPlace
	MarketPlaceAddress common.Address
	Maturity           *big.Int
}

func Deploy(e *Env) (*Dep, error) {
	maturity := big.NewInt(MATURITY)

	mpAddress, _, mpContract, mpErr := mocks.DeployMarketPlace(e.Owner.Opts, e.Blockchain)

	if mpErr != nil {
		return nil, mpErr
	}

	e.Blockchain.Commit()

	// NOTE ZcTokens will be deployed per spec...

	return &Dep{
		Maturity:           maturity,
		MarketPlaceAddress: mpAddress,
		MarketPlace:        mpContract,
	}, nil
}
