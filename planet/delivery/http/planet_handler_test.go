package http_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jarcoal/httpmock"
	"github.com/jsperandio/b2w-star-wars/domain"
	_mock "github.com/jsperandio/b2w-star-wars/domain/mocks"
	_phandler "github.com/jsperandio/b2w-star-wars/planet/delivery/http"
	_clnt "github.com/jsperandio/b2w-star-wars/planet/delivery/http/client"
	_planetsDeliveryHttpMiddleware "github.com/jsperandio/b2w-star-wars/planet/delivery/http/middleware"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestFindAll(t *testing.T) {
	hex_id := "5d2399ef96fb765873a24bae"
	sid, _ := primitive.ObjectIDFromHex(hex_id)

	mockUseCase := new(_mock.PlanetUsecase)
	mockedc := _clnt.NewRESTClient("https://swapi.dev/api/", 2, 2, 10)
	httpmock.ActivateNonDefault(mockedc.Client.GetClient())
	mockSwapiClient := _clnt.NewSwapi(mockedc)

	mockPlanet := domain.Planet{
		ID:          sid,
		Name:        "Test",
		Climate:     "Test",
		Terrain:     "Test",
		Appearances: 0,
	}

	mockSwapiPlanet := _clnt.SwapiPlanet{
		Name:           "Test",
		RotationPeriod: "1",
		OrbitalPeriod:  "1",
		Diameter:       "1",
		Climate:        "Test",
		Gravity:        "1",
		Terrain:        "Test",
		SurfaceWater:   "Test",
		Population:     "1",
		Residents:      []string{},
		Films:          []string{"Test", "test2"},
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

	mockListPlanet := make([]*domain.Planet, 0)
	mockListPlanet = append(mockListPlanet, &mockPlanet)

	respo.Results = append(respo.Results, mockSwapiPlanet)
	resp_json, _ := json.Marshal(respo)

	app := fiber.New()
	_planetsDeliveryHttpMiddleware.InitFiberMiddleware(app)

	t.Run("success", func(t *testing.T) {

		responder := httpmock.NewBytesResponder(200, resp_json)
		httpmock.RegisterResponder("GET", "https://swapi.dev/api/planets/?search="+mockPlanet.Name, responder)

		mockUseCase.On("FindAll").Return(mockListPlanet, nil).Once()

		_phandler.NewPlanetHandler(app, mockUseCase, mockSwapiClient)

		req, _ := http.NewRequest(
			"GET",
			"/api/v1/planets",
			nil,
		)

		res, err_http := app.Test(req, -1)

		body, err_read := ioutil.ReadAll(res.Body)
		mlist, err_json := json.Marshal(mockListPlanet)

		assert.NoError(t, err_http)
		assert.NoError(t, err_read)
		assert.NoError(t, err_json)
		assert.Equal(t, string(mlist), string(body))
		mockUseCase.AssertExpectations(t)
	})
	t.Run("error-usecase", func(t *testing.T) {
		responder := httpmock.NewBytesResponder(200, resp_json)
		httpmock.RegisterResponder("GET", "https://swapi.dev/api/planets/?search="+mockPlanet.Name, responder)

		mockUseCase.On("FindAll").Return(nil, errors.New("Unexpected")).Once()

		_phandler.NewPlanetHandler(app, mockUseCase, mockSwapiClient)

		req, _ := http.NewRequest(
			"GET",
			"/api/v1/planets",
			nil,
		)

		res, _ := app.Test(req, -1)
		assert.NotEqual(t, fiber.StatusOK, res.StatusCode)
		mockUseCase.AssertExpectations(t)
	})
}

func TestGetByName(t *testing.T) {
	hex_id := "5d2399ef96fb765873a24bae"
	sid, _ := primitive.ObjectIDFromHex(hex_id)

	mockUseCase := new(_mock.PlanetUsecase)
	mockedc := _clnt.NewRESTClient("https://swapi.dev/api/", 2, 2, 10)
	httpmock.ActivateNonDefault(mockedc.Client.GetClient())
	mockSwapiClient := _clnt.NewSwapi(mockedc)

	mockPlanet := domain.Planet{
		ID:          sid,
		Name:        "TestName",
		Climate:     "Test",
		Terrain:     "Test",
		Appearances: 0,
	}

	mockSwapiPlanet := _clnt.SwapiPlanet{
		Name:           "TestName",
		RotationPeriod: "1",
		OrbitalPeriod:  "1",
		Diameter:       "1",
		Climate:        "Test",
		Gravity:        "1",
		Terrain:        "Test",
		SurfaceWater:   "Test",
		Population:     "1",
		Residents:      []string{},
		Films:          []string{"Test", "test2"},
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

	respo.Results = append(respo.Results, mockSwapiPlanet)
	resp_json, _ := json.Marshal(respo)

	app := fiber.New()
	_planetsDeliveryHttpMiddleware.InitFiberMiddleware(app)

	t.Run("success", func(t *testing.T) {

		responder := httpmock.NewBytesResponder(200, resp_json)
		httpmock.RegisterResponder("GET", "https://swapi.dev/api/planets/?search="+mockPlanet.Name, responder)

		mockUseCase.On("GetByName", mockPlanet.Name).Return(&mockPlanet, nil).Once()

		_phandler.NewPlanetHandler(app, mockUseCase, mockSwapiClient)

		req, _ := http.NewRequest(
			"GET",
			"/api/v1/planet?name=TestName",
			nil,
		)

		res, err_http := app.Test(req, -1)

		body, err_read := ioutil.ReadAll(res.Body)
		mlist, err_json := json.Marshal(mockPlanet)

		assert.NoError(t, err_http)
		assert.NoError(t, err_read)
		assert.NoError(t, err_json)
		assert.Equal(t, string(mlist), string(body))
		mockUseCase.AssertExpectations(t)
	})
	t.Run("error-param", func(t *testing.T) {
		_phandler.NewPlanetHandler(app, mockUseCase, mockSwapiClient)

		req, _ := http.NewRequest(
			"GET",
			"/api/v1/planet?planet_name=",
			nil,
		)

		res, _ := app.Test(req, -1)
		assert.NotEqual(t, fiber.StatusOK, res.StatusCode)
		mockUseCase.AssertExpectations(t)
	})
	t.Run("error-usecase", func(t *testing.T) {
		mockUseCase.On("GetByName", mockPlanet.Name).Return(nil, errors.New("Unexpected")).Once()

		_phandler.NewPlanetHandler(app, mockUseCase, mockSwapiClient)

		req, _ := http.NewRequest(
			"GET",
			"/api/v1/planet?name=TestName",
			nil,
		)

		res, _ := app.Test(req, -1)
		assert.NotEqual(t, fiber.StatusOK, res.StatusCode)
		mockUseCase.AssertExpectations(t)
	})
}
