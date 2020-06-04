package main

import (
	"log"

	"github.com/Lukaesebrot/stacky/config"
)

func main() {
	// Load the current application configuration
	err := config.Load()
	if err != nil {
		log.Panicln(err)
	}
}
