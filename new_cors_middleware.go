package router

import (
	"net/http"

	"github.com/go-chi/cors"
)

type CorsOptions struct {
	// AllowedOrigins is a list of origins a cross-domain request can be executed from.
	// If the special "*" value is present in the list, all origins will be allowed.
	// An origin may contain a wildcard (*) to replace 0 or more characters
	// (i.e.: http://*.domain.com). Usage of wildcards implies a small performance penalty.
	// Only one wildcard can be used per origin.
	// Default value is ["*"]
	AllowedOrigins []string

	// AllowOriginFunc is a custom function to validate the origin. It takes the origin
	// as argument and returns true if allowed or false otherwise. If this option is
	// set, the content of AllowedOrigins is ignored.
	AllowOriginFunc func(r *http.Request, origin string) bool

	// AllowedMethods is a list of methods the client is allowed to use with
	// cross-domain requests. Default value is simple methods (HEAD, GET and POST).
	AllowedMethods []string

	// AllowedHeaders is list of non simple headers the client is allowed to use with
	// cross-domain requests.
	// If the special "*" value is present in the list, all headers will be allowed.
	// Default value is [] but "Origin" is always appended to the list.
	AllowedHeaders []string

	// ExposedHeaders indicates which headers are safe to expose to the API of a CORS
	// API specification
	ExposedHeaders []string

	// AllowCredentials indicates whether the request can include user credentials like
	// cookies, HTTP authentication or client side SSL certificates.
	AllowCredentials bool

	// MaxAge indicates how long (in seconds) the results of a preflight request
	// can be cached
	MaxAge int

	// OptionsPassthrough instructs preflight to let other potential next handlers to
	// process the OPTIONS method. Turn this on if your application handles OPTIONS.
	OptionsPassthrough bool

	// Debugging flag adds additional output to debug server side CORS issues
	Debug bool
}

// NewCorsMiddleware creates a new CORS middleware
//
// Example:
// <code>
//
//	router.NewCorsMiddleware(router.corsOptions{
//	    // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
//	    AllowedOrigins: []string{"https://*", "http://*"},
//	    // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
//	    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
//	    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
//	    ExposedHeaders:   []string{"Link"},
//	    AllowCredentials: false,
//	    MaxAge:           300, // Maximum value not ignored by any of major browsers
//	})
//
// </code>
func NewCorsMiddleware(options CorsOptions) Middleware {
	opts := cors.Options{
		AllowedOrigins:   options.AllowedOrigins,
		AllowOriginFunc:  options.AllowOriginFunc,
		AllowedMethods:   options.AllowedMethods,
		AllowedHeaders:   options.AllowedHeaders,
		ExposedHeaders:   options.ExposedHeaders,
		AllowCredentials: options.AllowCredentials,
		MaxAge:           options.MaxAge,
	}

	m := Middleware{
		Name:    "CORS Middleware",
		Handler: cors.Handler(opts),
	}

	return m
}
