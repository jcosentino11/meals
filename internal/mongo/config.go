package mongo

import (
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config - basic mongo connection configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
}

// NewConfigFromEnv initializes a Config struct using environment variables:
// MONGO_HOST,
// MONGO_PORT,
// MONGO_USERNAME (optional),
// MONGO_PASSWORD (optional)
func NewConfigFromEnv() Config {
	return Config{
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		User:     os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	}
}

// AuthProvided returns true if credentials are provided in the config
func (conf Config) AuthProvided() bool {
	return conf.User != "" && conf.Password != ""
}

// ClientOptions creates a new ClientOptions struct based on the config
func (conf Config) ClientOptions() *options.ClientOptions {
	opts := options.Client().ApplyURI(conf.connectionString())
	if conf.AuthProvided() {
		log.Printf("Using mongo authentication")
		opts.SetAuth(options.Credential{
			Username: conf.User,
			Password: conf.Password,
		})
	}
	return opts
}

func (conf Config) connectionString() string {
	return fmt.Sprintf("mongodb://%s:%s", conf.Host, conf.Port)
}
