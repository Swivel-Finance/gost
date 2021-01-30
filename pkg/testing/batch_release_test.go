package testing

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	// "github.com/swivel-finance/gost/test/contracts/fakes"
	"github.com/swivel-finance/gost/test/contracts/swivel"
	"github.com/swivel-finance/gost/test/contracts/tokens"
)

type batchReleaseTestSuite struct {
	suite.Suite
	Env           *Env
	Dep           *Dep
	Erc20         *tokens.Erc20Session
	CErc20        *tokens.CErc20Session
	Swivel        *swivel.SwivelSession // *Session objects are created by the go bindings
	OrderKeys     [][32]byte
	AgreementKeys [][32]byte
}

func (s *batchReleaseTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH * 2)) // each of the wallets in the env will begin with this balance
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
			From:     s.Env.Owner.Opts.From,
			Signer:   s.Env.Owner.Opts.Signer,
			GasPrice: big.NewInt(ONE_GWEI * 2),
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

	separator, _ := s.Swivel.DOMAIN()

	// create a fixed order. NOTE as the slice has no fixed length we must use append
	s.OrderKeys = append(s.OrderKeys, GenBytes32("fixedOrder"))

	// hash and sign
	hashOrder := NewHashOrder(s.OrderKeys[0], s.Env.User1.Opts.From, s.Dep.Erc20Address, false,
		ONE_ETH, (ONE_ETH / 20), 20, 1000, 41)

	orderHash, _ := s.Dep.HashFake.OrderTest(nil, hashOrder)
	messageHash, _ := s.Dep.HashFake.MessageTest(nil, separator, orderHash)
	sig, _ := crypto.Sign(messageHash[:], s.Env.User1.PK)
	vrs, _ := s.Dep.SigFake.SplitTest(nil, sig)

	// prepare the order to be filled
	order := NewSwivelOrderFromHashOrder(hashOrder)
	components := NewSwivelComponentsFromSigComponents(vrs)

	// fill the prepared fixed order
	s.AgreementKeys = append(s.AgreementKeys, GenBytes32("fixedAgreement"))
	filling := big.NewInt(ONE_ETH / 20) // filling all available volume

	// we can't use the session here as we need User2 to serve as Taker
	opts := &bind.TransactOpts{
		From:   s.Env.User2.Opts.From,
		Signer: s.Env.User2.Opts.Signer,
	}

	// we can send different transact opts directly to the deployed contract
	s.Dep.Swivel.FillFixed(opts, order, filling, s.AgreementKeys[0], components)
	s.Env.Blockchain.Commit()

	// prepare and fill a floating order to be released
	s.OrderKeys = append(s.OrderKeys, GenBytes32("floatingOrder"))

	// hash and sign NOTE the reuse of variables here, no need to := new ones
	hashOrder = NewHashOrder(s.OrderKeys[1], s.Env.User1.Opts.From, s.Dep.Erc20Address, true,
		ONE_ETH, (ONE_ETH / 20), 30, 1000, 43)

	orderHash, _ = s.Dep.HashFake.OrderTest(nil, hashOrder)
	messageHash, _ = s.Dep.HashFake.MessageTest(nil, separator, orderHash)
	sig, _ = crypto.Sign(messageHash[:], s.Env.User1.PK)
	vrs, _ = s.Dep.SigFake.SplitTest(nil, sig)

	// prepare the order to be filled
	order = NewSwivelOrderFromHashOrder(hashOrder)
	components = NewSwivelComponentsFromSigComponents(vrs)

	// fill the prepared floating order
	s.AgreementKeys = append(s.AgreementKeys, GenBytes32("floatingAgreement"))

	filling = big.NewInt(ONE_ETH) // filling all available volume

	opts = &bind.TransactOpts{
		From:   s.Env.User2.Opts.From,
		Signer: s.Env.User2.Opts.Signer,
	}

	s.Dep.Swivel.FillFloating(opts, order, filling, s.AgreementKeys[1], components)
	s.Env.Blockchain.Commit()
}

func (s *batchReleaseTestSuite) TestBatchRelease() {
	assert := assert.New(s.T())

	// fetch the first agreement that was made in the setup
	agreement1, _ := s.Swivel.Agreements(s.OrderKeys[0], s.AgreementKeys[0])
	agreement2, _ := s.Swivel.Agreements(s.OrderKeys[1], s.AgreementKeys[1])

	// s.T().Log(agreement1)
	// s.T().Log(agreement2)

	// assure it has not been released
	assert.False(agreement1.Released)
	assert.False(agreement2.Released)

	// prepare for releasing them
	rate := big.NewInt(1050000000000000000)
	rate = rate.Mul(rate, big.NewInt(100000000)) // 5%

	s.CErc20.ExchangeRateCurrentReturns(rate)
	s.Env.Blockchain.Commit()

	// redeemUnderlying must be stubbed to return 0 in order to proceed
	s.CErc20.RedeemUnderlyingReturns(big.NewInt(0))
	s.Env.Blockchain.Commit()

	// stub underlying to return true from transfer
	s.Erc20.TransferReturns(true)

	tx, err := s.Swivel.BatchRelease(s.OrderKeys, s.AgreementKeys)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// grab the updated agreements
	agreement1, _ = s.Swivel.Agreements(s.OrderKeys[0], s.AgreementKeys[0])
	agreement2, _ = s.Swivel.Agreements(s.OrderKeys[1], s.AgreementKeys[1])
	assert.True(agreement1.Released)
	assert.True(agreement2.Released)
}

func TestBatchReleaseSuite(t *test.T) {
	suite.Run(t, &batchReleaseTestSuite{})
}
