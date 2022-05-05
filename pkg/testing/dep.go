package testing

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/mocks"
)

type Dep struct {
	Erc20Address        common.Address
	Erc20               *mocks.Erc20
	YieldAddress        common.Address
	Yield               *mocks.Yield
	ZcTokenAddress      common.Address
	ZcToken             *mocks.ZcToken
	SwivelAddress       common.Address
	Swivel              *mocks.Swivel
	ElementTokenAddress common.Address
	ElementToken        *mocks.ElementToken
	ElementAddress      common.Address
	Element             *mocks.Element
}

func Deploy(e *Env) (*Dep, error) {
	// deploy the two mock tokens.
	ercAddress, _, ercContract, ercErr := mocks.DeployErc20(e.Owner.Opts, e.Blockchain)

	if ercErr != nil {
		return nil, ercErr
	}

	e.Blockchain.Commit()

	ytAddress, _, ytContract, ytErr := mocks.DeployYield(e.Owner.Opts, e.Blockchain)

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

	return &Dep{
		Erc20Address:        ercAddress,
		Erc20:               ercContract,
		YieldAddress:        ytAddress,
		Yield:               ytContract,
		ZcTokenAddress:      zcAddress,
		ZcToken:             zcContract,
		SwivelAddress:       swAddress,
		Swivel:              swContract,
		ElementTokenAddress: elementTokenAddress,
		ElementToken:        elementTokenContract,
		ElementAddress:      elementAddress,
		Element:             elementContract,
	}, nil
}
