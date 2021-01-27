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
)

type cancelTestSuite struct {
	suite.Suite
	Env        *Env
	Dep        *Dep
	Swivel     *swivel.SwivelSession
	Order      swivel.HashOrder
	Components swivel.SigComponents
}

func (s *cancelTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH))
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// bind owner to the swivel contract
	s.Swivel = &swivel.SwivelSession{
		Contract: s.Dep.Swivel,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	// put together an order, hash and sign it...
	// this represents the creation and storage of the offline order (and sig)
	orderKey := GenBytes32("order")
	principal := big.NewInt(5000)
	interest := big.NewInt(50)
	duration := big.NewInt(10000)
	expiry := big.NewInt(20000)
	nonce := big.NewInt(42)

	// we'll say User1 made a fixed order some time ago
	// NOTE fakes.* is being used here as the swivel package does not expose them
	// both are using the same sig && hash libs however
	hashOrder := fakes.HashOrder{
		Key:        orderKey,
		Maker:      s.Env.User1.Opts.From,
		Underlying: s.Dep.Erc20Address,
		Floating:   false,
		Principal:  principal,
		Interest:   interest,
		Duration:   duration,
		Expiry:     expiry,
		Nonce:      nonce,
	}

	orderHash, _ := s.Dep.HashFake.OrderTest(nil, hashOrder)
	// put the hashed order together with the eip712 domain and hash those
	separator, _ := s.Swivel.DOMAIN()
	messageHash, _ := s.Dep.HashFake.MessageTest(nil, separator, orderHash)
	// sign it with User1 private key
	sig, _ := crypto.Sign(messageHash[:], s.Env.User1.PK)
	// get the sig components
	vrs, _ := s.Dep.SigFake.SplitTest(nil, sig)

	// see sig_test.go
	if vrs.V < 27 {
		vrs.V += 27
	}

	// the order passed to cancel must be from the correct pkg
	s.Order = swivel.HashOrder{
		Key:        orderKey,
		Maker:      s.Env.User1.Opts.From,
		Underlying: s.Dep.Erc20Address,
		Floating:   false,
		Principal:  principal,
		Interest:   interest,
		Duration:   duration,
		Expiry:     expiry,
		Nonce:      nonce,
	}

	// like order the signature components must ref swivel
	s.Components = swivel.SigComponents{
		V: vrs.V,
		R: vrs.R,
		S: vrs.S,
	}
}

func (s *cancelTestSuite) TestCancel() {
	assert := assert.New(s.T())

	cancelled, _ := s.Swivel.Cancelled(s.Order.Key)
	assert.False(cancelled)

	tx, err := s.Swivel.Cancel(s.Order, s.Components)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	cancelled, _ = s.Swivel.Cancelled(s.Order.Key)
	assert.True(cancelled)
}

func TestCancelSuite(t *test.T) {
	suite.Run(t, &cancelTestSuite{})
}
