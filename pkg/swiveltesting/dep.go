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
	YearnVault           *mocks.YearnVault
	YearnVaultAddress    common.Address
	AavePool             *mocks.AavePool
	AavePoolAddress      common.Address
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
	erc20Address, _, erc20Contract, erc20Err := mocks.DeployErc20(e.Owner.Opts, e.Blockchain)

	if erc20Err != nil {
		return nil, erc20Err
	}

	e.Blockchain.Commit()

	ctAddress, _, ctContract, ctErr := mocks.DeployCompoundToken(e.Owner.Opts, e.Blockchain)

	if ctErr != nil {
		return nil, ctErr
	}

	e.Blockchain.Commit()

	erc4626Address, _, erc4626Contract, erc4626Err := mocks.DeployErc4626(e.Owner.Opts, e.Blockchain)

	if erc4626Err != nil {
		return nil, erc4626Err
	}

	e.Blockchain.Commit()

	yvAddress, _, yvContract, yvErr := mocks.DeployYearnVault(e.Owner.Opts, e.Blockchain)

	if yvErr != nil {
		return nil, yvErr
	}

	e.Blockchain.Commit()

	apAddress, _, apContract, apErr := mocks.DeployAavePool(e.Owner.Opts, e.Blockchain)

	if apErr != nil {
		return nil, apErr
	}

	e.Blockchain.Commit()

	marketAddress, _, marketContract, marketErr := mocks.DeployMarketPlace(e.Owner.Opts, e.Blockchain)

	if marketErr != nil {
		return nil, marketErr
	}

	e.Blockchain.Commit()

	// deploy swivel contract... TODO change the aaveAddr to the actual mock deployed pool when implemented
	swivelAddress, _, swivelContract, swivelErr := swivel.DeploySwivel(e.Owner.Opts, e.Blockchain, marketAddress, apAddress)

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
		Erc20:                erc20Contract,
		Erc20Address:         erc20Address,
		CompoundToken:        ctContract,
		CompoundTokenAddress: ctAddress,
		Erc4626:              erc4626Contract,
		Erc4626Address:       erc4626Address,
		YearnVault:           yvContract,
		YearnVaultAddress:    yvAddress,
		AavePool:             apContract,
		AavePoolAddress:      apAddress,
		MarketPlaceAddress:   marketAddress,
		MarketPlace:          marketContract,
		Maturity:             maturity,
		SwivelAddress:        swivelAddress,
		Swivel:               swivelContract,
	}, nil
}
