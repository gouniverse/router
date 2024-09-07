package router

import (
	"net/http"

	"github.com/gouniverse/responses"
)

// NewRouter creates a new router with the given global middlewares
// and routes.
//
// Parameters:
// - globalMiddlewares: A slice of middlewares that will be applied to all routes.
// - routes: A slice of Route structs that define the routes for the router.
//
// Returns:
// - chi.Mux: The newly created chi router.
func NewRouter(globalMiddlewares []Middleware, routes []Route) *http.ServeMux {
	mux := http.NewServeMux()

	for _, route := range routes {
		route.Middlewares = append(globalMiddlewares, route.Middlewares...)

		middlewareHandlers := []func(http.Handler) http.Handler{}

		for _, middleware := range route.Middlewares {
			middlewareHandlers = append(middlewareHandlers, middleware.Handler)
		}

		if len(route.Methods) > 0 {
			for _, method := range route.Methods {
				if method == "all" {
					mux.Handle(route.Path, handle(responses.HTMLHandler(route.Handler), middlewareHandlers))
				} else {
					mux.Handle(method+" "+route.Path, handle(responses.HTMLHandler(route.Handler), middlewareHandlers))
				}
			}
		} else {
			mux.Handle(route.Path, handle(responses.HTMLHandler(route.Handler), middlewareHandlers))
		}
	}

	return mux
}
