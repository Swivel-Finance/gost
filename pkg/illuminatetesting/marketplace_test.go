package illuminatetesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/build/marketplace"
	"github.com/swivel-finance/gost/test/mocks"
)

type marketplaceTestSuite struct {
	suite.Suite
	Env         *Env
	Dep         *Dep
	Erc20       *mocks.Erc20Session
	Pool        *mocks.PoolSession
	MarketPlace *marketplace.MarketPlaceSession
}

func (s *marketplaceTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.Erc20 = &mocks.Erc20Session{
		Contract: s.Dep.Erc20,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.Pool = &mocks.PoolSession{
		Contract: s.Dep.Pool,
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
}

func (s *marketplaceTestSuite) TestConstruction() {
	assert := assert.New(s.T())

	rAddr, _ := s.MarketPlace.Redeemer()

	assert.Equal(rAddr, s.Dep.RedeemerAddress)
}

func (s *marketplaceTestSuite) TestCreateMarket() {
	assert := assert.New(s.T())

	// stub the erc approve to return true or safe will revert
	s.Erc20.ApproveReturns(true)

	s.Env.Blockchain.Commit()

	maturity := big.NewInt(100000)

	// a fixed len array of 7 as illuminate is set in the method...
	principals := [8]common.Address{
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
		s.Dep.Erc20Address,
	}

	tx, err := s.MarketPlace.CreateMarket(s.Dep.Erc20Address, maturity, principals, "spam token", "SPAM", uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// we should see a market for that market-pair that has the new ZcToken address at [0] (as illuminate)
	// remember that you cannot fetch the entire array, but must provide the index of the principal you want...
	illPrincipal, err := s.MarketPlace.Markets(s.Dep.Erc20Address, maturity, big.NewInt(0))
	assert.Nil(err)
	// this will be the only one NOT the mock erc address...
	assert.NotEqual(illPrincipal, s.Dep.Erc20Address)

	// the first address in the array passed to the function should now be at [1] (as swivel)
	swivPrincipal, err := s.MarketPlace.Markets(s.Dep.Erc20Address, maturity, big.NewInt(1))
	assert.Nil(err)
	assert.Equal(swivPrincipal, s.Dep.Erc20Address)

	// should be the max amount approved in the contract...
	max, exp := big.NewInt(2), big.NewInt(256)
	max.Exp(max, exp, nil)
	max.Sub(max, big.NewInt(1))

	// we can see what was passed to approve on the last iteration of the for loop
	approved, err := s.Erc20.ApproveCalled(s.Dep.RedeemerAddress)
	assert.Nil(err)
	assert.Equal(approved, max)
}

func (s *marketplaceTestSuite) TestBuyPt() {
	assert := assert.New(s.T())

	s.Erc20.ApproveReturns(true)
	s.Env.Blockchain.Commit()

	maturity := big.NewInt(100000)
	amount := big.NewInt(1000000000000000000)
	returnAmount := new(big.Int).Sub(amount, big.NewInt(5))

	s.MarketPlace.SetPool(0, s.Dep.Erc20Address, maturity, s.Dep.PoolAddress)
	s.Env.Blockchain.Commit()

	s.Pool.UnderlyingReturns(s.Dep.Erc20Address)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferReturns(true)
	s.Env.Blockchain.Commit()

	s.Pool.BuyPTReturns(returnAmount)
	s.Env.Blockchain.Commit()

	s.Pool.BuyPTPreviewReturns(returnAmount)
	s.Env.Blockchain.Commit()

	tx, err := s.MarketPlace.BuyPT(0, s.Dep.Erc20Address, maturity, amount)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify methods were called as expected
	buyPtOut, err := s.Pool.BuyPTCalled(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.Equal(returnAmount, buyPtOut.PtOut)
	assert.Equal(amount, buyPtOut.Min)

	previewedAmount, err := s.Pool.BuyPTPreviewCalled()
	assert.Nil(err)
	assert.Equal(amount, previewedAmount)

	transferAmount, err := s.Erc20.TransferCalled(s.Dep.PoolAddress)
	assert.Nil(err)
	assert.Equal(amount, transferAmount)
}

func (s *marketplaceTestSuite) TestSellPt() {
	assert := assert.New(s.T())

	s.Erc20.ApproveReturns(true)
	s.Env.Blockchain.Commit()

	maturity := big.NewInt(100000)
	amount := big.NewInt(1000000000000000000)
	returnAmount := new(big.Int).Sub(amount, big.NewInt(5))

	s.MarketPlace.SetPool(0, s.Dep.Erc20Address, maturity, s.Dep.PoolAddress)
	s.Env.Blockchain.Commit()

	s.Pool.PTReturns(s.Dep.Erc20Address)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferReturns(true)
	s.Env.Blockchain.Commit()

	s.Pool.SellPTReturns(returnAmount)
	s.Env.Blockchain.Commit()

	s.Pool.SellPTPreviewReturns(returnAmount)
	s.Env.Blockchain.Commit()

	tx, err := s.MarketPlace.SellPT(0, s.Dep.Erc20Address, maturity, amount)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify methods were called as expected
	sellPtOut, err := s.Pool.SellPTCalled(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.Equal(returnAmount, sellPtOut)

	previewedAmount, err := s.Pool.SellPTPreviewCalled()
	assert.Nil(err)
	assert.Equal(amount, previewedAmount)

	transferAmount, err := s.Erc20.TransferCalled(s.Dep.PoolAddress)
	assert.Nil(err)
	assert.Equal(amount, transferAmount)
}

func (s *marketplaceTestSuite) TestBuyUnderlying() {
	assert := assert.New(s.T())

	s.Erc20.ApproveReturns(true)
	s.Env.Blockchain.Commit()

	maturity := big.NewInt(100000)
	amount := big.NewInt(1000000000000000000)
	returnAmount := new(big.Int).Sub(amount, big.NewInt(5))

	s.MarketPlace.SetPool(0, s.Dep.Erc20Address, maturity, s.Dep.PoolAddress)
	s.Env.Blockchain.Commit()

	s.Pool.PTReturns(s.Dep.Erc20Address)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferReturns(true)
	s.Env.Blockchain.Commit()

	s.Pool.BuyUnderlyingReturns(returnAmount)
	s.Env.Blockchain.Commit()

	s.Pool.BuyUnderlyingPreviewReturns(returnAmount)
	s.Env.Blockchain.Commit()

	tx, err := s.MarketPlace.BuyUnderlying(0, s.Dep.Erc20Address, maturity, amount)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify methods were called as expected
	buyPtOut, err := s.Pool.BuyUnderlyingCalled(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.Equal(returnAmount, buyPtOut.Amount)
	assert.Equal(amount, buyPtOut.Min)

	previewedAmount, err := s.Pool.BuyUnderlyingPreviewCalled()
	assert.Nil(err)
	assert.Equal(amount, previewedAmount)

	transferAmount, err := s.Erc20.TransferCalled(s.Dep.PoolAddress)
	assert.Nil(err)
	assert.Equal(amount, transferAmount)
}

func (s *marketplaceTestSuite) TestSellUnderlying() {
	assert := assert.New(s.T())

	s.Erc20.ApproveReturns(true)
	s.Env.Blockchain.Commit()

	maturity := big.NewInt(100000)
	amount := big.NewInt(1000000000000000000)
	returnAmount := new(big.Int).Sub(amount, big.NewInt(5))

	s.MarketPlace.SetPool(0, s.Dep.Erc20Address, maturity, s.Dep.PoolAddress)
	s.Env.Blockchain.Commit()

	s.Pool.UnderlyingReturns(s.Dep.Erc20Address)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferReturns(true)
	s.Env.Blockchain.Commit()

	s.Pool.SellUnderlyingReturns(returnAmount)
	s.Env.Blockchain.Commit()

	s.Pool.SellUnderlyingPreviewReturns(returnAmount)
	s.Env.Blockchain.Commit()

	tx, err := s.MarketPlace.SellUnderlying(0, s.Dep.Erc20Address, maturity, amount)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// verify methods were called as expected
	sellPtOut, err := s.Pool.SellUnderlyingCalled(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.Equal(returnAmount, sellPtOut)

	previewedAmount, err := s.Pool.SellUnderlyingPreviewCalled()
	assert.Nil(err)
	assert.Equal(amount, previewedAmount)

	transferAmount, err := s.Erc20.TransferCalled(s.Dep.PoolAddress)
	assert.Nil(err)
	assert.Equal(amount, transferAmount)
}

func TestIlluminateSuite(t *test.T) {
	suite.Run(t, &marketplaceTestSuite{})
}
