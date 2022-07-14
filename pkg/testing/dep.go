package testing

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/swivel-finance/gost/test/fakes"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/tokens"
)

type Dep struct {
	SigFakeAddress  common.Address
	SigFake         *fakes.SigFake // fake sig lib test contract
	HashFakeAddress common.Address
	HashFake        *fakes.HashFake // fake hash lib test contract
	Erc20Address    common.Address
	Erc20           *mocks.Erc20
	CErc20Address   common.Address
	CErc20          *mocks.CErc20
	FErc20Address   common.Address
	FErc20          *mocks.FErc20
	ZcTokenAddress  common.Address
	ZcToken         *tokens.ZcToken
}

func Deploy(e *Env) (*Dep, error) {
	// deploying the lib testing contract "fakes"
	// NOTE these _could_ be moved into their own package as they are not needed
	// for swivel to operate. TODO
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

	// deploy a real ZcToken in order to test the maturity requirement
	zctAddress, _, zctContract, zctErr := tokens.DeployZcToken(e.Owner.Opts, e.Blockchain, ercAddress, big.NewInt(MATURITY), "YoloToken", "YT", uint8(18))

	if zctErr != nil {
		return nil, zctErr
	}

	e.Blockchain.Commit()

	return &Dep{
		SigFakeAddress:  sigAddress,
		SigFake:         sigContract,
		HashFakeAddress: hashAddress,
		HashFake:        hashContract,
		Erc20Address:    ercAddress,
		Erc20:           ercContract,
		CErc20Address:   cercAddress,
		CErc20:          cercContract,
		FErc20Address:   fercAddress,
		FErc20:          fercContract,
		ZcTokenAddress:  zctAddress,
		ZcToken:         zctContract,
	}, nil
}
