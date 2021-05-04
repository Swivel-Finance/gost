package swiveltesting

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/swivel"
)

// TODO mock for marketplace...
type Dep struct {
	Erc20Address  common.Address
	Erc20         *mocks.Erc20 // mock erc20
	CErc20Address common.Address
	CErc20        *mocks.CErc20 // mock erc20
	SwivelAddress common.Address
	Swivel        *swivel.Swivel
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

	// deploy swivel contract...
	swivelAddress, _, swivelContract, swivelErr := swivel.DeploySwivel(e.Owner.Opts, e.Blockchain)

	if swivelErr != nil {
		return nil, swivelErr
	}

	e.Blockchain.Commit()

	return &Dep{
		Erc20Address:  ercAddress,
		Erc20:         ercContract,
		CErc20Address: cercAddress,
		CErc20:        cercContract,
		SwivelAddress: swivelAddress,
		Swivel:        swivelContract,
	}, nil
}
