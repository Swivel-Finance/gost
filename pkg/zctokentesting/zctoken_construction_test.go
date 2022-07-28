package zctokentesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	assertions "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/tokens"
)

type zcTokenCtorSuite struct {
	suite.Suite
	Env         *Env
	Dep         *Dep
	MarketPlace *mocks.MarketPlaceSession // *Session objects are created by the go bindings
	ZcToken     *tokens.ZcTokenSession
}

func (s *zcTokenCtorSuite) SetupSuite() {
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

	s.MarketPlace = &mocks.MarketPlaceSession{
		Contract: s.Dep.MarketPlace,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *zcTokenCtorSuite) TestConstruct() {
	assert := assertions.New(s.T())

	// simulate the args that might have been passed to the ctor
	p := uint8(1)
	u := common.HexToAddress("0x12345")
	m := s.Dep.Maturity
	c := common.HexToAddress("0x23456")
	r := s.Dep.MarketPlaceAddress
	n := "Awesome Token"
	symbol := "AT"
	d := uint8(18)

	// depoloy zctoken on demand
	_, _, contract, err := tokens.DeployZcToken(s.Env.Owner.Opts, s.Env.Blockchain, p, u, m, c, r, n, symbol, d)

	assert.Nil(err)
	assert.NotNil(contract)

	s.Env.Blockchain.Commit()

	zct := &tokens.ZcTokenSession{
		Contract: contract,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	proto, err := zct.Protocol()
	assert.Nil(err)
	assert.Equal(proto, p)

	under, err := zct.Underlying()
	assert.Nil(err)
	assert.Equal(under, u)

	maturity, err := zct.Maturity()
	assert.Nil(err)
	assert.Equal(maturity, m)

	cToken, err := zct.CToken()
	assert.Nil(err)
	assert.Equal(cToken, c)

	redeemer, err := zct.Redeemer()
	assert.Nil(err)
	assert.Equal(redeemer, r)
}

func TestZcTokenCtorSuite(t *test.T) {
	suite.Run(t, &zcTokenCtorSuite{})
}
