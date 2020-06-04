package v1

import (
	"github.com/Lukaesebrot/stacky/stacks"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
)

// endGetStack handles the GET '/stacks/{name}' endpoint
func endGetStack(ctx *fasthttp.RequestCtx) {
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

	// Respond with the stack information
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(successResponse(fasthttp.StatusOK, "", stack))
}
