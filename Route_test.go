package router

import (
	"net/http"
	"testing"
)

func TestRoute(t *testing.T) {
	route := Route{
		Path:    "/example",
		Methods: []string{"GET"},
		HTMLHandler: func(w http.ResponseWriter, r *http.Request) string {
			return "Hello, World!"
		},
		Middlewares: []Middleware{},
	}

	if route.Path != "/example" {
		t.Errorf("Expected path to be '/example', got '%s'", route.Path)
	}

	if len(route.Methods) != 1 {
		t.Errorf("Expected methods to be 1, got %d", len(route.Methods))
	}

	if route.Methods[0] != "GET" {
		t.Errorf("Expected method to be 'GET', got '%s'", route.Methods[0])
	}

	if route.HTMLHandler == nil {
		t.Error("Expected handler to be non-nil")
	}

	if len(route.Middlewares) != 0 {
		t.Error("Expected middlewares to be empty")
	}
}
