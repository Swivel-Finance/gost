package marketplacetesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	assertions "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/marketplace"
	"github.com/swivel-finance/gost/test/mocks"
)

type custodialExitSuite struct {
	suite.Suite
	Env           *Env
	Dep           *Dep
	Erc20         *mocks.Erc20Session
	CompoundToken *mocks.CompoundTokenSession
	EulerToken    *mocks.EulerTokenSession
	Creator       *mocks.CreatorSession
	ZcToken       *mocks.ZcTokenSession
	VaultTracker  *mocks.VaultTrackerSession
	MarketPlace   *marketplace.MarketPlaceSession // *Session objects are created by the go bindings
}

func (s *custodialExitSuite) SetupTest() {
	var err error
	assert := assertions.New(s.T())

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)
	assert.Nil(err)

	err = s.Env.Blockchain.AdjustTime(0) // set bc timestamp to 0
	assert.Nil(err)

	s.Env.Blockchain.Commit()

	s.Erc20 = &mocks.Erc20Session{
		Contract: s.Dep.Erc20,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.CompoundToken = &mocks.CompoundTokenSession{
		Contract: s.Dep.CompoundToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.EulerToken = &mocks.EulerTokenSession{
		Contract: s.Dep.EulerToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	// a creator mock to get VaultTracker and ZcToken deploys from
	s.Creator = &mocks.CreatorSession{
		Contract: s.Dep.Creator,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.ZcToken = &mocks.ZcTokenSession{
		Contract: s.Dep.ZcToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.VaultTracker = &mocks.VaultTrackerSession{
		Contract: s.Dep.VaultTracker,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.MarketPlace = &marketplace.MarketPlaceSession{
		Contract: s.Dep.MarketPlace,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	// the Marketplace needs to have a swivel addr set, we'll use the owner addr so as not to have to generate a new signer
	// and re-do all the calls...
	s.MarketPlace.SetSwivel(s.Env.Owner.Opts.From)
	s.Env.Blockchain.Commit()
}

func (s *custodialExitSuite) TestExitFailsWhenPaused() {
	assert := assertions.New(s.T())

	maturity := big.NewInt(123456789)

	tx, err := s.MarketPlace.Pause(uint8(1), true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	amount := big.NewInt(100)
	tx, err = s.MarketPlace.CustodialExit(uint8(1), s.Dep.Erc20Address, maturity, s.Env.Owner.Opts.From, s.Env.User1.Opts.From, amount)
	assert.Nil(tx)
	assert.NotNil(err)
	// TODO can we fetch the code out of the custom error reverts?
	// assert.Regexp("markets are paused", err.Error())

	// unpause so the other tests don't fail
	tx, err = s.MarketPlace.Pause(uint8(1), false)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()
}

func (s *custodialExitSuite) TestCustodialExit() {
	assert := assertions.New(s.T())

	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.EulerToken.UnderlyingAssetReturns(s.Dep.Erc20Address)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// return the deployed mocks
	tx, err = s.Creator.CreateReturns(s.Dep.ZcTokenAddress, s.Dep.VaultTrackerAddress)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	maturity := s.Dep.Maturity
	cToken := s.Dep.EulerTokenAddress

	tx, err = s.MarketPlace.CreateMarket(
		uint8(5),
		maturity,
		cToken,
		"awesome market",
		"AM",
	)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	ownerOpts := s.Env.Owner.Opts
	user1Opts := s.Env.User1.Opts

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(uint8(5), s.Dep.Erc20Address, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, cToken)

	// creator should recieve expected arguments
	createArgs, _ := s.Creator.CreateCalled(uint8(5))
	assert.Equal(createArgs.Underlying, s.Dep.Erc20Address)
	assert.Equal(createArgs.Maturity, maturity)
	// no need to test all of them, that is done elsewhere

	tx, err = s.ZcToken.BurnReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.VaultTracker.RemoveNotionalReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	amount := big.NewInt(100)
	tx, err = s.MarketPlace.CustodialExit(uint8(5), s.Dep.Erc20Address, maturity, ownerOpts.From, user1Opts.From, amount)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	mintAmount, err := s.ZcToken.BurnCalled(ownerOpts.From)
	assert.Nil(err)
	assert.Equal(amount, mintAmount)

	s.Env.Blockchain.Commit()

	addNotionalAmount, err := s.VaultTracker.RemoveNotionalCalled(user1Opts.From)
	assert.Nil(err)
	assert.Equal(amount, addNotionalAmount)
}

func (s *custodialExitSuite) TestCustodialInitiateBurnFails() {
	assert := assertions.New(s.T())

	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.CompoundToken.UnderlyingReturns(s.Dep.Erc20Address)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	maturity := s.Dep.Maturity
	cToken := s.Dep.CompoundTokenAddress

	tx, err = s.MarketPlace.CreateMarket(
		uint8(1),
		maturity,
		cToken,
		"awesome market",
		"AM",
	)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	ownerOpts := s.Env.Owner.Opts
	user1Opts := s.Env.User1.Opts

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(uint8(1), s.Dep.Erc20Address, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, cToken)

	tx, err = s.ZcToken.BurnReturns(false)
	assert.Nil(err)
	assert.NotNil(tx)

	tx, err = s.VaultTracker.AddNotionalReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	amount := big.NewInt(100)
	tx, err = s.MarketPlace.CustodialExit(uint8(1), s.Dep.Erc20Address, maturity, ownerOpts.From, user1Opts.From, amount)
	assert.NotNil(err)
	// TODO extract the custom error codes?
	// assert.Regexp("burn failed", err.Error())
	assert.Nil(tx)
	s.Env.Blockchain.Commit()
}

func (s *custodialExitSuite) TestCustodialInitiateRemoveNotionalFails() {
	assert := assertions.New(s.T())

	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.CompoundToken.UnderlyingReturns(s.Dep.Erc20Address)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	maturity := s.Dep.Maturity
	cToken := s.Dep.CompoundTokenAddress

	tx, err = s.MarketPlace.CreateMarket(
		uint8(1),
		maturity,
		cToken,
		"awesome market",
		"AM",
	)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	ownerOpts := s.Env.Owner.Opts
	user1Opts := s.Env.User1.Opts

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(uint8(1), s.Dep.Erc20Address, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, cToken)

	tx, err = s.ZcToken.BurnReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)

	tx, err = s.VaultTracker.RemoveNotionalReturns(false)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	amount := big.NewInt(100)
	tx, err = s.MarketPlace.CustodialExit(uint8(1), s.Dep.Erc20Address, maturity, ownerOpts.From, user1Opts.From, amount)
	assert.NotNil(err)
	// TODO extract the custom error codes?
	// assert.Regexp("remove notional failed", err.Error())
	assert.Nil(tx)
	s.Env.Blockchain.Commit()
}

func TestCustodialExitSuite(t *test.T) {
	suite.Run(t, &custodialExitSuite{})
}
