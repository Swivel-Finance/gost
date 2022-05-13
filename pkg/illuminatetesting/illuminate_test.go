package illuminatetesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/illuminate"
	"github.com/swivel-finance/gost/test/mocks"
)

type illuminateTestSuite struct {
	suite.Suite
	Env        *Env
	Dep        *Dep
	Erc20      *mocks.Erc20Session
	Illuminate *illuminate.IlluminateSession
}

func (s *illuminateTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.Erc20 = &mocks.Erc20Session{
		Contract: s.Dep.Erc20,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.Illuminate = &illuminate.IlluminateSession{
		Contract: s.Dep.Illuminate,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *illuminateTestSuite) TestConstruction() {
	assert := assert.New(s.T())

	rAddr, _ := s.Illuminate.Redeemer()

	assert.Equal(rAddr, s.Dep.RedeemerAddress)
}

func (s *illuminateTestSuite) TestCreateMarket() {
	assert := assert.New(s.T())

	// stub the erc approve to return true or safe will revert
	s.Erc20.ApproveReturns(true)

	s.Env.Blockchain.Commit()

	maturity := big.NewInt(100000)

	// a fixed len array of 7 as illuminate is set in the method...
	principals := [7]common.Address{
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
	}

	tx, err := s.Illuminate.CreateMarket(s.Dep.Erc20Address, maturity, principals, "spam token", "SPAM", uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// we should see a market for that market-pair that has the new ZcToken address at [0] (as illuminate)
	// remember that you cannot fetch the entire array, but must provide the index of the principal you want...
	illPrincipal, err := s.Illuminate.Markets(s.Dep.Erc20Address, maturity, big.NewInt(0))
	assert.Nil(err)
	// this will be the only one NOT the mock erc address...
	assert.NotEqual(illPrincipal, s.Dep.Erc20Address)

	// the first address in the array passed to the function should now be at [1] (as swivel)
	swivPrincipal, err := s.Illuminate.Markets(s.Dep.Erc20Address, maturity, big.NewInt(1))
	assert.Nil(err)
	assert.Equal(swivPrincipal, s.Dep.Erc20Address)

	// should be the max amount approved in the contract...
	max, exp := big.NewInt(2), big.NewInt(256)
	max.Exp(max, exp, nil)
	max.Sub(max, big.NewInt(1))

	// we can see what was passed to approve on the last iteration of the for loop
	approved, err := s.Erc20.ApproveCalled(s.Dep.RedeemerAddress)
	assert.Nil(err)
	assert.Equal(approved, max)
}

func TestIlluminateSuite(t *test.T) {
	suite.Run(t, &illuminateTestSuite{})
}
