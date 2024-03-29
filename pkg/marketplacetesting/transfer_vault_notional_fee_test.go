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

type vaultTransferFeeSuite struct {
	suite.Suite
	Env         *Env
	Dep         *Dep
	Erc20       *mocks.Erc20Session
	CErc20      *mocks.CErc20Session
	MarketPlace *marketplace.MarketPlaceSession
}

func (s *vaultTransferFeeSuite) SetupTest() {
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

	// the swivel address must be set, using owner here for ease of mocking
	_, err = s.MarketPlace.SetSwivelAddress(s.Env.Owner.Opts.From)
	assert.Nil(err)
	s.Env.Blockchain.Commit()
}

func (s *vaultTransferFeeSuite) TestVaultTransferFee() {
	assert := assertions.New(s.T())

	ownerOpts := s.Env.Owner.Opts
	user1Opts := s.Env.User1.Opts

	// make a market so that a vault is deployed
	underlying := s.Dep.Erc20Address
	maturity := big.NewInt(1234567)
	ctoken := s.Dep.CErc20Address

	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.CErc20.UnderlyingReturns(underlying)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.CreateMarket(
		maturity,
		ctoken,
		"awesome market",
		"AM",
	)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// find that vault, wrap it in a mock...
	market, err := s.MarketPlace.Markets(underlying, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, ctoken)

	vaultTrackerContract, err := mocks.NewVaultTracker(market.VaultAddr, s.Env.Blockchain)
	vaultTracker := &mocks.VaultTrackerSession{
		Contract: vaultTrackerContract,
		CallOpts: bind.CallOpts{From: ownerOpts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   ownerOpts.From,
			Signer: ownerOpts.Signer,
		},
	}

	// stub the mock vaulttracker...
	tx, err = vaultTracker.TransferNotionalFeeReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	amount := big.NewInt(100)

	tx, err = s.MarketPlace.TransferVaultNotionalFee(underlying, maturity, user1Opts.From, amount)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// marketplace should simply have passed the args thru...
	transferFee, err := vaultTracker.TransferNotionalFeeCalled(user1Opts.From)
	assert.Nil(err)
	assert.Equal(big.NewInt(100), transferFee)
}

func TestVaultTransferFeeSuite(t *test.T) {
	suite.Run(t, &vaultTransferFeeSuite{})
}
