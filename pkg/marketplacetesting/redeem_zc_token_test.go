package marketplacetesting

import (
	"math/big"
	test "testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	assertions "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/marketplace"
	"github.com/swivel-finance/gost/test/mocks"
)

type redeemZcTokenSuite struct {
	suite.Suite
	Env           *Env
	Dep           *Dep
	Erc20         *mocks.Erc20Session
	Erc4626       *mocks.Erc4626Session
	CompoundToken *mocks.CompoundTokenSession
	Creator       *mocks.CreatorSession
	ZcToken       *mocks.ZcTokenSession
	VaultTracker  *mocks.VaultTrackerSession
	MarketPlace   *marketplace.MarketPlaceSession // *Session objects are created by the go bindings
}

func (s *redeemZcTokenSuite) SetupTest() {
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

	s.Erc4626 = &mocks.Erc4626Session{
		Contract: s.Dep.Erc4626,
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

	// the swivel address must be set, use owner to accomodate onlySwivel calls...
	_, err = s.MarketPlace.SetSwivel(s.Env.Owner.Opts.From)
	assert.Nil(err)
	s.Env.Blockchain.Commit()
}

func (s *redeemZcTokenSuite) TestRedeemFailsWhenPaused() {
	assert := assertions.New(s.T())

	tx, err := s.MarketPlace.Pause(uint8(1), true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	underlying := s.Dep.Erc20Address
	maturity := big.NewInt(123456789)
	amount := big.NewInt(123456789)

	tx, err = s.MarketPlace.RedeemZcToken(uint8(1), underlying, maturity, s.Env.Owner.Opts.From, amount)

	assert.Nil(tx)
	assert.NotNil(err)
	// assert.Regexp("markets are paused", err.Error())

	// unpause so the other tests don't fail
	tx, err = s.MarketPlace.Pause(uint8(1), false)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()
}

func (s *redeemZcTokenSuite) TestRedeemZcTokenMaturedRequirementFails() {
	assert := assertions.New(s.T())
	underlying := s.Dep.Erc20Address
	maturity := s.Dep.Maturity
	cToken := s.Dep.CompoundTokenAddress

	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.CompoundToken.UnderlyingReturns(underlying)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.Creator.CreateReturns(s.Dep.ZcTokenAddress, s.Dep.VaultTrackerAddress)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

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

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(uint8(1), underlying, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, cToken)

	tx, err = s.VaultTracker.MatureVaultReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	amount := big.NewInt(123456789)
	tx, err = s.MarketPlace.RedeemZcToken(uint8(1), underlying, maturity, s.Env.Owner.Opts.From, amount)
	assert.NotNil(err)
	// TODO extract the custom error codes?
	// assert.Regexp("maturity not reached", err.Error())
	assert.Nil(tx)

	s.Env.Blockchain.Commit()
}

func (s *redeemZcTokenSuite) TestRedeemZcTokenNotMatured() {
	assert := assertions.New(s.T())
	underlying := s.Dep.Erc20Address
	maturity := s.Dep.Maturity
	cToken := s.Dep.CompoundTokenAddress

	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.CompoundToken.UnderlyingReturns(underlying)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.Creator.CreateReturns(s.Dep.ZcTokenAddress, s.Dep.VaultTrackerAddress)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

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

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(uint8(1), underlying, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, cToken)

	tx, err = s.ZcToken.BurnReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.VaultTracker.MatureVaultReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// move past the maturity
	err = s.Env.Blockchain.AdjustTime(MATURITY * time.Second)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	amount := big.NewInt(123456789)
	tx, err = s.MarketPlace.RedeemZcToken(uint8(1), underlying, maturity, s.Env.Owner.Opts.From, amount)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	burnAmount, err := s.ZcToken.BurnCalled(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.Equal(amount, burnAmount)
}

func (s *redeemZcTokenSuite) TestRedeemZcTokenNotMaturedBurnFails() {
	assert := assertions.New(s.T())
	underlying := s.Dep.Erc20Address
	maturity := s.Dep.Maturity
	cToken := s.Dep.CompoundTokenAddress

	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.CompoundToken.UnderlyingReturns(underlying)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.Creator.CreateReturns(s.Dep.ZcTokenAddress, s.Dep.VaultTrackerAddress)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

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

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(uint8(1), underlying, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, cToken)

	tx, err = s.ZcToken.BurnReturns(false)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.VaultTracker.MatureVaultReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// move past the maturity
	err = s.Env.Blockchain.AdjustTime(MATURITY * time.Second)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	amount := big.NewInt(123456789)
	tx, err = s.MarketPlace.RedeemZcToken(uint8(1), underlying, maturity, s.Env.Owner.Opts.From, amount)
	assert.NotNil(err)
	// TODO extract the custom error codes?
	// assert.Regexp("could not burn", err.Error())
	assert.Nil(tx)

	s.Env.Blockchain.Commit()
}

func (s *redeemZcTokenSuite) TestRedeemZcTokenMatured() {
	assert := assertions.New(s.T())
	underlying := s.Dep.Erc20Address
	maturity := s.Dep.Maturity
	cToken := s.Dep.Erc4626Address

	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.Erc4626.AssetReturns(underlying)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.Creator.CreateReturns(s.Dep.ZcTokenAddress, s.Dep.VaultTrackerAddress)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.CreateMarket(
		uint8(0), // Erc4626
		maturity,
		cToken,
		"awesome market",
		"AM",
	)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(uint8(0), underlying, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, cToken)

	tx, err = s.ZcToken.BurnReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.VaultTracker.MatureVaultReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// move past the maturity
	err = s.Env.Blockchain.AdjustTime(MATURITY * time.Second)
	assert.Nil(err)

	s.Env.Blockchain.Commit()

	rate := big.NewInt(223456789)
	tx, err = s.Erc4626.ConvertToAssetsReturns(rate)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.MatureMarket(uint8(0), underlying, maturity)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	amount := big.NewInt(123456789)
	tx, err = s.MarketPlace.RedeemZcToken(uint8(0), underlying, maturity, s.Env.Owner.Opts.From, amount)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	burnAmount, err := s.ZcToken.BurnCalled(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.Equal(amount, burnAmount)
}

func (s *redeemZcTokenSuite) TestRedeemZcTokenMaturedBurnFails() {
	assert := assertions.New(s.T())
	underlying := s.Dep.Erc20Address
	maturity := s.Dep.Maturity
	cToken := s.Dep.CompoundTokenAddress

	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.CompoundToken.UnderlyingReturns(underlying)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.Creator.CreateReturns(s.Dep.ZcTokenAddress, s.Dep.VaultTrackerAddress)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

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

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(uint8(1), underlying, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, cToken)

	tx, err = s.ZcToken.BurnReturns(false)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.VaultTracker.MatureVaultReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// move past the maturity
	err = s.Env.Blockchain.AdjustTime(MATURITY * time.Second)
	assert.Nil(err)

	s.Env.Blockchain.Commit()

	rate := big.NewInt(223456789)
	tx, err = s.CompoundToken.ExchangeRateCurrentReturns(rate)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.MatureMarket(uint8(1), underlying, maturity)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	amount := big.NewInt(123456789)
	tx, err = s.MarketPlace.RedeemZcToken(uint8(1), underlying, maturity, s.Env.Owner.Opts.From, amount)
	assert.NotNil(err)
	// TODO extract the custom error codes?
	// assert.Regexp("could not burn", err.Error())
	assert.Nil(tx)
}

func TestRedeemZcTokenSuite(t *test.T) {
	suite.Run(t, &redeemZcTokenSuite{})
}
