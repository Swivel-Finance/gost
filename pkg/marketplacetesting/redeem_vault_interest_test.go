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

type redeemVaultInterestSuite struct {
	suite.Suite
	Env         *Env
	Dep         *Dep
	Erc20       *mocks.Erc20Session
	CErc20      *mocks.CErc20Session
	MarketPlace *marketplace.MarketPlaceSession // *Session objects are created by the go bindings
}

func (s *redeemVaultInterestSuite) SetupTest() {
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

	s.CErc20 = &mocks.CErc20Session{
		Contract: s.Dep.CErc20,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.MarketPlace = &marketplace.MarketPlaceSession{
		Contract: s.Dep.MarketPlace,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	// the swivel address must be set (set to owner accomodating the onlySwivel calls)
	_, err = s.MarketPlace.SetSwivelAddress(s.Env.Owner.Opts.From)
	assert.Nil(err)
	s.Env.Blockchain.Commit()
}

func (s *redeemVaultInterestSuite) TestRedeemFailsWhenPaused() {
	assert := assertions.New(s.T())

	tx, err := s.MarketPlace.Pause(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	maturity := big.NewInt(123456789)

	tx, err = s.MarketPlace.RedeemVaultInterest(s.Dep.Erc20Address, maturity, s.Env.Owner.Opts.From)

	assert.Nil(tx)
	assert.NotNil(err)
	assert.Regexp("markets are paused", err.Error())

	// unpause so the other tests don't fail
	tx, err = s.MarketPlace.Pause(false)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()
}

func (s *redeemVaultInterestSuite) TestRedeemVaultInterest() {
	assert := assertions.New(s.T())
	maturity := s.Dep.Maturity
	ctokenAddr := s.Dep.CErc20Address

	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.CErc20.UnderlyingReturns(s.Dep.Erc20Address)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.CreateMarket(
		maturity,
		ctokenAddr,
		"awesome market",
		"AM",
	)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(s.Dep.Erc20Address, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, ctokenAddr)

	zcTokenContract, err := mocks.NewZcToken(market.ZcTokenAddr, s.Env.Blockchain)
	zcToken := &mocks.ZcTokenSession{
		Contract: zcTokenContract,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	zcMaturity, err := zcToken.Maturity()
	assert.Equal(maturity, zcMaturity)

	vaultTrackerContract, err := mocks.NewVaultTracker(market.VaultAddr, s.Env.Blockchain)
	vaultTracker := &mocks.VaultTrackerSession{
		Contract: vaultTrackerContract,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	vaultInterest := big.NewInt(123456789)
	tx, err = vaultTracker.RedeemInterestReturns(vaultInterest)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.RedeemVaultInterest(s.Dep.Erc20Address, maturity, s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// the vaulttracker mock should now retain the address it was passed
	address, err := vaultTracker.RedeemInterestCalled()
	assert.Nil(err)
	assert.NotNil(address)
	assert.Equal(address, s.Env.Owner.Opts.From)
}

func TestRedeemVaultInterestSuite(t *test.T) {
	suite.Run(t, &redeemVaultInterestSuite{})
}
