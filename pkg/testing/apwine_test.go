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

type apWineRouterTestSuite struct {
	suite.Suite
	Env          *Env
	Dep          *Dep
	APWineRouter *mocks.APWineRouterSession
}

func (s *apWineRouterTestSuite) SetupTest() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.APWineRouter = &mocks.APWineRouterSession{
		Contract: s.Dep.APWineRouter,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *apWineRouterTestSuite) TestSwapExactAmountIn() {
	assert := assert.New(s.T())

	s.APWineRouter.SwapExactAmountInReturns(big.NewInt(1000))
	s.Env.Blockchain.Commit()

	s.APWineRouter.SwapExactAmountIn(
		big.NewInt(1000),
		big.NewInt(2000),
		big.NewInt(3000),
		big.NewInt(4000),
		big.NewInt(5000),
		common.HexToAddress("0x0000000000000000000000000000000000000002"),
	)
	s.Env.Blockchain.Commit()

	id, err := s.APWineRouter.IdCalled()
	assert.Nil(err)
	assert.Equal(big.NewInt(1000), id)

	tokenIn, err := s.APWineRouter.TokenInCalled()
	assert.Nil(err)
	assert.Equal(big.NewInt(2000), tokenIn)

	amount, err := s.APWineRouter.AmountCalled()
	assert.Nil(err)
	assert.Equal(big.NewInt(3000), amount)

	tokenOut, err := s.APWineRouter.TokenOutCalled()
	assert.Nil(err)
	assert.Equal(big.NewInt(4000), tokenOut)

	minimumAmount, err := s.APWineRouter.MinimumAmountCalled()
	assert.Nil(err)
	assert.Equal(big.NewInt(5000), minimumAmount)

	to, err := s.APWineRouter.ToCalled()
	assert.Nil(err)
	assert.Equal(common.HexToAddress("0x0000000000000000000000000000000000000002"), to)
}

func TestAPWineSuite(t *test.T) {
	suite.Run(t, &apWineRouterTestSuite{})
}
