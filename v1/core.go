package v1

import routing "github.com/fasthttp/router"

// Initialize initializes the '/api/v1' routes
func Initialize(router *routing.Router) {
	// Initialize the API endpoints
	router.GET("/stacks/{name}", authenticated(endGetStack, 1))
}
