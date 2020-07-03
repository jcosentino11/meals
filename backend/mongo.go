package main

import (
	"context"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoConfig contains all the configuration needed to connect to a mongo cluster
type MongoConfig struct {
	MockDb   bool
	Hosts    string
	Port     string
	Username string
	Password string
}

// NewMongoClient constructs a new client based on configuration
func NewMongoClient(conf MongoConfig) MongoClient {
	if conf.MockDb {
		log.Println("Using mock database")
		return NewNoopMongoClient()
	}
	return NewBasicMongoClient(conf)
}

// MongoClient represents a mongo client implementation
type MongoClient interface {
	Initialize() error
	GetNumDatabases() (int, error)
}

// NoopClient performs no actions. Useful for mocking, or local running
type NoopClient struct {
}

// GetNumDatabases is no-op
func (n *NoopClient) GetNumDatabases() (int, error) {
	return 0, nil
}

// Initialize is no-op
func (n *NoopClient) Initialize() error {
	return nil
}

// NewNoopMongoClient constructs a no-op client
func NewNoopMongoClient() *NoopClient {
	return &NoopClient{}
}

// BasicMongoClient is a mongo client wrapper using a simple interface
type BasicMongoClient struct {
	conf             MongoConfig
	client           *mongo.Client
	operationTimeout time.Duration
}

// NewBasicMongoClient creates a new BasicClient based on the config
func NewBasicMongoClient(conf MongoConfig) *BasicMongoClient {
	return &BasicMongoClient{
		conf:             conf,
		operationTimeout: 2 * time.Second,
	}
}

// Initialize will initiate connection with mongo
func (c *BasicMongoClient) Initialize() error {
	// initiate connection
	log.Println("Initializing mongo connection...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, c.ClientOptions())
	if err != nil {
		return err
	}
	// wait for successful ping
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	log.Println("Mongo connection successful")
	c.client = client
	return nil
}

// ClientOptions creates a new ClientOptions struct based on the config
func (c *BasicMongoClient) ClientOptions() *options.ClientOptions {
	if c.conf.Hosts == "" {
		log.Fatalf("Mongo hosts must be provided in config")
	}
	opts := options.Client().ApplyURI(c.connectionString())
	if c.authProvided() {
		log.Printf("Using mongo authentication")
		opts.SetAuth(options.Credential{
			Username: c.conf.Username,
			Password: c.conf.Password,
		})
	}
	return opts
}

func (c *BasicMongoClient) connectionString() string {
	hosts := c.parseHosts()
	if c.conf.Port != "" {
		for i := range hosts {
			hosts[i] += ":" + c.conf.Port
		}
	}
	return "mongodb://" + strings.Join(hosts, ",")
}

func (c *BasicMongoClient) parseHosts() []string {
	if c.conf.Hosts == "" {
		log.Fatalf("Mongo hosts must be provided in config")
	}
	return strings.Split(c.conf.Hosts, ",")
}

// AuthProvided returns true if credentials are provided in the config
func (c *BasicMongoClient) authProvided() bool {
	return c.conf.Username != "" && c.conf.Password != ""
}

// GetNumDatabases reports the number of databases within the mongo cluster.
// This is just a dummy function to test out basic integration with mongo,
// and will be removed once actual features are implemented :)
func (c *BasicMongoClient) GetNumDatabases() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.operationTimeout)
	defer cancel()
	result, err := c.client.ListDatabases(ctx, bson.D{})
	if err != nil {
		log.Printf("failed to list databases: %s", err)
		return 0, err
	}
	return len(result.Databases), nil
}
