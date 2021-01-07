package testing

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"math/big"
	test "testing"
)

type tokenTestSuite struct {
	suite.Suite
	Env *Env
	Dep *Dep
}

var deployErr error

// SetupSuite serves as a 'beforeAll', hydrating both the Env and Dep objects
func (s *tokenTestSuite) SetupSuite() {
	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, deployErr = Deploy(s.Env)

	if deployErr != nil {
		panic(deployErr)
	}
}

func (s *tokenTestSuite) TestTotalSupply() {
	supply, err := s.Dep.Token.TotalSupply(nil) // NOTE that nil can be passed for *CallOpts
	// if you wanna log some stuff...
	s.T().Logf("Token initial supply is: %v", Commafy(supply))
	// Cmp returns 0 if the two are equal
	assert.Nil(s.T(), err)
	res := supply.Cmp(big.NewInt(ONE_ETH))
	assert.Equal(s.T(), res, 0)
}

func TestSuite(t *test.T) {
	suite.Run(t, &tokenTestSuite{})
}
