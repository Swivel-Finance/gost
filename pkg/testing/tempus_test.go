package testing

import "github.com/stretchr/testify/suite"

type tempusTestSuite struct {
	suite.Suite
	Env    *Env
	Dep    *Dep
	Tempus *mocks.Tempus
}
