package vaulttrackertesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	assertions "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/vaulttracker"
)

type trackerCtorSuite struct {
	suite.Suite
	Env          *Env
	Dep          *Dep
	VaultTracker *vaulttracker.VaultTrackerSession // *Session objects are created by the go bindings
}

func (s *trackerCtorSuite) SetupSuite() {
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

func (s *trackerCtorSuite) TestAdmin() {
	assert := assertions.New(s.T())
	addr, err := s.VaultTracker.Admin()
	assert.Nil(err)
	assert.Equal(addr, s.Env.Owner.Opts.From)
}

func (s *trackerCtorSuite) TestCTokenAddress() {
	assert := assertions.New(s.T())
	addr, err := s.VaultTracker.CTokenAddress()
	assert.Nil(err)
	assert.Equal(s.Dep.CErc20Address, addr)
}

func (s *trackerCtorSuite) TestMaturity() {
	assert := assertions.New(s.T())
	maturity, err := s.VaultTracker.Maturity()
	assert.Nil(err)
	assert.Equal(maturity, s.Dep.Maturity)
}

func TestTrackerCtorSuite(t *test.T) {
	suite.Run(t, &trackerCtorSuite{})
}
