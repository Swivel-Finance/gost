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

type pendleDataTestSuite struct {
	suite.Suite
	Env        *Env
	Dep        *Dep
	PendleData *mocks.PendleDataSession
}

func (s *pendleDataTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.PendleData = &mocks.PendleDataSession{
		Contract: s.Dep.PendleData,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *pendleDataTestSuite) TestIsValidXYT() {
	assert := assert.New(s.T())

	s.PendleData.IsValidXYTReturns(true)
	s.Env.Blockchain.Commit()

	s.PendleData.IsValidXYT([32]byte{1, 2, 3}, common.BigToAddress(big.NewInt(1)), big.NewInt(5))
	s.Env.Blockchain.Commit()

	byteReturn, err := s.PendleData.IsValidXYTId()
	assert.NoError(err)
	assert.Equal([32]byte{1, 2, 3}, byteReturn)

	addressReturn, err := s.PendleData.IsValidXYTUnderlying()
	assert.NoError(err)
	assert.Equal(common.BigToAddress(big.NewInt(1)), addressReturn)

	amountReturn, err := s.PendleData.IsValidXYTMaturity()
	assert.NoError(err)
	assert.Equal(big.NewInt(5), amountReturn)
}

func (s *pendleDataTestSuite) TestIsValidOT() {
	assert := assert.New(s.T())

	s.PendleData.IsValidOTReturns(true)
	s.Env.Blockchain.Commit()

	s.PendleData.IsValidOT([32]byte{1, 2, 3}, common.BigToAddress(big.NewInt(1)), big.NewInt(5))
	s.Env.Blockchain.Commit()

	byteReturn, err := s.PendleData.IsValidOTId()
	assert.NoError(err)
	assert.Equal([32]byte{1, 2, 3}, byteReturn)

	addressReturn, err := s.PendleData.IsValidOTUnderlying()
	assert.NoError(err)
	assert.Equal(common.BigToAddress(big.NewInt(1)), addressReturn)

	amountReturn, err := s.PendleData.IsValidOTMaturity()
	assert.NoError(err)
	assert.Equal(big.NewInt(5), amountReturn)
}

func (s *pendleDataTestSuite) TestXYTTokens() {
	assert := assert.New(s.T())

	s.PendleData.XytTokensReturns(common.BigToAddress(big.NewInt(1)))
	s.Env.Blockchain.Commit()

	s.PendleData.XytTokens([32]byte{1, 2, 3}, common.BigToAddress(big.NewInt(1)), big.NewInt(5))
	s.Env.Blockchain.Commit()

	byteReturn, err := s.PendleData.XytTokensCalledId()
	assert.NoError(err)
	assert.Equal([32]byte{1, 2, 3}, byteReturn)

	underlyingReturn, err := s.PendleData.XytTokensCalledUnderlying()
	assert.NoError(err)
	assert.Equal(common.BigToAddress(big.NewInt(1)), underlyingReturn)

	maturityReturn, err := s.PendleData.XytTokensCalledMaturity()
	assert.NoError(err)
	assert.Equal(big.NewInt(5), maturityReturn)
}
func TestPendleDataSuite(t *test.T) {
	suite.Run(t, &pendleDataTestSuite{})
}
