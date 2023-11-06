package router

// RoutesPrependMiddlewares prepends the given middlewares to the Middlewares
// field of each Route in the provided slice.
//
// Parameters:
// - routes: A slice of Route structs representing the routes.
// - middlewares: A slice of middlewares to be prepended to each Route.
//
// Returns:
// - A slice of Route structs with the updated Middlewares field.
func RoutesPrependMiddlewares(routes []Route, middlewares []Middleware) []Route {
	for index := range routes {
		routes[index].Middlewares = middlewares
	}

	return routes
}
