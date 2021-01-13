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
	Env *Env
	Dep *Dep
}

var hashTestDeployError error

// SetupSuite serves as a 'beforeAll', hydrating both the Env and Dep objects
func (s *hashTestSuite) SetupSuite() {
	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, hashTestDeployError = Deploy(s.Env)

	if hashTestDeployError != nil {
		panic(hashTestDeployError)
	}
}

func (s *hashTestSuite) TestDomain() {
	separator, err := s.Dep.HashTest.DomainTest(
		nil,
		"Swivel Finance",
		"1.0.0",
		big.NewInt(5),
		common.HexToAddress("0xaBc123"),
	)

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), separator)
}

func TestHashSuite(t *test.T) {
	suite.Run(t, &hashTestSuite{})
}
