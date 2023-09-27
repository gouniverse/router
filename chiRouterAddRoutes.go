package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gouniverse/responses"
)

// chiRouterAddRoutes is a function that takes a chi.Router and a slice of Routes
// as input and sets up the routes on the chiRouter.
//
// Parameters:
// - chiRouter: The chi.Router to set up the routes on.
// - routes: The slice of Route containing the routes to be set up.
//
// Returns:
// - Nothing
func chiRouterAddRoutes(chiRouter chi.Router, routes []Route) {
	for _, route := range routes {
		middlewares := []func(http.Handler) http.Handler{}
		for _, middleware := range route.Middlewares {
			middlewares = append(middlewares, middleware.Handler)
		}

		if len(route.Methods) > 0 {
			for _, method := range route.Methods {
				if method == "all" {
					chiRouter.Handle(route.Path, handle(responses.HTMLHandler(route.Handler), middlewares))
				} else {
					chiRouter.Method(method, route.Path, handle(responses.HTMLHandler(route.Handler), middlewares))
				}
			}
		} else {
			chiRouter.Handle(route.Path, handle(responses.HTMLHandler(route.Handler), middlewares))
		}
	}
}
