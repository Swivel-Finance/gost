package creatortesting

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/creator"
)

type Dep struct {
	MarketPlaceAddress common.Address
	CreatorAddress     common.Address
	Creator            *creator.Creator
	Maturity           *big.Int
}

func Deploy(e *Env) (*Dep, error) {
	maturity := big.NewInt(MATURITY)
	mpAddress := common.HexToAddress("0x12345")

	cAddress, _, cContract, err := creator.DeployCreator(e.Owner.Opts, e.Blockchain)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	// NOTE ZcTokens will be deployed per spec...

	return &Dep{
		Maturity:           maturity,
		MarketPlaceAddress: mpAddress,
		CreatorAddress:     cAddress,
		Creator:            cContract,
	}, nil
}
