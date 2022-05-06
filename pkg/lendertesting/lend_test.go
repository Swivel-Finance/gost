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
	Env          *Env
	Dep          *Dep
	Erc20        *mocks.Erc20Session
	Illuminate   *mocks.IlluminateSession
	Yield        *mocks.YieldSession
	ZcToken      *mocks.ZcTokenSession
	Swivel       *mocks.SwivelSession
	ElementToken *mocks.ElementTokenSession
	Element      *mocks.ElementSession
	Pendle       *mocks.PendleSession
	Sushi        *mocks.SushiSession
	Lender       *lender.LenderSession
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

	s.Illuminate = &mocks.IlluminateSession{
		Contract: s.Dep.Illuminate,
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

	s.Yield = &mocks.YieldSession{
		Contract: s.Dep.Yield,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.Swivel = &mocks.SwivelSession{
		Contract: s.Dep.Swivel,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.ElementToken = &mocks.ElementTokenSession{
		Contract: s.Dep.ElementToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.Element = &mocks.ElementSession{
		Contract: s.Dep.Element,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.Pendle = &mocks.PendleSession{
		Contract: s.Dep.Pendle,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.Sushi = &mocks.SushiSession{
		Contract: s.Dep.Sushi,
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

	ORDERS[0].Maker = s.Env.User1.Opts.From
	ORDERS[1].Maker = s.Env.User2.Opts.From
	ORDERS[0].Underlying = s.Dep.Erc20Address
	ORDERS[1].Underlying = s.Dep.Erc20Address
}

func (s *lendTestSuite) TestLendIlluminate() {
	assert := assert.New(s.T())

	maturity := big.NewInt(100000)
	amountLent := big.NewInt(5000)
	sellBasePreview := big.NewInt(4000)

	s.Erc20.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferReturns(true)
	s.Env.Blockchain.Commit()

	s.Yield.BaseReturns(s.Dep.Erc20Address)
	s.Env.Blockchain.Commit()

	s.Illuminate.MarketsReturns([8]common.Address{
		common.HexToAddress("0x0"),
		common.HexToAddress("0x1"),
		common.HexToAddress("0x2"),
		common.HexToAddress("0x3"),
		common.HexToAddress("0x4"),
		common.HexToAddress("0x5"),
		common.HexToAddress("0x6"),
		common.HexToAddress("0x7"),
	})

	s.Yield.MaturityReturns(uint32(maturity.Uint64()))
	s.Env.Blockchain.Commit()

	s.Yield.SellBasePreviewReturns(amountLent)
	s.Env.Blockchain.Commit()

	s.Yield.SellBaseReturns(sellBasePreview)
	s.Env.Blockchain.Commit()

	tx, err := s.Lender.Lend2(0, s.Dep.Erc20Address, maturity, s.Dep.YieldAddress, amountLent)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// verify that mocks were called as expected
	yieldMaturity, err := s.Yield.Maturity()
	assert.Nil(err)
	assert.Equal(uint32(maturity.Uint64()), yieldMaturity)

	transferFromRes, err := s.Erc20.TransferFromCalled(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.Equal(amountLent, transferFromRes.Amount)
	assert.Equal(s.Dep.LenderAddress, transferFromRes.To)

	yieldSellBasePreview, err := s.Yield.SellBasePreviewCalled()
	assert.Nil(err)
	assert.Equal(amountLent, yieldSellBasePreview)

	transferRes, err := s.Erc20.TransferCalled(s.Dep.YieldAddress)
	assert.Nil(err)
	assert.Equal(amountLent, transferRes)

	yieldSellBase, err := s.Yield.SellBaseCalled(s.Dep.LenderAddress)
	assert.Nil(err)
	assert.Equal(amountLent, yieldSellBase)
}

func (s *lendTestSuite) TestLendYield() {
	assert := assert.New(s.T())

	maturity := big.NewInt(100000)
	amountLent := big.NewInt(5000)
	sellBasePreview := big.NewInt(4000)

	s.Erc20.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferReturns(true)
	s.Env.Blockchain.Commit()

	s.Yield.BaseReturns(s.Dep.Erc20Address)
	s.Env.Blockchain.Commit()

	markets := [8]common.Address{
		s.Dep.ZcTokenAddress,
		common.HexToAddress("0x1"),
		common.HexToAddress("0x2"),
		common.HexToAddress("0x3"),
		common.HexToAddress("0x4"),
		common.HexToAddress("0x5"),
		common.HexToAddress("0x6"),
		common.HexToAddress("0x7"),
	}
	s.Illuminate.MarketsReturns(markets)
	s.Env.Blockchain.Commit()

	s.Yield.MaturityReturns(uint32(maturity.Uint64()))
	s.Env.Blockchain.Commit()

	s.Yield.SellBasePreviewReturns(amountLent)
	s.Env.Blockchain.Commit()

	s.Yield.SellBaseReturns(sellBasePreview)
	s.Env.Blockchain.Commit()

	s.ZcToken.MintReturns(true)
	s.Env.Blockchain.Commit()

	tx, err := s.Lender.Lend2(2, s.Dep.Erc20Address, maturity, s.Dep.YieldAddress, amountLent)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that mocks were called as expected
	yieldMaturity, err := s.Yield.Maturity()
	assert.Nil(err)
	assert.Equal(uint32(maturity.Uint64()), yieldMaturity)

	transferFromRes, err := s.Erc20.TransferFromCalled(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.Equal(amountLent, transferFromRes.Amount)
	assert.Equal(s.Dep.LenderAddress, transferFromRes.To)

	yieldSellBasePreview, err := s.Yield.SellBasePreviewCalled()
	assert.Nil(err)
	assert.Equal(amountLent, yieldSellBasePreview)

	transferRes, err := s.Erc20.TransferCalled(s.Dep.YieldAddress)
	assert.Nil(err)
	assert.Equal(amountLent, transferRes)

	yieldSellBase, err := s.Yield.SellBaseCalled(s.Dep.LenderAddress)
	assert.Nil(err)
	assert.Equal(amountLent, yieldSellBase)

	mint, err := s.ZcToken.MintCalled(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.Equal(amountLent, mint)
}

func (s *lendTestSuite) TestLendSwivel() {
	assert := assert.New(s.T())

	tx, err := s.Swivel.InitiateReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.Lender.Lend0(3, s.Dep.Erc20Address, TEST_MATURITY, s.Dep.YieldAddress, ORDERS, AMOUNTS, COMPONENTS)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that mocks were called as expected
	amountResult, err := s.Swivel.InitiateCalledAmount(ORDERS[0].Maker)
	assert.Nil(err)
	assert.Equal(AMOUNTS[0], amountResult)

	amountResult, err = s.Swivel.InitiateCalledAmount(ORDERS[1].Maker)
	assert.Nil(err)
	assert.Equal(AMOUNTS[1], amountResult)

	signatureResult, err := s.Swivel.InitiateCalledSignature(ORDERS[0].Maker)
	assert.Nil(err)
	assert.Equal(COMPONENTS[0].V, signatureResult)

	signatureResult, err = s.Swivel.InitiateCalledSignature(ORDERS[1].Maker)
	assert.Nil(err)
	assert.Equal(COMPONENTS[1].V, signatureResult)
}

func (s *lendTestSuite) TestLendElement() {
	assert := assert.New(s.T())

	s.Illuminate.MarketsReturns([8]common.Address{
		s.Dep.ElementTokenAddress,
		s.Dep.ElementTokenAddress,
		s.Dep.ElementTokenAddress,
		s.Dep.ElementTokenAddress,
		s.Dep.ElementTokenAddress,
		s.Dep.ElementTokenAddress,
		s.Dep.ElementTokenAddress,
		s.Dep.ElementTokenAddress,
	})
	s.Env.Blockchain.Commit()

	maturity := big.NewInt(100000)
	elementPoolId := [32]byte{1}
	amount := big.NewInt(10000)
	returnAmount := big.NewInt(100)
	deadline := big.NewInt(9999)

	tx, err := s.ElementToken.UnderlyingReturns(s.Dep.Erc20Address)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.ElementToken.UnlockTimestampReturns(maturity)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.Erc20.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	tx, err = s.Lender.Lend1(3, s.Dep.Erc20Address, maturity, s.Dep.ElementAddress, elementPoolId, amount, returnAmount, deadline)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that mocks were called as expected
	elementTokenMaturity, err := s.ElementToken.UnlockTimestamp()
	assert.Nil(err)
	assert.Equal(maturity, elementTokenMaturity)

	elementTokenUnderlying, err := s.ElementToken.Underlying()
	assert.Nil(err)
	assert.Equal(s.Dep.Erc20Address, elementTokenUnderlying)

	elementDeadline, err := s.Element.Deadline()
	assert.Nil(err)
	assert.Equal(deadline, elementDeadline)

	elementReturn, err := s.Element.Return()
	assert.Nil(err)
	assert.Equal(returnAmount, elementReturn)
}

func (s *lendTestSuite) TestLendPendle() {
	assert := assert.New(s.T())
	maturity := big.NewInt(100000)

	s.Sushi.SwapExactTokensForTokensReturns([]*big.Int{big.NewInt(1), big.NewInt(2)})
	s.Env.Blockchain.Commit()

	s.Pendle.MaturityReturns(maturity)
	s.Env.Blockchain.Commit()

	s.Pendle.UnderlyingReturns(s.Dep.Erc20Address)
	s.Env.Blockchain.Commit()

	s.ZcToken.MintReturns(true)
	s.Env.Blockchain.Commit()

	// TODO: This should be a helper that always returns a given address
	s.MarketPlace.MarketsReturns([8]common.Address{
		s.Dep.ZcTokenAddress,
		s.Dep.PendleAddress,
		s.Dep.PendleAddress,
		s.Dep.PendleAddress,
		s.Dep.PendleAddress,
		s.Dep.PendleAddress,
		s.Dep.PendleAddress,
		s.Dep.PendleAddress,
	})

	amount := big.NewInt(100)
	minimumBought := big.NewInt(50)
	deadline := big.NewInt(1000000000)

	tx, err := s.Lender.Lend(4, s.Dep.Erc20Address, maturity, amount, minimumBought, deadline)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that mocks were called as expected
	in, err := s.Sushi.InCalled()
	assert.NoError(err)
	assert.Equal(amount, in)

	out, err := s.Sushi.OutMinimumCalled()
	assert.NoError(err)
	assert.Equal(minimumBought, out)

	address, err := s.Sushi.PathCalled(big.NewInt(0))
	assert.NoError(err)
	assert.Equal(s.Dep.Erc20Address, address)

	address, err = s.Sushi.PathCalled(big.NewInt(1))
	assert.NoError(err)
	assert.Equal(s.Dep.PendleAddress, address)

	to, err := s.Sushi.ToCalled()
	assert.NoError(err)
	assert.Equal(s.Dep.LenderAddress, to)

	calledDeadline, err := s.Sushi.DeadlineCalled()
	assert.NoError(err)
	assert.Equal(deadline, calledDeadline)
}

func TestLendSuite(t *test.T) {
	suite.Run(t, &lendTestSuite{})
}
