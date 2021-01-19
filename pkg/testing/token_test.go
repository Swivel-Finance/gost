package testing

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/contracts/tokens"
)

type tokenTestSuite struct {
	suite.Suite
	Env   *Env
	Dep   *Dep
	Erc20 *tokens.Erc20Session // *Session objects are created by the go bindings
}

func (s *tokenTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	s.Erc20 = &tokens.Erc20Session{
		Contract: s.Dep.Erc20,
		CallOpts: bind.CallOpts{From: s.Env.User1.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.User1.Opts.From,
			Signer: s.Env.User1.Opts.Signer,
		},
	}
}

func (s *tokenTestSuite) TestApprove() {
	assert := assert.New(s.T())
	// set approve to return true
	tx, err := s.Erc20.ApproveReturns(true)
	assert.NotNil(tx)
	assert.Nil(err)
	// nothing happens witout manually 'mining'...
	s.Env.Blockchain.Commit()
	// do an actual approval
	address := common.HexToAddress("0xaBC")
	amount := big.NewInt(ONE_ETH)
	tx, err = s.Erc20.Approve(address, amount)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()
	// check the args we passed
	stored, err := s.Erc20.ApprovedArgs(address)
	assert.Nil(err)
	assert.Equal(amount, stored)
}

func TestTokenSuite(t *test.T) {
	suite.Run(t, &tokenTestSuite{})
}
