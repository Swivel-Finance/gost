package marketplacetesting

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/marketplace"
	"github.com/swivel-finance/gost/test/mocks"
)

type Dep struct {
	Erc20        *mocks.Erc20
	Erc20Address common.Address

	CompoundToken        *mocks.CompoundToken
	CompoundTokenAddress common.Address

	Erc4626        *mocks.Erc4626
	Erc4626Address common.Address

	MarketPlaceAddress common.Address
	MarketPlace        *marketplace.MarketPlace

	Maturity *big.Int

	SwivelAddress common.Address
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

	// deploy contract...
	marketAddress, _, marketContract, marketErr := marketplace.DeployMarketPlace(e.Owner.Opts, e.Blockchain)

	if marketErr != nil {
		return nil, marketErr
	}

	e.Blockchain.Commit()

	return &Dep{
		MarketPlaceAddress:   marketAddress,
		MarketPlace:          marketContract,
		Erc20Address:         erc20Address,
		Erc20:                erc20Contract,
		CompoundToken:        ctContract,
		CompoundTokenAddress: ctAddress,
		Erc4626:              erc4626Contract,
		Erc4626Address:       erc4626Address,
		SwivelAddress:        common.HexToAddress("0x123aBc"),
		Maturity:             maturity,
	}, nil
}
