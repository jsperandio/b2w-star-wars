// Code generated by mockery (devel). DO NOT EDIT.

package mocks

import (
	domain "github.com/jsperandio/b2w-star-wars/domain"
	mock "github.com/stretchr/testify/mock"
)

// PlanetUsecase is an autogenerated mock type for the PlanetUsecase type
type PlanetUsecase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *PlanetUsecase) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields:
func (_m *PlanetUsecase) FindAll() ([]*domain.Planet, error) {
	ret := _m.Called()

	var r0 []*domain.Planet
	if rf, ok := ret.Get(0).(func() []*domain.Planet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Planet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *PlanetUsecase) GetByID(id string) (*domain.Planet, error) {
	ret := _m.Called(id)

	var r0 *domain.Planet
	if rf, ok := ret.Get(0).(func(string) *domain.Planet); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Planet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: title
func (_m *PlanetUsecase) GetByName(title string) (*domain.Planet, error) {
	ret := _m.Called(title)

	var r0 *domain.Planet
	if rf, ok := ret.Get(0).(func(string) *domain.Planet); ok {
		r0 = rf(title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Planet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0
func (_m *PlanetUsecase) Store(_a0 *domain.Planet) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Planet) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
