package router

import (
	"github.com/go-chi/chi/v5/middleware"
)

func NewGetHeadMiddleware() Middleware {
	m := Middleware{
		Name:    "GetHead Middleware",
		Handler: middleware.GetHead,
	}

	return m
}
