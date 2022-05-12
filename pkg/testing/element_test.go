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

type elementTestSuite struct {
	suite.Suite
	Env     *Env
	Dep     *Dep
	Element *mocks.ElementSession
}

func (s *elementTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.Element = &mocks.ElementSession{
		Contract: s.Dep.Element,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *elementTestSuite) TestDeploy() {
	assert := assert.New(s.T())

	deadline := big.NewInt(1000)
	limit := big.NewInt(2000)
	fundManagement := mocks.FundManagement{
		Sender:              s.Env.Owner.Opts.From,
		Recipient:           common.HexToAddress("0x0000000000000000000000000000000000000001"),
		FromInternalBalance: true,
		ToInternalBalance:   true,
	}
	singleSwap := mocks.SingleSwap{
		Amount:   big.NewInt(1000),
		AssetIn:  common.BigToAddress(big.NewInt(1)),
		AssetOut: common.BigToAddress(big.NewInt(2)),
	}

	s.Element.Swap(singleSwap, fundManagement, limit, deadline)
	s.Env.Blockchain.Commit()

	swap, err := s.Element.SwapCalled(fundManagement.Sender)
	assert.NoError(err)
	assert.Equal(deadline, swap.Deadline)
	assert.Equal(limit, swap.Limit)
	assert.Equal(fundManagement.Recipient, swap.Recipient)
	assert.Equal(singleSwap.Amount, swap.SwapAmount)
}

func TestElementSuite(t *test.T) {
	suite.Run(t, &elementTestSuite{})
}
