package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// chiRouterAddMiddlewares is a function that takes a chi.Router and a slice
// of middlewares as input and sets up the middlewares on the chiRouter.
//
// Parameters:
// - chiRouter: The chi.Router to set up the routes on.
// - middlewares: The slice of middlewares containing the middlewares to be set up.
//
// Returns:
// - Nothing
func chiRouterAddMiddlewares(chiRouter chi.Router, middlewares []func(http.Handler) http.Handler) {
	for _, middleware := range middlewares {
		chiRouter.Use(middleware)
	}
}
