package creatortesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	assertions "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/creator"
)

type creatorCtorSuite struct {
	suite.Suite
	Env     *Env
	Dep     *Dep
	Creator *creator.CreatorSession
}

func (s *creatorCtorSuite) SetupSuite() {
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
}

func (s *creatorCtorSuite) TestAdmin() {
	assert := assertions.New(s.T())

	admin, err := s.Creator.Admin()
	assert.Nil(err)
	assert.Equal(s.Env.Owner.Opts.From, admin)
}

func TestCreatorCtorSuite(t *test.T) {
	suite.Run(t, &creatorCtorSuite{})
}
