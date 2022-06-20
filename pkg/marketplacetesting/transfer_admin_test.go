package marketplacetesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/marketplace"
)

type transferAdminSuite struct {
	suite.Suite
	Env              *Env
	Dep              *Dep
	MarketPlaceOwner *marketplace.MarketPlaceSession // *Session objects are created by the go bindings
	MarketPlaceUser  *marketplace.MarketPlaceSession
}

func (s *transferAdminSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	s.MarketPlaceOwner = &marketplace.MarketPlaceSession{
		Contract: s.Dep.MarketPlace,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	// we'll bind a second instance to user so we cann change it back without having to pass the opts directly...
	s.MarketPlaceUser = &marketplace.MarketPlaceSession{
		Contract: s.Dep.MarketPlace,
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
	addr, err := s.MarketPlaceUser.Admin()
	assert.Nil(err)
	assert.Equal(addr, s.Env.Owner.Opts.From)

	// transfer it to User1, must be done from owner (current admin)
	tx, err := s.MarketPlaceOwner.SetAdmin(s.Env.User1.Opts.From)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// admin address should have changed
	addr, err = s.MarketPlaceOwner.Admin()
	assert.Nil(err)
	assert.Equal(addr, s.Env.User1.Opts.From)

	// change it back so everything doesn't fail... (must be done from user, as that's the admin now)
	tx, err = s.MarketPlaceUser.SetAdmin(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// admin address is owner again...
	addr, err = s.MarketPlaceUser.Admin()
	assert.Nil(err)
	assert.Equal(addr, s.Env.Owner.Opts.From)
}

func TestTransferAdminSuite(t *test.T) {
	suite.Run(t, &transferAdminSuite{})
}
