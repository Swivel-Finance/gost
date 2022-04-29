package lendertesting

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/mocks"
)

type Dep struct {
	Erc20Address       common.Address
	Erc20              *mocks.Erc20
	YieldTokenAddress  common.Address
	YieldToken         *mocks.YieldToken
	MarketPlaceAddress common.Address
	MarketPlace        *mocks.MarketPlace
	ZcTokenAddress     common.Address
	ZcToken            *mocks.ZcToken
}

func Deploy(e *Env) (*Dep, error) {
	// deploy the two mock tokens.
	ercAddress, _, ercContract, ercErr := mocks.DeployErc20(e.Owner.Opts, e.Blockchain)

	if ercErr != nil {
		return nil, ercErr
	}

	e.Blockchain.Commit()

	ytAddress, _, ytContract, ytErr := mocks.DeployYieldToken(e.Owner.Opts, e.Blockchain)

	if ytErr != nil {
		return nil, ytErr
	}

	e.Blockchain.Commit()

	mpAddress, _, mpContract, mpErr := mocks.DeployMarketPlace(e.Owner.Opts, e.Blockchain)

	if mpErr != nil {
		return nil, mpErr
	}

	e.Blockchain.Commit()

	zcAddress, _, zcContract, zcErr := mocks.DeployZcToken(e.Owner.Opts, e.Blockchain)

	if zcErr != nil {
		return nil, zcErr
	}

	e.Blockchain.Commit()

	return &Dep{
		Erc20Address:       ercAddress,
		Erc20:              ercContract,
		YieldTokenAddress:  ytAddress,
		YieldToken:         ytContract,
		MarketPlaceAddress: mpAddress,
		MarketPlace:        mpContract,
		ZcTokenAddress:     zcAddress,
		ZcToken:            zcContract,
	}, nil
}
