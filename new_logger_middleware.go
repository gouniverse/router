package router

import (
	"github.com/go-chi/chi/v5/middleware"
)

func NewLoggerMiddleware() Middleware {
	m := Middleware{
		Name:    "Logger Middleware",
		Handler: middleware.Logger,
	}

	return m
}
