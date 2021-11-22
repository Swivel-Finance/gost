package vaulttrackertesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	assertions "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/vaulttracker"
)

type matureVaultSuite struct {
	suite.Suite
	Env          *Env
	Dep          *Dep
	VaultTracker *vaulttracker.VaultTrackerSession // *Session objects are created by the go bindings
}

func (s *matureVaultSuite) SetupTest() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)
	if err != nil {
		panic(err)
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

func (s *matureVaultSuite) TestMatureVault() {
	assert := assertions.New(s.T())

	rate1 := big.NewInt(123456789)

	tx, err := s.VaultTracker.MatureVault(rate1)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	rate2, err := s.VaultTracker.MaturityRate()
	assert.Nil(err)
	assert.Equal(rate2, rate1)
}

func TestTrackerMatureVaultSuite(t *test.T) {
	suite.Run(t, &matureVaultSuite{})
}
