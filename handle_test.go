package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandle(t *testing.T) {
	// Define a simple handler function
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	// Create a mock request and response for testing
	req := httptest.NewRequest("GET", "/example", nil)
	rec := httptest.NewRecorder()

	// Create a handler using the handle function
	handler := handle(http.HandlerFunc(handlerFunc), []func(http.Handler) http.Handler{
		// Define your middleware functions here
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Perform middleware logic
				// For example, set a custom header
				w.Header().Set("X-Custom-Header", "Hello")
				// Call the next handler in the chain
				next.ServeHTTP(w, r)
			})
		},
	})

	// Serve the request using the handler
	handler.ServeHTTP(rec, req)

	// Assert the response status code
	if rec.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Result().StatusCode)
	}

	// Assert the custom header set by the middleware
	if rec.Header().Get("X-Custom-Header") != "Hello" {
		t.Errorf("Expected custom header value 'Hello', got '%s'", rec.Header().Get("X-Custom-Header"))
	}
}
