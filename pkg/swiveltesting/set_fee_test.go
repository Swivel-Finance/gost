package swiveltesting

import (
	"math/big"
	"strings"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/swivel"
)

type setFeeSuite struct {
	suite.Suite
	Env    *Env
	Dep    *Dep
	Swivel *swivel.SwivelSession // *Session objects are created by the go bindings
}

func (s *setFeeSuite) SetupSuite() {
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

func (s *setFeeSuite) TestSetFee() {
	assert := assert.New(s.T())

	// inspect the original fee first
	feenominator, err := s.Swivel.Feenominators(big.NewInt(1))
	assert.Nil(err)
	assert.Equal(feenominator, uint16(600))

	tx, err := s.Swivel.SetFee(uint16(1), uint16(500))
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// feenominator should have changed...
	feenominator, err = s.Swivel.Feenominators(big.NewInt(1))
	assert.Nil(err)
	assert.Equal(feenominator, uint16(500))
}

func (s *setFeeSuite) TestSetFeeFails() {
	assert := assert.New(s.T())

	// any feenominator lower than 33 should revert
	tx, err := s.Swivel.SetFee(uint16(2), uint16(25))
	assert.Nil(tx)
	assert.NotNil(err)
	assert.True(strings.Contains(err.Error(), "fee too high"))
}

func TestSetFeeSuite(t *test.T) {
	suite.Run(t, &setFeeSuite{})
}
