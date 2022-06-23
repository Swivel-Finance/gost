package swiveltesting

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/fakes"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/swivel"
)

// TODO mock for marketplace...
type Dep struct {
	SigFakeAddress       common.Address
	SigFake              *fakes.SigFake // fake sig lib test contract
	HashFakeAddress      common.Address
	HashFake             *fakes.HashFake // fake hash lib test contract
	Erc20                *mocks.Erc20
	Erc20Address         common.Address
	CompoundToken        *mocks.CompoundToken
	CompoundTokenAddress common.Address
	Erc4626              *mocks.Erc4626
	Erc4626Address       common.Address
	MarketPlaceAddress   common.Address
	MarketPlace          *mocks.MarketPlace // mock marketplace
	Maturity             *big.Int
	SwivelAddress        common.Address
	Swivel               *swivel.Swivel
}

func Deploy(e *Env) (*Dep, error) {
	maturity := big.NewInt(MATURITY)
	// deploy the fakes so we can access the libs from tests
	sigAddress, _, sigContract, sigErr := fakes.DeploySigFake(e.Owner.Opts, e.Blockchain)

	if sigErr != nil {
		return nil, sigErr
	}

	e.Blockchain.Commit()

	hashAddress, _, hashContract, hashErr := fakes.DeployHashFake(e.Owner.Opts, e.Blockchain)

	if hashErr != nil {
		return nil, hashErr
	}

	e.Blockchain.Commit()

	// deploy the two mock tokens.
	ercAddress, _, ercContract, ercErr := mocks.DeployErc20(e.Owner.Opts, e.Blockchain)

	if ercErr != nil {
		return nil, ercErr
	}

	e.Blockchain.Commit()

	ctAddress, _, ctContract, ctErr := mocks.DeployCompoundToken(e.Owner.Opts, e.Blockchain)

	if ctErr != nil {
		return nil, ctErr
	}

	e.Blockchain.Commit()

	tvAddress, _, tvContract, tvErr := mocks.DeployErc4626(e.Owner.Opts, e.Blockchain)

	if tvErr != nil {
		return nil, tvErr
	}

	e.Blockchain.Commit()

	marketAddress, _, marketContract, marketErr := mocks.DeployMarketPlace(e.Owner.Opts, e.Blockchain)

	if marketErr != nil {
		return nil, marketErr
	}

	e.Blockchain.Commit()

	// deploy swivel contract...
	swivelAddress, _, swivelContract, swivelErr := swivel.DeploySwivel(e.Owner.Opts, e.Blockchain, marketAddress)

	// TODO call marketPlace swivel contract address setter when implemented...

	if swivelErr != nil {
		return nil, swivelErr
	}

	e.Blockchain.Commit()

	return &Dep{
		SigFakeAddress:       sigAddress,
		SigFake:              sigContract,
		HashFakeAddress:      hashAddress,
		HashFake:             hashContract,
		Erc20:                ercContract,
		Erc20Address:         ercAddress,
		CompoundToken:        ctContract,
		CompoundTokenAddress: ctAddress,
		Erc4626:              tvContract,
		Erc4626Address:       tvAddress,
		MarketPlaceAddress:   marketAddress,
		MarketPlace:          marketContract,
		Maturity:             maturity,
		SwivelAddress:        swivelAddress,
		Swivel:               swivelContract,
	}, nil
}
