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

func TestHandleWithMiddleware(t *testing.T) {
	// Create a mock request and response for testing
	req := httptest.NewRequest("GET", "/example", nil)
	rec := httptest.NewRecorder()

	// Define a middleware function that sets a custom header
	middleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Perform middleware logic
			// For example, set a custom header
			w.Header().Set("X-Custom-Header", "Hello")
			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}

	// Define a handler function
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		// Assert that the middleware has been invoked
		if header := w.Header().Get("X-Custom-Header"); header != "Hello" {
			t.Errorf("Expected custom header value 'Hello', got '%s'", header)
		}
	}

	// Create a handler using the handle function with the middleware
	handler := handle(http.HandlerFunc(handlerFunc), []func(http.Handler) http.Handler{middleware})

	// Serve the request using the handler
	handler.ServeHTTP(rec, req)
}

func TestHandleWithMultipleMiddlewares(t *testing.T) {
	// Create a mock request and response for testing
	req := httptest.NewRequest("GET", "/example", nil)
	rec := httptest.NewRecorder()

	// Define the first middleware function that sets a custom header
	middleware1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Perform middleware logic
			// For example, set a custom header
			w.Header().Set("X-Custom-Header1", "Hello")

			w.Write([]byte("Middleware1 says hello! "))
			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}

	// Define the second middleware function that modifies the response body
	middleware2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Perform middleware logic
			// For example, modify the response body
			w.Header().Set("X-Custom-Header2", "World")

			w.Write([]byte("Middleware2 says world! "))
			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}

	// Define the handler function
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		// Assert that the first middleware has been invoked
		if header := w.Header().Get("X-Custom-Header1"); header != "Hello" {
			t.Errorf("Expected custom header value 'Hello', got '%s'", header)
		}

		// Assert that the second middleware has been invoked
		if body := rec.Body.String(); body != "Middleware2 says world! Middleware1 says hello! " {
			t.Errorf("Expected response body 'Middleware2 says world! Middleware1 says hello! ', got '%s'", body)
		}

		// Write the final response body
		w.Write([]byte("Hello, World!"))
	}

	// Create a handler using the handle function with the middlewares
	handler := handle(http.HandlerFunc(handlerFunc), []func(http.Handler) http.Handler{middleware1, middleware2})

	// Serve the request using the handler
	handler.ServeHTTP(rec, req)

	// Assert the final response body
	if body := rec.Body.String(); body != "Middleware2 says world! Middleware1 says hello! Hello, World!" {
		t.Errorf("Expected response body 'Middleware2 says world! Middleware1 says hello! Hello, World!', got '%s'", body)
	}
}
