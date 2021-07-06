package http_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	_clnt "github.com/jsperandio/b2w-star-wars/planet/delivery/http/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SwapiSuite struct {
	sc _clnt.SwapiClient
	suite.Suite
}

// The SetupSuite method will be run by testify once, at the very
// start of the testing suite, before any tests are run.
func (ss *SwapiSuite) SetupSuite() {
	mockedc := _clnt.NewRESTClient("https://swapi.dev/api/", 2, 2, 10)
	httpmock.ActivateNonDefault(mockedc.Client.GetClient())

	ss.sc = _clnt.NewSwapi(mockedc)
}

// The TearDownTest method will be run after every test in the suite.
func (ss *SwapiSuite) TearDownTest() {
	httpmock.Reset()
}

// The TearDownSuite method will be run by testify once, at the very
// end of the testing suite, after all tests have been run.
func (ss *SwapiSuite) TearDownSuite() {
	httpmock.DeactivateAndReset()
}

//TestGetPlanetByNameErrorHttp check if errors are propagating for http
func (ss *SwapiSuite) TestGetPlanetByNameErrorHttp() {
	assrt := assert.New(ss.T())

	responder := httpmock.NewStringResponder(500, "")
	httpmock.RegisterResponder("GET", "https://swapi.dev/api/planets/?search=test_http", responder)

	_, err := ss.sc.GetPlanetByName("test_http")

	assrt.Equal("HTTP status `code >= 400`", err.Error())
}

//TestGetPlanetByNameCacheHit check if cache layer are all ok
func (ss *SwapiSuite) TestGetPlanetByNameCacheHit() {
	assrt := assert.New(ss.T())

	tswp := _clnt.SwapiPlanet{
		Name:           "test",
		RotationPeriod: "1",
		OrbitalPeriod:  "1",
		Diameter:       "1",
		Climate:        "test",
		Gravity:        "12",
		Terrain:        "test",
		SurfaceWater:   "test",
		Population:     "1",
		Residents:      []string{},
		Films:          []string{},
		Created:        time.Time{},
		Edited:         time.Time{},
		URL:            "",
	}

	ss.sc.SetCache("test", tswp)

	swp, _ := ss.sc.GetPlanetByName("test")

	assrt.Equal(&tswp, swp)
}

//TestGetPlanetByNameSucess check if return sucess
func (ss *SwapiSuite) TestGetPlanetByNameSucess() {
	assrt := assert.New(ss.T())

	swp := _clnt.SwapiPlanet{
		Name:           "test",
		RotationPeriod: "1",
		OrbitalPeriod:  "1",
		Diameter:       "1",
		Climate:        "test",
		Gravity:        "1",
		Terrain:        "test",
		SurfaceWater:   "test",
		Population:     "1",
		Residents:      []string{},
		Films:          []string{},
		Created:        time.Time{},
		Edited:         time.Time{},
		URL:            "",
	}

	respo := _clnt.Response{
		Count:    1,
		Next:     nil,
		Previous: nil,
		Results:  []_clnt.SwapiPlanet{},
	}
	respo.Results = append(respo.Results, swp)
	json, _ := json.Marshal(respo)

	responder := httpmock.NewBytesResponder(200, json)
	httpmock.RegisterResponder("GET", "https://swapi.dev/api/planets/?search=test_sucess", responder)

	swp_resp, _ := ss.sc.GetPlanetByName("test_sucess")

	assrt.Equal(&swp, swp_resp)
}

//TestGetPlanetByNameSucessButZero check if return sucess but has 0 count
func (ss *SwapiSuite) TestGetPlanetByNameSucessButZero() {
	assrt := assert.New(ss.T())

	respo := _clnt.Response{
		Count:    0,
		Next:     nil,
		Previous: nil,
		Results:  []_clnt.SwapiPlanet{},
	}
	json, _ := json.Marshal(respo)

	responder := httpmock.NewBytesResponder(200, json)
	httpmock.RegisterResponder("GET", "https://swapi.dev/api/planets/?search=test_sucess_but_zero", responder)

	_, err := ss.sc.GetPlanetByName("test_sucess_but_zero")

	assrt.Equal("ok, but no occurrences", err.Error())
}

// TestConfigSuite is the function to kick off the test suite
func TestSwapiSuite(t *testing.T) {
	suite.Run(t, new(SwapiSuite))
}
