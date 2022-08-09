package vaulttrackertesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	assertions "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/vaulttracker"
)

type ratesSuite struct {
	suite.Suite
	Env           *Env
	Dep           *Dep
	CompoundToken *mocks.CompoundTokenSession
	VaultTracker  *vaulttracker.VaultTrackerSession // *Session objects are created by the go bindings
}

func (s *ratesSuite) SetupTest() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)
	if err != nil {
		panic(err)
	}

	s.CompoundToken = &mocks.CompoundTokenSession{
		Contract: s.Dep.CompoundToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.VaultTracker = &vaulttracker.VaultTrackerSession{
		Contract: s.Dep.VaultTracker,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *ratesSuite) TestRatesWhenMaturity0() {
	assert := assertions.New(s.T())

	rate1 := big.NewInt(123456789)

	tx, err := s.CompoundToken.ExchangeRateCurrentReturns(rate1)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// at this point the vault's maturityRate is 0
	mRate, xRate, err := s.VaultTracker.Rates()
	assert.Nil(err)
	assert.NotNil(mRate)
	assert.NotNil(xRate)
	// when maturity is 0 we will get exchange rate twice
	assert.Equal(mRate, rate1)
	assert.Equal(xRate, rate1)
}

func (s *ratesSuite) TestRatesWhenMaturityNot0() {
	assert := assertions.New(s.T())

	rate1 := big.NewInt(123456789)
	rate2 := big.NewInt(234567890)

	tx, err := s.CompoundToken.ExchangeRateCurrentReturns(rate1)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// we can use matureVault to set a maturityRate
	tx, err = s.VaultTracker.MatureVault(rate2)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	mRate, xRate, err := s.VaultTracker.Rates()
	assert.Nil(err)
	assert.NotNil(mRate)
	assert.NotNil(xRate)
	// when maturity is !0 we will get each
	assert.Equal(mRate, rate2)
	assert.Equal(xRate, rate1)
}

func TestRatesSuite(t *test.T) {
	suite.Run(t, &ratesSuite{})
}
