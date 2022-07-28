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
	Maturity             *big.Int
	SwivelAddress        common.Address
	ZcToken              *mocks.ZcToken
	ZcTokenAddress       common.Address
	VaultTracker         *mocks.VaultTracker
	VaultTrackerAddress  common.Address
}

func Deploy(e *Env) (*Dep, error) {
	maturity := big.NewInt(MATURITY)

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

	erc20Address, _, erc20Contract, erc20Err := mocks.DeployErc20(e.Owner.Opts, e.Blockchain)

	if erc20Err != nil {
		return nil, erc20Err
	}

	e.Blockchain.Commit()

	yvAddress, _, yvContract, yvErr := mocks.DeployYearnVault(e.Owner.Opts, e.Blockchain)

	if yvErr != nil {
		return nil, yvErr
	}

	e.Blockchain.Commit()

	atAddress, _, atContract, atErr := mocks.DeployAaveToken(e.Owner.Opts, e.Blockchain)

	if atErr != nil {
		return nil, atErr
	}

	e.Blockchain.Commit()

	apAddress, _, apContract, apErr := mocks.DeployAavePool(e.Owner.Opts, e.Blockchain)

	if apErr != nil {
		return nil, apErr
	}

	e.Blockchain.Commit()

	etAddress, _, etContract, etErr := mocks.DeployEulerToken(e.Owner.Opts, e.Blockchain)

	if etErr != nil {
		return nil, etErr
	}

	e.Blockchain.Commit()

	crAddress, _, crContract, crErr := mocks.DeployCreator(e.Owner.Opts, e.Blockchain)

	if crErr != nil {
		return nil, crErr
	}

	e.Blockchain.Commit()

	marketAddress, _, marketContract, marketErr := marketplace.DeployMarketPlace(e.Owner.Opts, e.Blockchain, crAddress)

	if marketErr != nil {
		return nil, marketErr
	}

	e.Blockchain.Commit()

	swivAddress := common.HexToAddress("0x123aBc")

	// NOTE the public properties on the mock can be changed per spec
	zcAddress, _, zcContract, zcErr := mocks.DeployZcToken(e.Owner.Opts, e.Blockchain, uint8(1), erc20Address, big.NewInt(MATURITY), ctAddress, marketAddress, "", "", uint8(18))

	if zcErr != nil {
		return nil, zcErr
	}

	e.Blockchain.Commit()

	// NOTE same as ZcToken WRT properties...
	vtAddress, _, vtContract, vtErr := mocks.DeployVaultTracker(e.Owner.Opts, e.Blockchain, uint8(1), big.NewInt(MATURITY), ctAddress, swivAddress)

	if vtErr != nil {
		return nil, vtErr
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
		SwivelAddress:        swivAddress,
		Maturity:             maturity,
		ZcToken:              zcContract,
		ZcTokenAddress:       zcAddress,
		VaultTracker:         vtContract,
		VaultTrackerAddress:  vtAddress,
	}, nil
}
