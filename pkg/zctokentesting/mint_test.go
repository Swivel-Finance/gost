package zctokentesting

import (
	"math/big"
	test "testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	assertions "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/tokens"
)

type zcTokenMintSuite struct {
	suite.Suite
	Env     *Env
	Dep     *Dep
	ZcToken *tokens.ZcTokenSession
}

func (s *zcTokenMintSuite) SetupSuite() {
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
}

func (s *zcTokenMintSuite) TestMint() {
	assert := assertions.New(s.T())

	// simulate the args that might have been passed to the ctor
	p := uint8(1)
	u := common.HexToAddress("0x12345")
	m := s.Dep.Maturity
	c := common.HexToAddress("0x23456")
	// use the owner address as redeemer so our call is authorized
	r := s.Env.Owner.Opts.From
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

	// succeed as we have not moved past maturity
	tx, err := zct.Mint(s.Env.User1.Opts.From, big.NewInt(100))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// fail as we have moved past maturity
	err = s.Env.Blockchain.AdjustTime(MATURITY * time.Second)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	tx, err = zct.Mint(s.Env.User1.Opts.From, big.NewInt(100))
	assert.Nil(tx)
	assert.NotNil(err)
}

func TestZcTokenMintSuite(t *test.T) {
	suite.Run(t, &zcTokenMintSuite{})
}
