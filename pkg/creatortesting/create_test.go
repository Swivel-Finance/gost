package creatortesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	assertions "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/creator"
)

type createSuite struct {
	suite.Suite
	Env        *Env
	Dep        *Dep
	Creator    *creator.CreatorSession
	Protocol   uint8
	Underlying common.Address
	Maturity   *big.Int
	CToken     common.Address
	Swivel     common.Address
	Name       string
	Symbol     string
	Decimals   uint8
}

func (s *createSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)
	if err != nil {
		panic(err)
	}

	s.Creator = &creator.CreatorSession{
		Contract: s.Dep.Creator,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.Protocol = uint8(1)
	s.Underlying = common.HexToAddress("0x234567")
	s.Maturity = s.Dep.Maturity
	s.CToken = common.HexToAddress("0x345678")
	s.Swivel = common.HexToAddress("0x456789")
	s.Name = "Awesome Token"
	s.Symbol = "AT"
	s.Decimals = uint8(18)
}

func (s *createSuite) testCreateFailsNoMarketPlace() {
	assert := assertions.New(s.T())

	tx, err := s.Creator.Create(s.Protocol, s.Underlying, s.Maturity, s.CToken, s.Swivel, s.Name, s.Symbol, s.Decimals)
	assert.Nil(tx)
	assert.NotNil(err)
}

func (s *createSuite) testCreate() {
	assert := assertions.New(s.T())

	// set the marketplace...
	tx, err := s.Creator.SetMarketPlace(s.Dep.MarketPlaceAddress)
	assert.Nil(err)
	assert.NotNil(tx)

	mPlace, err := s.Creator.MarketPlace()
	assert.Nil(err)
	assert.Equal(mPlace, s.Dep.MarketPlaceAddress)

	tx, err = s.Creator.Create(s.Protocol, s.Underlying, s.Maturity, s.CToken, s.Swivel, s.Name, s.Symbol, s.Decimals)
	assert.Nil(tx)
	assert.NotNil(err)

	// TODO check the params passed to the deployed contracts here
}

func TestCreateSuite(t *test.T) {
	suite.Run(t, &createSuite{})
}
