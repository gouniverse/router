package router

import (
	"github.com/go-chi/chi/v5/middleware"
)

// NewGetHeadMiddleware automatically route undefined HEAD requests to GET handlers.
func NewGetHeadMiddleware() Middleware {
	m := Middleware{
		Name:    "GetHead Middleware",
		Handler: middleware.GetHead,
	}

	return m
}
