package testing

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/mocks"
)

type Dep struct {
<<<<<<< HEAD
	Erc20Address       common.Address
	Erc20              *mocks.Erc20
	YieldTokenAddress  common.Address
	YieldToken         *mocks.YieldToken
	ZcTokenAddress     common.Address
	ZcToken            *mocks.ZcToken
	SwivelTokenAddress common.Address
	SwivelToken        *mocks.SwivelToken
=======
	Erc20Address      common.Address
	Erc20             *mocks.Erc20
	YieldTokenAddress common.Address
	YieldToken        *mocks.YieldToken
	ZcTokenAddress    common.Address
	ZcToken           *mocks.ZcToken
>>>>>>> 85f4eb3d1741965abee2064ff7b959bc24f0cf17
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

<<<<<<< HEAD
	swivelAddress, _, swivelContract, swivelErr := mocks.DeploySwivelToken(e.Owner.Opts, e.Blockchain)

	if swivelErr != nil {
		return nil, swivelErr
	}

	e.Blockchain.Commit()

	return &Dep{
		Erc20Address:       ercAddress,
		Erc20:              ercContract,
		YieldTokenAddress:  ytAddress,
		YieldToken:         ytContract,
		ZcTokenAddress:     zcAddress,
		ZcToken:            zcContract,
		SwivelTokenAddress: swivelAddress,
		SwivelToken:        swivelContract,
=======
	return &Dep{
		Erc20Address:      ercAddress,
		Erc20:             ercContract,
		YieldTokenAddress: ytAddress,
		YieldToken:        ytContract,
		ZcTokenAddress:    zcAddress,
		ZcToken:           zcContract,
>>>>>>> 85f4eb3d1741965abee2064ff7b959bc24f0cf17
	}, nil
}
