package testing

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/mocks"
)

type yieldTestSuite struct {
	suite.Suite
	Env   *Env
	Dep   *Dep
	Yield *mocks.YieldSession
}

func (s *yieldTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.Yield = &mocks.YieldSession{
		Contract: s.Dep.Yield,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *yieldTestSuite) TestSellBase() {
	assert := assert.New(s.T())
	tx, err := s.Yield.SellBaseReturns(big.NewInt(1000))
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	amount := big.NewInt(ONE_ETH)
	// fake user1 sell base of ONE_ETH
	tx, err = s.Yield.SellBase(
		s.Env.User1.Opts.From,
		amount,
	)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	amountSold, err := s.Yield.SellBaseCalled(s.Env.User1.Opts.From)
	assert.Nil(err)
	assert.Equal(amount, amountSold)
}

func (s *yieldTestSuite) TestSellBasePreview() {
	assert := assert.New(s.T())
	amount := big.NewInt(ONE_ETH)
	tx, err := s.Yield.SellBasePreviewReturns(amount)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	tx, err = s.Yield.SellBasePreview(amount)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	preview, err := s.Yield.SellBasePreviewCalled()
	assert.Nil(err)
	assert.Equal(amount, preview)

}

func TestYieldTokenSuite(t *test.T) {
	suite.Run(t, &yieldTestSuite{})
}
