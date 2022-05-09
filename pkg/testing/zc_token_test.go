package testing

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/mocks"
)

type zcTokenTestSuite struct {
	suite.Suite
	Env     *Env
	Dep     *Dep
	ZcToken *mocks.ZcTokenSession
}

func (s *zcTokenTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.ZcToken = &mocks.ZcTokenSession{
		Contract: s.Dep.ZcToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *zcTokenTestSuite) TestMint() {
	assert := assert.New(s.T())
	tx, err := s.ZcToken.MintReturns(true)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	amount := big.NewInt(ONE_ETH)
	// fake user1 mint ONE_ETH
	tx, err = s.ZcToken.Mint(
		s.Env.User1.Opts.From,
		amount,
	)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	amountMinted, err := s.ZcToken.BalanceOf(s.Env.User1.Opts.From)
	assert.Nil(err)
	assert.Equal(amount, amountMinted)
}

func (s *zcTokenTestSuite) TestBalanceOf() {
	assert := assert.New(s.T())
	address := common.BigToAddress(big.NewInt(12432))
	amount := big.NewInt(ONE_ETH)

	tx, err := s.ZcToken.BalanceOfReturns(address, amount)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	// fake user1 mint ONE_ETH
	tx, err = s.ZcToken.Mint(
		s.Env.User1.Opts.From,
		amount,
	)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	balance, err := s.ZcToken.Balances(s.Env.User1.Opts.From)
	s.Env.Blockchain.Commit()

	assert.Nil(err)
	assert.Equal(amount, balance)
}

func TestZcTokenSuite(t *test.T) {
	suite.Run(t, &zcTokenTestSuite{})
}
