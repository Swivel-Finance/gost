package swiveltesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	assertions "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/swivel"
)

type redeemVaultInterestSuite struct {
	suite.Suite
	Env         *Env
	Dep         *Dep
	Erc20       *mocks.Erc20Session
	AavePool    *mocks.AavePoolSession
	MarketPlace *mocks.MarketPlaceSession
	Swivel      *swivel.SwivelSession
}

func (s *redeemVaultInterestSuite) SetupTest() {
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

	s.AavePool = &mocks.AavePoolSession{
		Contract: s.Dep.AavePool,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.MarketPlace = &mocks.MarketPlaceSession{
		Contract: s.Dep.MarketPlace,
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

func (s *redeemVaultInterestSuite) TestRedeemVaultInterest() {
	assert := assertions.New(s.T())
	underlying := s.Dep.Erc20Address
	maturity := s.Dep.Maturity

	// stub the underlying to return true or the Safe lib will revert
	tx, err := s.Erc20.TransferReturns(true)
	assert.NotNil(tx)
	assert.Nil(err)

	tx, err = s.AavePool.WithdrawReturns(big.NewInt(10000))
	assert.NotNil(tx)
	assert.Nil(err)

	s.Env.Blockchain.Commit()

	// cToken addres not important in this case...
	tx, err = s.MarketPlace.CTokenAndAdapterAddressReturns(common.HexToAddress("0x123"), common.HexToAddress("0x234"))
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	redeemed := big.NewInt(12345)
	tx, err = s.MarketPlace.RedeemVaultInterestReturns(redeemed)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	tx, err = s.Swivel.RedeemVaultInterest(uint8(4), underlying, maturity)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// underlying tranfer should have been called with an amount redeemed
	transferred, err := s.Erc20.TransferCalled(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.Equal(redeemed, transferred)

	// check the args sent to the aave pool mock
	txArgs, err := s.AavePool.WithdrawCalled(underlying)
	assert.Nil(err)
	assert.NotNil(txArgs)
	assert.Equal(txArgs.To, s.Dep.SwivelAddress)
	assert.Equal(txArgs.Amount, redeemed)
}

func (s *redeemVaultInterestSuite) TestRedeemVaultInterestUnderlyingFails() {
	assert := assertions.New(s.T())
	underlying := s.Dep.Erc20Address
	maturity := s.Dep.Maturity

	tx, err := s.AavePool.WithdrawReturns(big.NewInt(0))
	assert.NotNil(tx)
	assert.Nil(err)

	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.CTokenAndAdapterAddressReturns(common.HexToAddress("0x123"), common.HexToAddress("0x234"))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.RedeemVaultInterestReturns(big.NewInt(12345))
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	tx, err = s.Swivel.RedeemVaultInterest(uint8(4), underlying, maturity)
	assert.NotNil(err)
	assert.Regexp("transfer failed", err.Error())
	assert.Nil(tx)
}

func TestRedeemVaultInterestSuite(t *test.T) {
	suite.Run(t, &redeemVaultInterestSuite{})
}
