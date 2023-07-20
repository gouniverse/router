package router

import (
	"net/http"
)

type Route struct {
	Path        string
	Methods     []string
	Handler     func(w http.ResponseWriter, r *http.Request) string
	Middlewares []func(http.Handler) http.Handler
}