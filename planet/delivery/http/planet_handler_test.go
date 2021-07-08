package http_test

import (
	"bytes"
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

// Test try with some default Fiber Test
// To-Do Refactor to testify suite ?

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

func TestGetById(t *testing.T) {
	hex_id := "5d2399ef96fb765873a24bae"
	sid, _ := primitive.ObjectIDFromHex(hex_id)

	mockUseCase := new(_mock.PlanetUsecase)
	mockedc := _clnt.NewRESTClient("https://swapi.dev/api/", 2, 2, 10)
	httpmock.ActivateNonDefault(mockedc.Client.GetClient())
	mockSwapiClient := _clnt.NewSwapi(mockedc)

	mockPlanet := domain.Planet{
		ID:          sid,
		Name:        "TestId",
		Climate:     "Test",
		Terrain:     "Test",
		Appearances: 0,
	}

	mockSwapiPlanet := _clnt.SwapiPlanet{
		Name:           "TestId",
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

		mockUseCase.On("GetByID", hex_id).Return(&mockPlanet, nil).Once()

		_phandler.NewPlanetHandler(app, mockUseCase, mockSwapiClient)

		req, _ := http.NewRequest(
			"GET",
			"/api/v1/planet/"+hex_id,
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
			"/api/v1/planet/",
			nil,
		)

		res, _ := app.Test(req, -1)
		assert.NotEqual(t, fiber.StatusOK, res.StatusCode)
		mockUseCase.AssertExpectations(t)
	})
	t.Run("error-usecase", func(t *testing.T) {
		mockUseCase.On("GetByID", hex_id).Return(nil, errors.New("Unexpected")).Once()

		_phandler.NewPlanetHandler(app, mockUseCase, mockSwapiClient)

		req, _ := http.NewRequest(
			"GET",
			"/api/v1/planet/"+hex_id,
			nil,
		)

		res, _ := app.Test(req, -1)
		assert.NotEqual(t, fiber.StatusOK, res.StatusCode)
		mockUseCase.AssertExpectations(t)
	})
}

func TestStore(t *testing.T) {
	mockUseCase := new(_mock.PlanetUsecase)
	mockedc := _clnt.NewRESTClient("https://swapi.dev/api/", 2, 2, 10)
	httpmock.ActivateNonDefault(mockedc.Client.GetClient())
	mockSwapiClient := _clnt.NewSwapi(mockedc)

	var jsonStr = []byte(`{"name": "Bespin","climate": "temperate","terrain": "gas giant"}`)

	mockPlanet := domain.Planet{}

	json.Unmarshal(jsonStr, &mockPlanet)

	app := fiber.New()
	_planetsDeliveryHttpMiddleware.InitFiberMiddleware(app)

	t.Run("success", func(t *testing.T) {

		mockUseCase.On("Store", &mockPlanet).Return(nil).Once()

		_phandler.NewPlanetHandler(app, mockUseCase, mockSwapiClient)

		req, _ := http.NewRequest(
			"POST",
			"/api/v1/planet",
			bytes.NewBuffer(jsonStr),
		)
		req.Header.Set("Content-Type", "application/json")

		res, _ := app.Test(req, -1)

		assert.Equal(t, fiber.StatusCreated, res.StatusCode)
		mockUseCase.AssertExpectations(t)
	})
	t.Run("error-parse", func(t *testing.T) {
		jsonErrorStr := []byte(`{"name" = "Bespin","climate"= "temperate","terrain"= "gas giant"}`)
		_phandler.NewPlanetHandler(app, mockUseCase, mockSwapiClient)

		req, _ := http.NewRequest(
			"POST",
			"/api/v1/planet",
			bytes.NewBuffer(jsonErrorStr),
		)
		req.Header.Set("Content-Type", "application/json")

		res, _ := app.Test(req, -1)

		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})
	t.Run("error-invalid-struct", func(t *testing.T) {
		jsonInvalid := []byte(`{"planet_name" : "Bespin","climate": "temperate","terrain": "gas giant"}`)
		_phandler.NewPlanetHandler(app, mockUseCase, mockSwapiClient)

		req, _ := http.NewRequest(
			"POST",
			"/api/v1/planet",
			bytes.NewBuffer(jsonInvalid),
		)
		req.Header.Set("Content-Type", "application/json")

		res, _ := app.Test(req, -1)

		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})
	t.Run("error-usecase", func(t *testing.T) {

		mockUseCase.On("Store", &mockPlanet).Return(domain.ErrConflict).Once()

		_phandler.NewPlanetHandler(app, mockUseCase, mockSwapiClient)

		req, _ := http.NewRequest(
			"POST",
			"/api/v1/planet",
			bytes.NewBuffer(jsonStr),
		)
		req.Header.Set("Content-Type", "application/json")

		res, _ := app.Test(req, -1)

		assert.NotEqual(t, fiber.StatusCreated, res.StatusCode)
		assert.Equal(t, fiber.StatusConflict, res.StatusCode)
		mockUseCase.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	hex_id := "5d2399ef96fb765873a24bae"

	mockUseCase := new(_mock.PlanetUsecase)
	mockedc := _clnt.NewRESTClient("https://swapi.dev/api/", 2, 2, 10)
	httpmock.ActivateNonDefault(mockedc.Client.GetClient())
	mockSwapiClient := _clnt.NewSwapi(mockedc)

	var jsonStr = []byte(`{"name": "Bespin","climate": "temperate","terrain": "gas giant"}`)

	mockPlanet := domain.Planet{}

	json.Unmarshal(jsonStr, &mockPlanet)

	app := fiber.New()
	_planetsDeliveryHttpMiddleware.InitFiberMiddleware(app)

	t.Run("success", func(t *testing.T) {

		mockUseCase.On("Delete", hex_id).Return(nil).Once()

		_phandler.NewPlanetHandler(app, mockUseCase, mockSwapiClient)

		req, _ := http.NewRequest(
			"DELETE",
			"/api/v1/planet/"+hex_id,
			nil,
		)
		res, _ := app.Test(req, -1)

		assert.Equal(t, fiber.StatusNoContent, res.StatusCode)
		mockUseCase.AssertExpectations(t)
	})
	t.Run("error-param", func(t *testing.T) {
		_phandler.NewPlanetHandler(app, mockUseCase, mockSwapiClient)

		req, _ := http.NewRequest(
			"DELETE",
			"/api/v1/planet",
			nil,
		)

		res, _ := app.Test(req, -1)
		assert.NotEqual(t, fiber.StatusNoContent, res.StatusCode)
		mockUseCase.AssertExpectations(t)
	})
	t.Run("error-usecase", func(t *testing.T) {

		mockUseCase.On("Delete", hex_id).Return(domain.ErrNotFound).Once()

		_phandler.NewPlanetHandler(app, mockUseCase, mockSwapiClient)

		req, _ := http.NewRequest(
			"DELETE",
			"/api/v1/planet/"+hex_id,
			nil,
		)
		res, _ := app.Test(req, -1)

		assert.Equal(t, fiber.StatusNotFound, res.StatusCode)
		mockUseCase.AssertExpectations(t)
	})

}
