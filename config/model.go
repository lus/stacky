package config

// Config represents an application configuration
type Config struct {
	MongoDBURI      string
	MongoDBDatabase string
	AuthKeys        map[string]int
}
