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

var ORDERS = []mocks.Order{}
var AMOUNTS = []*big.Int{}
var COMPONENTS = []mocks.Components{}

type swivelTestSuite struct {
	suite.Suite
	Env    *Env
	Dep    *Dep
	Swivel *mocks.SwivelSession
}

func (s *swivelTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.Swivel = &mocks.SwivelSession{
		Contract: s.Dep.Swivel,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	ORDERS = []mocks.Order{
		{
			Maker:      s.Env.User1.Opts.From,
			Underlying: common.BigToAddress(big.NewInt(1)),
			Vault:      true,
			Exit:       false,
			Principal:  big.NewInt(11111),
			Premium:    big.NewInt(22222),
			Maturity:   big.NewInt(33333),
			Expiry:     big.NewInt(44444),
		},
		{
			Maker:      s.Env.User2.Opts.From,
			Underlying: common.BigToAddress(big.NewInt(2)),
			Vault:      false,
			Exit:       true,
			Principal:  big.NewInt(99999),
			Premium:    big.NewInt(88888),
			Maturity:   big.NewInt(77777),
			Expiry:     big.NewInt(66666),
		},
	}

	AMOUNTS = []*big.Int{
		big.NewInt(111),
		big.NewInt(222),
	}

	COMPONENTS = []mocks.Components{
		{
			V: 100,
			R: [32]byte{1, 2, 3},
			S: [32]byte{4, 5, 6},
		},
		{
			V: 200,
			R: [32]byte{7, 8, 9},
			S: [32]byte{10, 11, 12},
		},
	}
}

func (s *swivelTestSuite) TestInitiate() {
	assert := assert.New(s.T())
	tx, err := s.Swivel.Initiate(ORDERS, AMOUNTS, COMPONENTS)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	initiateAmount, err := s.Swivel.InitiateCalledAmount(s.Env.User1.Opts.From)
	assert.Nil(err)
	assert.Equal(AMOUNTS[0], initiateAmount)

	initiateAmount, err = s.Swivel.InitiateCalledAmount(s.Env.User2.Opts.From)
	assert.Nil(err)
	assert.Equal(AMOUNTS[1], initiateAmount)
}

func (s *swivelTestSuite) TestRedeemZcToken() {
	addr := common.BigToAddress(big.NewInt(1))
	amount := big.NewInt(2)

	assert := assert.New(s.T())
	tx, err := s.Swivel.RedeemZcToken(addr, amount, big.NewInt(3))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	redeemAmount, err := s.Swivel.RedeemZcTokenCalledAmount(addr)
	assert.Equal(amount, redeemAmount)
}

func TestSwivelSuite(t *test.T) {
	suite.Run(t, &swivelTestSuite{})
}
