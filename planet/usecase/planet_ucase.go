package usecase

import (
	"github.com/jsperandio/b2w-star-wars/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type planetUsecase struct {
	planetRepo domain.PlanetRepository
}

// NewplanetUsecase will create new an planetUsecase object representation of domain.planetUsecase interface
func NewPlanetUsecase(a domain.PlanetRepository) domain.PlanetUsecase {
	return &planetUsecase{
		planetRepo: a,
	}
}

func (a *planetUsecase) FindAll() (result []*domain.Planet, err error) {
	result, err = a.planetRepo.FindAll()

	if err != nil {
		return nil, err
	}

	return result, err
}

func (a *planetUsecase) GetByID(id string) (res *domain.Planet, err error) {
	var sid primitive.ObjectID

	sid, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, domain.ErrBadParamInput
	}

	res, err = a.planetRepo.GetByID(sid)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (a *planetUsecase) GetByName(name string) (res *domain.Planet, err error) {
	res, err = a.planetRepo.GetByName(name)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (a *planetUsecase) Store(m *domain.Planet) (err error) {
	existedplanet, _ := a.GetByName(m.Name)

	if existedplanet != (&domain.Planet{}) {
		return domain.ErrConflict
	}

	err = a.planetRepo.Store(m)
	return nil
}

func (a *planetUsecase) Delete(id string) (err error) {
	var sid primitive.ObjectID

	sid, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		return domain.ErrBadParamInput
	}

	existedplanet, err := a.planetRepo.GetByID(sid)

	if err != nil {
		return err
	}
	if existedplanet.Name == "" {
		return domain.ErrNotFound
	}
	return a.planetRepo.Delete(sid)
}
