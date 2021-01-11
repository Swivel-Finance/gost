package testing

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type sigTestSuite struct {
	suite.Suite
	Env *Env
	Dep *Dep
}

var deployErr error

// SetupSuite serves as a 'beforeAll', hydrating both the Env and Dep objects
func (s *sigTestSuite) SetupSuite() {
	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, deployErr = Deploy(s.Env)

	if deployErr != nil {
		panic(deployErr)
	}
}

func (s *sigTestSuite) TestSplit() {
	assert := assert.New(s.T())

	msg := []byte("Yo Dawg, heard u liek unit tests")
	hash := crypto.Keccak256Hash(msg)

	// sign with user1...
	sig, err := crypto.Sign(hash.Bytes(), s.Env.User1.PK)

	assert.Nil(err)
	assert.NotNil(sig)

	vrs, err := s.Dep.Sig.Split(nil, sig)

	assert.Nil(err)
	assert.NotNil(vrs)
}

func (s *sigTestSuite) TestRecoverVRS() {
	assert := assert.New(s.T())

	msg := []byte("Yo Dawg, heard u liek unit tests")
	hash := crypto.Keccak256Hash(msg)

	// sign with user1...
	sig, err := crypto.Sign(hash.Bytes(), s.Env.User1.PK)
	// the go bindings return a struct here
	vrs, err := s.Dep.Sig.Split(nil, sig)

	// crypto.Sign will produce a split whose V is 0 or 1
	v := vrs.V
	// 27 or 28 are acceptable
	if v < 27 {
		v += 27
	}
	// Recover0 allows for them to be passed indiviually if needed
	// NOTE: the peculiar name is the result of the golang implementation of overloading
	addr, err := s.Dep.Sig.Recover(nil, hash, v, vrs.R, vrs.S)

	assert.Nil(err)
	assert.NotNil(addr)
	assert.Equal(addr, s.Env.User1.Opts.From)
}

func TestSuite(t *test.T) {
	suite.Run(t, &sigTestSuite{})
}
