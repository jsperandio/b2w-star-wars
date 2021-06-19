package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Planet ...
type Planet struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" validate:"required" bson:"name"`
	Climate     string             `json:"climate" validate:"required" bson:"climate"`
	Terrain     string             `json:"terrain" validate:"required" bson:"terrain"`
	Appearances int                `json:"appearances" bson:"-"`
}

// PlanetUsecase represent the Planet usecases
type PlanetUsecase interface {
	FindAll() (result []*Planet, err error)
	GetByID(id string) (*Planet, error)
	GetByName(title string) (*Planet, error)
	Store(*Planet) error
	Delete(id string) error
}

// PlanetRepository represent the planet's repository contract
type PlanetRepository interface {
	FindAll() (result []*Planet, err error)
	GetByID(id primitive.ObjectID) (*Planet, error)
	GetByName(name string) (*Planet, error)
	Store(a *Planet) error
	Delete(id primitive.ObjectID) error
}
