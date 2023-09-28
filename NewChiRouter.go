package router

import (
	"github.com/go-chi/chi/v5"
)

// NewChiRouter creates a new chi router with the given global middlewares
// and routes.
//
// Parameters:
// - globalMiddlewares: A slice of middlewares that will be applied to all routes.
// - routes: A slice of Route structs that define the routes for the router.
//
// Returns:
// - chi.Mux: The newly created chi router.
func NewChiRouter(globalMiddlewares []Middleware, routes []Route) *chi.Mux {
	chiRouter := chi.NewRouter()
	chiRouterAddMiddlewares(chiRouter, globalMiddlewares)
	chiRouterAddRoutes(chiRouter, routes)
	return chiRouter
}
