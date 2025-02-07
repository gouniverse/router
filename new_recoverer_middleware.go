package router

import (
	"github.com/go-chi/chi/v5/middleware"
)

func NewRecovererMiddleware() Middleware {
	m := Middleware{
		Name:    "Recoverer Middleware",
		Handler: middleware.Recoverer,
	}

	return m

}
