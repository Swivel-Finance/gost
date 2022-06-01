package lendertesting

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/lender"
	"github.com/swivel-finance/gost/test/mocks"
)

type Dep struct {
	Erc20Address        common.Address
	Erc20               *mocks.Erc20
	YieldAddress        common.Address
	Yield               *mocks.Yield
	MarketPlaceAddress  common.Address
	MarketPlace         *mocks.MarketPlace
	LenderAddress       common.Address
	Lender              *lender.Lender
	ZcTokenAddress      common.Address
	ZcToken             *mocks.ZcToken
	SwivelAddress       common.Address
	Swivel              *mocks.Swivel
	ElementTokenAddress common.Address
	ElementToken        *mocks.ElementToken
	ElementAddress      common.Address
	Element             *mocks.Element
	PendleTokenAddress  common.Address
	PendleToken         *mocks.PendleToken
	PendleAddress       common.Address
	Pendle              *mocks.Pendle
	TempusAddress       common.Address
	Tempus              *mocks.Tempus
	SenseAddress        common.Address
	Sense               *mocks.Sense
	SenseTokenAddress   common.Address
	SenseToken          *mocks.SenseToken
	APWineToken         *mocks.APWineToken
	APWineTokenAddress  common.Address
	APWine              *mocks.APWine
	APWineAddress       common.Address
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

	zcAddress, _, zcContract, zcErr := mocks.DeployZcToken(e.Owner.Opts, e.Blockchain)

	if zcErr != nil {
		return nil, zcErr
	}

	e.Blockchain.Commit()

	elementTokenAddress, _, elementTokenContract, elementTokenErr := mocks.DeployElementToken(e.Owner.Opts, e.Blockchain)

	if elementTokenErr != nil {
		return nil, elementTokenErr
	}

	e.Blockchain.Commit()

	elementAddress, _, elementContract, elementErr := mocks.DeployElement(e.Owner.Opts, e.Blockchain)

	if elementErr != nil {
		return nil, elementErr
	}

	e.Blockchain.Commit()

	pAddress, _, pContract, pErr := mocks.DeployPendleToken(e.Owner.Opts, e.Blockchain)

	if pErr != nil {
		return nil, pErr
	}

	e.Blockchain.Commit()

	peAddress, _, peContract, peErr := mocks.DeployPendle(e.Owner.Opts, e.Blockchain)

	if peErr != nil {
		return nil, peErr
	}

	e.Blockchain.Commit()

	tAddress, _, tContract, tErr := mocks.DeployTempus(e.Owner.Opts, e.Blockchain)

	if tErr != nil {
		return nil, tErr
	}

	e.Blockchain.Commit()

	seAddress, _, seContract, seErr := mocks.DeploySense(e.Owner.Opts, e.Blockchain)

	if seErr != nil {
		return nil, seErr
	}

	e.Blockchain.Commit()

	senseTokenAddress, _, senseTokenContract, senseTokenErr := mocks.DeploySenseToken(e.Owner.Opts, e.Blockchain)

	if senseTokenErr != nil {
		return nil, senseTokenErr
	}

	e.Blockchain.Commit()

	apAddress, _, apContract, apErr := mocks.DeployAPWineToken(e.Owner.Opts, e.Blockchain)

	if apErr != nil {
		return nil, apErr
	}

	e.Blockchain.Commit()

	aprAddress, _, aprContract, aprErr := mocks.DeployAPWine(e.Owner.Opts, e.Blockchain)

	if aprErr != nil {
		return nil, aprErr
	}

	e.Blockchain.Commit()

	lenderAddress, _, lender, lenderErr := lender.DeployLender(e.Owner.Opts, e.Blockchain, swAddress, peAddress, tAddress)
	if lenderErr != nil {
		return nil, lenderErr
	}

	e.Blockchain.Commit()

	return &Dep{
		Erc20Address:        ercAddress,
		Erc20:               ercContract,
		YieldAddress:        ytAddress,
		Yield:               ytContract,
		MarketPlaceAddress:  mpAddress,
		MarketPlace:         mpContract,
		LenderAddress:       lenderAddress,
		Lender:              lender,
		ZcTokenAddress:      zcAddress,
		ZcToken:             zcContract,
		SwivelAddress:       swAddress,
		Swivel:              swContract,
		ElementTokenAddress: elementTokenAddress,
		ElementToken:        elementTokenContract,
		ElementAddress:      elementAddress,
		Element:             elementContract,
		PendleTokenAddress:  pAddress,
		PendleToken:         pContract,
		PendleAddress:       peAddress,
		Pendle:              peContract,
		TempusAddress:       tAddress,
		Tempus:              tContract,
		SenseAddress:        seAddress,
		Sense:               seContract,
		SenseTokenAddress:   senseTokenAddress,
		SenseToken:          senseTokenContract,
		APWineToken:         apContract,
		APWineTokenAddress:  apAddress,
		APWineAddress:       aprAddress,
		APWine:              aprContract,
	}, nil
}
