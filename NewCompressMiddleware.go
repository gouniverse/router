package router

import (
	"github.com/go-chi/chi/v5/middleware"
)

// NewCompressMiddleware compresses response
// body of a given content types to a data format based
// on Accept-Encoding request header. It uses a given
// compression level.
//
// NOTE: make sure to set the Content-Type header on your response
// otherwise this middleware will not compress the response body. For ex, in
// your handler you should set w.Header().Set("Content-Type", http.DetectContentType(yourBody))
// or set it manually.
//
// Passing a compression level of 5 is sensible value
func NewCompressMiddleware(level int, types ...string) Middleware {
	m := Middleware{
		Name:    "Compress Middleware",
		Handler: middleware.Compress(level, types...),
	}

	return m
}
