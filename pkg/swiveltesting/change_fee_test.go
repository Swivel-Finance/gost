package swiveltesting

import (
	"math/big"
	test "testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/swivel"
)

type changeFeeSuite struct {
	suite.Suite
	Env    *Env
	Dep    *Dep
	Swivel *swivel.SwivelSession // *Session objects are created by the go bindings
}

func (s *changeFeeSuite) SetupSuite() {
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

func (s *changeFeeSuite) TestChangeFeeFailsNotScheduled() {
	assert := assert.New(s.T())

	// there is no fee change scheduled as of yet
	when, err := s.Swivel.FeeChange()
	assert.Nil(err)
	assert.Equal(when.Cmp(big.NewInt(0)), 0)

	// a new set of feenominators
	fees := [4]uint16{uint16(100), uint16(200), uint16(300), uint16(400)}

	tx, err := s.Swivel.ChangeFee(fees)
	assert.Nil(tx)
	assert.NotNil(err)
}

func (s *changeFeeSuite) TestChangeFeeSuccess() {
	assert := assert.New(s.T())

	// there is no fee change scheduled as of yet
	when, err := s.Swivel.FeeChange()
	assert.Nil(err)
	assert.Equal(when.Cmp(big.NewInt(0)), 0)

	// schedule one
	tx, err := s.Swivel.ScheduleFeeChange()
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// fee change date is set now
	when, err = s.Swivel.FeeChange()
	assert.Nil(err)
	// s.T().Log(when)
	assert.Equal(when.Cmp(big.NewInt(0)), 1) // no longer 0

	// we need to move past the when date now (the hold is 259310)
	err = s.Env.Blockchain.AdjustTime(259310 * time.Second)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	// note that indexes 1, 4 won't change
	fees := [4]uint16{uint16(100), uint16(600), uint16(100), uint16(200)}

	tx, err = s.Swivel.ChangeFee(fees)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// check the new fee denominators
	fee1, _ := s.Swivel.Feenominators(big.NewInt(0))
	assert.Equal(fee1, uint16(100))
	fee2, _ := s.Swivel.Feenominators(big.NewInt(1))
	assert.Equal(fee2, uint16(600))
	fee3, _ := s.Swivel.Feenominators(big.NewInt(2))
	assert.Equal(fee3, uint16(100))
	fee4, _ := s.Swivel.Feenominators(big.NewInt(3))
	assert.Equal(fee4, uint16(200))

	// the changeFee should be reset
	when, err = s.Swivel.FeeChange()
	assert.Nil(err)
	assert.Equal(when.Cmp(big.NewInt(0)), 0)
}

func TestChangeFeeSuite(t *test.T) {
	suite.Run(t, &changeFeeSuite{})
}
