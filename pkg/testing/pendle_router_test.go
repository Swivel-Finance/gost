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

type pendleRouterTestSuite struct {
	suite.Suite
	Env          *Env
	Dep          *Dep
	PendleRouter *mocks.PendleRouterSession
}

func (s *pendleRouterTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.PendleRouter = &mocks.PendleRouterSession{
		Contract: s.Dep.PendleRouter,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *pendleRouterTestSuite) TestData() {
	assert := assert.New(s.T())
	example := common.HexToAddress("0x0000000000000000000000000000000000000001")
	s.PendleRouter.DataReturns(example)
	s.Env.Blockchain.Commit()

	data, err := s.PendleRouter.Data()
	assert.NoError(err)
	assert.Equal(example, data)
}

func (s *pendleRouterTestSuite) TestTokenizeYield() {
	assert := assert.New(s.T())

	returnAmount := big.NewInt(100)
	s.PendleRouter.TokenizeYieldReturns(returnAmount)
	s.Env.Blockchain.Commit()

	id := [32]byte{1, 2, 3}
	underlying := common.HexToAddress("0x0000000000000000000000000000000000000001")
	maturity := big.NewInt(50)
	amount := big.NewInt(101)
	destination := common.HexToAddress("0x0000000000000000000000000000000000000002")

	s.PendleRouter.TokenizeYield(id, underlying, maturity, amount, destination)
	s.Env.Blockchain.Commit()

	returnedId, err := s.PendleRouter.TokenizeYieldIdCalled()
	assert.NoError(err)
	assert.Equal(id, returnedId)

	returnedUnderlying, err := s.PendleRouter.TokenizeYieldUnderlyingCalled()
	assert.NoError(err)
	assert.Equal(underlying, returnedUnderlying)

	returnedMaturity, err := s.PendleRouter.TokenizeYieldMaturityCalled()
	assert.NoError(err)
	assert.Equal(maturity, returnedMaturity)

	returnedAmount, err := s.PendleRouter.TokenizeYieldAmountCalled()
	assert.NoError(err)
	assert.Equal(amount, returnedAmount)

	returnedToken, err := s.PendleRouter.TokenizeYieldTokenCalled()
	assert.NoError(err)
	assert.Equal(destination, returnedToken)

	assert.Equal(1, 2)
}

func TestPendleRouterSuite(t *test.T) {
	suite.Run(t, &pendleRouterTestSuite{})
}
