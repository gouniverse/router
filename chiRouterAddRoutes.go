package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
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
func chiRouterAddRoutes(chiRouter chi.Router, routes []RouteInterface) {
	for _, route := range routes {
		middlewares := []func(http.Handler) http.Handler{}
		for _, middleware := range route.GetMiddlewares() {
			middlewares = append(middlewares, middleware.Handler)
		}

		if len(route.GetMethods()) > 0 {
			for _, method := range route.GetMethods() {
				chiRouter.Method(method, route.GetPath(), handleFunc(route.GetHandler(), middlewares))
			}
		} else {
			chiRouter.Handle(route.GetPath(), handleFunc(route.GetHandler(), middlewares))
		}
	}
}
