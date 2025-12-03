package api

import (
	"http-template/middlewares"
	"net/http"
)

// handleHealthGET returns "OK"
func handleHealthGET(ctx *middlewares.AppContext) {
	ctx.WriteText(http.StatusOK, "OK")
}
