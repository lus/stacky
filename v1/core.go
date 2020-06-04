package v1

import routing "github.com/fasthttp/router"

// Initialize initializes the '/api/v1' routes
func Initialize(router *routing.Router) {
	// Initialize the API endpoints
	router.GET("/stacks/{name}", authenticated(endGetStack, 1))
	router.PUT("/stacks", authenticated(endPutStack, 2))
	router.DELETE("/stacks/{name}", authenticated(endDeleteStack, 2))
	router.GET("/stacks/{name}/hosts/best", authenticated(endGetStackBestHost, 1))
	router.GET("/stacks/{name}/hosts", authenticated(endGetStackHosts, 1))
	router.PUT("/stacks/{name}/hosts", authenticated(endPutStackHost, 2))
	router.DELETE("/stacks/{name}/hosts", authenticated(endDeleteStackHost, 2))
}
