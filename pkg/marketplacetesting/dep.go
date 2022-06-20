package marketplacetesting

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/marketplace"
	"github.com/swivel-finance/gost/test/mocks"
)

type Dep struct {
	Erc20        *mocks.Erc20
	Erc20Address common.Address

	Compounding        *mocks.Compound
	CompoundingAddress common.Address

	MarketPlaceAddress common.Address
	MarketPlace        *marketplace.MarketPlace

	Maturity *big.Int

	SwivelAddress        common.Address
	CompoundTokenAddress common.Address
}

func Deploy(e *Env) (*Dep, error) {
	maturity := big.NewInt(MATURITY)
	compAddress, _, compContract, compErr := mocks.DeployCompound(e.Owner.Opts, e.Blockchain)

	if compErr != nil {
		return nil, compErr
	}

	e.Blockchain.Commit()

	ercAddress, _, ercContract, ercErr := mocks.DeployErc20(e.Owner.Opts, e.Blockchain)

	if ercErr != nil {
		return nil, ercErr
	}

	e.Blockchain.Commit()

	// deploy contract...
	marketAddress, _, marketContract, marketErr := marketplace.DeployMarketPlace(e.Owner.Opts, e.Blockchain)

	if marketErr != nil {
		return nil, marketErr
	}

	e.Blockchain.Commit()

	return &Dep{
		MarketPlaceAddress:   marketAddress,
		MarketPlace:          marketContract,
		CompoundingAddress:   compAddress,
		Compounding:          compContract,
		Erc20Address:         ercAddress,
		Erc20:                ercContract,
		SwivelAddress:        common.HexToAddress("0x123aBc"),
		CompoundTokenAddress: common.HexToAddress("0x234Bcd"),
		Maturity:             maturity,
	}, nil
}
