package testing

import (
	"math/big"
	test "testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/tokens"
)

type zcTokenSuite struct {
	suite.Suite
	Env     *Env
	Dep     *Dep
	Erc20   *mocks.Erc20Session // *Session objects are created by the go bindings
	ZcToken *tokens.ZcTokenSession
}

func (s *zcTokenSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	s.Erc20 = &mocks.Erc20Session{
		Contract: s.Dep.Erc20,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.ZcToken = &tokens.ZcTokenSession{
		Contract: s.Dep.ZcToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *zcTokenSuite) TestName() {
	assert := assert.New(s.T())

	name, err := s.ZcToken.Name()
	assert.Nil(err)
	assert.Equal("YoloToken", name)
}

func (s *zcTokenSuite) TestSymbol() {
	assert := assert.New(s.T())

	sym, err := s.ZcToken.Symbol()
	assert.Nil(err)
	assert.Equal("YT", sym)
}

func (s *zcTokenSuite) TestDecimals() {
	assert := assert.New(s.T())

	dec, err := s.ZcToken.Decimals()
	assert.Nil(err)
	assert.Equal(uint8(18), dec)
}

func (s *zcTokenSuite) TestMintFailsWhenMature() {
	assert := assert.New(s.T())

	// move past the maturity in order to create the fail condition
	err := s.Env.Blockchain.AdjustTime(MATURITY * time.Second)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	tx, err := s.ZcToken.Mint(s.Env.User1.Opts.From, big.NewInt(1000000))
	// should have an err...
	assert.Nil(tx)
	assert.NotNil(err)

	// error should be that the maturity has been exceeded
	assert.Regexp("maturity reached", err.Error())
}

func TestZcTokenSuite(t *test.T) {
	suite.Run(t, &zcTokenSuite{})
}
