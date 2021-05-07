package swiveltesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/swivel"
)

// yeah, just make it an acronym...
type IVFZISuite struct {
	suite.Suite
	Env         *Env
	Dep         *Dep
	Erc20       *mocks.Erc20Session
	MarketPlace *mocks.MarketPlaceSession
	Swivel      *swivel.SwivelSession
}

func (s *IVFZISuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH))
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	s.Erc20 = &mocks.Erc20Session{
		Contract: s.Dep.Erc20,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.MarketPlace = &mocks.MarketPlaceSession{
		Contract: s.Dep.MarketPlace,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.Swivel = &swivel.SwivelSession{
		Contract: s.Dep.Swivel,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *swivelCtorSuite) TestMarketPlaceAddress() {
	assert := assert.New(s.T())
	addr, err := s.Swivel.MarketPlaceAddr()
	assert.Nil(err)
	assert.Equal(addr, s.Dep.MarketPlaceAddress)
}

func (s *IVFZISuite) TestIVFZI() {
	assert := assert.New(s.T())

	// stub underlying (erc20) transferfrom to return true
	tx, err := s.Erc20.TransferFromReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)

	// and approve
	tx, err = s.Erc20.ApproveReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)

	// and cTokenAddress getter
	ctAddr := common.HexToAddress("0xImmaCtoken")
	tx, err = s.MarketPlace.MarketCTokenAddressReturns(ctAddr)
	assert.Nil(err)
	assert.NotNil(tx)

	// and the ctoken mint

	// TODO order
	// TODO signature
	// call it...
}

func TestIVFZISuite(t *test.T) {
	suite.Run(t, &IVFZISuite{})
}
