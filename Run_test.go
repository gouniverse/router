package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestRun(t *testing.T) {
	// Create a new Chi router
	router := chi.NewRouter()

	// Define a sample route
	route := Route{
		Path:    "/example",
		Methods: []string{"GET"},
		Handler: func(w http.ResponseWriter, r *http.Request) string {
			return "Hello, World!"
		},
		Middlewares: []func(http.Handler) http.Handler{},
	}

	// Create a slice of routes
	routes := []Route{route}

	// Run the Run function
	Run(router, routes)

	// Create a mock request and response for testing
	req := httptest.NewRequest("GET", "/example", nil)
	rec := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(rec, req)

	// Assert the response body
	expectedBody := "Hello, World!"
	if rec.Body.String() != expectedBody {
		t.Errorf("Expected response body '%s', got '%s'", expectedBody, rec.Body.String())
	}
}
