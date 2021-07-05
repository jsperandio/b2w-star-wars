package http_test

import (
	_clnt "github.com/jsperandio/b2w-star-wars/planet/delivery/http/client"
	"github.com/stretchr/testify/suite"
)

type SwapiSuite struct {
	sc _clnt.SwapiClient
	suite.Suite
}

// The SetupSuite method will be run by testify once, at the very
// start of the testing suite, before any tests are run.
func (ss *SwapiSuite) SetupSuite() {
	ss.sc = _clnt.NewSwapi(nil)

}
