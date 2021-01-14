package testing

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type hashTestSuite struct {
	suite.Suite
	Env       *Env
	Dep       *Dep
	Separator [32]byte // keep a ref to prevent re calculating
}

// helper to return an array that can be used for order params
// TODO move to a helper if warranted...
func orderParams() [6]*big.Int {
	// NOTE: none of the actual numbers used matter here for the purpose of this test.
	return [6]*big.Int{
		big.NewInt(5),         // rate
		big.NewInt(1000),      // principal
		big.NewInt(100),       // interest
		big.NewInt(123456),    // duration
		big.NewInt(123456789), // expiry
		big.NewInt(1234),      // nonce
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

func (s *hashTestSuite) TestDomain() {
	var err error

	assert := assert.New(s.T())

	s.Separator, err = s.Dep.HashTest.DomainTest(
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

	hash, err := s.Dep.HashTest.OrderTest(
		nil,
		GenBytes32("abc123"),            // key
		s.Env.Owner.Opts.From,           // owner
		common.HexToAddress("0xbcd234"), // underlying
		true,                            // fixed
		orderParams(),
	)

	assert.Nil(err)
	assert.NotNil(hash)
	assert.Equal(len(hash), 32)
}

func (s *hashTestSuite) TestMessage() {
	assert := assert.New(s.T())
	// get a hashed order
	orderHash, _ := s.Dep.HashTest.OrderTest(
		nil,
		GenBytes32("bcd234"),
		s.Env.Owner.Opts.From,
		common.HexToAddress("0xcde345"),
		false,
		orderParams(),
	)

	messageHash, err := s.Dep.HashTest.MessageTest(nil, s.Separator, orderHash)

	assert.Nil(err)
	assert.NotNil(messageHash)
	assert.Equal(len(messageHash), 32)
}

func TestHashSuite(t *test.T) {
	suite.Run(t, &hashTestSuite{})
}
