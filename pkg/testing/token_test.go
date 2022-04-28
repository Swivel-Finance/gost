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

type tokenTestSuite struct {
	suite.Suite
	Env   *Env
	Dep   *Dep
	Erc20 *mocks.Erc20Session // *Session objects are created by the go bindings
}

func (s *tokenTestSuite) SetupSuite() {
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
}

func (s *tokenTestSuite) TestName() {
	assert := assert.New(s.T())
	tx, err := s.Erc20.NameReturns("awesomeToken")
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	stored, err := s.Erc20.Name()
	assert.Nil(err)
	assert.Equal("awesomeToken", stored)
}

func (s *tokenTestSuite) TestSymbol() {
	assert := assert.New(s.T())
	tx, err := s.Erc20.SymbolReturns("AT")
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	stored, err := s.Erc20.Symbol()
	assert.Nil(err)
	assert.Equal("AT", stored)
}

func (s *tokenTestSuite) TestDecimals() {
	assert := assert.New(s.T())
	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	stored, err := s.Erc20.Decimals()
	assert.Nil(err)
	assert.Equal(uint8(18), stored)
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
	stored, err := s.Erc20.ApproveCalled(address)
	assert.Nil(err)
	assert.Equal(amount, stored)
}

func (s *tokenTestSuite) TestTransfer() {
	assert := assert.New(s.T())
	tx, err := s.Erc20.TransferReturns(true)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	amount := big.NewInt(ONE_ETH)
	// fake transfer to user2
	tx, err = s.Erc20.Transfer(
		s.Env.User2.Opts.From,
		amount,
	)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	// mapping uses the passed address as key
	stored, err := s.Erc20.TransferCalled(s.Env.User2.Opts.From)
	assert.Nil(err)
	assert.Equal(stored, amount)
}

func (s *tokenTestSuite) TestTransferFrom() {
	assert := assert.New(s.T())
	tx, err := s.Erc20.TransferFromReturns(true)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	amount := big.NewInt(ONE_ETH)
	// fake transfer from user1 to owner
	tx, err = s.Erc20.TransferFrom(
		s.Env.User1.Opts.From,
		s.Env.Owner.Opts.From,
		amount,
	)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	// mapping uses the from address as key
	stored, err := s.Erc20.TransferFromCalled(s.Env.User1.Opts.From)
	assert.Nil(err)
	assert.Equal(stored.To, s.Env.Owner.Opts.From)
	assert.Equal(stored.Amount, amount)
}

func TestTokenSuite(t *test.T) {
	suite.Run(t, &tokenTestSuite{})
}
