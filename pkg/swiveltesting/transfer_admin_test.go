package swiveltesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/swivel"
)

type transferAdminSuite struct {
	suite.Suite
	Env         *Env
	Dep         *Dep
	SwivelOwner *swivel.SwivelSession // *Session objects are created by the go bindings
	SwivelUser  *swivel.SwivelSession
}

func (s *transferAdminSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.SwivelOwner = &swivel.SwivelSession{
		Contract: s.Dep.Swivel,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	// using a session bound to each so we can switch the admin back...
	s.SwivelUser = &swivel.SwivelSession{
		Contract: s.Dep.Swivel,
		CallOpts: bind.CallOpts{From: s.Env.User1.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.User1.Opts.From,
			Signer: s.Env.User1.Opts.Signer,
		},
	}
}

func (s *transferAdminSuite) TestTransferAdmin() {
	assert := assert.New(s.T())

	// check the current address, it should be owner
	// NOTE don't confuse the 2 bound sessions for the contract, admin is the same in both...
	addr, err := s.SwivelUser.Admin()
	assert.Nil(err)
	assert.Equal(addr, s.Env.Owner.Opts.From)

	// transfer it to User1, must be done from owner (current admin)
	tx, err := s.SwivelOwner.SetAdmin(s.Env.User1.Opts.From)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// admin address should have changed
	addr, err = s.SwivelOwner.Admin()
	assert.Nil(err)
	assert.Equal(addr, s.Env.User1.Opts.From)

	// change it back so everything doesn't fail... (must be done from user, as that's the admin now)
	tx, err = s.SwivelUser.SetAdmin(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// admin address is owner again...
	addr, err = s.SwivelUser.Admin()
	assert.Nil(err)
	assert.Equal(addr, s.Env.Owner.Opts.From)
}

func TestTransferAdminSuite(t *test.T) {
	suite.Run(t, &transferAdminSuite{})
}
