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

type tempusTestSuite struct {
	suite.Suite
	Env    *Env
	Dep    *Dep
	Tempus *mocks.TempusSession
}

func (s *tempusTestSuite) SetupTest() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.Tempus = &mocks.TempusSession{
		Contract: s.Dep.Tempus,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *tempusTestSuite) TestMaturityTime() {
	assert := assert.New(s.T())

	maturity := big.NewInt(14204)
	s.Tempus.MaturityTimeReturns(maturity)

	s.Env.Blockchain.Commit()

	maturityTime, err := s.Tempus.MaturityTime()
	assert.NoError(err)
	assert.Equal(maturity, maturityTime)
}

func (s *tempusTestSuite) TestYieldBearingToken() {
	assert := assert.New(s.T())

	token := common.BigToAddress(big.NewInt(1))
	s.Tempus.YieldBearingTokenReturns(token)

	s.Env.Blockchain.Commit()

	returnedToken, err := s.Tempus.YieldBearingToken()
	assert.NoError(err)
	assert.Equal(token, returnedToken)
}

func (s *tempusTestSuite) TestDepositAndFix() {
	assert := assert.New(s.T())
	s.Tempus.DepositAndFixReturns(big.NewInt(1000))

	tx, err := s.Tempus.DepositAndFix(
		common.BigToAddress(big.NewInt(1)),
		common.BigToAddress(big.NewInt(2)),
		big.NewInt(1000),
		true,
		big.NewInt(2000),
		big.NewInt(1420),
	)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	amm, err := s.Tempus.TempusAMMCalled()
	assert.NoError(err)
	assert.Equal(common.BigToAddress(big.NewInt(1)), amm)

	pool, err := s.Tempus.TempusPoolCalled()
	assert.NoError(err)
	assert.Equal(common.BigToAddress(big.NewInt(2)), pool)

	amount, err := s.Tempus.AmountCalled()
	assert.NoError(err)
	assert.Equal(big.NewInt(1000), amount)

	isBackingToken, err := s.Tempus.IsBackingTokenCalled()
	assert.NoError(err)
	assert.Equal(true, isBackingToken)

	minimumRate, err := s.Tempus.MinimumReturnCalled()
	assert.NoError(err)
	assert.Equal(big.NewInt(2000), minimumRate)

	deadline, err := s.Tempus.DeadlineCalled()
	assert.NoError(err)
	assert.Equal(big.NewInt(1420), deadline)
}

func TestTempusSuite(t *test.T) {
	suite.Run(t, &tempusTestSuite{})
}
