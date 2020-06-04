package api

import (
	v1 "github.com/Lukaesebrot/stacky/v1"
	routing "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// Serve serves the REST API
func Serve() {
	// Define the API routes
	router := routing.New()
	v1.Initialize(router.Group("/api/v1"))

	// Start the API server
	panic(fasthttp.ListenAndServe(":8080", router.Handler))
}
