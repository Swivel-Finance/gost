package swiveltesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	// "github.com/swivel-finance/gost/internal/helpers"
	// "github.com/swivel-finance/gost/test/fakes"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/swivel"
)

type splitCombineSuite struct {
	suite.Suite
	Env         *Env
	Dep         *Dep
	Erc20       *mocks.Erc20Session
	Compound    *mocks.CompoundSession
	MarketPlace *mocks.MarketPlaceSession
	Swivel      *swivel.SwivelSession
}

func (s *splitCombineSuite) SetupTest() {
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

	s.Compound = &mocks.CompoundSession{
		Contract: s.Dep.Compound,
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

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.Swivel = &swivel.SwivelSession{
		Contract: s.Dep.Swivel,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *splitCombineSuite) TestSplit() {
	assert := assert.New(s.T())

	// stub the underlying to return true or the Safe lib will revert
	tx, err := s.Erc20.TransferFromReturns(true)
	assert.NotNil(tx)
	assert.Nil(err)

	tx, err = s.Compound.DepositReturns(big.NewInt(0))
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	// the marketplace mock...
	tx, err = s.MarketPlace.CTokenAndAdapterAddressReturns(s.Dep.CompoundTokenAddress, s.Dep.CompoundAddress) // must use the actual dep addr here
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.MintZcTokenAddingNotionalReturns(true)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	maturity := big.NewInt(123456789)
	amount := big.NewInt(1000)

	tx, err = s.Swivel.SplitUnderlying(uint8(1), s.Dep.Erc20Address, maturity, amount)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// we just care that the args were passed thru...
	splitArgs, err := s.MarketPlace.MintZcTokenAddingNotionalCalled(uint8(1))
	assert.Nil(err)
	assert.NotNil(splitArgs)
	assert.Equal(splitArgs.Maturity, maturity)
	assert.Equal(splitArgs.Amount, amount)
	// msg.sender will be owner here
	assert.Equal(splitArgs.One, s.Env.Owner.Opts.From)
}

func (s *splitCombineSuite) TestCombine() {
	assert := assert.New(s.T())

	// stub the underlying to return true or the Safe lib will revert
	tx, err := s.Erc20.TransferReturns(true)
	assert.NotNil(tx)
	assert.Nil(err)

	// stub (Cerc20) to return 0
	tx, err = s.Compound.RedeemUnderlyingReturns(big.NewInt(0))
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	// the marketplace mock...
	tx, err = s.MarketPlace.CTokenAndAdapterAddressReturns(s.Dep.CompoundTokenAddress, s.Dep.CompoundAddress) // must use the actual dep addr here
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.BurnZcTokenRemovingNotionalReturns(true)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	maturity := big.NewInt(123456789)
	amount := big.NewInt(1000)

	tx, err = s.Swivel.CombineTokens(uint8(1), s.Dep.Erc20Address, maturity, amount)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// we just care that the args were passed thru...
	combineArgs, err := s.MarketPlace.BurnZcTokenRemovingNotionalCalled(uint8(1))
	assert.Nil(err)
	assert.NotNil(combineArgs)
	assert.Equal(combineArgs.Maturity, maturity)
	assert.Equal(combineArgs.Amount, amount)
	// msg.sender will be owner here
	assert.Equal(combineArgs.One, s.Env.Owner.Opts.From)
}

func TestSplitCombineSuite(t *test.T) {
	suite.Run(t, &splitCombineSuite{})
}
