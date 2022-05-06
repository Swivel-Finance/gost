package redeemertesting

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/redeemer"
)

type Dep struct {
	Erc20Address      common.Address
	Erc20             *mocks.Erc20
	YieldAddress      common.Address
	Yield             *mocks.Yield
	ZcTokenAddress    common.Address
	ZcToken           *mocks.ZcToken
	SwivelAddress     common.Address
	Swivel            *mocks.Swivel
	IlluminateAddress common.Address
	Illuminate        *mocks.Illuminate
	RedeemerAddress   common.Address
	Redeemer          *redeemer.Redeemer
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

	swivelAddress, _, swivelContract, swivelErr := mocks.DeploySwivel(e.Owner.Opts, e.Blockchain)

	if swivelErr != nil {
		return nil, swivelErr
	}

	e.Blockchain.Commit()

	mpAddress, _, mpContract, mpErr := mocks.DeployIlluminate(e.Owner.Opts, e.Blockchain)

	if mpErr != nil {
		return nil, mpErr
	}

	e.Blockchain.Commit()

	redeemerAddress, _, redeemerContract, redeemerErr := redeemer.DeployRedeemer(e.Owner.Opts, e.Blockchain, mpAddress)

	if redeemerErr != nil {
		return nil, redeemerErr
	}

	e.Blockchain.Commit()

	return &Dep{
		Erc20Address:      ercAddress,
		Erc20:             ercContract,
		YieldAddress:      ytAddress,
		Yield:             ytContract,
		ZcTokenAddress:    zcAddress,
		ZcToken:           zcContract,
		SwivelAddress:     swivelAddress,
		Swivel:            swivelContract,
		IlluminateAddress: mpAddress,
		Illuminate:        mpContract,
		RedeemerAddress:   redeemerAddress,
		Redeemer:          redeemerContract,
	}, nil
}
