package illuminatetesting

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/build/marketplace"
	"github.com/swivel-finance/gost/test/mocks"
)

type Dep struct {
	Erc20Address       common.Address
	Erc20              *mocks.Erc20
	RedeemerAddress    common.Address
	MarketPlaceAddress common.Address
	MarketPlace        *marketplace.MarketPlace
}

func Deploy(e *Env) (*Dep, error) {
	// deploy mock tokens.
	ercAddress, _, ercContract, ercErr := mocks.DeployErc20(e.Owner.Opts, e.Blockchain)

	if ercErr != nil {
		return nil, ercErr
	}

	e.Blockchain.Commit()

	rAddr := common.HexToAddress("0x123")

	mpAddress, _, mpContract, mpErr := marketplace.DeployMarketPlace(e.Owner.Opts, e.Blockchain, rAddr)
	if mpErr != nil {
		return nil, mpErr
	}

	e.Blockchain.Commit()

	return &Dep{
		Erc20Address:       ercAddress,
		Erc20:              ercContract,
		RedeemerAddress:    rAddr,
		MarketPlaceAddress: mpAddress,
		MarketPlace:        mpContract,
	}, nil
}
