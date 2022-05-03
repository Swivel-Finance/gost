package lendertesting

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/lender"
	"github.com/swivel-finance/gost/test/mocks"
)

type Dep struct {
	Erc20Address       common.Address
	Erc20              *mocks.Erc20
	YieldTokenAddress  common.Address
	YieldToken         *mocks.YieldToken
	MarketPlaceAddress common.Address
	MarketPlace        *mocks.MarketPlace
	LenderAddress      common.Address
	Lender             *lender.Lender
	ZcTokenAddress     common.Address
	ZcToken            *mocks.ZcToken
	SwivelAddress      common.Address
	Swivel             *mocks.Swivel
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

	swAddress, _, swContract, swErr := mocks.DeploySwivel(e.Owner.Opts, e.Blockchain)

	if swErr != nil {
		return nil, swErr
	}

	e.Blockchain.Commit()

	lenderAddress, _, lender, lenderErr := lender.DeployLender(e.Owner.Opts, e.Blockchain, mpAddress, swAddress)
	if lenderErr != nil {
		return nil, lenderErr
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
		LenderAddress:      lenderAddress,
		Lender:             lender,
		ZcTokenAddress:     zcAddress,
		ZcToken:            zcContract,
		SwivelAddress:      swAddress,
		Swivel:             swContract,
	}, nil
}
