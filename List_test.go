package router

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestList(t *testing.T) {
	checkUserAuthenticatedMiddleware := Middleware{
		Name: "Check if User is Authenticated",
	}

	middleware1 := Middleware{
		Name: "middleware1",
	}

	middleware2 := Middleware{
		Name: "middleware2",
	}

	globalMiddleware1 := Middleware{
		Name: "global_middleware1",
	}

	routes := []Route{
		// Example of simple "Hello world" endpoint
		{
			Name: "Home",
			Path: "/",
			Handler: func(w http.ResponseWriter, r *http.Request) string {
				return "Hello world"
			},
			Middlewares: []Middleware{middleware1},
		},
		{
			Name:    "Example",
			Path:    "/example",
			Methods: []string{http.MethodGet, http.MethodPost},
			Handler: func(w http.ResponseWriter, r *http.Request) string {
				return "Hello, World!"
			},
			Middlewares: []Middleware{middleware1, middleware2},
		},
		// Example of POST route
		{
			Name:    "Submit Form",
			Path:    "/form-submit",
			Methods: []string{http.MethodPost},
			Handler: func(w http.ResponseWriter, r *http.Request) string {
				return "Form submitted"
			},
		},
		// Example of route with local middlewares
		{
			Name:        "User Dashboard",
			Path:        "/user/dashboard",
			Middlewares: []Middleware{checkUserAuthenticatedMiddleware},
			Handler: func(w http.ResponseWriter, r *http.Request) string {
				return "Welcome to your dashboard"
			},
		},
		{
			Name: "Catch All. Page Not Found",
			Path: "/*",
			Handler: func(w http.ResponseWriter, r *http.Request) string {
				return "Page not found"
			},
		},
	}

	globalMiddlewares := []Middleware{globalMiddleware1}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	List(globalMiddlewares, routes)
	t.Log(buf.String())
}
