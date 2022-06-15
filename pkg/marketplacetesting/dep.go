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

	CompoundToken        *mocks.CompoundToken
	CompoundTokenAddress common.Address

	MarketPlaceAddress common.Address
	MarketPlace        *marketplace.MarketPlace

	Maturity *big.Int

	SwivelAddress common.Address
}

func Deploy(e *Env) (*Dep, error) {
	maturity := big.NewInt(MATURITY)
	ctAddress, _, ctContract, ctErr := mocks.DeployCompoundToken(e.Owner.Opts, e.Blockchain)

	if ctErr != nil {
		return nil, ctErr
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
		CompoundTokenAddress: ctAddress,
		CompoundToken:        ctContract,
		Erc20Address:         ercAddress,
		Erc20:                ercContract,
		SwivelAddress:        common.HexToAddress("0xAbC123"),
		Maturity:             maturity,
	}, nil
}
