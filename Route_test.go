package router

import (
	"net/http"
	"testing"
)

func TestRoute(t *testing.T) {
	route := Route{
		Path:    "/example",
		Methods: []string{"GET"},
		Handler: func(w http.ResponseWriter, r *http.Request) string {
			return "Hello, World!"
		},
		Middlewares: []func(http.Handler) http.Handler{},
	}

	if route.Path != "/example" {
		t.Errorf("Expected path to be '/example', got '%s'", route.Path)
	}

}
