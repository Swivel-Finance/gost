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

type pendleYieldTokenTestSuite struct {
	suite.Suite
	Env              *Env
	Dep              *Dep
	PendleYieldToken *mocks.PendleYieldTokenSession
}

func (s *pendleYieldTokenTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.PendleYieldToken = &mocks.PendleYieldTokenSession{
		Contract: s.Dep.PendleYieldToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *pendleYieldTokenTestSuite) TestUnderlyingAsset() {
	assert := assert.New(s.T())
	address := common.HexToAddress("0x0000000000000000000000000000000000000001")
	s.PendleYieldToken.UnderlyingAssetReturns(address)
	s.Env.Blockchain.Commit()

	underlying, err := s.PendleYieldToken.UnderlyingAsset()
	assert.NoError(err)
	assert.Equal(address, underlying)
}

func (s *pendleYieldTokenTestSuite) TestUnderlyingYieldToken() {
	assert := assert.New(s.T())
	address := common.HexToAddress("0x0000000000000000000000000000000000000001")
	s.PendleYieldToken.UnderlyingYieldTokenReturns(address)
	s.Env.Blockchain.Commit()

	underlying, err := s.PendleYieldToken.UnderlyingYieldToken()
	assert.NoError(err)
	assert.Equal(address, underlying)
}

func TestPendleYieldTokenSuite(t *test.T) {
	suite.Run(t, &pendleYieldTokenTestSuite{})
}
