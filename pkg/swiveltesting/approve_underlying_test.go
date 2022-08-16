package swiveltesting

import (
	"math/big"
	test "testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/swivel"
)

type approveUnderlyingSuite struct {
	suite.Suite
	Env    *Env
	Dep    *Dep
	Erc20  *mocks.Erc20Session
	Swivel *swivel.SwivelSession
}

func (s *approveUnderlyingSuite) SetupTest() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH))
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	s.Erc20 = &mocks.Erc20Session{
		Contract: s.Dep.Erc20,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.Swivel = &swivel.SwivelSession{
		Contract: s.Dep.Swivel,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *approveUnderlyingSuite) TestScheduleApprove() {
	assert := assert.New(s.T())

	tokenAddress := common.HexToAddress("0xiamatoken") // we don't need an actual token here

	tx, err := s.Swivel.ScheduleApproval(tokenAddress)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	hold, _ := s.Swivel.Approvals(tokenAddress)
	assert.Equal(1, hold.Cmp(big.NewInt(259200))) // hold should be greater than the hold constant
}

func (s *approveUnderlyingSuite) TestApprove() {
	assert := assert.New(s.T())

	// stub the underlying to return true or the Safe lib will revert
	tx, err := s.Erc20.ApproveReturns(true)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	uTokens := []common.Address{s.Dep.Erc20Address}
	cTokens := []common.Address{s.Dep.CompoundTokenAddress}

	// reset time to 0
	err = s.Env.Blockchain.AdjustTime(0)
	if err != nil {
		panic(err)
	}
	s.Env.Blockchain.Commit()

	// schedule the approval
	tx, err = s.Swivel.ScheduleApproval(s.Dep.Erc20Address) // use the mock here
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// move past the hold
	err = s.Env.Blockchain.AdjustTime(259200 * time.Second)
	if err != nil {
		panic(err)
	}
	s.Env.Blockchain.Commit()

	// do it
	tx, err = s.Swivel.ApproveUnderlying(uTokens, cTokens)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// we just care that the args were passed thru...
	amount, err := s.Erc20.ApproveCalled(s.Dep.CompoundTokenAddress)
	assert.Nil(err)
	assert.NotNil(amount)
	// it should be for the max sol integer
	max := new(big.Int)
	max = max.Exp(big.NewInt(2), big.NewInt(256), nil).Sub(max, big.NewInt(1))
	assert.Equal(amount, max)
}

func TestApproveUnderlyingSuite(t *test.T) {
	suite.Run(t, &approveUnderlyingSuite{})
}
