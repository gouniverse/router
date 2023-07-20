package router

import "net/http"

// handle applies a list of middlewares to an HTTP handler.
//
// The function takes an `http.Handler` and a slice of functions that
// transform an `http.Handler` and returns an `http.Handler`. It applies
// each middleware function to the original handler in the order they are
// provided, returning the final transformed handler.
//
// Parameters:
//   - `handler` (http.Handler): The original HTTP handler to be transformed.
//   - `middlewares` ([]func(http.Handler) http.Handler): A slice of functions
//     that transform an `http.Handler` and returns an `http.Handler`. The
//     middlewares are applied to the original handler in the order they are
//     provided.
//
// Returns:
//   - (http.Handler): The final transformed HTTP handler after applying
//     all the middlewares.
func handle(handler http.Handler, middlewares []func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
