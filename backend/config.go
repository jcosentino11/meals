package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	// EnvAuthEnabled is the name of env var that
	// controls whether auth is enabled for endpoints
	EnvAuthEnabled = "APP_AUTH_ENABLED"
	// EnvAuthEnabledDefault is the default value for EnvAuthEnabled
	EnvAuthEnabledDefault = true

	// EnvMockDb will use a no-op db client when true
	EnvMockDb = "APP_MOCK_DB"
	// EnvMockDbDefault is the default value for EnvMockDb
	EnvMockDbDefault = false

	// EnvMongoHosts is a comma-separated list of mongo hosts to connect to.\
	// This property is REQUIRED!
	EnvMongoHosts = "MONGO_HOSTS"

	// EnvMongoPort is the port to use to connect to mongo
	EnvMongoPort = "MONGO_PORT"
	// EnvMongoPortDefault is the default value for EnvMongoPort
	EnvMongoPortDefault = "27017"

	// EnvMongoUsername is the auth username for mongo
	EnvMongoUsername = "MONGO_USERNAME"
	// EnvMongoUsernameDefault is the default value for EnvMongoUsername
	EnvMongoUsernameDefault = ""

	// EnvMongoPassword is the auth password for mongo
	EnvMongoPassword = "MONGO_PASSWORD"
	// EnvMongoPasswordDefault is the default value for EnvMongoPassword
	EnvMongoPasswordDefault = ""

	// EnvAuthSecret is the secret used during jwt verification.
	// REQUIRED when AuthEnabled is set.
	EnvAuthSecret = "AUTH_SECRET"
)

// Config holds application-level configuration
type Config struct {
	AuthEnabled   bool
	MockDb        bool
	MongoHosts    string
	MongoPort     string
	MongoUsername string
	MongoPassword string
	AuthSecret    string
}

// NewConfigFromEnv creates a new Config from environment variables
func NewConfigFromEnv() Config {
	loadEnvFile()
	return Config{
		AuthEnabled:   GetenvBool(EnvAuthEnabled, EnvAuthEnabledDefault),
		MockDb:        GetenvBool(EnvMockDb, EnvMockDbDefault),
		MongoHosts:    Getenv(EnvMongoHosts, ""),
		MongoPort:     Getenv(EnvMongoPort, EnvMongoPortDefault),
		MongoUsername: Getenv(EnvMongoUsername, EnvMongoUsernameDefault),
		MongoPassword: Getenv(EnvMongoPassword, EnvMongoPasswordDefault),
		AuthSecret:    Getenv(EnvAuthSecret, ""),
	}
}

func loadEnvFile() {
	err := godotenv.Load()
	if err == nil {
		log.Print("Loaded vars from .env")
	}
}

func authEnabled() bool {
	defaultVal := true
	val := os.Getenv(EnvAuthEnabled)
	if val == "" {
		return defaultVal
	}

	parsedVal, err := strconv.ParseBool(val)
	if err != nil {
		return defaultVal
	}

	return parsedVal
}

func mockDb() bool {
	defaultVal := false
	val := os.Getenv(EnvMockDb)
	if val == "" {
		return defaultVal
	}

	parsedVal, err := strconv.ParseBool(val)
	if err != nil {
		return defaultVal
	}

	return parsedVal
}
