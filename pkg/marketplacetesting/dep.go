package marketplacetesting

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/marketplace"
	"github.com/swivel-finance/gost/test/mocks"
)

type Dep struct {
	Erc20                *mocks.Erc20
	Erc20Address         common.Address
	CompoundToken        *mocks.CompoundToken
	CompoundTokenAddress common.Address
	Erc4626              *mocks.Erc4626
	Erc4626Address       common.Address
	YearnVault           *mocks.YearnVault
	YearnVaultAddress    common.Address
	AaveToken            *mocks.AaveToken
	AaveTokenAddress     common.Address
	AavePool             *mocks.AavePool
	AavePoolAddress      common.Address
	EulerToken           *mocks.EulerToken
	EulerTokenAddress    common.Address
	CreatorAddress       common.Address
	Creator              *mocks.Creator
	MarketPlaceAddress   common.Address
	MarketPlace          *marketplace.MarketPlace
	SwivelAddress        common.Address
	Swivel               *mocks.Swivel
	Maturity             *big.Int
	ZcToken              *mocks.ZcToken
	ZcTokenAddress       common.Address
	VaultTracker         *mocks.VaultTracker
	VaultTrackerAddress  common.Address
}

func Deploy(e *Env) (*Dep, error) {
	maturity := big.NewInt(MATURITY)

	ctAddress, _, ctContract, err := mocks.DeployCompoundToken(e.Owner.Opts, e.Blockchain)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	erc4626Address, _, erc4626Contract, err := mocks.DeployErc4626(e.Owner.Opts, e.Blockchain)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	erc20Address, _, erc20Contract, err := mocks.DeployErc20(e.Owner.Opts, e.Blockchain)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	yvAddress, _, yvContract, err := mocks.DeployYearnVault(e.Owner.Opts, e.Blockchain)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	atAddress, _, atContract, err := mocks.DeployAaveToken(e.Owner.Opts, e.Blockchain)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	apAddress, _, apContract, err := mocks.DeployAavePool(e.Owner.Opts, e.Blockchain)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	etAddress, _, etContract, err := mocks.DeployEulerToken(e.Owner.Opts, e.Blockchain)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	crAddress, _, crContract, err := mocks.DeployCreator(e.Owner.Opts, e.Blockchain)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	swivAddress, _, swivContract, err := mocks.DeploySwivel(e.Owner.Opts, e.Blockchain)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	marketAddress, _, marketContract, err := marketplace.DeployMarketPlace(e.Owner.Opts, e.Blockchain, crAddress)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	// NOTE the public properties on the mock can be changed per spec
	zcAddress, _, zcContract, err := mocks.DeployZcToken(e.Owner.Opts, e.Blockchain, uint8(1), erc20Address, big.NewInt(MATURITY), ctAddress, marketAddress, "", "", uint8(18))

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	// NOTE same as ZcToken WRT properties...
	vtAddress, _, vtContract, err := mocks.DeployVaultTracker(e.Owner.Opts, e.Blockchain, uint8(1), big.NewInt(MATURITY), ctAddress, swivAddress, marketAddress)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	return &Dep{
		CreatorAddress:       crAddress,
		Creator:              crContract,
		MarketPlaceAddress:   marketAddress,
		MarketPlace:          marketContract,
		Erc20Address:         erc20Address,
		Erc20:                erc20Contract,
		CompoundToken:        ctContract,
		CompoundTokenAddress: ctAddress,
		Erc4626:              erc4626Contract,
		Erc4626Address:       erc4626Address,
		YearnVault:           yvContract,
		YearnVaultAddress:    yvAddress,
		AaveToken:            atContract,
		AaveTokenAddress:     atAddress,
		AavePool:             apContract,
		AavePoolAddress:      apAddress,
		EulerToken:           etContract,
		EulerTokenAddress:    etAddress,
		Maturity:             maturity,
		ZcToken:              zcContract,
		ZcTokenAddress:       zcAddress,
		VaultTracker:         vtContract,
		VaultTrackerAddress:  vtAddress,
		SwivelAddress:        swivAddress,
		Swivel:               swivContract,
	}, nil
}
