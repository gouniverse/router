package router

import (
	"net/http"

	"golang.org/x/exp/slices"
)

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
	slices.Reverse(middlewares)
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

// handleFunc applies a list of middlewares to an HTTP handler function.
//
// The function takes an `http.HandlerFunc` and a slice of functions that
// transform an `http.HandlerFunc` and returns an `http.HandlerFunc`. It applies
// each middleware function to the original handler function in the order they
// are provided, returning the final transformed handler function.
//
// Parameters:
//   - `handler` (http.HandlerFunc): The original HTTP handler function to be
//     transformed.
//   - `middlewares` ([]func(http.HandlerFunc) http.HandlerFunc): A slice of
//     functions that transform an `http.HandlerFunc` and returns an
//     `http.HandlerFunc`. The middlewares are applied to the original handler
//     function in the order they are provided.
//
// Returns:
//   - (http.HandlerFunc): The final transformed HTTP handler function after
//     applying all the middlewares.
func handleFunc(handlerFunc http.HandlerFunc, middlewares []func(http.Handler) http.Handler) http.Handler {
	slices.Reverse(middlewares)
	handler := handlerfuncToHandler(handlerFunc)
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func handlerfuncToHandler(hf http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hf(w, r)
	})
}
