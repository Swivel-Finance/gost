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

func (s *pendleTestSuite) TestSwapExactTokensForTokens() {
	assert := assert.New(s.T())
	returnedArray := []*big.Int{big.NewInt(1), big.NewInt(2)}
	tx, err := s.Pendle.SwapExactTokensForTokensReturns(returnedArray)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	amount := big.NewInt(1)
	minimumBought := big.NewInt(2)
	path := []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")}
	to := common.HexToAddress("0x3")
	deadline := big.NewInt(4)

	tx, err = s.Pendle.SwapExactTokensForTokens(amount, minimumBought, path, to, deadline)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	swap, err := s.Pendle.SwapExactTokensForTokensCalled(to)
	assert.Equal(amount, swap.Amount)
	assert.Equal(minimumBought, swap.MinimumBought)
	assert.Equal(deadline, swap.Deadline)
}

func TestPendleSuite(t *test.T) {
	suite.Run(t, &pendleTestSuite{})
}
