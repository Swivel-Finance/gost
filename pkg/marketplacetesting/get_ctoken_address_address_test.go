package marketplacetesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	assertions "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/marketplace"
	"github.com/swivel-finance/gost/test/mocks"
)

type cTokenAddrSuite struct {
	suite.Suite
	Env           *Env
	Dep           *Dep
	Erc20         *mocks.Erc20Session
	CompoundToken *mocks.CompoundTokenSession
	MarketPlace   *marketplace.MarketPlaceSession // *Session objects are created by the go bindings
}

func (s *cTokenAddrSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	s.Erc20 = &mocks.Erc20Session{
		Contract: s.Dep.Erc20,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.CErc20 = &mocks.CErc20Session{
		Contract: s.Dep.CErc20,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.MarketPlace = &marketplace.MarketPlaceSession{
		Contract: s.Dep.MarketPlace,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *cTokenAddrSuite) TestSetCTokenAddr() {
	assert := assertions.New(s.T())
	// the swivel address must be set
	_, err := s.MarketPlace.SetSwivelAddress(s.Dep.SwivelAddress)
	assert.Nil(err)
	s.Env.Blockchain.Commit()
	// addresses can be BS in this test...
	underlying := s.Dep.Erc20Address
	maturity := big.NewInt(123456789)
	cTokenAddr := s.Dep.CErc20Address

	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.CErc20.UnderlyingReturns(underlying)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.CreateMarket(
		maturity,
		cTokenAddr,
		"awesome market",
		"AM",
	)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	cTokenAddrC, err := s.MarketPlace.CTokenAddress(underlying, maturity)
	assert.Nil(err)
	assert.Equal(cTokenAddr, cTokenAddrC)
}

func TestCTokenAddrSuite(t *test.T) {
	suite.Run(t, &cTokenAddrSuite{})
}
