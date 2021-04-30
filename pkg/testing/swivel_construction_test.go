package testing

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/contracts/swivel"
)

type swivelCtorSuite struct {
	suite.Suite
	Env    *Env
	Dep    *Dep
	Swivel *swivel.SwivelSession // *Session objects are created by the go bindings
}

func (s *swivelCtorSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.Swivel = &swivel.SwivelSession{
		Contract: s.Dep.Swivel,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *swivelCtorSuite) TestName() {
	assert := assert.New(s.T())
	name, err := s.Swivel.NAME()
	assert.Nil(err)
	assert.Equal(name, "Swivel Finance")
}

func (s *swivelCtorSuite) TestVersion() {
	assert := assert.New(s.T())
	verz, err := s.Swivel.VERSION()
	assert.Nil(err)
	assert.Equal(verz, "2.0.0")
}

func (s *swivelCtorSuite) TestAdmin() {
	assert := assert.New(s.T())
	addr, err := s.Swivel.Admin()
	assert.Nil(err)
	assert.Equal(addr, s.Env.Owner.Opts.From)
}

func (s *swivelCtorSuite) TestDomain() {
	assert := assert.New(s.T())
	separator, err := s.Swivel.DOMAIN()
	assert.Nil(err)
	assert.Equal(32, len(separator))
}

func TestSwivelCtorSuite(t *test.T) {
	suite.Run(t, &swivelCtorSuite{})
}
