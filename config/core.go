package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/Lukaesebrot/stacky/static"
	"github.com/joho/godotenv"
)

// CurrentConfig defines the current application configuration
var CurrentConfig *Config

// Load loads the current application configuration
func Load() error {
	// Load the .env file if the application runs in development mode
	if static.Mode == "dev" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	// Read the auth keys
	authKeysRaw := strings.Split(os.Getenv("STACKY_AUTH_KEYS"), ",")
	authKeys := make(map[string]int, len(authKeysRaw))
	for _, authKeyRaw := range authKeysRaw {
		split := strings.Split(authKeyRaw, ":")
		i, err := strconv.Atoi(strings.Join(split[1:], ""))
		if err != nil {
			return err
		}
		authKeys[split[0]] = i
	}

	// Load the current application configuration
	CurrentConfig = &Config{
		MongoDBURI:      os.Getenv("STACKY_MONGODB_URI"),
		MongoDBDatabase: os.Getenv("STACKY_MONGODB_DATABASE"),
		AuthKeys:        authKeys,
	}
	return nil
}
