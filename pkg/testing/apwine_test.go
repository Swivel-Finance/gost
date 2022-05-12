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

type apWineTestSuite struct {
	suite.Suite
	Env    *Env
	Dep    *Dep
	APWine *mocks.APWineSession
}

func (s *apWineTestSuite) SetupTest() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.APWine = &mocks.APWineSession{
		Contract: s.Dep.APWine,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *apWineTestSuite) TestSwapExactAmountIn() {
	assert := assert.New(s.T())

	s.APWine.SwapExactAmountInReturns(big.NewInt(1000))
	s.Env.Blockchain.Commit()

	to := common.HexToAddress("0x0000000000000000000000000000000000000002")
	s.APWine.SwapExactAmountIn(
		big.NewInt(1000),
		big.NewInt(2000),
		big.NewInt(3000),
		big.NewInt(4000),
		big.NewInt(5000),
		to,
	)
	s.Env.Blockchain.Commit()

	swap, err := s.APWine.SwapExactAmountInCalled(to)
	assert.NoError(err)
	assert.Equal(big.NewInt(1000), swap.Id)
	assert.Equal(big.NewInt(2000), swap.TokenIn)
	assert.Equal(big.NewInt(3000), swap.Amount)
	assert.Equal(big.NewInt(4000), swap.TokenOut)
	assert.Equal(big.NewInt(5000), swap.MinimumAmount)
}

func TestAPWineSuite(t *test.T) {
	suite.Run(t, &apWineTestSuite{})
}
