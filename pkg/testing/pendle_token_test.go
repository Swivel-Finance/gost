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

type pendleTokenTestSuite struct {
	suite.Suite
	Env         *Env
	Dep         *Dep
	PendleToken *mocks.PendleTokenSession
}

func (s *pendleTokenTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.PendleToken = &mocks.PendleTokenSession{
		Contract: s.Dep.PendleToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *pendleTokenTestSuite) TestExpiry() {
	assert := assert.New(s.T())

	expiry := big.NewInt(10)
	tx, err := s.PendleToken.ExpiryReturns(expiry)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	pendleMaturity, err := s.PendleToken.Expiry()
	assert.Nil(err)
	assert.Equal(expiry, pendleMaturity)
}

func (s *pendleTokenTestSuite) TestYieldToken() {
	assert := assert.New(s.T())

	yieldToken := common.HexToAddress("0x0000000000000000000000000000000000000001")
	tx, err := s.PendleToken.YieldTokenReturns(yieldToken)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	pendleUnderlying, err := s.PendleToken.YieldToken()
	assert.Nil(err)
	assert.Equal(yieldToken, pendleUnderlying)
}

func TestPendleTokenSuite(t *test.T) {
	suite.Run(t, &pendleTokenTestSuite{})
}
