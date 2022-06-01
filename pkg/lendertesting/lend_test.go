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
	MarketPlace  *mocks.MarketPlaceSession
	Yield        *mocks.YieldSession
	ZcToken      *mocks.ZcTokenSession
	Swivel       *mocks.SwivelSession
	ElementToken *mocks.ElementTokenSession
	Element      *mocks.ElementSession
	PendleToken  *mocks.PendleTokenSession
	Pendle       *mocks.PendleSession
	Tempus       *mocks.TempusSession
	Sense        *mocks.SenseSession
	SenseToken   *mocks.SenseTokenSession
	APWineToken  *mocks.APWineTokenSession
	APWine       *mocks.APWineSession
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

	s.SenseToken = &mocks.SenseTokenSession{
		Contract: s.Dep.SenseToken,
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

	s.APWine = &mocks.APWineSession{
		Contract: s.Dep.APWine,
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

	s.Lender.SetMarketPlaceAddress(s.Dep.MarketPlaceAddress)
	s.Env.Blockchain.Commit()

	ORDERS[0].Maker = s.Env.User1.Opts.From
	ORDERS[1].Maker = s.Env.User2.Opts.From
	ORDERS[0].Underlying = s.Dep.Erc20Address
	ORDERS[1].Underlying = s.Dep.Erc20Address
}

func (s *lendTestSuite) TestLendIlluminate() {
	assert := assert.New(s.T())

	maturity := big.NewInt(100000)
	amountLent := big.NewInt(5000)
	fee := new(big.Int).Div(amountLent, big.NewInt(FEENOMINATOR))
	lentAfterFees := new(big.Int).Sub(amountLent, fee)
	sellBasePreview := big.NewInt(4000)

	s.Erc20.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferReturns(true)
	s.Env.Blockchain.Commit()

	s.Yield.BaseReturns(s.Dep.Erc20Address)
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
	assert.Equal(lentAfterFees, yieldSellBasePreview)

	transferRes, err := s.Erc20.TransferCalled(s.Dep.YieldAddress)
	assert.Nil(err)
	assert.Equal(amountLent, transferRes)

	yieldSellBase, err := s.Yield.SellBaseCalled(s.Dep.LenderAddress)
	assert.Nil(err)
	assert.Equal(amountLent, yieldSellBase)

	transferredAmount, err := s.Erc20.TransferCalled(s.Dep.YieldAddress)
	assert.Nil(err)
	assert.Equal(amountLent, transferredAmount)
}

func (s *lendTestSuite) TestLendYield() {
	assert := assert.New(s.T())

	maturity := big.NewInt(100000)
	amountLent := big.NewInt(5000)
	fee := new(big.Int).Div(amountLent, big.NewInt(FEENOMINATOR))
	lentAfterFees := new(big.Int).Sub(amountLent, fee)
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
	s.MarketPlace.MarketsReturns(markets)
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
	assert.Equal(lentAfterFees, yieldSellBasePreview)

	transferRes, err := s.Erc20.TransferCalled(s.Dep.YieldAddress)
	assert.Nil(err)
	assert.Equal(amountLent, transferRes)

	yieldSellBase, err := s.Yield.SellBaseCalled(s.Dep.LenderAddress)
	assert.Nil(err)
	assert.Equal(amountLent, yieldSellBase)

	mint, err := s.ZcToken.MintCalled(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.Equal(amountLent, mint)

	transferredAmount, err := s.Erc20.TransferCalled(s.Dep.YieldAddress)
	assert.Nil(err)
	assert.Equal(amountLent, transferredAmount)
}

func (s *lendTestSuite) TestLendSwivel() {
	assert := assert.New(s.T())

	tx, err := s.Swivel.InitiateReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferReturns(true)
	s.Env.Blockchain.Commit()

	s.Erc20.ApproveReturns(true)
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
	s.MarketPlace.MarketsReturns(markets)
	s.Env.Blockchain.Commit()

	s.Yield.MaturityReturns(uint32(TEST_MATURITY.Uint64()))
	s.Env.Blockchain.Commit()

	s.Yield.SellBasePreviewReturns(TOTAL_AMOUNT)
	s.Env.Blockchain.Commit()

	s.Yield.SellBaseReturns(new(big.Int).Div(TOTAL_AMOUNT, big.NewInt(2)))
	s.Env.Blockchain.Commit()

	s.ZcToken.MintReturns(true)
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

	transferredAmount, err := s.Erc20.TransferCalled(s.Dep.YieldAddress)
	assert.Nil(err)
	assert.Equal(TOTAL_AMOUNT, transferredAmount)
}

func (s *lendTestSuite) TestLendElement() {
	assert := assert.New(s.T())

	s.MarketPlace.MarketsReturns([8]common.Address{
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
	limit := big.NewInt(100)
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

	tx, err = s.Lender.Lend3(3, s.Dep.Erc20Address, maturity, s.Dep.ElementAddress, elementPoolId, amount, limit, deadline)
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

	swap, err := s.Element.SwapCalled(s.Dep.LenderAddress)
	assert.Equal(deadline, swap.Deadline)
	assert.Equal(limit, swap.Limit)
	assert.Equal(s.Dep.LenderAddress, swap.Recipient)
	fee := new(big.Int).Div(amount, big.NewInt(FEENOMINATOR))
	amountAfterFees := new(big.Int).Sub(amount, fee)
	assert.Equal(amountAfterFees, swap.SwapAmount)
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
	s.MarketPlace.MarketsReturns([8]common.Address{
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
	fee := new(big.Int).Div(amount, big.NewInt(FEENOMINATOR))
	lent := new(big.Int).Sub(amount, fee)
	minimumBought := big.NewInt(50)
	deadline := big.NewInt(1000000000)

	tx, err := s.Lender.Lend0(4, s.Dep.Erc20Address, maturity, lent, minimumBought, deadline)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that mocks were called as expected
	swap, err := s.Pendle.SwapExactTokensForTokensCalled(s.Dep.LenderAddress)
	assert.Nil(err)
	assert.Equal(deadline, swap.Deadline)
	assert.Equal(lent, swap.Amount)
	assert.Equal(minimumBought, swap.MinimumBought)
}

func (s *lendTestSuite) TestLendTempus() {
	assert := assert.New(s.T())

	s.MarketPlace.MarketsReturns([8]common.Address{
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

	s.ZcToken.BalanceOfReturns(big.NewInt(100))
	s.Env.Blockchain.Commit()

	amount := big.NewInt(1032)
	fee := new(big.Int).Div(amount, big.NewInt(FEENOMINATOR))
	lent := new(big.Int).Sub(amount, fee)
	minimumReturn := big.NewInt(312)
	amm := common.HexToAddress("0x4321")
	pool := common.HexToAddress("0x1234")
	deadline := big.NewInt(9999)

	tx, err := s.Lender.Lend5(5, s.Dep.Erc20Address, maturity, amount, minimumReturn, amm, pool, deadline)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that mocks were called as expected
	deposit, err := s.Tempus.DepositAndFixCalled(lent)
	assert.Nil(err)
	assert.Equal(amm, deposit.Amm)
	assert.Equal(pool, deposit.Pool)
	assert.Equal(deadline, deposit.Deadline)
	assert.True(deposit.Bt)
	assert.Equal(minimumReturn, deposit.MinimumReturned)
}

func (s *lendTestSuite) TestLendSense() {
	assert := assert.New(s.T())
	s.MarketPlace.MarketsReturns([8]common.Address{
		s.Dep.ZcTokenAddress,
		s.Dep.SenseTokenAddress,
		s.Dep.SenseTokenAddress,
		s.Dep.SenseTokenAddress,
		s.Dep.SenseTokenAddress,
		s.Dep.SenseTokenAddress,
		s.Dep.SenseTokenAddress,
		s.Dep.SenseTokenAddress,
	})
	s.Env.Blockchain.Commit()

	s.SenseToken.UnderlyingReturns(s.Dep.Erc20Address)
	s.Env.Blockchain.Commit()

	s.SenseToken.UnderlyingReturns(s.Dep.Erc20Address)
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
	fee := new(big.Int).Div(amount, big.NewInt(FEENOMINATOR))
	lent := new(big.Int).Sub(amount, fee)
	minimumBought := big.NewInt(34)

	tx, err := s.Lender.Lend2(6, s.Dep.Erc20Address, maturity, s.Dep.SenseAddress, adapter, amount, minimumBought)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that mocks were called as expected
	swap, err := s.Sense.SwapUnderlyingForPTsCalled(adapter)
	assert.NoError(err)
	assert.Equal(maturity, swap.Maturity)
	assert.Equal(lent, swap.Amount)
	assert.Equal(minimumBought, swap.MinimumBought)
}

func (s *lendTestSuite) TestAPWineSense() {
	assert := assert.New(s.T())
	s.MarketPlace.MarketsReturns([8]common.Address{
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

	s.APWine.SwapExactAmountInReturns(big.NewInt(12345))
	s.Env.Blockchain.Commit()

	s.ZcToken.MintReturns(true)
	s.Env.Blockchain.Commit()

	maturity := big.NewInt(12094201240)
	amount := big.NewInt(100000)
	fee := new(big.Int).Div(amount, big.NewInt(FEENOMINATOR))
	lent := new(big.Int).Sub(amount, fee)
	minimumAmount := big.NewInt(34)
	id := big.NewInt(1000)

	tx, err := s.Lender.Lend(7, s.Dep.Erc20Address, maturity, amount, minimumAmount, s.Dep.APWineAddress, id)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that mocks were called as expected
	swap, err := s.APWine.SwapExactAmountInCalled(s.Dep.LenderAddress)
	assert.Equal(id, swap.Id)
	assert.Equal(big.NewInt(1), swap.TokenIn)
	assert.Equal(lent, swap.Amount)
	assert.True(swap.TokenOut.Cmp(big.NewInt(0)) == 0)
	assert.Equal(minimumAmount, swap.MinimumAmount)
}

func TestLendSuite(t *test.T) {
	suite.Run(t, &lendTestSuite{})
}
