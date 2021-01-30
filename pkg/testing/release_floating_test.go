package testing

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/contracts/fakes"
	"github.com/swivel-finance/gost/test/contracts/swivel"
	"github.com/swivel-finance/gost/test/contracts/tokens"
)

type releaseFloatingTestSuite struct {
	suite.Suite
	Env          *Env
	Dep          *Dep
	Erc20        *tokens.Erc20Session
	CErc20       *tokens.CErc20Session
	Swivel       *swivel.SwivelSession // *Session objects are created by the go bindings
	OrderKey     [32]byte
	AgreementKey [32]byte
}

func (s *releaseFloatingTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	s.Erc20 = &tokens.Erc20Session{
		Contract: s.Dep.Erc20,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.CErc20 = &tokens.CErc20Session{
		Contract: s.Dep.CErc20,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.Swivel = &swivel.SwivelSession{
		Contract: s.Dep.Swivel,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	// get an agreement made and stored
	s.Erc20.ApproveReturns(true)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	// create a correct-ish rate as this test will do the maths...
	rate := big.NewInt(1000000000000000000)
	rate = rate.Mul(rate, big.NewInt(100000000)) // 1e26 total

	s.CErc20.ExchangeRateCurrentReturns(rate)
	s.Env.Blockchain.Commit()

	minted := big.NewInt(ONE_GWEI)
	s.CErc20.MintReturns(minted)
	s.Env.Blockchain.Commit()

	s.OrderKey = GenBytes32("floatingOrder")
	principal := big.NewInt(ONE_ETH)
	interest := big.NewInt(ONE_ETH / 20)
	duration := big.NewInt(1) // so we don't have to adjust time
	expiry := big.NewInt(100)
	nonce := big.NewInt(42)

	hashOrder := fakes.HashOrder{
		Key:        s.OrderKey,
		Maker:      s.Env.User1.Opts.From,
		Underlying: s.Dep.Erc20Address,
		Floating:   true,
		Principal:  principal,
		Interest:   interest,
		Duration:   duration,
		Expiry:     expiry,
		Nonce:      nonce,
	}

	orderHash, _ := s.Dep.HashFake.OrderTest(nil, hashOrder)
	separator, _ := s.Swivel.DOMAIN()
	messageHash, _ := s.Dep.HashFake.MessageTest(nil, separator, orderHash)
	sig, _ := crypto.Sign(messageHash[:], s.Env.User1.PK)
	vrs, _ := s.Dep.SigFake.SplitTest(nil, sig)
	// see sig_test.go
	if vrs.V < 27 {
		vrs.V += 27
	}

	order := swivel.HashOrder{
		Key:        s.OrderKey,
		Maker:      s.Env.User1.Opts.From,
		Underlying: s.Dep.Erc20Address,
		Floating:   true,
		Principal:  principal,
		Interest:   interest,
		Duration:   duration,
		Expiry:     expiry,
		Nonce:      nonce,
	}

	components := swivel.SigComponents{
		V: vrs.V,
		R: vrs.R,
		S: vrs.S,
	}

	s.AgreementKey = GenBytes32("floatingAgreement")
	filling := big.NewInt(ONE_ETH) // filling all available volume (principal in floating...)

	// we can't use the session here as we need User2 to serve as Taker
	opts := &bind.TransactOpts{
		From:   s.Env.User2.Opts.From,
		Signer: s.Env.User2.Opts.Signer,
	}
	// we can send different transact opts directly to the deployed contract
	s.Dep.Swivel.FillFloating(opts, order, filling, s.AgreementKey, components)
	s.Env.Blockchain.Commit()
}

func (s *releaseFloatingTestSuite) TestReleaseFloating() {
	assert := assert.New(s.T())
	// t := s.T() // t is the pointer to the current testing.T instance

	// fetch our agreement that was made in the setup
	agreement, _ := s.Swivel.Agreements(s.OrderKey, s.AgreementKey)
	// assure it has not been released
	assert.False(agreement.Released)

	// we should have user1, user2 as maker/taker
	assert.Equal(agreement.Maker, s.Env.User1.Opts.From)
	assert.Equal(agreement.Taker, s.Env.User2.Opts.From)

	// whats p && i
	// t.Log(agreement.Principal)
	// t.Log(agreement.Interest)

	// TODO helper for generating a rate
	rate := big.NewInt(1050000000000000000)
	rate = rate.Mul(rate, big.NewInt(100000000)) // 5%

	s.CErc20.ExchangeRateCurrentReturns(rate)
	s.Env.Blockchain.Commit()

	// redeemUnderlying must be stubbed to return 0 in order to proceed
	s.CErc20.RedeemUnderlyingReturns(big.NewInt(0))
	s.Env.Blockchain.Commit()

	// stub underlying to return true from transfer
	s.Erc20.TransferReturns(true)

	tx, err := s.Swivel.ReleaseFloating(s.OrderKey, s.AgreementKey)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// what was  redeem... called with
	redeemed, _ := s.CErc20.RedeemedUnderlyingArgs()
	assert.True(redeemed.Cmp(big.NewInt(0)) > 0) // Cmp returns 1 when >
	// t.Log(redeemed)

	// what was transferred?
	makers, _ := s.Erc20.TransferredArgs(agreement.Maker)
	assert.True(makers.Cmp(big.NewInt(0)) > 0)
	// t.Log(makers)
	takers, _ := s.Erc20.TransferredArgs(agreement.Taker)
	assert.True(takers.Cmp(big.NewInt(0)) > 0)
	// t.Log(takers)

	updated, _ := s.Swivel.Agreements(s.OrderKey, s.AgreementKey)
	assert.True(updated.Released)
}

func TestReleaseFloatingSuite(t *test.T) {
	suite.Run(t, &releaseFloatingTestSuite{})
}
