package app

import (
	"os"
	"strconv"

	"github.com/jcosentino11/meals/internal/mongo"
)

const (
	// EnvFirebaseKeyFile is the name of env var that
	// points to the firebase secret key
	EnvFirebaseKeyFile    = "APP_FIREBASE_KEY_FILE"
	// EnvAuthEnabled is the name of env var that
	// controls whether auth is enabled for endpoints
	EnvAuthEnabled = "APP_AUTH_ENABLED"
)

// Config holds application-level configuration
type Config struct {
	AuthEnabled bool
	AuthFile    string
	MongoConfig mongo.Config
}

// NewConfigFromEnv creates a new Config from environment variables
func NewConfigFromEnv() Config {
	return Config{
		AuthEnabled: authEnabled(),
		AuthFile:    os.Getenv(EnvFirebaseKeyFile),
		MongoConfig: mongo.NewConfigFromEnv(),
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
