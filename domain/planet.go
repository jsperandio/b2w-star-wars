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
	// Fetch(cursor string, num int64) ([]Planet, string, error)
	FindAll() (result []*Planet, err error)
	GetByID(id string) (*Planet, error)
	// Update( ar *Planet) error
	GetByName(title string) (*Planet, error)
	Store(*Planet) error
	Delete(id string) error
}

// PlanetRepository represent the planet's repository contract
type PlanetRepository interface {
	// Fetch(cursor string, num int64) (res []Planet, nextCursor string, err error)
	FindAll() (result []*Planet, err error)
	GetByID(id primitive.ObjectID) (*Planet, error)
	GetByName(name string) (*Planet, error)
	// Update( ar *Planet) error
	Store(a *Planet) error
	Delete(id primitive.ObjectID) error
}
