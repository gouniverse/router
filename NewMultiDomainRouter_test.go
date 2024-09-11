package router

// func TestNewRouterWithDomain(t *testing.T) {
// 	// Create sample routes
// 	routes := []Route{
// 		{

// 			Domain: "example.com",
// 			Path:   "/",
// 			Handler: func(w http.ResponseWriter, r *http.Request) string {
// 				w.WriteHeader(http.StatusOK)
// 				return "Domain: example.com, Method: GET, Path: /"
// 			},
// 		},
// 		{
// 			Domain: "example.com",
// 			Path:   "/hello",
// 			Handler: func(w http.ResponseWriter, r *http.Request) string {
// 				w.WriteHeader(http.StatusOK)
// 				return "Domain: example.com, Method: GET, Path: /hello"
// 			},
// 		},
// 		{
// 			Domain: "another.com",
// 			Path:   "/world",
// 			Handler: func(w http.ResponseWriter, r *http.Request) string {
// 				w.WriteHeader(http.StatusOK)
// 				return "Domain: another.com, Method: GET, Path: /world"
// 			},
// 		},
// 		{
// 			Domain: "another.com",
// 			Path:   "/notfound",
// 			Handler: func(w http.ResponseWriter, r *http.Request) string {
// 				w.WriteHeader(http.StatusOK)
// 				return "Domain: another.com, Method: GET, Path: /world"
// 			},
// 		},
// 	}

// 	// Create the router
// 	router := NewMultiDomainRouter([]Middleware{}, routes)

// 	// Test requests for different domains and paths
// 	tests := []struct {
// 		Domain     string
// 		Method     string
// 		Path       string
// 		StatusCode int
// 	}{
// 		{"example.com", "GET", "/", http.StatusOK},
// 		{"example.com", "GET", "/hello", http.StatusOK},
// 		{"another.com", "GET", "/world", http.StatusOK},
// 		{"example.com", "GET", "/notfound", http.StatusNotFound},
// 		{"another.com", "GET", "/", http.StatusNotFound},
// 		{"another.com", "GET", "/notfound", http.StatusNotFound},
// 	}

// 	for _, test := range tests {
// 		req := httptest.NewRequest(test.Method, "http://"+test.Domain+test.Path, nil)
// 		w := httptest.NewRecorder()
// 		router.ServeHTTP(w, req)
// 		res := w.Result()

// 		if res.StatusCode != test.StatusCode {
// 			t.Errorf("At domain %s, method %s, path %s, expected status code %d, got %d, body: %s", test.Domain, test.Method, test.Path, test.StatusCode, res.StatusCode, w.Body.String())
// 		}
// 	}
// }
