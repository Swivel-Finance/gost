package illuminatetesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func TestIlluminateSuite(t *test.T) {
	suite.Run(t, &illuminateTestSuite{})
}
