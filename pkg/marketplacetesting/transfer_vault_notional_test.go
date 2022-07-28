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

type vaultTransferSuite struct {
	suite.Suite
	Env           *Env
	Dep           *Dep
	Erc20         *mocks.Erc20Session
	CompoundToken *mocks.CompoundTokenSession
	Creator       *mocks.CreatorSession
	ZcToken       *mocks.ZcTokenSession
	VaultTracker  *mocks.VaultTrackerSession
	MarketPlace   *marketplace.MarketPlaceSession // *Session objects are created by the go bindings
}

func (s *vaultTransferSuite) SetupTest() {
	var err error
	assert := assertions.New(s.T())

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)
	assert.Nil(err)

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

	s.VaultTracker = &mocks.VaultTrackerSession{
		Contract: s.Dep.VaultTracker,
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

	// the swivel address must be set
	_, err = s.MarketPlace.SetSwivel(s.Dep.SwivelAddress)
	assert.Nil(err)
	s.Env.Blockchain.Commit()
}

func (s *vaultTransferSuite) TestTransferFailsWhenPaused() {
	assert := assertions.New(s.T())

	underlying := s.Dep.Erc20Address
	maturity := big.NewInt(123456789)

	tx, err := s.MarketPlace.Pause(uint8(1), true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	amount := big.NewInt(100)
	tx, err = s.MarketPlace.TransferVaultNotional(uint8(1), underlying, maturity, s.Env.User1.Opts.From, amount)
	assert.Nil(tx)
	assert.NotNil(err)
	// assert.Regexp("markets are paused", err.Error())

	// unpause so the other tests don't fail
	tx, err = s.MarketPlace.Pause(uint8(1), false)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()
}

func (s *vaultTransferSuite) TestVaultTransfer() {
	assert := assertions.New(s.T())

	ownerOpts := s.Env.Owner.Opts
	user1Opts := s.Env.User1.Opts

	// make a market so that a vault is deployed
	underlying := s.Dep.Erc20Address
	maturity := big.NewInt(1234567)
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

	// find that vault, wrap it in a mock...
	market, err := s.MarketPlace.Markets(uint8(1), underlying, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, cToken)

	// stub the mock vaulttracker...
	tx, err = s.VaultTracker.TransferNotionalFromReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	amount := big.NewInt(100)

	tx, err = s.MarketPlace.TransferVaultNotional(uint8(1), underlying, maturity, user1Opts.From, amount)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// marketplace should have called transfer with the msg.sender as owner
	transferArgs, err := s.VaultTracker.TransferNotionalFromCalled(ownerOpts.From)
	assert.Nil(err)
	assert.Equal(user1Opts.From, transferArgs.To)
	assert.Equal(amount, transferArgs.Amount)
}

func TestVaultTransferSuite(t *test.T) {
	suite.Run(t, &vaultTransferSuite{})
}
