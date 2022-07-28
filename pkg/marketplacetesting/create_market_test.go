package marketplacetesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/marketplace"
	"github.com/swivel-finance/gost/test/mocks"
)

type createMarketSuite struct {
	suite.Suite
	Env           *Env
	Dep           *Dep
	Erc20         *mocks.Erc20Session
	CompoundToken *mocks.CompoundTokenSession
	Creator       *mocks.CreatorSession
	MarketPlace   *marketplace.MarketPlaceSession // *Session objects are created by the go bindings
}

func (s *createMarketSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// we need to mock the decimals() and underlying() calls...
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

	s.Creator = &mocks.CreatorSession{
		Contract: s.Dep.Creator,
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

	// the swivel address must be set
	s.MarketPlace.SetSwivel(s.Dep.SwivelAddress)
	s.Env.Blockchain.Commit()
}

func (s *createMarketSuite) TestCreateFailsWhenPaused() {
	assert := assert.New(s.T())

	maturity := big.NewInt(123456789)
	cToken := s.Dep.CompoundTokenAddress

	tx, err := s.MarketPlace.Pause(uint8(1), true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// we'd need to stub the underlying() and decimals() call but this will fail before that...
	// also no need to stub creator here, paused() occurs first...

	tx, err = s.MarketPlace.CreateMarket(
		uint8(1),
		maturity,
		cToken,
		"nope",
		"NM",
	)

	assert.Nil(tx)
	assert.NotNil(err)
	// TODO these are now simply "execution reverted", can we get the args?
	// assert.Regexp("markets are paused", err.Error())

	// unpause so the other tests don't fail
	tx, err = s.MarketPlace.Pause(uint8(1), false)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()
}

func (s *createMarketSuite) TestCreateMarket18Decimals() {
	assert := assert.New(s.T())

	underlying := s.Dep.Erc20Address
	cToken := s.Dep.CompoundTokenAddress

	// stub underlying (erc20) methods...
	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.CompoundToken.UnderlyingReturns(underlying)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// stub the return for Creator.create so the market is created with zctoken and vaulttracker addresses
	tx, err = s.Creator.CreateReturns(s.Dep.ZcTokenAddress, s.Dep.VaultTrackerAddress)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	maturity := big.NewInt(123456789)

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
	assert.Equal(market.ZcToken, s.Dep.ZcTokenAddress)
	assert.Equal(market.VaultTracker, s.Dep.VaultTrackerAddress)

	// creator should recieve expected arguments
	createArgs, _ := s.Creator.CreateCalled(uint8(1))
	assert.Equal(createArgs.Underlying, underlying)
	assert.Equal(createArgs.Maturity, maturity)
	assert.Equal(createArgs.CToken, cToken)
	assert.Equal(createArgs.Swivel, s.Dep.SwivelAddress)
	assert.Equal(createArgs.Name, "awesome market")
	assert.Equal(createArgs.Symbol, "AM")
	assert.Equal(createArgs.Decimals, uint8(18)) // gotten from underlying
}

func (s *createMarketSuite) TestCreateMarket6Decimals() {
	assert := assert.New(s.T())

	underlying := s.Dep.Erc20Address
	cToken := s.Dep.CompoundTokenAddress

	// stub underlying (erc20) methods...
	tx, err := s.Erc20.DecimalsReturns(uint8(6))
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

	maturity := big.NewInt(123456781)

	tx, err = s.MarketPlace.CreateMarket(
		uint8(1),
		maturity,
		cToken,
		"awesomer market",
		"ARM",
	)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(uint8(1), underlying, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, cToken)
	assert.Equal(market.ZcToken, s.Dep.ZcTokenAddress)
	assert.Equal(market.VaultTracker, s.Dep.VaultTrackerAddress)

	// creator should recieve expected arguments
	createArgs, _ := s.Creator.CreateCalled(uint8(1))
	assert.Equal(createArgs.Underlying, underlying)
	assert.Equal(createArgs.Maturity, maturity)
	assert.Equal(createArgs.CToken, cToken)
	assert.Equal(createArgs.Swivel, s.Dep.SwivelAddress)
	assert.Equal(createArgs.Name, "awesomer market")
	assert.Equal(createArgs.Symbol, "ARM")
	assert.Equal(createArgs.Decimals, uint8(6)) // gotten from underlying
}

func TestCreateMarketSuite(t *test.T) {
	suite.Run(t, &createMarketSuite{})
}
