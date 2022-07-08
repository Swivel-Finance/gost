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

type custodialInitiateSuite struct {
	suite.Suite
	Env           *Env
	Dep           *Dep
	Erc20         *mocks.Erc20Session
	CompoundToken *mocks.CompoundTokenSession
	YearnVault    *mocks.YearnVaultSession
	MarketPlace   *marketplace.MarketPlaceSession // *Session objects are created by the go bindings
}

func (s *custodialInitiateSuite) SetupTest() {
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

	s.YearnVault = &mocks.YearnVaultSession{
		Contract: s.Dep.YearnVault,
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

	s.MarketPlace.SetSwivel(s.Env.Owner.Opts.From)
	s.Env.Blockchain.Commit()
}

func (s *custodialInitiateSuite) TestInitiateFailsWhenPaused() {
	assert := assertions.New(s.T())

	underlying := s.Dep.Erc20Address
	maturity := big.NewInt(123456789)

	tx, err := s.MarketPlace.Pause(uint8(1), true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	amount := big.NewInt(100)
	tx, err = s.MarketPlace.CustodialInitiate(uint8(1), underlying, maturity, s.Env.Owner.Opts.From, s.Env.User1.Opts.From, amount)

	assert.Nil(tx)
	assert.NotNil(err)
	// assert.Regexp("markets are paused", err.Error())

	// unpause so the other tests don't fail
	tx, err = s.MarketPlace.Pause(uint8(1), false)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()
}

func (s *custodialInitiateSuite) TestCustodialInitiate() {
	assert := assertions.New(s.T())

	underlying := s.Dep.Erc20Address
	maturity := s.Dep.Maturity
	// use yearn protocol for this test
	ctoken := s.Dep.YearnVaultAddress

	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.YearnVault.UnderlyingReturns(underlying)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.CreateMarket(
		uint8(3), // yearn protocol
		maturity,
		ctoken,
		"awesome market",
		"AM",
	)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	ownerOpts := s.Env.Owner.Opts
	user1Opts := s.Env.User1.Opts

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(uint8(3), underlying, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, ctoken)

	zcTokenContract, err := mocks.NewZcToken(market.ZcToken, s.Env.Blockchain)
	zcToken := &mocks.ZcTokenSession{
		Contract: zcTokenContract,
		CallOpts: bind.CallOpts{From: ownerOpts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   ownerOpts.From,
			Signer: ownerOpts.Signer,
		},
	}

	zcMaturity, err := zcToken.Maturity()
	assert.Equal(maturity, zcMaturity)

	tx, err = zcToken.MintReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)

	vaultTrackerContract, err := mocks.NewVaultTracker(market.VaultTracker, s.Env.Blockchain)
	vaultTracker := &mocks.VaultTrackerSession{
		Contract: vaultTrackerContract,
		CallOpts: bind.CallOpts{From: ownerOpts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   ownerOpts.From,
			Signer: ownerOpts.Signer,
		},
	}

	tx, err = vaultTracker.AddNotionalReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	amount := big.NewInt(100)
	tx, err = s.MarketPlace.CustodialInitiate(uint8(3), underlying, maturity, ownerOpts.From, user1Opts.From, amount)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	mintAmount, err := zcToken.MintCalled(ownerOpts.From)
	assert.Nil(err)
	assert.Equal(amount, mintAmount)

	s.Env.Blockchain.Commit()

	addNotionalAmount, err := vaultTracker.AddNotionalCalled(user1Opts.From)
	assert.Nil(err)
	assert.Equal(amount, addNotionalAmount)

	s.Env.Blockchain.Commit()
}

func (s *custodialInitiateSuite) TestCustodialInitiateMintFails() {
	assert := assertions.New(s.T())
	underlying := s.Dep.Erc20Address
	maturity := s.Dep.Maturity
	ctoken := s.Dep.CompoundTokenAddress

	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.CompoundToken.UnderlyingReturns(underlying)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

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

	ownerOpts := s.Env.Owner.Opts
	user1Opts := s.Env.User1.Opts

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(uint8(1), underlying, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, ctoken)

	zcTokenContract, err := mocks.NewZcToken(market.ZcToken, s.Env.Blockchain)
	zcToken := &mocks.ZcTokenSession{
		Contract: zcTokenContract,
		CallOpts: bind.CallOpts{From: ownerOpts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   ownerOpts.From,
			Signer: ownerOpts.Signer,
		},
	}

	zcMaturity, err := zcToken.Maturity()
	assert.Equal(maturity, zcMaturity)

	tx, err = zcToken.MintReturns(false)
	assert.Nil(err)
	assert.NotNil(tx)

	vaultTrackerContract, err := mocks.NewVaultTracker(market.VaultTracker, s.Env.Blockchain)
	vaultTracker := &mocks.VaultTrackerSession{
		Contract: vaultTrackerContract,
		CallOpts: bind.CallOpts{From: ownerOpts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   ownerOpts.From,
			Signer: ownerOpts.Signer,
		},
	}

	tx, err = vaultTracker.AddNotionalReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	amount := big.NewInt(100)
	tx, err = s.MarketPlace.CustodialInitiate(uint8(1), underlying, maturity, ownerOpts.From, user1Opts.From, amount)
	assert.NotNil(err)
	// TODO extract the custom error codes?
	// assert.Regexp("mint failed", err.Error())
	assert.Nil(tx)

	s.Env.Blockchain.Commit()
}

func (s *custodialInitiateSuite) TestCustodialInitiateAddNotionalFails() {
	assert := assertions.New(s.T())
	underlying := s.Dep.Erc20Address
	maturity := s.Dep.Maturity
	ctoken := s.Dep.CompoundTokenAddress

	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.CompoundToken.UnderlyingReturns(underlying)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

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

	ownerOpts := s.Env.Owner.Opts
	user1Opts := s.Env.User1.Opts

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(uint8(1), underlying, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, ctoken)

	zcTokenContract, err := mocks.NewZcToken(market.ZcToken, s.Env.Blockchain)
	zcToken := &mocks.ZcTokenSession{
		Contract: zcTokenContract,
		CallOpts: bind.CallOpts{From: ownerOpts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   ownerOpts.From,
			Signer: ownerOpts.Signer,
		},
	}

	zcMaturity, err := zcToken.Maturity()
	assert.Equal(maturity, zcMaturity)

	tx, err = zcToken.MintReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)

	vaultTrackerContract, err := mocks.NewVaultTracker(market.VaultTracker, s.Env.Blockchain)
	vaultTracker := &mocks.VaultTrackerSession{
		Contract: vaultTrackerContract,
		CallOpts: bind.CallOpts{From: ownerOpts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   ownerOpts.From,
			Signer: ownerOpts.Signer,
		},
	}

	tx, err = vaultTracker.AddNotionalReturns(false)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	amount := big.NewInt(100)
	tx, err = s.MarketPlace.CustodialInitiate(uint8(1), underlying, maturity, ownerOpts.From, user1Opts.From, amount)
	assert.NotNil(err)
	// TODO extract the custom error codes?
	// assert.Regexp("add notional failed", err.Error())
	assert.Nil(tx)

	s.Env.Blockchain.Commit()
}

func TestCustodialInitiateSuite(t *test.T) {
	suite.Run(t, &custodialInitiateSuite{})
}
