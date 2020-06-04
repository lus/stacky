package config

import (
	"os"

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

	// Load the current application configuration
	CurrentConfig = &Config{
		MongoDBURI:      os.Getenv("STACKY_MONGODB_URI"),
		MongoDBDatabase: os.Getenv("STACKY_MONGODB_DATABASE"),
	}
	return nil
}
