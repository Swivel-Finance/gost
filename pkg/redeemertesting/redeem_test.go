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
	Env        *Env
	Dep        *Dep
	Erc20      *mocks.Erc20Session
	Illuminate *mocks.IlluminateSession
	Yield      *mocks.YieldSession
	ZcToken    *mocks.ZcTokenSession
	Swivel     *mocks.SwivelSession
	Router     *mocks.RouterSession
	Redeemer   *redeemer.RedeemerSession
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

	s.Router = &mocks.RouterSession{
		Contract: s.Dep.Router,
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

func (s *redeemTestSuite) TestRedeemIlluminateTempusApwine() {
	assert := assert.New(s.T())

	amount := big.NewInt(1000)
	maturity := big.NewInt(9999999)

	markets := [8]common.Address{
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
	}
	s.Illuminate.MarketsReturns(markets)
	s.Env.Blockchain.Commit()

	s.Erc20.BalanceOfReturns(amount)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	tx, err := s.Redeemer.Redeem0(0, s.Dep.Erc20Address, maturity, s.Dep.IlluminateAddress)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify mocks were called as expected
	transfer, err := s.Erc20.TransferFromCalled(s.Dep.IlluminateAddress)
	assert.NoError(err)
	assert.Equal(amount, transfer.Amount)
	assert.True(transfer.To == s.Dep.RedeemerAddress)

	calledAmount, err := s.Router.AmountCalled()
	assert.NoError(err)
	assert.Equal(amount, calledAmount)

	calledYieldAmount, err := s.Router.YieldAmountCalled()
	assert.NoError(err)
	assert.Equal(calledYieldAmount.Cmp(big.NewInt(0)), 0)

	calledOwner, err := s.Router.OwnerCalled()
	assert.NoError(err)
	assert.Equal(s.Dep.IlluminateAddress, calledOwner)

	calledRecipient, err := s.Router.RecipientCalled()
	assert.NoError(err)
	assert.Equal(s.Dep.RedeemerAddress, calledRecipient)
}

func TestRedeemSuite(t *test.T) {
	suite.Run(t, &redeemTestSuite{})
}
