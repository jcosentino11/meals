package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type dbClient interface {
	GetNumDatabases() int
}

// BasicClient is a mongo client wrapper using a simple interface
type BasicClient struct {
	m *mongo.Client
	operationTimeout time.Duration
}

// NewBasicClient creates a new BasicClient based on the config
func NewBasicClient(conf Config) (*BasicClient, error) {
	client, err := initMongoClient(conf)
	if err != nil {
		return nil, err
	}
	return &BasicClient{
		m: client,
		operationTimeout: 2*time.Second,
	}, nil
}

func initMongoClient(conf Config) (*mongo.Client, error) {
	// initiate connection
	log.Println("Initializing mongo connection...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, conf.ClientOptions())
	if err != nil {
		return nil, err
	}
	// wait for successful ping
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	log.Println("Mongo connection successful")
	return client, nil
}

// GetNumDatabases reports the number of databases within the mongo cluster.
// This is just a dummy function to test out basic integration with mongo,
// and will be removed once actual features are implemented :)
func (c *BasicClient) GetNumDatabases() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.operationTimeout)
	defer cancel()
	result, err := c.m.ListDatabases(ctx, bson.D{})
	if err != nil {
		log.Printf("failed to list databases: %s", err)
		return 0, err
	}
	return len(result.Databases), nil
}
