package usecase_test

import (
	"errors"
	"testing"

	"github.com/jsperandio/b2w-star-wars/domain"
	_mock "github.com/jsperandio/b2w-star-wars/domain/mocks"
	_ucase "github.com/jsperandio/b2w-star-wars/planet/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TO-DO
// Refactor default test to test suite

func TestFindAll(t *testing.T) {
	hex_id := "5d2399ef96fb765873a24bae"
	sid, _ := primitive.ObjectIDFromHex(hex_id)

	mockPlanetRepo := new(_mock.PlanetRepository)
	mockPlanet := domain.Planet{
		ID:          sid,
		Name:        "Test",
		Climate:     "Test",
		Terrain:     "Test",
		Appearances: 0,
	}

	mockListPlanet := make([]*domain.Planet, 0)
	mockListPlanet = append(mockListPlanet, &mockPlanet)

	t.Run("success", func(t *testing.T) {
		mockPlanetRepo.On("FindAll").Return(mockListPlanet, nil).Once()

		u := _ucase.NewPlanetUsecase(mockPlanetRepo)

		list, err := u.FindAll()

		assert.NoError(t, err)
		assert.Len(t, list, len(mockListPlanet))

		mockPlanetRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockPlanetRepo.On("FindAll").Return(nil, errors.New("Unexpexted Error")).Once()

		u := _ucase.NewPlanetUsecase(mockPlanetRepo)

		list, err := u.FindAll()

		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockPlanetRepo.AssertExpectations(t)
	})

}

func TestGetByID(t *testing.T) {
	hex_id := "5d2399ef96fb765873a24bae"
	invalid_hex := hex_id + "2021"
	sid, _ := primitive.ObjectIDFromHex(hex_id)

	mockPlanetRepo := new(_mock.PlanetRepository)
	mockPlanet := domain.Planet{
		ID:          sid,
		Name:        "Test",
		Climate:     "Test",
		Terrain:     "Test",
		Appearances: 0,
	}

	t.Run("success", func(t *testing.T) {
		mockPlanetRepo.On("GetByID", mock.AnythingOfType("primitive.ObjectID")).Return(&mockPlanet, nil).Once()

		u := _ucase.NewPlanetUsecase(mockPlanetRepo)

		plnt, err := u.GetByID(hex_id)

		assert.NoError(t, err)
		assert.NotNil(t, plnt)

		mockPlanetRepo.AssertExpectations(t)
	})
	t.Run("error-param", func(t *testing.T) {
		u := _ucase.NewPlanetUsecase(mockPlanetRepo)

		plnt, err := u.GetByID(invalid_hex)

		assert.Error(t, err)
		assert.EqualError(t, domain.ErrBadParamInput, err.Error())
		assert.Nil(t, plnt)

		mockPlanetRepo.AssertExpectations(t)
	})
	t.Run("error-repo", func(t *testing.T) {
		mockPlanetRepo.On("GetByID", mock.AnythingOfType("primitive.ObjectID")).Return(nil, errors.New("Unexpected")).Once()

		u := _ucase.NewPlanetUsecase(mockPlanetRepo)

		plnt, err := u.GetByID(hex_id)

		assert.Error(t, err)
		assert.Nil(t, plnt)

		mockPlanetRepo.AssertExpectations(t)
	})

}

func TestGetByName(t *testing.T) {
	hex_id := "5d2399ef96fb765873a24bae"
	sid, _ := primitive.ObjectIDFromHex(hex_id)

	mockPlanetRepo := new(_mock.PlanetRepository)
	mockPlanet := domain.Planet{
		ID:          sid,
		Name:        "Test",
		Climate:     "Test",
		Terrain:     "Test",
		Appearances: 0,
	}

	t.Run("success", func(t *testing.T) {
		mockPlanetRepo.On("GetByName", mock.AnythingOfType("string")).Return(&mockPlanet, nil).Once()

		u := _ucase.NewPlanetUsecase(mockPlanetRepo)

		plnt, err := u.GetByName(mockPlanet.Name)

		assert.NoError(t, err)
		assert.NotNil(t, plnt)

		mockPlanetRepo.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		mockPlanetRepo.On("GetByName", mock.AnythingOfType("string")).Return(nil, errors.New("Unexpected")).Once()

		u := _ucase.NewPlanetUsecase(mockPlanetRepo)

		plnt, err := u.GetByName(mockPlanet.Name)

		assert.Error(t, err)
		assert.Nil(t, plnt)

		mockPlanetRepo.AssertExpectations(t)
	})

}

func TestStore(t *testing.T) {
	hex_id := "5d2399ef96fb765873a24bae"
	sid, _ := primitive.ObjectIDFromHex(hex_id)

	mockPlanetRepo := new(_mock.PlanetRepository)
	mockPlanet := domain.Planet{
		ID:          sid,
		Name:        "Test",
		Climate:     "Test",
		Terrain:     "Test",
		Appearances: 0,
	}

	t.Run("success", func(t *testing.T) {
		tempMockPlanet := mockPlanet
		mockPlanetRepo.On("GetByName", mock.AnythingOfType("string")).Return(&domain.Planet{}, domain.ErrNotFound).Once()
		mockPlanetRepo.On("Store", mock.AnythingOfType("*domain.Planet")).Return(nil).Once()

		u := _ucase.NewPlanetUsecase(mockPlanetRepo)

		err := u.Store(&tempMockPlanet)

		assert.NoError(t, err)
		assert.Equal(t, mockPlanet.Name, tempMockPlanet.Name)
		mockPlanetRepo.AssertExpectations(t)
	})
	t.Run("alread-exists", func(t *testing.T) {
		existing_planet := mockPlanet
		mockPlanetRepo.On("GetByName", mock.AnythingOfType("string")).Return(&existing_planet, nil).Once()

		u := _ucase.NewPlanetUsecase(mockPlanetRepo)

		err := u.Store(&mockPlanet)

		assert.Error(t, err)
		mockPlanetRepo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	hex_id := "5d2399ef96fb765873a24bae"
	invalid_hex := hex_id + "2021"
	sid, _ := primitive.ObjectIDFromHex(hex_id)

	mockPlanetRepo := new(_mock.PlanetRepository)
	mockPlanet := domain.Planet{
		ID:          sid,
		Name:        "Test",
		Climate:     "Test",
		Terrain:     "Test",
		Appearances: 0,
	}

	t.Run("success", func(t *testing.T) {
		mockPlanetRepo.On("GetByID", mock.AnythingOfType("primitive.ObjectID")).Return(&mockPlanet, nil).Once()

		mockPlanetRepo.On("Delete", mock.AnythingOfType("primitive.ObjectID")).Return(nil).Once()

		u := _ucase.NewPlanetUsecase(mockPlanetRepo)

		err := u.Delete(hex_id)

		assert.NoError(t, err)
		mockPlanetRepo.AssertExpectations(t)
	})
	t.Run("error-param", func(t *testing.T) {
		u := _ucase.NewPlanetUsecase(mockPlanetRepo)

		err := u.Delete(invalid_hex)

		assert.Error(t, err)
		assert.EqualError(t, domain.ErrBadParamInput, err.Error())

		mockPlanetRepo.AssertExpectations(t)
	})
	t.Run("planet-not-exist", func(t *testing.T) {
		mockPlanetRepo.On("GetByID", mock.AnythingOfType("primitive.ObjectID")).Return(&domain.Planet{}, nil).Once()

		u := _ucase.NewPlanetUsecase(mockPlanetRepo)

		err := u.Delete(hex_id)

		assert.Error(t, err)
		mockPlanetRepo.AssertExpectations(t)
	})
	t.Run("error-in-db", func(t *testing.T) {
		mockPlanetRepo.On("GetByID", mock.AnythingOfType("primitive.ObjectID")).Return(nil, errors.New("Unexpected Error")).Once()

		u := _ucase.NewPlanetUsecase(mockPlanetRepo)

		err := u.Delete(hex_id)

		assert.Error(t, err)
		mockPlanetRepo.AssertExpectations(t)
	})

}
