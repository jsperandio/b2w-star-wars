package http_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	_clnt "github.com/jsperandio/b2w-star-wars/planet/delivery/http/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RestClientSuite struct {
	rc *_clnt.RESTClient
	suite.Suite
}

// The SetupSuite method will be run by testify once, at the very
// start of the testing suite, before any tests are run.
func (rcs *RestClientSuite) SetupSuite() {
	rcs.rc = _clnt.NewRESTClient("https://test.com/result", 2, 2, 10)
	httpmock.ActivateNonDefault(rcs.rc.Client.GetClient())
}

// The TearDownTest method will be run after every test in the suite.
func (rcs *RestClientSuite) TearDownTest() {
	httpmock.Reset()
}

// The TearDownSuite method will be run by testify once, at the very
// end of the testing suite, after all tests have been run.
func (rcs *RestClientSuite) TearDownSuite() {
	httpmock.DeactivateAndReset()
}

// TestGet check if the Get function on client run for all type of rest API
func (rcs *RestClientSuite) TestGet() {
	assrt := assert.New(rcs.T())

	type BodyTest struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}

	body_test := `{"message": "Test Result", "code": 200}`

	responder := httpmock.NewStringResponder(200, body_test)
	httpmock.RegisterResponder("GET", rcs.rc.ApiUrl, responder)

	intfc_ref := &BodyTest{}
	resp, _ := rcs.rc.Get("", intfc_ref)

	bdytst := resp.Result().(*BodyTest)

	assrt.Equal(1, httpmock.GetTotalCallCount())
	assrt.NotEqual(true, resp.IsError())
	assrt.Equal(BodyTest{
		Message: "Test Result",
		Code:    200,
	}, *bdytst)

}

// TestGet check if the Get function on client run for all type of rest API
func (rcs *RestClientSuite) TestGetError() {
	assrt := assert.New(rcs.T())

	type BodyTest struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}

	responder := httpmock.NewStringResponder(500, "")
	httpmock.RegisterResponder("GET", rcs.rc.ApiUrl+"_error", responder)

	intfc_ref := &BodyTest{}
	resp, _ := rcs.rc.Get("_error", intfc_ref)

	assrt.Equal(1, httpmock.GetTotalCallCount())
	assrt.Equal(500, resp.StatusCode())
	assrt.Equal(true, resp.IsError())
}

// TestConfigSuite is the function to kick off the test suite
func TestRestClientSuite(t *testing.T) {
	suite.Run(t, new(RestClientSuite))
}
