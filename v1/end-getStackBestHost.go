package v1

import (
	"github.com/Lukaesebrot/stacky/stacks"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
)

// endGetStackBestHost handles the GET '/stacks/{name}/hosts/best' endpoint
func endGetStackBestHost(ctx *fasthttp.RequestCtx) {
	// Retrieve the stack
	stack, err := stacks.Retrieve(ctx.UserValue("name").(string))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			ctx.SetBody(errorResponse(fasthttp.StatusNotFound, "the requested stack couldn't be found", nil))
			return
		}
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody(errorResponse(fasthttp.StatusInternalServerError, err.Error(), nil))
		return
	}

	// Define the best host
	bestHost, err := stack.CalculateBestHost()
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusExpectationFailed)
		ctx.SetBody(errorResponse(fasthttp.StatusExpectationFailed, err.Error(), nil))
		return
	}

	// Respond with the best host
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(successResponse(fasthttp.StatusOK, "", bestHost))
}
