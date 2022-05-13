package illuminatetesting

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/illuminate"
	"github.com/swivel-finance/gost/test/mocks"
)

type Dep struct {
	Erc20Address      common.Address
	Erc20             *mocks.Erc20
	RedeemerAddress   common.Address
	IlluminateAddress common.Address
	Illuminate        *illuminate.Illuminate
}

func Deploy(e *Env) (*Dep, error) {
	// deploy mock tokens.
	ercAddress, _, ercContract, ercErr := mocks.DeployErc20(e.Owner.Opts, e.Blockchain)

	if ercErr != nil {
		return nil, ercErr
	}

	e.Blockchain.Commit()

	rAddr := common.HexToAddress("0x123")

	illAddress, _, illContract, illErr := illuminate.DeployIlluminate(e.Owner.Opts, e.Blockchain, rAddr)
	if illErr != nil {
		return nil, illErr
	}

	e.Blockchain.Commit()

	return &Dep{
		Erc20Address:      ercAddress,
		Erc20:             ercContract,
		RedeemerAddress:   rAddr,
		IlluminateAddress: illAddress,
		Illuminate:        illContract,
	}, nil
}
