package mongodb

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

const (
	mongoConnTimeout  = 10 * time.Second
	mongoQueryTimeout = 10 * time.Second
)

// MongoRepository represents a mongodb repository
type MongoRepository struct {
	DB     *mongo.Database
	client *mongo.Client
}

// NewMongoAppRepository creates a mongo repo for app
func NewMongoAppRepository(dsn string) (*MongoRepository, error) {
	log.WithField("dsn", dsn).Debug("Trying to connect to MongoDB...")

	ctx, cancel := context.WithTimeout(context.Background(), mongoConnTimeout)
	defer cancel()

	fmt.Printf(dsn)

	connString, err := connstring.Parse(dsn)
	if err != nil {
		return nil, fmt.Errorf("could not parse mongodb connection string: %w", err)
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, fmt.Errorf("could not connect to mongodb: %w", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("could not ping mongodb after connect: %w", err)
	}

	mongoDB := client.Database(connString.Database)
	return &MongoRepository{
		DB:     mongoDB,
		client: client,
	}, nil
}

// Close terminates underlying mongo connection.
func (r *MongoRepository) Close() error {
	return r.client.Disconnect(context.TODO())
}
