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

	in := big.NewInt(1)
	out := big.NewInt(2)
	path := []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")}
	to := common.HexToAddress("0x3")
	deadline := big.NewInt(4)

	tx, err = s.Pendle.SwapExactTokensForTokens(in, out, path, to, deadline)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	inCalled, err := s.Pendle.InCalled()
	assert.Nil(err)
	assert.Equal(in, inCalled)

	outCalled, err := s.Pendle.OutMinimumCalled()
	assert.Nil(err)
	assert.Equal(out, outCalled)

	var pathCalled [2]common.Address
	path0, err := s.Pendle.PathCalled(big.NewInt(0))
	pathCalled[0] = path0
	path1, err := s.Pendle.PathCalled(big.NewInt(1))
	pathCalled[1] = path1
	assert.Nil(err)
	assert.Equal(path[0], pathCalled[0])
	assert.Equal(path[1], pathCalled[1])
	assert.Equal(len(path), len(pathCalled))

	toCalled, err := s.Pendle.ToCalled()
	assert.Nil(err)
	assert.Equal(to, toCalled)

	deadlineCalled, err := s.Pendle.DeadlineCalled()
	assert.Nil(err)
	assert.Equal(deadline, deadlineCalled)
}

func TestSushiSuite(t *test.T) {
	suite.Run(t, &pendleTestSuite{})
}
