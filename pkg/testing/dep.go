package testing

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/mocks"
)

type Dep struct {
	Erc20Address            common.Address
	Erc20                   *mocks.Erc20
	YieldTokenAddress       common.Address
	YieldToken              *mocks.YieldToken
	ZcTokenAddress          common.Address
	ZcToken                 *mocks.ZcToken
	SwivelAddress           common.Address
	Swivel                  *mocks.Swivel
	ElementTokenAddress     common.Address
	ElementToken            *mocks.ElementToken
	ElementAddress          common.Address
	Element                 *mocks.Element
	PendleYieldToken        *mocks.PendleYieldToken
	PendleYieldTokenAddress common.Address
	PendleRouter            *mocks.PendleRouter
	PendleRouterAddress     common.Address
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

	zcAddress, _, zcContract, zcErr := mocks.DeployZcToken(e.Owner.Opts, e.Blockchain)

	if zcErr != nil {
		return nil, zcErr
	}

	e.Blockchain.Commit()

	swAddress, _, swContract, swErr := mocks.DeploySwivel(e.Owner.Opts, e.Blockchain)

	if swErr != nil {
		return nil, swErr
	}

	e.Blockchain.Commit()

	elementTokenAddress, _, elementTokenContract, elementTokenErr := mocks.DeployElementToken(e.Owner.Opts, e.Blockchain)

	if elementTokenErr != nil {
		return nil, elementTokenErr
	}

	elementAddress, _, elementContract, elementErr := mocks.DeployElement(e.Owner.Opts, e.Blockchain)

	if elementErr != nil {
		return nil, elementErr
	}

	e.Blockchain.Commit()

	pytAddress, _, pytContract, pytErr := mocks.DeployPendleYieldToken(e.Owner.Opts, e.Blockchain)

	if pytErr != nil {
		return nil, pytErr
	}

	e.Blockchain.Commit()

	prAddress, _, prContract, prErr := mocks.DeployPendleRouter(e.Owner.Opts, e.Blockchain)

	if prErr != nil {
		return nil, prErr
	}

	e.Blockchain.Commit()

	return &Dep{
		Erc20Address:            ercAddress,
		Erc20:                   ercContract,
		YieldTokenAddress:       ytAddress,
		YieldToken:              ytContract,
		ZcTokenAddress:          zcAddress,
		ZcToken:                 zcContract,
		SwivelAddress:           swAddress,
		Swivel:                  swContract,
		ElementTokenAddress:     elementTokenAddress,
		ElementToken:            elementTokenContract,
		ElementAddress:          elementAddress,
		Element:                 elementContract,
		PendleYieldTokenAddress: pytAddress,
		PendleYieldToken:        pytContract,
		PendleRouter:            prContract,
		PendleRouterAddress:     prAddress,
	}, nil
}
