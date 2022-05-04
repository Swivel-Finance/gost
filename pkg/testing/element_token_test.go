package testing

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/mocks"
)

type elementTokenTestSuite struct {
	suite.Suite
	Env          *Env
	Dep          *Dep
	ElementToken *mocks.ElementTokenSession
}

func (s *elementTokenTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.ElementToken = &mocks.ElementTokenSession{
		Contract: s.Dep.ElementToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *elementTokenTestSuite) TestMaturity() {
	assert := assert.New(s.T())

	// set the maturity to 1
	tx, err := s.ElementToken.UnlockTimestampReturns(big.NewInt(1111))
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	maturity, err := s.ElementToken.UnlockTimestamp()
	assert.Nil(err)
	assert.Equal(big.NewInt(1111), maturity)
}

func (s *elementTokenTestSuite) TestUnderlying() {
	assert := assert.New(s.T())
	underlying := common.HexToAddress("0x0123")

	tx, err := s.ElementToken.UnderlyingReturns(underlying)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	underlyingResult, err := s.ElementToken.Underlying()
	assert.Nil(err)
	assert.Equal(underlying, underlyingResult)
}

func TestElementTokenSuite(t *test.T) {
	suite.Run(t, &elementTokenTestSuite{})
}
