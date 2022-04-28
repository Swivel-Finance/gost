package lendertesting

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/mocks"
)

type Dep struct {
	Erc20Address      common.Address
	Erc20             *mocks.Erc20
	YieldTokenAddress common.Address
	YieldToken        *mocks.YieldToken
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

	return &Dep{
		Erc20Address:      ercAddress,
		Erc20:             ercContract,
		YieldTokenAddress: ytAddress,
		YieldToken:        ytContract,
	}, nil
}