package router

import (
	"github.com/go-chi/chi/v5/middleware"
)

func NewCompressMiddleware(level int, types ...string) Middleware {
	m := Middleware{
		Name:    "Compress Middleware",
		Handler: middleware.Compress(level, types...),
	}

	return m
}
