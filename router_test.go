package router

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewRouter(t *testing.T) {
	// Create a new router
	router := NewRouter([]Middleware{}, []Route{
		{
			Path:    "/example",
			Methods: []string{"GET"},
			HTMLHandler: func(w http.ResponseWriter, r *http.Request) string {
				return "Hello, World!"
			},
			Middlewares: []Middleware{},
		},
	})

	// Create a mock request and response for testing
	req := httptest.NewRequest("GET", "/example", nil)
	rec := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(rec, req)

	// Assert the response status code
	if rec.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Result().StatusCode)
	}

	// Assert the response body
	expectedBody := "Hello, World!"
	if rec.Body.String() != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, rec.Body.String())
	}
}

func TestNewRouterWithMiddleware(t *testing.T) {
	// Define a mock middleware function that modifies the response
	middleware := Middleware{
		Handler: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Modify the response by setting a custom header
				w.Header().Set("X-Middleware", "Invoked")
				// Call the next handler in the chain
				next.ServeHTTP(w, r)
			})
		},
	}

	router := NewRouter([]Middleware{}, []Route{
		{
			Path:    "/example",
			Methods: []string{"GET"},
			HTMLHandler: func(w http.ResponseWriter, r *http.Request) string {
				// Assert that the middleware has been invoked
				if header := w.Header().Get("X-Middleware"); header != "Invoked" {
					t.Errorf("Expected custom header value 'Invoked', got '%s'", header)
				}
				return "Hello, World!"
			},
			Middlewares: []Middleware{middleware},
		},
	})

	// Create a mock request and response for testing
	req := httptest.NewRequest("GET", "/example", nil)
	rec := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(rec, req)

	// Assert the response status code
	if rec.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Result().StatusCode)
	}

	// Assert the response body
	expectedBody := "Hello, World!"
	if rec.Body.String() != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, rec.Body.String())
	}

}

func TestNewRouterWithGlobalMiddleware(t *testing.T) {
	// Define a mock middleware function that modifies the response
	globalMiddleware := Middleware{
		Handler: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Modify the response by setting a custom header
				w.Header().Set("X-GlobalMiddleware", "Invoked")
				// Call the next handler in the chain
				next.ServeHTTP(w, r)
			})
		},
	}

	// Define a mock middleware function that modifies the response
	middleware := Middleware{
		Handler: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Modify the response by setting a custom header
				w.Header().Set("X-Middleware", "Invoked")
				// Call the next handler in the chain
				next.ServeHTTP(w, r)
			})
		},
	}

	router := NewRouter([]Middleware{globalMiddleware}, []Route{
		{
			Path:    "/example",
			Methods: []string{"GET"},
			HTMLHandler: func(w http.ResponseWriter, r *http.Request) string {
				// Assert that the global middleware has been invoked
				if header := w.Header().Get("X-GlobalMiddleware"); header != "Invoked" {
					t.Errorf("Expected custom header value 'Invoked', got '%s'", header)
				}
				// Assert that the middleware has been invoked
				if header := w.Header().Get("X-Middleware"); header != "Invoked" {
					t.Errorf("Expected custom header value 'Invoked', got '%s'", header)
				}
				return "Hello, World!"
			},
			Middlewares: []Middleware{middleware},
		},
	})

	// Create a mock request and response for testing
	req := httptest.NewRequest("GET", "/example", nil)
	rec := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(rec, req)

	// Assert the response status code
	if rec.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Result().StatusCode)
	}

	// Assert the response body
	expectedBody := "Hello, World!"
	if rec.Body.String() != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, rec.Body.String())
	}

}

// Add a new table test for the above

func TestNewRouterFull(t *testing.T) {
	table := []struct {
		Methods      []string
		Path         string
		ResponseCode int
		ExpectedBody string
	}{
		{Path: "/", ResponseCode: http.StatusNotFound, ExpectedBody: "404 page not found"},
		{Path: "/example", ResponseCode: http.StatusOK, ExpectedBody: "Hello, World!"},
		{Path: "/nonexisting", ResponseCode: http.StatusNotFound, ExpectedBody: "404 page not found"},
	}

	// Create a new router
	router := NewRouter([]Middleware{}, []Route{
		{
			Path:    "/example",
			Methods: []string{"GET"},
			HTMLHandler: func(w http.ResponseWriter, r *http.Request) string {
				return "Hello, World!"
			},
			Middlewares: []Middleware{},
		},
	})

	for _, testRow := range table {

		// Create a mock request and response for testing
		req := httptest.NewRequest("GET", testRow.Path, nil)
		rec := httptest.NewRecorder()

		// Serve the request using the router
		router.ServeHTTP(rec, req)

		// Assert the response status code
		if rec.Result().StatusCode != testRow.ResponseCode {
			t.Errorf("Expected status code %d, got %d", testRow.ResponseCode, rec.Result().StatusCode)
		}

		// Assert the response body
		expectedBody := testRow.ExpectedBody
		if strings.TrimSpace(rec.Body.String()) != expectedBody {
			t.Errorf("Expected body '%s', got '%s'", expectedBody, strings.TrimSpace(rec.Body.String()))
		}

	}

}
