package lendertesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/lender"
	"github.com/swivel-finance/gost/test/mocks"
)

type lendTestSuite struct {
	suite.Suite
	Env         *Env
	Dep         *Dep
	Erc20       *mocks.Erc20Session
	MarketPlace *mocks.MarketPlaceSession
	YieldToken  *mocks.YieldTokenSession
	ZcToken     *mocks.ZcTokenSession
	Lender      *lender.LenderSession
}

func (s *lendTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.Erc20 = &mocks.Erc20Session{
		Contract: s.Dep.Erc20,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.MarketPlace = &mocks.MarketPlaceSession{
		Contract: s.Dep.MarketPlace,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.ZcToken = &mocks.ZcTokenSession{
		Contract: s.Dep.ZcToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.YieldToken = &mocks.YieldTokenSession{
		Contract: s.Dep.YieldToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.Lender = &lender.LenderSession{
		Contract: s.Dep.Lender,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *lendTestSuite) TestLendIlluminate() {
	assert := assert.New(s.T())

	maturity := big.NewInt(100000)
	amountLent := big.NewInt(5000)
	returnValue := big.NewInt(4000)

	tx, err := s.Erc20.TransferFromReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	tx, err = s.Erc20.TransferReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	s.MarketPlace.MarketsReturns([8]common.Address{
		common.HexToAddress("0x0"),
		common.HexToAddress("0x1"),
		common.HexToAddress("0x2"),
		common.HexToAddress("0x3"),
		common.HexToAddress("0x4"),
		common.HexToAddress("0x5"),
		common.HexToAddress("0x6"),
		common.HexToAddress("0x7"),
	})

	s.YieldToken.MaturityReturns(uint32(maturity.Uint64()))
	s.Env.Blockchain.Commit()

	s.YieldToken.SellBasePreviewReturns(returnValue)
	s.Env.Blockchain.Commit()

	tx, err = s.Lender.Lend(0, s.Dep.Erc20Address, maturity, s.Dep.YieldTokenAddress, amountLent)
	assert.Nil(err)
	assert.NotNil(tx)
}

func (s *lendTestSuite) TestLendYield() {
}

func TestLendSuite(t *test.T) {
	suite.Run(t, &lendTestSuite{})
}
