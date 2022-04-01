package testing

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/mocks"
)

type Dep struct {
	Erc20Address  common.Address
	Erc20         *mocks.Erc20
	CErc20Address common.Address
	CErc20        *mocks.CErc20
	FErc20Address common.Address
	FErc20        *mocks.FErc20
}

func Deploy(e *Env) (*Dep, error) {
	// deploy the two mock tokens.
	ercAddress, _, ercContract, ercErr := mocks.DeployErc20(e.Owner.Opts, e.Blockchain)

	if ercErr != nil {
		return nil, ercErr
	}

	e.Blockchain.Commit()

	cercAddress, _, cercContract, cercErr := mocks.DeployCErc20(e.Owner.Opts, e.Blockchain)

	if cercErr != nil {
		return nil, cercErr
	}

	e.Blockchain.Commit()

	fercAddress, _, fercContract, fercErr := mocks.DeployFErc20(e.Owner.Opts, e.Blockchain)

	if fercErr != nil {
		return nil, fercErr
	}

	e.Blockchain.Commit()

	return &Dep{
		Erc20Address:  ercAddress,
		Erc20:         ercContract,
		CErc20Address: cercAddress,
		CErc20:        cercContract,
		FErc20Address: fercAddress,
		FErc20:        fercContract,
	}, nil
}
