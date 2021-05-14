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

type addNotionalNoVaultSuite struct {
	suite.Suite
	Env          *Env
	Dep          *Dep
	CErc20       *mocks.CErc20Session
	VaultTracker *vaulttracker.VaultTrackerSession // *Session objects are created by the go bindings
}

func (s *addNotionalNoVaultSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)
	if err != nil {
		panic(err)
	}

	s.CErc20 = &mocks.CErc20Session{
		Contract: s.Dep.CErc20,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.VaultTracker = &vaulttracker.VaultTrackerSession{
		Contract: s.Dep.VaultTracker,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *addNotionalNoVaultSuite) TestAddNotional() {
	assert := assertions.New(s.T())

	rate1 := big.NewInt(123456789)
	tx, err := s.CErc20.ExchangeRateCurrentReturns(rate1)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// no vault found for Owner
	vault, err := s.VaultTracker.Vaults(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.NotNil(vault)
	assert.Equal(vault.Redeemable.Cmp(ZERO), 0)
	assert.Equal(vault.Notional.Cmp(ZERO), 0)
	assert.Equal(vault.ExchangeRate.Cmp(ZERO),0)

	// call AddNotional for Owner with no vault
	caller := s.Env.Owner.Opts.From
	amount1 := big.NewInt(10000000)
	redeemable1 := ZERO
	tx, err = s.VaultTracker.AddNotional(caller, amount1)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// found vault for Owner
	vault, err = s.VaultTracker.Vaults(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.NotNil(vault)
	assert.Equal(amount1, vault.Notional)
	assert.Equal(rate1, vault.ExchangeRate)
	assert.Equal(vault.Redeemable.Cmp(redeemable1), 0)
}

func TestAddNotionalNoVaultSuite(t *test.T) {
	suite.Run(t, &addNotionalNoVaultSuite{})
}
