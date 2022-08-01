package zctokentesting

import (
	"math/big"
	test "testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	assertions "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/tokens"
)

type zcTokenWithdrawSuite struct {
	suite.Suite
	Env             *Env
	Dep             *Dep
	MarketPlace     *mocks.MarketPlaceSession // *Session objects are created by the go bindings
	ZcToken         *tokens.ZcTokenSession
	ZcTokenContract *tokens.ZcToken
	ZcTokenAddress  common.Address
	Protocol        uint8
	Underlying      common.Address
	CToken          common.Address
	Name            string
	Symbol          string
	Decimals        uint8
}

func (s *zcTokenWithdrawSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)
	if err != nil {
		panic(err)
	}

	err = s.Env.Blockchain.AdjustTime(0) // set bc timestamp to 0
	if err != nil {
		panic(err)
	}
	s.Env.Blockchain.Commit()

	s.MarketPlace = &mocks.MarketPlaceSession{
		Contract: s.Dep.MarketPlace,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	// simulate the args that might have been passed to a zctoken ctor
	s.Protocol = uint8(1)
	s.Underlying = common.HexToAddress("0x12345")
	s.CToken = common.HexToAddress("0x23456")
	s.Name = "Awesome Token"
	s.Symbol = "AT"
	s.Decimals = uint8(18)

	// m := s.Dep.Maturity
	// r := s.Dep.MarketPlaceAddress

	// depoloy zctoken on demand
	addr, _, contract, err := tokens.DeployZcToken(s.Env.Owner.Opts, s.Env.Blockchain, s.Protocol, s.Underlying, s.Dep.Maturity, s.CToken,
		s.Dep.MarketPlaceAddress, s.Name, s.Symbol, s.Decimals)

	if err != nil {
		panic(err)
	}
	s.Env.Blockchain.Commit()

	s.ZcTokenAddress = addr
	s.ZcTokenContract = contract

	s.ZcToken = &tokens.ZcTokenSession{
		Contract: contract,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *zcTokenWithdrawSuite) TestPreviewWithdraw() {
	assert := assertions.New(s.T())

	// we need to move past the maturity or revert
	err := s.Env.Blockchain.AdjustTime(MATURITY * time.Second)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	// the previewWithdraw will need a market present with access to .maturityRate > .exchangeRate
	xRate := big.NewInt(5)
	mRate := big.NewInt(10)

	tx, err := s.MarketPlace.ExchangeRateReturns(xRate)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.CreateMarket(s.Protocol, s.Underlying, s.Dep.Maturity, s.CToken, s.ZcTokenAddress, common.HexToAddress("0x34567"), mRate)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	uAmount := big.NewInt(100)
	prev, err := s.ZcToken.PreviewWithdraw(uAmount)
	assert.Nil(err)
	assert.NotNil(prev)

	shouldBe := uAmount.Mul(uAmount, mRate).Div(uAmount, xRate)
	assert.Equal(prev, shouldBe)
}

func (s *zcTokenWithdrawSuite) TestWithdrawIsSender() {
	assert := assertions.New(s.T())

	// we need to move past the maturity or revert
	// err := s.Env.Blockchain.AdjustTime(MATURITY * time.Second)
	// assert.Nil(err)
	// s.Env.Blockchain.Commit()

	// the previewWithdraw will need a market present with access to .maturityRate > .exchangeRate
	// xRate := big.NewInt(5)
	// mRate := big.NewInt(10)

	// tx, err := s.MarketPlace.ExchangeRateReturns(xRate)
	// assert.Nil(err)
	// assert.NotNil(tx)
	// s.Env.Blockchain.Commit()

	// tx, err = s.MarketPlace.CreateMarket(s.Protocol, s.Underlying, s.Dep.Maturity, s.CToken, s.ZcTokenAddress, common.HexToAddress("0x34567"), mRate)
	// assert.Nil(err)
	// assert.NotNil(tx)
	// s.Env.Blockchain.Commit()

	tx, err := s.MarketPlace.AuthRedeemReturns(big.NewInt(1))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	uAmount := big.NewInt(100)

	// No allowance check should be performed here
	tx, err = s.ZcToken.Withdraw(uAmount, s.Env.User1.Opts.From, s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// check that authRedeem is called with expected args
	args, _ := s.MarketPlace.AuthRedeemCalled(s.Protocol)
	assert.Equal(args.Underlying, s.Underlying)
	assert.Equal(args.Maturity, s.Dep.Maturity)
	assert.Equal(args.One, s.Env.Owner.Opts.From)
	assert.Equal(args.Two, s.Env.User1.Opts.From)
	assert.Equal(args.Amount, big.NewInt(200)) // <- is the "preview amount"
}

func (s *zcTokenWithdrawSuite) TestWithdrawIsNotSender() {
	assert := assertions.New(s.T())

	// we need to move past the maturity or revert
	// err := s.Env.Blockchain.AdjustTime(MATURITY * time.Second)
	// assert.Nil(err)
	// s.Env.Blockchain.Commit()

	// the previewWithdraw will need a market present with access to .maturityRate > .exchangeRate
	// xRate := big.NewInt(5)
	// mRate := big.NewInt(10)

	// tx, err := s.MarketPlace.ExchangeRateReturns(xRate)
	// assert.Nil(err)
	// assert.NotNil(tx)
	// s.Env.Blockchain.Commit()

	// tx, err = s.MarketPlace.CreateMarket(s.Protocol, s.Underlying, s.Dep.Maturity, s.CToken, s.ZcTokenAddress, common.HexToAddress("0x34567"), mRate)
	// assert.Nil(err)
	// assert.NotNil(tx)
	// s.Env.Blockchain.Commit()

	tx, err := s.MarketPlace.AuthRedeemReturns(big.NewInt(1))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	uAmount := big.NewInt(100)

	// No allowance made yet -- should fail at this point
	tx, err = s.ZcToken.Withdraw(uAmount, s.Env.Owner.Opts.From, s.Env.User1.Opts.From)
	assert.NotNil(err)
	assert.Nil(tx)

	// provide an allowance such that allowance > preview amount. We'll use the lower level contract so we can set msg.sender
	// to user1 as the session is bound to owner
	opts := &bind.TransactOpts{
		From:   s.Env.User1.Opts.From,
		Signer: s.Env.User1.Opts.Signer,
	}

	tx, err = s.ZcTokenContract.Approve(opts, s.Env.Owner.Opts.From, big.NewInt(400))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// should be able to withdraw now...
	tx, err = s.ZcToken.Withdraw(uAmount, s.Env.Owner.Opts.From, s.Env.User1.Opts.From)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// check that authRedeem is called with expected args
	args, _ := s.MarketPlace.AuthRedeemCalled(s.Protocol)
	assert.Equal(args.Underlying, s.Underlying)
	assert.Equal(args.Maturity, s.Dep.Maturity)
	assert.Equal(args.Two, s.Env.Owner.Opts.From)
	assert.Equal(args.One, s.Env.User1.Opts.From)
	assert.NotEqual(args.Amount, big.NewInt(0))
}

func TestZcTokenWithdrawSuite(t *test.T) {
	suite.Run(t, &zcTokenWithdrawSuite{})
}
