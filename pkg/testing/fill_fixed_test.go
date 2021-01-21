package testing

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/contracts/fakes"
	"github.com/swivel-finance/gost/test/contracts/swivel"
	"github.com/swivel-finance/gost/test/contracts/tokens"
)

type fillFixedTestSuite struct {
	suite.Suite
	Env    *Env
	Dep    *Dep
	Erc20  *tokens.Erc20Session
	CErc20 *tokens.CErc20Session
	Swivel *swivel.SwivelSession // *Session objects are created by the go bindings
}

func (s *fillFixedTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// bind the owner wallet to the two tokens
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

	// bind user1 to the swivel contract
	s.Swivel = &swivel.SwivelSession{
		Contract: s.Dep.Swivel,
		CallOpts: bind.CallOpts{From: s.Env.User1.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.User1.Opts.From,
			Signer: s.Env.User1.Opts.Signer,
		},
	}
}

func (s *fillFixedTestSuite) TestFillFixed() {
	assert := assert.New(s.T())

	// stub the returns from the tokens
	tx, err := s.Erc20.ApproveReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.Erc20.TransferFromReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	rate := big.NewInt(123456789)
	tx, err = s.CErc20.ExchangeRateCurrentReturns(rate)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	minted := big.NewInt(ONE_GWEI)
	tx, err = s.CErc20.MintReturns(minted)
	assert.NotNil(tx)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

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

	orderHash, err := s.Dep.HashFake.OrderTest(nil, hashOrder)
	assert.Nil(err)
	assert.NotNil(orderHash)

	// put the hashed order together with the eip712 domain and hash those
	separator, _ := s.Swivel.DOMAIN()
	messageHash, err := s.Dep.HashFake.MessageTest(nil, separator, orderHash)
	assert.Nil(err)
	assert.NotNil(messageHash)

	// sign it with User1 private key
	sig, err := crypto.Sign(messageHash[:], s.Env.User1.PK)
	assert.Nil(err)
	assert.NotNil(sig)

	// get the sig components
	vrs, err := s.Dep.SigFake.SplitTest(nil, sig)
	assert.Nil(err)
	assert.NotNil(vrs)

	// see sig_test.go
	if vrs.V < 27 {
		vrs.V += 27
	}

	// the order passed to fillFixed must be from the correct pkg
	order := swivel.HashOrder{
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
	components := swivel.SigComponents{
		V: vrs.V,
		R: vrs.R,
		S: vrs.S,
	}

	// NOW we are ready to call our method...
	agreementKey := GenBytes32("agreement")
	filling := big.NewInt(25) // filling half the available volume

	_, err = s.Swivel.FillFixed(order, filling, agreementKey, components)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// test the agreement was stored
	agreement, err := s.Swivel.Agreements(orderKey, agreementKey)
	assert.Nil(err)
	assert.NotNil(agreement)
}

func TestSwivelSuite(t *test.T) {
	suite.Run(t, &fillFixedTestSuite{})
}
