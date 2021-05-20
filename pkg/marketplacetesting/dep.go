package marketplacetesting

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/marketplace"
	"github.com/swivel-finance/gost/test/mocks"
)

// TODO mock for marketplace...
type Dep struct {
	CErc20        *mocks.CErc20
	CErc20Address common.Address

	MarketPlaceAddress common.Address
	MarketPlace        *marketplace.MarketPlace

	Maturity *big.Int
}

func Deploy(e *Env) (*Dep, error) {
	maturity := big.NewInt(MATURITY)
	cercAddress, _, cercContract, cercErr := mocks.DeployCErc20(e.Owner.Opts, e.Blockchain)

	if cercErr != nil {
		return nil, cercErr
	}

	e.Blockchain.Commit()

	// deploy contract...
	marketAddress, _, marketContract, marketErr := marketplace.DeployMarketPlace(e.Owner.Opts, e.Blockchain)

	if marketErr != nil {
		return nil, marketErr
	}

	e.Blockchain.Commit()

	return &Dep{
		MarketPlaceAddress: marketAddress,
		MarketPlace:        marketContract,
		CErc20Address:      cercAddress,
		CErc20:             cercContract,
		Maturity:           maturity,
	}, nil
}
