package v1

import (
	"github.com/Lukaesebrot/stacky/stacks"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
)

// endDeleteStack handles the DELETE '/stacks/{name}' endpoint
func endDeleteStack(ctx *fasthttp.RequestCtx) {
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

	// Delete the current stack
	err = stack.Delete()
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody(errorResponse(fasthttp.StatusInternalServerError, err.Error(), nil))
		return
	}

	// Respond with a success message
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(successResponse(fasthttp.StatusOK, "the given stack got deleted", nil))
}
