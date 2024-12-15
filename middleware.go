package router

import "net/http"

type Middleware struct {
	Name    string
	Handler func(http.Handler) http.Handler
}
