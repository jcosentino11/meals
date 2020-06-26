package app

import (
	"os"
	"strconv"

	"github.com/jcosentino11/meals/internal/mongo"
)

const (
	// EnvFirebaseKeyFile is the name of env var that
	// points to the firebase secret key
	EnvFirebaseKeyFile = "APP_FIREBASE_KEY_FILE"
	// EnvAuthEnabled is the name of env var that
	// controls whether auth is enabled for endpoints
	EnvAuthEnabled = "APP_AUTH_ENABLED"
	// EnvMockDb will use a no-op db client when true
	EnvMockDb = "APP_MOCK_DB"
)

// Config holds application-level configuration
type Config struct {
	AuthEnabled bool
	AuthFile    *string
	MockDb      bool
	MongoConfig mongo.Config
}

// NewConfigFromEnv creates a new Config from environment variables
func NewConfigFromEnv() Config {
	return Config{
		AuthEnabled: authEnabled(),
		AuthFile:    authFile(),
		MockDb:      mockDb(),
		MongoConfig: mongo.NewConfigFromEnv(),
	}
}

func authFile() *string {
	prop, exists := os.LookupEnv(EnvFirebaseKeyFile)
	if exists {
		return &prop
	}
	return nil
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
