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

type authRedeemSuite struct {
	suite.Suite
	Env           *Env
	Dep           *Dep
	Erc20         *mocks.Erc20Session
	Erc4626       *mocks.Erc4626Session
	CompoundToken *mocks.CompoundTokenSession
	Creator       *mocks.CreatorSession
	ZcToken       *mocks.ZcTokenSession
	VaultTracker  *mocks.VaultTrackerSession
	Swivel        *mocks.SwivelSession
	MarketPlace   *marketplace.MarketPlaceSession // *Session objects are created by the go bindings
}

func (s *authRedeemSuite) SetupTest() {
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

	s.Swivel = &mocks.SwivelSession{
		Contract: s.Dep.Swivel,
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

	// authRedeem requires that the swivel address be the actual mock
	_, err = s.MarketPlace.SetSwivel(s.Dep.SwivelAddress)
	assert.Nil(err)
	s.Env.Blockchain.Commit()
}

func (s *authRedeemSuite) TestAuthRedeemFailsOnPaused() {
	assert := assertions.New(s.T())

	underlying := s.Dep.Erc20Address
	maturity := s.Dep.Maturity
	amount := big.NewInt(123456789)
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

	// a market must be created in order for the zcToken authorized check to succeed
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

	// stub the returns for zctoken and vaulttracker
	tx, err = s.ZcToken.BurnReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.VaultTracker.MatureVaultReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// move past the maturity date
	err = s.Env.Blockchain.AdjustTime(MATURITY * time.Second)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	// force the paued modifier to fail here
	tx, err = s.MarketPlace.Pause(uint8(1), true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.ZcToken.AuthRedeem(uint8(1), underlying, maturity, s.Env.Owner.Opts.From, s.Env.User1.Opts.From, amount)
	assert.Nil(tx)
	assert.NotNil(err)
	s.Env.Blockchain.Commit()

	// unpause so the other tests don't fail
	tx, err = s.MarketPlace.Pause(uint8(1), false)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()
}

func (s *authRedeemSuite) TestAuthRedeem() {
	assert := assertions.New(s.T())

	underlying := s.Dep.Erc20Address
	maturity := s.Dep.Maturity
	amount := big.NewInt(123456789)
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

	// a market must be created in order for the zcToken authorized check to succeed
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

	// stub the returns for zctoken and vaulttracker
	tx, err = s.ZcToken.BurnReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.VaultTracker.MatureVaultReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// move past the maturity date
	err = s.Env.Blockchain.AdjustTime(MATURITY * time.Second)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	// we'll use the exposed `authRedeem` method on the ZcToken mock which will call it's redeemer (MarketPlace) thus preserving the sender we need
	// and be passed directly thru to the MarketPlace.AuthRedeem method we are testing
	tx, err = s.ZcToken.AuthRedeem(uint8(1), underlying, maturity, s.Env.Owner.Opts.From, s.Env.User1.Opts.From, amount)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// the swivel mock should have seen the call
	args, err := s.Swivel.AuthRedeemCalled(uint8(1))
	assert.Nil(err)
	assert.NotNil(args)
	assert.Equal(args.Underlying, underlying)
	assert.Equal(args.One, cToken)
	assert.Equal(args.Two, s.Env.User1.Opts.From)
	assert.Equal(args.Amount, amount)
}

func TestAuthRedeemSuite(t *test.T) {
	suite.Run(t, &authRedeemSuite{})
}
