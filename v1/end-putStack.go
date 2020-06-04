package v1

import (
	"strings"

	"github.com/Lukaesebrot/stacky/stacks"
	"github.com/valyala/fasthttp"
)

// endPutStack handles the PUT '/stacks' endpoint
func endPutStack(ctx *fasthttp.RequestCtx) {
	// Validate the inputs
	name := string(ctx.QueryArgs().Peek("name"))
	hosts := strings.Split(string(ctx.QueryArgs().Peek("host")), ",")
	if name == "" {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody(errorResponse(fasthttp.StatusBadRequest, "you have to specify a stack name", nil))
		return
	}

	// Create the stack
	stack, err := stacks.Create(name, hosts...)
	if err != nil {
		if err == stacks.ErrStackAlreadyExists {
			ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
			ctx.SetBody(errorResponse(fasthttp.StatusUnprocessableEntity, err.Error(), nil))
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
