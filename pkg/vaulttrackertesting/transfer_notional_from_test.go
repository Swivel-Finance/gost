package vaulttrackertesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	assertions "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/vaulttracker"
)

type transferNotionalFromSuite struct {
	suite.Suite
	Env          *Env
	Dep          *Dep
	AaveToken    *mocks.AaveTokenSession
	AavePool     *mocks.AavePoolSession
	VaultTracker *vaulttracker.VaultTrackerSession // *Session objects are created by the go bindings
}

func (s *transferNotionalFromSuite) SetupTest() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)
	if err != nil {
		panic(err)
	}

	err = s.Env.Blockchain.AdjustTime(0) // set bc timestamp to 0
	if err != nil {
		panic(err)
	}
	s.Env.Blockchain.Commit()

	s.AaveToken = &mocks.AaveTokenSession{
		Contract: s.Dep.AaveToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.AavePool = &mocks.AavePoolSession{
		Contract: s.Dep.AavePool,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	// the exchangerate chain for Aave must be stubbed here or we'll revert
	_, err = s.AaveToken.UNDERLYINGReturns(common.HexToAddress("0x1234"))
	if err != nil {
		panic(err)
	}
	s.Env.Blockchain.Commit()

	_, err = s.AaveToken.POOLReturns(s.Dep.AavePoolAddress)
	if err != nil {
		panic(err)
	}
	s.Env.Blockchain.Commit()

	rate1 := big.NewInt(123456789)
	_, err = s.AavePool.GetReserveNormalizedIncomeReturns(rate1)
	if err != nil {
		panic(err)
	}
	s.Env.Blockchain.Commit()

	// deploy a vaultTracker with Protocols.Aave
	_, _, contract, err := vaulttracker.DeployVaultTracker(s.Env.Owner.Opts, s.Env.Blockchain, uint8(4),
		big.NewInt(MATURITY), s.Dep.AaveTokenAddress, s.Dep.SwivelAddress)

	if err != nil {
		// panic(err)
		s.T().Log(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.VaultTracker = &vaulttracker.VaultTrackerSession{
		Contract: contract,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *transferNotionalFromSuite) TestTransferNotionalFrom() {
	assert := assertions.New(s.T())

	tx, err := s.AaveToken.UNDERLYINGReturns(common.HexToAddress("0x1234"))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.AaveToken.POOLReturns(s.Dep.AavePoolAddress)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	rate1 := big.NewInt(123456789)
	tx, err = s.AavePool.GetReserveNormalizedIncomeReturns(rate1)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// call AddNotional for Owner with no ownerVault
	amount1 := big.NewInt(1000)
	tx, err = s.VaultTracker.AddNotional(s.Env.Owner.Opts.From, amount1)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// call AddNotional for User1 with no ownerVault
	amount2 := big.NewInt(1000)
	tx, err = s.VaultTracker.AddNotional(s.Env.User1.Opts.From, amount2)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	tx, err = s.VaultTracker.TransferNotionalFrom(s.Env.Owner.Opts.From, s.Env.User1.Opts.From, big.NewInt(100))
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	ownerVault, err := s.VaultTracker.Vaults(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.NotNil(ownerVault)
	assert.Equal(big.NewInt(900), ownerVault.Notional)
	assert.Equal(rate1, ownerVault.ExchangeRate)
	assert.Equal(ownerVault.Redeemable.Cmp(ZERO), 0)

	user1Vault, err := s.VaultTracker.Vaults(s.Env.User1.Opts.From)
	assert.Nil(err)
	assert.NotNil(user1Vault)
	assert.Equal(big.NewInt(1100), user1Vault.Notional)
	assert.Equal(rate1, user1Vault.ExchangeRate)
	assert.Equal(user1Vault.Redeemable.Cmp(ZERO), 0)
}

func (s *transferNotionalFromSuite) TestTransferNotionalFromAmountExceedsFail() {
	assert := assertions.New(s.T())

	tx, err := s.AaveToken.UNDERLYINGReturns(common.HexToAddress("0x1234"))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.AaveToken.POOLReturns(s.Dep.AavePoolAddress)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	rate1 := big.NewInt(123456789)
	tx, err = s.AavePool.GetReserveNormalizedIncomeReturns(rate1)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// call AddNotional for Owner with no ownerVault
	amount1 := big.NewInt(1000)
	tx, err = s.VaultTracker.AddNotional(s.Env.Owner.Opts.From, amount1)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// call AddNotional for User1 with no ownerVault
	amount2 := big.NewInt(1000)
	tx, err = s.VaultTracker.AddNotional(s.Env.User1.Opts.From, amount2)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	tx, err = s.VaultTracker.TransferNotionalFrom(s.Env.Owner.Opts.From, s.Env.User1.Opts.From, big.NewInt(2000))
	assert.NotNil(err)
	// TODO custom err codes?
	// assert.Regexp("amount exceeds available balance", err.Error())
	assert.Nil(tx)

	s.Env.Blockchain.Commit()
}

func TestTrackerTransferNotionalFromSuite(t *test.T) {
	suite.Run(t, &transferNotionalFromSuite{})
}
