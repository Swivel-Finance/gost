package marketplacetesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	ctoken := s.Dep.CompoundTokenAddress

	tx, err := s.MarketPlace.Pause(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.CreateMarket(
		uint8(1),
		maturity,
		ctoken,
		"nope",
		"NM",
	)

	assert.Nil(tx)
	assert.NotNil(err)
	assert.Regexp("markets are paused", err.Error())

	// unpause so the other tests don't fail
	tx, err = s.MarketPlace.Pause(false)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()
}

func (s *createMarketSuite) TestCreateMarket18Decimals() {
	assert := assert.New(s.T())

	underlying := s.Dep.Erc20Address
	ctoken := s.Dep.CompoundTokenAddress

	// stub underlying (erc20) methods...
	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.CompoundToken.UnderlyingReturns(underlying)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	maturity := big.NewInt(123456789)

	tx, err = s.MarketPlace.CreateMarket(
		uint8(1),
		maturity,
		ctoken,
		"awesome market",
		"AM",
	)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(uint8(1), underlying, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, ctoken)
	assert.NotEqual(market.ZcToken, common.HexToAddress("0x0"))
	assert.NotEqual(market.VaultTracker, common.HexToAddress("0x0"))

	zcTokenContract, err := mocks.NewZcToken(market.ZcToken, s.Env.Blockchain)
	zcToken := &mocks.ZcTokenSession{
		Contract: zcTokenContract,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	decimals, err := zcToken.Decimals()
	assert.Equal(decimals, uint8(18))
}

func (s *createMarketSuite) TestCreateMarket6Decimals() {
	assert := assert.New(s.T())

	underlying := s.Dep.Erc20Address
	ctoken := s.Dep.CompoundTokenAddress
	// stub underlying (erc20) methods...
	tx, err := s.Erc20.DecimalsReturns(uint8(6))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.CompoundToken.UnderlyingReturns(underlying)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	maturity := big.NewInt(123456781)

	tx, err = s.MarketPlace.CreateMarket(
		uint8(1),
		maturity,
		ctoken,
		"awesomer market",
		"ARM",
	)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(uint8(1), underlying, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, ctoken)
	assert.NotEqual(market.ZcToken, common.HexToAddress("0x0"))
	assert.NotEqual(market.VaultTracker, common.HexToAddress("0x0"))

	zcTokenContract, err := mocks.NewZcToken(market.ZcToken, s.Env.Blockchain)
	zcToken := &mocks.ZcTokenSession{
		Contract: zcTokenContract,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	decimals, err := zcToken.Decimals()
	assert.Equal(decimals, uint8(6))
}

func TestCreateMarketSuite(t *test.T) {
	suite.Run(t, &createMarketSuite{})
}
