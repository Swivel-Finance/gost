package testing

import (
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/mocks"
)

type swivelTokenTestSuite struct {
	suite.Suite
	Env        *Env
	Dep        *Dep
	YieldToken *mocks.SwivelTokenSession
}
