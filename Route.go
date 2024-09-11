package router

import (
	"net/http"
)

type Route struct {
	// Domain      string
	Path        string
	Methods     []string // optional, default all methods
	Handler     func(w http.ResponseWriter, r *http.Request) string
	Middlewares []Middleware
	Name        string // optional, default empty string
}
