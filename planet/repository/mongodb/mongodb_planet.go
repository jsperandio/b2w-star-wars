package mongodb

import (
	"context"

	"github.com/jsperandio/b2w-star-wars/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	log "github.com/sirupsen/logrus"
)

const (
	collectionName string = "planet"
)

// MongoRepository represents a mongo_db repository
type MongoDbPlanetRepository struct {
	collection *mongo.Collection
}

// NewMongoRepository creates a mongo planet definition repo
func NewMongoDbPlanetRepository(db *mongo.Database) domain.PlanetRepository {
	return &MongoDbPlanetRepository{collection: db.Collection(collectionName)}
}

// FindAll fetches all the planets definitions available
func (r *MongoDbPlanetRepository) FindAll() (result []*domain.Planet, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), mongoQueryTimeout)
	defer cancel()

	cur, err := r.collection.Find(ctx, bson.M{}, options.Find().SetSort(bson.D{{Key: "name", Value: 1}}))

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		u := new(domain.Planet)

		if err := cur.Decode(u); err != nil {
			return nil, err
		}

		result = append(result, u)
	}

	return result, cur.Err()
}

func (r *MongoDbPlanetRepository) GetByID(id primitive.ObjectID) (*domain.Planet, error) {
	return r.findOneByQuery(bson.M{"_id": id})
}

// FindByplanetName find an planet by planet name case insensitive
func (r *MongoDbPlanetRepository) GetByName(name string) (*domain.Planet, error) {
	return r.findOneByQuery(
		bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: "^" + name + "$", Options: "i"}}},
	)
}

func (r *MongoDbPlanetRepository) findOneByQuery(query interface{}) (*domain.Planet, error) {
	var result domain.Planet

	ctx, cancel := context.WithTimeout(context.Background(), mongoQueryTimeout)
	defer cancel()

	err := r.collection.FindOne(ctx, query).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, domain.ErrNotFound
	}

	return &result, err
}

// Add adds an planet to the repository
func (r *MongoDbPlanetRepository) Store(planet *domain.Planet) error {
	ctx, cancel := context.WithTimeout(context.Background(), mongoQueryTimeout)

	defer cancel()

	planet.ID = primitive.NewObjectID()

	if err := r.collection.FindOneAndUpdate(
		ctx,
		bson.M{"name": planet.Name},
		bson.M{"$set": planet},
		options.FindOneAndUpdate().SetUpsert(true),
	).Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}

		log.WithField("Name", planet.Name).Error("There was an error adding the planet")
		return err
	}

	log.WithField("Name", planet.Name).Debug("planet added")
	return nil
}

// Remove an planet from the repository
func (r *MongoDbPlanetRepository) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), mongoQueryTimeout)
	defer cancel()

	res, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if res.DeletedCount < 1 {
		return domain.ErrNotFound
	}

	log.Debug("planet removed")
	return nil
}
