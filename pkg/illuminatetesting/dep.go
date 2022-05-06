package illuminatetesting

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/illuminate"
	"github.com/swivel-finance/gost/test/mocks"
)

type Dep struct {
	Erc20Address      common.Address
	Erc20             *mocks.Erc20
	IlluminateAddress common.Address
	Illuminate        *illuminate.Illuminate
}

func Deploy(e *Env) (*Dep, error) {
	// deploy the two mock tokens.
	ercAddress, _, ercContract, ercErr := mocks.DeployErc20(e.Owner.Opts, e.Blockchain)

	if ercErr != nil {
		return nil, ercErr
	}

	e.Blockchain.Commit()

	illAddress, _, illContract, illErr := illuminate.DeployIlluminate(e.Owner.Opts, e.Blockchain)
	if illErr != nil {
		return nil, illErr
	}

	e.Blockchain.Commit()

	return &Dep{
		Erc20Address:      ercAddress,
		Erc20:             ercContract,
		IlluminateAddress: illAddress,
		Illuminate:        illContract,
	}, nil
}
