package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewNakedDomainToWwwMiddleware(t *testing.T) {
	tests := []struct {
		name        string
		host        string
		exclude     []string
		wantCode    int
		redirectUrl string
		body        string
	}{
		{"Redirect Naked Domain", "example.com", nil, http.StatusPermanentRedirect, "https://www.example.com/", `<a href="https://www.example.com/">Permanent Redirect</a>.`},
		{"No Redirect for WWW Domain", "www.example.com", nil, http.StatusOK, "", `No redirect for "www.example.com"`},
		{"No Redirect for Excluded Domain", "example.com", []string{"example.com"}, http.StatusOK, "", `No redirect for "example.com"`},
		{"No Redirect for Localhost", "localhost", []string{"localhost"}, http.StatusOK, "", `No redirect for "localhost"`},
		{"No Redirect for 127.0.0.1", "127.0.0.1", []string{"127.0.0.1"}, http.StatusOK, "", `No redirect for "127.0.0.1"`},
		{"No Redirect for 127.0.0.1 with Port", "127.0.0.1:8080", []string{"127.0.0.1"}, http.StatusOK, "", `No redirect for "127.0.0.1`},
		{"Handle Empty Host Excludes List", "example.com", []string{}, http.StatusPermanentRedirect, "https://www.example.com/", `<a href="https://www.example.com/">Permanent Redirect</a>.`},
		{"Handle Invalid Hostname", "", nil, http.StatusOK, "", `No redirect for ""`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			req.Host = tt.host
			rec := httptest.NewRecorder()

			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				text := fmt.Sprintf("No redirect for %q", r.Host)
				w.Write([]byte(text))
			})

			middleware := NewNakedDomainToWwwMiddleware(tt.exclude)
			middleware.Handler(next).ServeHTTP(rec, req)

			code := rec.Code
			host := rec.Header().Get("Location")
			body := rec.Body.String()

			if code != tt.wantCode {
				t.Errorf("Expected status code %d, got %d", tt.wantCode, code)
			}
			if host != tt.redirectUrl {
				t.Errorf("Expected host %q, got %q", tt.redirectUrl, host)
			}

			if !strings.Contains(body, tt.body) {
				t.Errorf("Expected body %q, got %q", tt.body, body)
			}
		})
	}
}

// 	// Create a mock request and response for testing
// 	req := httptest.NewRequest("GET", "/example", nil)
// 	rec := httptest.NewRecorder()

// 	// Create a new NakedDomainToWwwMiddleware
// 	middleware := NewNakedDomainToWwwMiddleware([]string{"localhost"})

// 	// Serve the request using the middleware
// 	middleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello, World!"))
// 	})).ServeHTTP(rec, req)

// 	// Assert the response
// 	if rec.Result().StatusCode != 200 {
// 		t.Errorf("Expected status code %d, got %d", 200, rec.Result().StatusCode)
// 	}
// }
