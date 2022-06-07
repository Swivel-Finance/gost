package redeemertesting

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/lender"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/redeemer"
)

type Dep struct {
	Erc20Address        common.Address
	Erc20               *mocks.Erc20
	YieldAddress        common.Address
	Yield               *mocks.Yield
	YieldToken          *mocks.YieldToken
	YieldTokenAddress   common.Address
	ZcTokenAddress      common.Address
	ZcToken             *mocks.ZcToken
	SwivelAddress       common.Address
	Swivel              *mocks.Swivel
	MarketplaceAddress  common.Address
	Marketplace         *mocks.MarketPlace
	APWineAddress       common.Address
	APWine              *mocks.APWine
	APWineTokenAddress  common.Address
	APWineToken         *mocks.APWineToken
	TempusAddress       common.Address
	Tempus              *mocks.Tempus
	TempusTokenAddress  common.Address
	TempusToken         *mocks.TempusToken
	RedeemerAddress     common.Address
	Redeemer            *redeemer.Redeemer
	Pendle              *mocks.Pendle
	PendleAddress       common.Address
	PendleToken         *mocks.PendleToken
	PendleTokenAddress  common.Address
	Sense               *mocks.Sense
	SenseAddress        common.Address
	SenseToken          *mocks.SenseToken
	SenseTokenAddress   common.Address
	Element             *mocks.Element
	ElementAddress      common.Address
	ElementToken        *mocks.ElementToken
	ElementTokenAddress common.Address
	Notional            *mocks.Notional
	NotionalAddress     common.Address
	Lender              *lender.Lender
	LenderAddress       common.Address
}

func Deploy(e *Env) (*Dep, error) {
	// deploy the two mock tokens.
	ercAddress, _, ercContract, ercErr := mocks.DeployErc20(e.Owner.Opts, e.Blockchain)

	if ercErr != nil {
		return nil, ercErr
	}

	e.Blockchain.Commit()

	yAddress, _, yContract, yErr := mocks.DeployYield(e.Owner.Opts, e.Blockchain)

	if yErr != nil {
		return nil, yErr
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

	swivelAddress, _, swivelContract, swivelErr := mocks.DeploySwivel(e.Owner.Opts, e.Blockchain)

	if swivelErr != nil {
		return nil, swivelErr
	}

	e.Blockchain.Commit()

	mpAddress, _, mpContract, mpErr := mocks.DeployMarketPlace(e.Owner.Opts, e.Blockchain)

	if mpErr != nil {
		return nil, mpErr
	}

	e.Blockchain.Commit()

	apAddress, _, apContract, apErr := mocks.DeployAPWine(e.Owner.Opts, e.Blockchain)

	if apErr != nil {
		return nil, apErr
	}

	e.Blockchain.Commit()

	aptAddress, _, aptContract, aptErr := mocks.DeployAPWineToken(e.Owner.Opts, e.Blockchain)

	if aptErr != nil {
		return nil, aptErr
	}

	e.Blockchain.Commit()

	tAddress, _, tContract, tErr := mocks.DeployTempus(e.Owner.Opts, e.Blockchain)

	if tErr != nil {
		return nil, tErr
	}

	e.Blockchain.Commit()

	ttAddress, _, ttContract, ttErr := mocks.DeployTempusToken(e.Owner.Opts, e.Blockchain)

	if ttErr != nil {
		return nil, ttErr
	}

	e.Blockchain.Commit()

	pAddress, _, pContract, pErr := mocks.DeployPendle(e.Owner.Opts, e.Blockchain)

	if pErr != nil {
		return nil, pErr
	}

	e.Blockchain.Commit()

	ptAddress, _, ptContract, ptErr := mocks.DeployPendleToken(e.Owner.Opts, e.Blockchain)

	if ptErr != nil {
		return nil, ptErr
	}

	e.Blockchain.Commit()

	sAddress, _, sContract, sErr := mocks.DeploySense(e.Owner.Opts, e.Blockchain)

	if sErr != nil {
		return nil, sErr
	}

	e.Blockchain.Commit()

	stAddress, _, stContract, stErr := mocks.DeploySenseToken(e.Owner.Opts, e.Blockchain)

	if stErr != nil {
		return nil, stErr
	}

	e.Blockchain.Commit()

	eAddress, _, eContract, eErr := mocks.DeployElement(e.Owner.Opts, e.Blockchain)

	if eErr != nil {
		return nil, eErr
	}

	e.Blockchain.Commit()

	etAddress, _, etContract, etErr := mocks.DeployElementToken(e.Owner.Opts, e.Blockchain)

	if etErr != nil {
		return nil, etErr
	}

	e.Blockchain.Commit()

	nAddress, _, nContract, nErr := mocks.DeployNotional(e.Owner.Opts, e.Blockchain)

	if nErr != nil {
		return nil, nErr
	}

	e.Blockchain.Commit()

	lenderAddress, _, lenderContract, lenderErr := lender.DeployLender(e.Owner.Opts, e.Blockchain, swivelAddress, pAddress, tAddress)

	if lenderErr != nil {
		return nil, lenderErr
	}

	e.Blockchain.Commit()

	redeemerAddress, _, redeemerContract, redeemerErr := redeemer.DeployRedeemer(e.Owner.Opts, e.Blockchain, lenderAddress, swivelAddress, pAddress, tAddress, apAddress)

	if redeemerErr != nil {
		return nil, redeemerErr
	}

	e.Blockchain.Commit()

	return &Dep{
		Erc20Address:        ercAddress,
		Erc20:               ercContract,
		YieldAddress:        yAddress,
		Yield:               yContract,
		YieldToken:          ytContract,
		YieldTokenAddress:   ytAddress,
		ZcTokenAddress:      zcAddress,
		ZcToken:             zcContract,
		SwivelAddress:       swivelAddress,
		Swivel:              swivelContract,
		MarketplaceAddress:  mpAddress,
		Marketplace:         mpContract,
		APWine:              apContract,
		APWineAddress:       apAddress,
		APWineToken:         aptContract,
		APWineTokenAddress:  aptAddress,
		Tempus:              tContract,
		TempusAddress:       tAddress,
		TempusToken:         ttContract,
		TempusTokenAddress:  ttAddress,
		Pendle:              pContract,
		PendleAddress:       pAddress,
		PendleToken:         ptContract,
		PendleTokenAddress:  ptAddress,
		Sense:               sContract,
		SenseAddress:        sAddress,
		SenseToken:          stContract,
		SenseTokenAddress:   stAddress,
		Element:             eContract,
		ElementAddress:      eAddress,
		ElementToken:        etContract,
		ElementTokenAddress: etAddress,
		Notional:            nContract,
		NotionalAddress:     nAddress,
		Lender:              lenderContract,
		LenderAddress:       lenderAddress,
		Redeemer:            redeemerContract,
		RedeemerAddress:     redeemerAddress,
	}, nil
}
