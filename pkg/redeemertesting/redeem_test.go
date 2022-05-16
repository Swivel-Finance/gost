package redeemertesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/redeemer"
)

type redeemTestSuite struct {
	suite.Suite
	Env          *Env
	Dep          *Dep
	Erc20        *mocks.Erc20Session
	Illuminate   *mocks.IlluminateSession
	Yield        *mocks.YieldSession
	ZcToken      *mocks.ZcTokenSession
	Swivel       *mocks.SwivelSession
	APWine       *mocks.APWineSession
	APWineToken  *mocks.APWineTokenSession
	Tempus       *mocks.TempusSession
	TempusToken  *mocks.TempusTokenSession
	Pendle       *mocks.PendleSession
	PendleToken  *mocks.PendleTokenSession
	Sense        *mocks.SenseSession
	SenseToken   *mocks.SenseTokenSession
	Element      *mocks.ElementSession
	ElementToken *mocks.ElementTokenSession
	Redeemer     *redeemer.RedeemerSession
}

func (s *redeemTestSuite) SetupSuite() {
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

	s.APWine = &mocks.APWineSession{
		Contract: s.Dep.APWine,
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

	s.Tempus = &mocks.TempusSession{
		Contract: s.Dep.Tempus,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.TempusToken = &mocks.TempusTokenSession{
		Contract: s.Dep.TempusToken,
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

	s.PendleToken = &mocks.PendleTokenSession{
		Contract: s.Dep.PendleToken,
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

	s.Element = &mocks.ElementSession{
		Contract: s.Dep.Element,
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

	s.Redeemer = &redeemer.RedeemerSession{
		Contract: s.Dep.Redeemer,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *redeemTestSuite) TestAPWineRedeem() {
	assert := assert.New(s.T())

	amount := big.NewInt(1000)
	maturity := big.NewInt(9999999)
	owner := s.Env.User1.Opts.From
	principal := uint8(7)

	s.Illuminate.MarketsReturns([8]common.Address{
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
	})
	s.Env.Blockchain.Commit()

	s.APWineToken.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	s.APWineToken.BalanceOfReturns(amount)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferReturns(true)
	s.Env.Blockchain.Commit()

	tx, err := s.Redeemer.Redeem0(principal, s.Dep.Erc20Address, maturity, owner)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that the mocked functions were called as expected
	amountCalled, err := s.APWine.WithdrawCalled(owner)
	assert.NoError(err)
	assert.Equal(amount, amountCalled)

	vaultCalled, err := s.APWineToken.BalanceOfCalled()
	assert.NoError(err)
	assert.Equal(owner, vaultCalled)

	underlyingTransfer, err := s.Erc20.TransferCalled(owner)
	assert.NoError(err)
	assert.Equal(amount, underlyingTransfer)
}

func (s *redeemTestSuite) TestTempusRedeem() {
	assert := assert.New(s.T())

	amount := big.NewInt(1000)
	maturity := big.NewInt(9999999)
	owner := s.Env.User1.Opts.From
	prinicipal := uint8(5)

	s.Illuminate.MarketsReturns([8]common.Address{
		s.Dep.TempusTokenAddress,
		s.Dep.TempusTokenAddress,
		s.Dep.TempusTokenAddress,
		s.Dep.TempusTokenAddress,
		s.Dep.TempusTokenAddress,
		s.Dep.TempusTokenAddress,
		s.Dep.TempusTokenAddress,
		s.Dep.TempusTokenAddress,
	})
	s.Env.Blockchain.Commit()

	s.TempusToken.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	s.TempusToken.BalanceOfReturns(amount)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferReturns(true)
	s.Env.Blockchain.Commit()

	tx, err := s.Redeemer.Redeem0(prinicipal, s.Dep.Erc20Address, maturity, owner)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that the mocked functions were called as expected
	redeemCall, err := s.Tempus.RedeemToBackingCalled(owner)
	assert.NoError(err)
	assert.Equal(amount, redeemCall.Amount)
	assert.Equal(maturity, redeemCall.Maturity)
	assert.Equal(s.Dep.Erc20Address, redeemCall.Underlying)

	underlyingTransfer, err := s.Erc20.TransferCalled(owner)
	assert.NoError(err)
	assert.Equal(amount, underlyingTransfer)
}

func (s *redeemTestSuite) TestIlluminateRedeem() {
	assert := assert.New(s.T())

	amount := big.NewInt(1000)
	maturity := big.NewInt(9999999)
	owner := s.Env.User1.Opts.From
	principal := uint8(0)

	s.Illuminate.MarketsReturns([8]common.Address{
		s.Dep.ZcTokenAddress,
		s.Dep.ZcTokenAddress,
		s.Dep.ZcTokenAddress,
		s.Dep.ZcTokenAddress,
		s.Dep.ZcTokenAddress,
		s.Dep.ZcTokenAddress,
		s.Dep.ZcTokenAddress,
		s.Dep.ZcTokenAddress,
	})
	s.Env.Blockchain.Commit()

	s.ZcToken.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	s.ZcToken.BalanceOfReturns(amount)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferReturns(true)
	s.Env.Blockchain.Commit()

	tx, err := s.Redeemer.Redeem0(principal, s.Dep.Erc20Address, maturity, owner)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that the mocked functions were called as expected
	burnCall, err := s.ZcToken.BurnCalled(owner)
	assert.NoError(err)
	assert.Equal(amount, burnCall)

	underlyingTransfer, err := s.Erc20.TransferCalled(owner)
	assert.NoError(err)
	assert.Equal(amount, underlyingTransfer)
}

func (s *redeemTestSuite) TestPendleRedeem() {
	assert := assert.New(s.T())

	amount := big.NewInt(1000)
	forgeId := [32]byte{3, 3, 4, 2}
	maturity := big.NewInt(9999999)
	principal := uint8(4)

	s.Illuminate.MarketsReturns([8]common.Address{
		s.Dep.PendleTokenAddress,
		s.Dep.PendleTokenAddress,
		s.Dep.PendleTokenAddress,
		s.Dep.PendleTokenAddress,
		s.Dep.PendleTokenAddress,
		s.Dep.PendleTokenAddress,
		s.Dep.PendleTokenAddress,
		s.Dep.PendleTokenAddress,
	})

	s.PendleToken.BalanceOfReturns(amount)
	s.Env.Blockchain.Commit()

	s.PendleToken.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	tx, err := s.Redeemer.Redeem1(principal, s.Dep.Erc20Address, maturity, forgeId)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that the mocked functions were called as expected
	redeemCall, err := s.Pendle.RedeemAfterExpiryCalled(s.Dep.Erc20Address)
	assert.NoError(err)
	assert.Equal(forgeId, redeemCall.ForgeId)
	assert.Equal(maturity, redeemCall.Maturity)

	underlyingTransfer, err := s.PendleToken.TransferFromCalled(s.Dep.IlluminateAddress)
	assert.NoError(err)
	assert.Equal(amount, underlyingTransfer.Amount)
	assert.Equal(s.Dep.RedeemerAddress, underlyingTransfer.To)
}

func (s *redeemTestSuite) TestSenseRedeem() {
	assert := assert.New(s.T())

	amount := big.NewInt(1000)
	adapter := s.Env.User1.Opts.From
	maturity := big.NewInt(9999999)
	principal := uint8(4)

	s.Illuminate.MarketsReturns([8]common.Address{
		s.Dep.SenseTokenAddress,
		s.Dep.SenseTokenAddress,
		s.Dep.SenseTokenAddress,
		s.Dep.SenseTokenAddress,
		s.Dep.SenseTokenAddress,
		s.Dep.SenseTokenAddress,
		s.Dep.SenseTokenAddress,
		s.Dep.SenseTokenAddress,
	})

	s.SenseToken.BalanceOfReturns(amount)
	s.Env.Blockchain.Commit()

	s.SenseToken.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	tx, err := s.Redeemer.Redeem2(principal, s.Dep.Erc20Address, maturity, s.Dep.SenseAddress, adapter)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that the mocked functions were called as expected
	redeemCall, err := s.Sense.RedeemCalled(adapter)
	assert.NoError(err)
	assert.Equal(maturity, redeemCall.Maturity)
	assert.Equal(amount, redeemCall.Amount)

	underlyingTransfer, err := s.SenseToken.TransferFromCalled(s.Dep.IlluminateAddress)
	assert.NoError(err)
	assert.Equal(amount, underlyingTransfer.Amount)
	assert.Equal(s.Dep.RedeemerAddress, underlyingTransfer.To)
}

func (s *redeemTestSuite) TestSwivelRedeem() {
	assert := assert.New(s.T())

	amount := big.NewInt(1000)
	maturity := big.NewInt(9999999)
	principal := uint8(1)

	s.Illuminate.MarketsReturns([8]common.Address{
		s.Dep.ZcTokenAddress,
		s.Dep.ZcTokenAddress,
		s.Dep.ZcTokenAddress,
		s.Dep.ZcTokenAddress,
		s.Dep.ZcTokenAddress,
		s.Dep.ZcTokenAddress,
		s.Dep.ZcTokenAddress,
		s.Dep.ZcTokenAddress,
	})

	s.ZcToken.BalanceOfReturns(amount)
	s.Env.Blockchain.Commit()

	s.ZcToken.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	tx, err := s.Redeemer.Redeem(principal, s.Dep.Erc20Address, maturity)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that the mocked functions were called as expected
	call, err := s.Swivel.RedeemZcTokenCalled(s.Dep.Erc20Address)
	assert.NoError(err)
	assert.Equal(maturity, call.Maturity)
	assert.Equal(amount, call.Amount)

	underlyingTransfer, err := s.ZcToken.TransferFromCalled(s.Dep.IlluminateAddress)
	assert.NoError(err)
	assert.Equal(amount, underlyingTransfer.Amount)
	assert.Equal(s.Dep.RedeemerAddress, underlyingTransfer.To)
}

func (s *redeemTestSuite) TestElementRedeem() {
	assert := assert.New(s.T())

	amount := big.NewInt(1000)
	maturity := big.NewInt(9999999)
	principal := uint8(3)

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

	s.ElementToken.BalanceOfReturns(amount)
	s.Env.Blockchain.Commit()

	s.ElementToken.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	tx, err := s.Redeemer.Redeem(principal, s.Dep.Erc20Address, maturity)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that the mocked functions were called as expected
	withdrawnAmount, err := s.ElementToken.WithdrawPrincipalCalled(s.Dep.IlluminateAddress)
	assert.NoError(err)
	assert.Equal(amount, withdrawnAmount)

	underlyingTransfer, err := s.ElementToken.TransferFromCalled(s.Dep.IlluminateAddress)
	assert.NoError(err)
	assert.Equal(amount, underlyingTransfer.Amount)
	assert.Equal(s.Dep.RedeemerAddress, underlyingTransfer.To)
}

func (s *redeemTestSuite) TestYieldRedeem() {
	assert := assert.New(s.T())

	amount := big.NewInt(1000)
	maturity := big.NewInt(9999999)
	principal := uint8(3)

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

	s.ElementToken.BalanceOfReturns(amount)
	s.Env.Blockchain.Commit()

	s.ElementToken.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	tx, err := s.Redeemer.Redeem(principal, s.Dep.Erc20Address, maturity)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify that the mocked functions were called as expected
	withdrawnAmount, err := s.ElementToken.WithdrawPrincipalCalled(s.Dep.IlluminateAddress)
	assert.NoError(err)
	assert.Equal(amount, withdrawnAmount)

	underlyingTransfer, err := s.ElementToken.TransferFromCalled(s.Dep.IlluminateAddress)
	assert.NoError(err)
	assert.Equal(amount, underlyingTransfer.Amount)
	assert.Equal(s.Dep.RedeemerAddress, underlyingTransfer.To)
}

func TestRedeemSuite(t *test.T) {
	suite.Run(t, &redeemTestSuite{})
}
