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

type senseTestSuite struct {
	suite.Suite
	Env   *Env
	Dep   *Dep
	Sense *mocks.SenseSession
}

func (s *senseTestSuite) SetupTest() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.Sense = &mocks.SenseSession{
		Contract: s.Dep.Sense,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *senseTestSuite) TestSwapUnderlyingForPTs() {
	assert := assert.New(s.T())
	s.Sense.SwapUnderlyingForPTsReturns(big.NewInt(1000))
	s.Env.Blockchain.Commit()

	adapter := common.BigToAddress(big.NewInt(1))
	tx, err := s.Sense.SwapUnderlyingForPTs(
		adapter,
		big.NewInt(1000),
		big.NewInt(2000),
		big.NewInt(3000),
	)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	swapUnderlying, err := s.Sense.SwapUnderlyingForPTsCalled(adapter)

	assert.NoError(err)
	assert.Equal(big.NewInt(1000), swapUnderlying.Maturity)
	assert.Equal(big.NewInt(2000), swapUnderlying.Amount)
	assert.Equal(big.NewInt(3000), swapUnderlying.MinimumBought)
}

func TestSenseSuite(t *test.T) {
	suite.Run(t, &senseTestSuite{})
}