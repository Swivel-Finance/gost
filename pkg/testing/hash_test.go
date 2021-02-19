package testing

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/contracts/fakes"
)

type hashTestSuite struct {
	suite.Suite
	Env       *Env
	Dep       *Dep
	Separator [32]byte // keep a ref to prevent re calculating
}

// helper to return a hydrated order. TODO move to a helper...
func order(m common.Address, f bool) fakes.HashOrder { // abigen defined
	// NOTE: none of the actual numbers used matter here for the purpose of this test.
	return fakes.HashOrder{
		Key:        GenBytes32("abc123"),
		Maker:      m,
		Underlying: common.HexToAddress("0xbcd234"),
		Floating:   f,
		Principal:  big.NewInt(1000),
		Interest:   big.NewInt(100),
		Duration:   big.NewInt(123456),
		Expiry:     big.NewInt(123456789),
	}
}

// SetupSuite serves as a 'beforeAll', hydrating both the Env and Dep objects
func (s *hashTestSuite) SetupSuite() {
	// declared because we can't use := here (s.Dep already defined)
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}
}

// NOTE: present just to calculate the hashed values stored in the contract
func (s *hashTestSuite) TestTypeHashes() {
	domainTypehash, _ := s.Dep.HashFake.DomainTypeHash(nil)
	orderTypehash, _ := s.Dep.HashFake.OrderTypeHash(nil)
	s.T().Logf("Domain Typehash: %v", hexutil.Encode(domainTypehash[:]))
	s.T().Logf("Order Typehash: %v", hexutil.Encode(orderTypehash[:]))
}

func (s *hashTestSuite) TestDomain() {
	var err error

	assert := assert.New(s.T())

	s.Separator, err = s.Dep.HashFake.DomainTest(
		nil,
		"Swivel Finance",
		"1.0.0",
		big.NewInt(5),
		common.HexToAddress("0xaBc123"),
	)

	assert.Nil(err)
	assert.NotNil(s.Separator)
	assert.Equal(len(s.Separator), 32)
}

func (s *hashTestSuite) TestOrder() {
	assert := assert.New(s.T())

	order := order(s.Env.Owner.Opts.From, false)

	hash, err := s.Dep.HashFake.OrderTest(nil, order)

	assert.Nil(err)
	assert.NotNil(hash)
	assert.Equal(len(hash), 32)
}

func (s *hashTestSuite) TestMessage() {
	assert := assert.New(s.T())
	// get a hashed order

	order := order(s.Env.Owner.Opts.From, false)
	orderHash, _ := s.Dep.HashFake.OrderTest(nil, order)

	messageHash, err := s.Dep.HashFake.MessageTest(nil, s.Separator, orderHash)

	assert.Nil(err)
	assert.NotNil(messageHash)
	assert.Equal(len(messageHash), 32)
}

func TestHashSuite(t *test.T) {
	suite.Run(t, &hashTestSuite{})
}
