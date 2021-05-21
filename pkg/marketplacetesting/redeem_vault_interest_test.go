package marketplacetesting

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	assertions "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/marketplace"
	"github.com/swivel-finance/gost/test/mocks"
	"math/big"
	test "testing"
)

type redeemVaultInterestTest struct {
	suite.Suite
	Env         *Env
	Dep         *Dep
	Erc20Addr   common.Address
	Erc20       *mocks.Erc20Session
	CErc20      *mocks.CErc20Session
	MarketPlace *marketplace.MarketPlaceSession // *Session objects are created by the go bindings
}

func (s *redeemVaultInterestTest) SetupTest() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	err = s.Env.Blockchain.AdjustTime(0) // set bc timestamp to 0
	if err != nil {
		panic(err)
	}
	s.Env.Blockchain.Commit()

	ercAddress, _, ercContract, ercErr := mocks.DeployErc20(s.Env.Owner.Opts, s.Env.Blockchain)

	if ercErr != nil {
		panic(err)
	}

	s.Env.Blockchain.Commit()

	s.Erc20Addr = ercAddress
	s.Erc20 = &mocks.Erc20Session{
		Contract: ercContract,
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

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.MarketPlace = &marketplace.MarketPlaceSession{
		Contract: s.Dep.MarketPlace,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *redeemVaultInterestTest) TestRedeemVaultInterest() {
	assert := assertions.New(s.T())
	maturity := s.Dep.Maturity
	ctokenAddr := s.Dep.CErc20Address

	tx, err := s.MarketPlace.CreateMarket(
		s.Erc20Addr,
		maturity,
		ctokenAddr,
		"awesome market",
		"AM",
	)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(s.Erc20Addr, maturity)
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

	tx, err = s.CErc20.RedeemUnderlyingReturns(ZERO)
	assert.NotNil(tx)
	assert.Nil(err)

	s.Env.Blockchain.Commit()

	tx, err = s.Erc20.TransferReturns(true)
	assert.NotNil(tx)
	assert.Nil(err)

	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.RedeemVaultInterest(s.Erc20Addr, maturity)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	redeemUnderlying, err := s.CErc20.RedeemUnderlyingCalled()
	assert.Nil(err)
	assert.Equal(vaultInterest, redeemUnderlying)

	s.Env.Blockchain.Commit()

	transfer, err := s.Erc20.TransferCalled(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.Equal(vaultInterest, transfer)

	s.Env.Blockchain.Commit()
}

func (s *redeemVaultInterestTest) TestRedeemVaultInterestRedeemUnderlyingFails() {
	assert := assertions.New(s.T())
	maturity := s.Dep.Maturity
	ctokenAddr := s.Dep.CErc20Address

	tx, err := s.MarketPlace.CreateMarket(
		s.Erc20Addr,
		maturity,
		ctokenAddr,
		"awesome market",
		"AM",
	)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(s.Erc20Addr, maturity)
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

	tx, err = s.CErc20.RedeemUnderlyingReturns(big.NewInt(1))
	assert.NotNil(tx)
	assert.Nil(err)

	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.RedeemVaultInterest(s.Erc20Addr, maturity)
	assert.NotNil(err)
	assert.Regexp("redemption from Compound Failed", err.Error())
	assert.Nil(tx)
}

func (s *redeemVaultInterestTest) TestRedeemVaultInterestTransferFails() {
	assert := assertions.New(s.T())
	maturity := s.Dep.Maturity
	ctokenAddr := s.Dep.CErc20Address

	tx, err := s.MarketPlace.CreateMarket(
		s.Erc20Addr,
		maturity,
		ctokenAddr,
		"awesome market",
		"AM",
	)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(s.Erc20Addr, maturity)
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

	tx, err = s.CErc20.RedeemUnderlyingReturns(ZERO)
	assert.NotNil(tx)
	assert.Nil(err)

	s.Env.Blockchain.Commit()

	tx, err = s.Erc20.TransferReturns(false)
	assert.NotNil(tx)
	assert.Nil(err)

	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.RedeemVaultInterest(s.Erc20Addr, maturity)
	assert.NotNil(err)
	assert.Regexp("transfer of redeemable failed", err.Error())
	assert.Nil(tx)
}

func TestRedeemVaultInterestTest(t *test.T) {
	suite.Run(t, &redeemVaultInterestTest{})
}
