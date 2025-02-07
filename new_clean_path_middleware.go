package router

import (
	"github.com/go-chi/chi/v5/middleware"
)

// NewCleanPathMiddleware will clean out double slash mistakes from a user's request path.
// For example, if a user requests /users//1 or //users////1 will both be treated as: /users/1
func NewCleanPathMiddleware() Middleware {
	m := Middleware{
		Name:    "Clean Path Middleware",
		Handler: middleware.CleanPath,
	}

	return m
}
