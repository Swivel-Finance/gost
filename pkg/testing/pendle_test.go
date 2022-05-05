package testing

import (
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/mocks"
)

type pendleTestSuite struct {
	suite.Suite
	Env    *Env
	Dep    *Dep
	Pendle *mocks.PendleSession
}
