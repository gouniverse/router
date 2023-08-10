package router

// RoutesPrependPath prepends the given path to the path of each route
// in the provided slice.
//
// Parameters:
// - routes: a slice of Route structs representing the routes.
// - path: a string representing the path to prepend to each route's path.
//
// Returns:
// - a slice of Route structs with the updated paths.
func RoutesPrependPath(routes []Route, path string) []Route {
	for index, route := range routes {
		path := path + route.Path
		routes[index].Path = path
	}

	return routes
}
