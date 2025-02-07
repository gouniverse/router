package router

// import (
// 	"net/http"

// 	"github.com/gouniverse/responses"
// )

// // NewMultiDomainRouter creates a new router which routes with the given global middlewares
// // and routes for each domain.
// //
// // Parameters:
// // - globalMiddlewares: A slice of middlewares that will be applied to all routes.
// // - routes: A slice of Route structs that define the routes for the router.
// //
// // Returns:
// // - http.ServeMux: The newly created router.
// func NewMultiDomainRouter(globalMiddlewares []Middleware, routes []Route) *http.ServeMux {
// 	mux := http.NewServeMux()

// 	// Group routes by domain
// 	domainRoutes := map[string][]Route{}
// 	for _, route := range routes {
// 		domainRoutes[route.Domain] = append(domainRoutes[route.Domain], route)
// 	}

// 	for domain, routes := range domainRoutes {
// 		mux.Handle(domain+"/", createSubrouter(globalMiddlewares, routes))
// 	}

// 	return mux
// }

// func createSubrouter(globalMiddlewares []Middleware, routes []Route) *http.ServeMux {
// 	mux := http.NewServeMux()

// 	// Set a NotFound handler for the entire subrouter
// 	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	for _, route := range routes {
// 	// 		if matchRoute(route, r) {
// 	// 			// Handle the route
// 	// 			// ...
// 	// 			return
// 	// 		}
// 	// 	}

// 	// 	// Handle not found
// 	// 	http.NotFoundHandler().ServeHTTP(w, r)
// 	// })

// 	for _, route := range routes {

// 		// Prepend global middlewares
// 		route.Middlewares = append(globalMiddlewares, route.Middlewares...)

// 		// Create slice of middleware handlers
// 		middlewareHandlers := []func(http.Handler) http.Handler{}

// 		for _, middleware := range route.Middlewares {
// 			middlewareHandlers = append(middlewareHandlers, middleware.Handler)
// 		}

// 		// Handle route
// 		if len(route.Methods) > 0 {
// 			for _, method := range route.Methods {
// 				if method == "all" {
// 					mux.Handle(route.Path, handle(responses.HTMLHandler(route.Handler), middlewareHandlers))
// 				} else {
// 					mux.Handle(method+" "+route.Path, handle(responses.HTMLHandler(route.Handler), middlewareHandlers))
// 				}
// 			}
// 		} else {
// 			mux.Handle(route.Path, handle(responses.HTMLHandler(route.Handler), middlewareHandlers))
// 		}
// 	}

// 	// Set a NotFound handler for the entire subrouter
// 	// if !slices.Contains(paths, "/") {
// 	// 	mux.Handle("/", http.NotFoundHandler())
// 	// }

// 	return mux
// }

// func matchRoute(route Route, r *http.Request) bool {
// 	method := r.Method
// 	path := r.URL.Path

// 	if len(route.Methods) > 0 {
// 		for _, method := range route.Methods {
// 			if method == "all" {
// 				return true
// 			} else if method == r.Method {
// 				return true
// 			}
// 		}
// 	} else {
// 		return true
// 	}

// 	return false
// }
