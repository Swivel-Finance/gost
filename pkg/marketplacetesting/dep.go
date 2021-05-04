package marketplacetesting

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/marketplace"
)

// TODO mock for marketplace...
type Dep struct {
	MarketPlaceAddress common.Address
	MarketPlace        *marketplace.MarketPlace
}

func Deploy(e *Env) (*Dep, error) {
	// deploy contract...
	marketAddress, _, marketContract, marketErr := marketplace.DeployMarketPlace(e.Owner.Opts, e.Blockchain)

	if marketErr != nil {
		return nil, marketErr
	}

	e.Blockchain.Commit()

	return &Dep{
		MarketPlaceAddress: marketAddress,
		MarketPlace:        marketContract,
	}, nil
}
