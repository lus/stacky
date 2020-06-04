package main

import (
	"log"

	"github.com/Lukaesebrot/stacky/config"
	"github.com/Lukaesebrot/stacky/database"
)

func main() {
	// Load the current application configuration
	log.Println("Loading the application configuration...")
	err := config.Load()
	if err != nil {
		log.Panicln(err)
	}

	// Connect to the configured MongoDB host
	log.Println("Connecting to the configured MongoDB host...")
	err = database.Connect(config.CurrentConfig.MongoDBURI)
	if err != nil {
		log.Panicln(err)
	}
}
