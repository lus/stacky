package v1

import (
	"github.com/Lukaesebrot/stacky/config"
	"github.com/valyala/fasthttp"
)

// authenticated works as a wrapper for authenticated endpoints
func authenticated(endpoint fasthttp.RequestHandler, minLevel int) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		// Define the API token
		token := string(ctx.QueryArgs().Peek("token"))

		// Validate the API token
		validTokens := config.CurrentConfig.AuthKeys
		if validTokens[token] < minLevel {
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			ctx.SetBody(errorResponse(fasthttp.StatusUnauthorized, "you are not authorized to use this endpoint", nil))
			return
		}

		// call the endpoint
		endpoint(ctx)
	}
}
