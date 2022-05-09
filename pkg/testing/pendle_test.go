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

type pendleTestSuite struct {
	suite.Suite
	Env    *Env
	Dep    *Dep
	Pendle *mocks.PendleSession
}

func (s *pendleTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.Pendle = &mocks.PendleSession{
		Contract: s.Dep.Pendle,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *pendleTestSuite) TestExpiry() {
	assert := assert.New(s.T())

	expiry := big.NewInt(10)
	tx, err := s.Pendle.ExpiryReturns(expiry)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	pendleMaturity, err := s.Pendle.Expiry()
	assert.Nil(err)
	assert.Equal(expiry, pendleMaturity)
}

func (s *pendleTestSuite) TestYieldToken() {
	assert := assert.New(s.T())

	yieldToken := common.HexToAddress("0x0000000000000000000000000000000000000001")
	tx, err := s.Pendle.YieldTokenReturns(yieldToken)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	pendleUnderlying, err := s.Pendle.YieldToken()
	assert.Nil(err)
	assert.Equal(yieldToken, pendleUnderlying)
}

func TestPendleSuite(t *test.T) {
	suite.Run(t, &pendleTestSuite{})
}
