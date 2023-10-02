package router

import (
	"github.com/go-chi/chi/v5/middleware"
)

func NewRedirectSlashesMiddleware() Middleware {
	m := Middleware{
		Name:    "Redirect Slashes Middleware",
		Handler: middleware.RedirectSlashes,
	}

	return m
}
