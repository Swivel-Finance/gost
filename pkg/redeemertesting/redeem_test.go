package redeemertesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/redeemer"
)

type redeemTestSuite struct {
	suite.Suite
	Env         *Env
	Dep         *Dep
	Erc20       *mocks.Erc20Session
	Illuminate  *mocks.IlluminateSession
	Yield       *mocks.YieldSession
	ZcToken     *mocks.ZcTokenSession
	Swivel      *mocks.SwivelSession
	APWine      *mocks.APWineSession
	APWineToken *mocks.APWineTokenSession
	Redeemer    *redeemer.RedeemerSession
}

func (s *redeemTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.Erc20 = &mocks.Erc20Session{
		Contract: s.Dep.Erc20,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.Illuminate = &mocks.IlluminateSession{
		Contract: s.Dep.Illuminate,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.ZcToken = &mocks.ZcTokenSession{
		Contract: s.Dep.ZcToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.Yield = &mocks.YieldSession{
		Contract: s.Dep.Yield,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.Swivel = &mocks.SwivelSession{
		Contract: s.Dep.Swivel,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.APWine = &mocks.APWineSession{
		Contract: s.Dep.APWine,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.APWineToken = &mocks.APWineTokenSession{
		Contract: s.Dep.APWineToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.Redeemer = &redeemer.RedeemerSession{
		Contract: s.Dep.Redeemer,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *redeemTestSuite) TestAPWineRedeem() {
	assert := assert.New(s.T())

	amount := big.NewInt(1000)
	maturity := big.NewInt(9999999)
	vault := common.HexToAddress("0x0000000000000000000000000000000000000002")
	apwinePrincipal := uint8(7)

	s.Illuminate.MarketsReturns([8]common.Address{
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
		s.Dep.APWineTokenAddress,
	})
	s.Env.Blockchain.Commit()

	s.APWineToken.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	s.APWineToken.BalanceOfReturns(vault, amount)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferReturns(true)
	s.Env.Blockchain.Commit()

	tx, err := s.Redeemer.Redeem0(apwinePrincipal, s.Dep.Erc20Address, maturity, vault)
	assert.NoError(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()
}

func (s *redeemTestSuite) TestTempusRedeem() {

}

func (s *redeemTestSuite) TestIlluminateRedeem() {

}

func TestRedeemSuite(t *test.T) {
	suite.Run(t, &redeemTestSuite{})
}
