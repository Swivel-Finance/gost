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
	PendleToken  *mocks.PendleTokenSession
	Pendle       *mocks.PendleSession
	Tempus       *mocks.TempusSession
	Sense        *mocks.SenseSession
	SenseAdapter *mocks.SenseAdapterSession
	APWineToken  *mocks.APWineTokenSession
	APWineRouter *mocks.APWineRouterSession
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

	s.PendleToken = &mocks.PendleTokenSession{
		Contract: s.Dep.PendleToken,
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

	s.Tempus = &mocks.TempusSession{
		Contract: s.Dep.Tempus,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.Sense = &mocks.SenseSession{
		Contract: s.Dep.Sense,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.SenseAdapter = &mocks.SenseAdapterSession{
		Contract: s.Dep.SenseAdapter,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.APWineToken = &mocks.APWineTokenSession{
		Contract: s.Dep.APWineToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.APWineRouter = &mocks.APWineRouterSession{
		Contract: s.Dep.APWineRouter,
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

	tx, err := s.Lender.Lend4(0, s.Dep.Erc20Address, maturity, s.Dep.YieldAddress, amountLent)
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

	tx, err := s.Lender.Lend4(2, s.Dep.Erc20Address, maturity, s.Dep.YieldAddress, amountLent)
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

	tx, err = s.Lender.Lend1(3, s.Dep.Erc20Address, TEST_MATURITY, s.Dep.YieldAddress, ORDERS, AMOUNTS, COMPONENTS)
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

	tx, err = s.Lender.Lend3(3, s.Dep.Erc20Address, maturity, s.Dep.ElementAddress, elementPoolId, amount, returnAmount, deadline)
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

	s.Pendle.SwapExactTokensForTokensReturns([]*big.Int{big.NewInt(1), big.NewInt(2)})
	s.Env.Blockchain.Commit()

	s.PendleToken.ExpiryReturns(maturity)
	s.Env.Blockchain.Commit()

	s.PendleToken.YieldTokenReturns(s.Dep.Erc20Address)
	s.Env.Blockchain.Commit()

	s.ZcToken.MintReturns(true)
	s.Env.Blockchain.Commit()

	// TODO: This should be a helper that always returns a given address
	s.Illuminate.MarketsReturns([8]common.Address{
		s.Dep.ZcTokenAddress,
		s.Dep.PendleTokenAddress,
		s.Dep.PendleTokenAddress,
		s.Dep.PendleTokenAddress,
		s.Dep.PendleTokenAddress,
		s.Dep.PendleTokenAddress,
		s.Dep.PendleTokenAddress,
		s.Dep.PendleTokenAddress,
	})

	amount := big.NewInt(100)
	minimumBought := big.NewInt(50)
	deadline := big.NewInt(1000000000)

	tx, err := s.Lender.Lend0(4, s.Dep.Erc20Address, maturity, amount, minimumBought, deadline)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that mocks were called as expected
	in, err := s.Pendle.InCalled()
	assert.NoError(err)
	assert.Equal(amount, in)

	out, err := s.Pendle.OutMinimumCalled()
	assert.NoError(err)
	assert.Equal(minimumBought, out)

	address, err := s.Pendle.PathCalled(big.NewInt(0))
	assert.NoError(err)
	assert.Equal(s.Dep.Erc20Address, address)

	address, err = s.Pendle.PathCalled(big.NewInt(1))
	assert.NoError(err)
	assert.Equal(s.Dep.PendleTokenAddress, address)

	to, err := s.Pendle.ToCalled()
	assert.NoError(err)
	assert.Equal(s.Dep.LenderAddress, to)

	calledDeadline, err := s.Pendle.DeadlineCalled()
	assert.NoError(err)
	assert.Equal(deadline, calledDeadline)
}

func (s *lendTestSuite) TestLendTempus() {
	assert := assert.New(s.T())

	s.Illuminate.MarketsReturns([8]common.Address{
		s.Dep.ZcTokenAddress,
		s.Dep.TempusAddress,
		s.Dep.TempusAddress,
		s.Dep.TempusAddress,
		s.Dep.TempusAddress,
		s.Dep.TempusAddress,
		s.Dep.TempusAddress,
		s.Dep.TempusAddress,
	})
	s.Env.Blockchain.Commit()

	maturity := big.NewInt(12094201240)
	s.Tempus.MaturityTimeReturns(maturity)
	s.Env.Blockchain.Commit()

	s.Tempus.YieldBearingTokenReturns(s.Dep.Erc20Address)
	s.Env.Blockchain.Commit()

	s.Tempus.DepositAndFixReturns(big.NewInt(102))
	s.Env.Blockchain.Commit()

	s.Erc20.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	s.ZcToken.MintReturns(true)
	s.Env.Blockchain.Commit()

	s.ZcToken.BalanceOfReturns(s.Dep.LenderAddress, big.NewInt(100))
	s.Env.Blockchain.Commit()

	amount := big.NewInt(1032)
	minimumReturn := big.NewInt(312)
	amm := common.HexToAddress("0x4321")
	pool := common.HexToAddress("0x1234")
	deadline := big.NewInt(9999)

	tx, err := s.Lender.Lend5(5, s.Dep.Erc20Address, maturity, amount, minimumReturn, amm, pool, deadline)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that mocks were called as expected
	ammCalled, err := s.Tempus.TempusAMMCalled()
	assert.NoError(err)
	assert.Equal(amm, ammCalled)

	poolCalled, err := s.Tempus.TempusPoolCalled()
	assert.NoError(err)
	assert.Equal(pool, poolCalled)

	amountCalled, err := s.Tempus.AmountCalled()
	assert.NoError(err)
	assert.Equal(amount, amountCalled)

	isTokenBacking, err := s.Tempus.IsBackingTokenCalled()
	assert.NoError(err)
	assert.True(isTokenBacking)

	minimumReturnCalled, err := s.Tempus.MinimumReturnCalled()
	assert.NoError(err)
	assert.Equal(minimumReturn, minimumReturnCalled)

	deadlineCalled, err := s.Tempus.DeadlineCalled()
	assert.NoError(err)
	assert.Equal(deadline, deadlineCalled)
}

func (s *lendTestSuite) TestLendSense() {
	assert := assert.New(s.T())
	s.Illuminate.MarketsReturns([8]common.Address{
		s.Dep.ZcTokenAddress,
		s.Dep.SenseAddress,
		s.Dep.SenseAddress,
		s.Dep.SenseAddress,
		s.Dep.SenseAddress,
		s.Dep.SenseAddress,
		s.Dep.SenseAddress,
		s.Dep.SenseAddress,
	})
	s.Env.Blockchain.Commit()

	s.SenseAdapter.UnderlyingReturns(s.Dep.Erc20Address)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	s.Sense.SwapUnderlyingForPTsReturns(big.NewInt(12345))
	s.Env.Blockchain.Commit()

	s.ZcToken.MintReturns(true)
	s.Env.Blockchain.Commit()

	maturity := big.NewInt(12094201240)
	adapter := common.HexToAddress("0x1234")
	amount := big.NewInt(1032)
	minimumBought := big.NewInt(34)

	tx, err := s.Lender.Lend2(6, s.Dep.Erc20Address, maturity, s.Dep.SenseAddress, adapter, amount, minimumBought)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that mocks were called as expected
	adapterCalled, err := s.Sense.SenseAdapterCalled()
	assert.NoError(err)
	assert.Equal(adapter, adapterCalled)

	maturityCalled, err := s.Sense.MaturityCalled()
	assert.NoError(err)
	assert.Equal(maturity, maturityCalled)

	amountCalled, err := s.Sense.AmountCalled()
	assert.NoError(err)
	assert.Equal(amount, amountCalled)

	minimumBoughtCalled, err := s.Sense.MinimumBoughtCalled()
	assert.NoError(err)
	assert.Equal(minimumBought, minimumBoughtCalled)
}

func (s *lendTestSuite) TestAPWineSense() {
	assert := assert.New(s.T())
	s.Illuminate.MarketsReturns([8]common.Address{
		s.Dep.ZcTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
	})
	s.Env.Blockchain.Commit()

	s.APWineToken.GetPTAddressReturns(s.Dep.Erc20Address)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	s.APWineRouter.SwapExactAmountInReturns(big.NewInt(12345))
	s.Env.Blockchain.Commit()

	s.ZcToken.MintReturns(true)
	s.Env.Blockchain.Commit()

	maturity := big.NewInt(12094201240)
	amount := big.NewInt(1032)
	minimumAmount := big.NewInt(34)
	id := big.NewInt(1000)

	tx, err := s.Lender.Lend(7, s.Dep.Erc20Address, maturity, amount, minimumAmount, s.Dep.APWineRouterAddress, id)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that mocks were called as expected
	idCalled, err := s.APWineRouter.IdCalled()
	assert.NoError(err)
	assert.Equal(id, idCalled)

	tokenInCalled, err := s.APWineRouter.TokenInCalled()
	assert.NoError(err)
	assert.Equal(big.NewInt(1), tokenInCalled)

	amountCalled, err := s.APWineRouter.AmountCalled()
	assert.NoError(err)
	assert.Equal(amount, amountCalled)

	tokenOutCalled, err := s.APWineRouter.TokenOutCalled()
	assert.NoError(err)
	assert.True(big.NewInt(0).Cmp(tokenOutCalled) == 0)

	minimumAmountCalled, err := s.APWineRouter.MinimumAmountCalled()
	assert.NoError(err)
	assert.Equal(minimumAmount, minimumAmountCalled)

	toCalled, err := s.APWineRouter.ToCalled()
	assert.NoError(err)
	assert.Equal(s.Dep.LenderAddress, toCalled)
}

func TestLendSuite(t *test.T) {
	suite.Run(t, &lendTestSuite{})
}
