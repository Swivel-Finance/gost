package testing

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/contracts/swivel"
	"github.com/swivel-finance/gost/test/contracts/tokens"
)

type batchFillFloatingTestSuite struct {
	suite.Suite
	Env    *Env
	Dep    *Dep
	Erc20  *tokens.Erc20Session
	CErc20 *tokens.CErc20Session
	Swivel *swivel.SwivelSession // *Session objects are created by the go bindings
}

func (s *batchFillFloatingTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// bind the owner wallet to the two tokens and swivel
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
			From:   s.Env.User1.Opts.From,
			Signer: s.Env.User1.Opts.Signer,
		},
	}
}

func (s *batchFillFloatingTestSuite) TestBatchFillFloating() {
	assert := assert.New(s.T())

	// stub the returns from the tokens
	s.Erc20.ApproveReturns(true)
	s.Env.Blockchain.Commit()

	s.Erc20.TransferFromReturns(true)
	s.Env.Blockchain.Commit()

	rate := big.NewInt(123456789)
	s.CErc20.ExchangeRateCurrentReturns(rate)
	s.Env.Blockchain.Commit()

	minted := big.NewInt(ONE_GWEI)
	s.CErc20.MintReturns(minted)
	s.Env.Blockchain.Commit()

	// we'll say User1 made fixed orders some time ago
	sigOrder := NewHashOrder(GenBytes32("order1"), s.Env.User1.Opts.From, s.Dep.Erc20Address, true, 5000, 50,
		10000, 20000, 42)

	sigOrderHash, _ := s.Dep.HashFake.OrderTest(nil, sigOrder)
	// put the hashed order together with the eip712 domain and hash those
	separator, _ := s.Swivel.DOMAIN()
	sigOrderMessageHash, _ := s.Dep.HashFake.MessageTest(nil, separator, sigOrderHash)
	// sign it with User1 private key
	sig, _ := crypto.Sign(sigOrderMessageHash[:], s.Env.User1.PK)
	// get the sig components
	vrs, _ := s.Dep.SigFake.SplitTest(nil, sig)

	// the type must be swivel order to be passed to contract as arg
	order1 := NewSwivelOrderFromHashOrder(sigOrder)
	// like order the signature components must ref swivel
	components1 := NewSwivelComponentsFromSigComponents(vrs)

	// make a second order
	sigOrder = NewHashOrder(GenBytes32("order2"), s.Env.User1.Opts.From, s.Dep.Erc20Address, true, 6000, 60,
		10000, 20000, 44)

	sigOrderHash, _ = s.Dep.HashFake.OrderTest(nil, sigOrder)
	sigOrderMessageHash, _ = s.Dep.HashFake.MessageTest(nil, separator, sigOrderHash)
	// sign it with User1 private key
	sig, _ = crypto.Sign(sigOrderMessageHash[:], s.Env.User1.PK)
	// get the sig components
	vrs, _ = s.Dep.SigFake.SplitTest(nil, sig)

	// the type must be swivel order to be passed to contract as arg
	order2 := NewSwivelOrderFromHashOrder(sigOrder)
	// like order the signature components must ref swivel
	components2 := NewSwivelComponentsFromSigComponents(vrs)

	// NOW we are ready to call our method...
	agreementKey := GenBytes32("batchAgreement")
	orders := []swivel.HashOrder{order1, order2}
	filling := []*big.Int{big.NewInt(25), big.NewInt(30)}
	sigs := []swivel.SigComponents{components1, components2}

	tx, err := s.Swivel.BatchFillFloating(orders, filling, agreementKey, sigs)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// test the agreements were stored
	agreement1, err := s.Swivel.Agreements(order1.Key, agreementKey)
	assert.Nil(err)
	assert.NotNil(agreement1)
	// agreement.principal is the taker volume amount
	assert.Equal(agreement1.Principal.Cmp(big.NewInt(25)), 0)

	agreement2, err := s.Swivel.Agreements(order2.Key, agreementKey)
	assert.Nil(err)
	assert.NotNil(agreement2)
	assert.Equal(agreement2.Principal.Cmp(big.NewInt(30)), 0)
}

func TestBatchFloatingSuite(t *test.T) {
	suite.Run(t, &batchFillFloatingTestSuite{})
}
