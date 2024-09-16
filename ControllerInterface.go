package router

import "net/http"

// ControllerInterface is an interface for controllers with idiomatic behavior.
// It will not add any headers to the response by default.
type ControllerInterface interface {
	// Handler is the single entry point for the controller.
	//
	// Parameters:
	//   w - The http.ResponseWriter object.
	//   r - The http.Request object.
	//
	// Returns:
	//   void - No return value.
	Handler(w http.ResponseWriter, r *http.Request)
}

// HTMLControllerInterface is an interface for controllers that return HTML.
// It will automatically add the "Content-Type: text/html" header to the response.
type HTMLControllerInterface interface {
	// Handler is the single entry point for the controller.
	//
	// Parameters:
	//   w - The http.ResponseWriter object.
	//   r - The http.Request object.
	//
	// Returns:
	//   string - The response to be sent back to the client.
	Handler(w http.ResponseWriter, r *http.Request) string
}

// JSONControllerInterface is an interface for controllers that return JSON.
// It will automatically add the "Content-Type: application/json" header to the response.
type JSONControllerInterface interface {
	// Handler is the single entry point for the controller.
	//
	// Parameters:
	//   w - The http.ResponseWriter object.
	//   r - The http.Request object.
	//
	// Returns:
	//   string - The response to be sent back to the client.
	Handler(w http.ResponseWriter, r *http.Request) string
}

// TextControllerInterface is an interface for controllers that return plain text.
// It will automatically add the "Content-Type: text/plain" header to the response.
type TextControllerInterface interface {
	// Handler is the single entry point for the controller.
	//
	// Parameters:
	//   w - The http.ResponseWriter object.
	//   r - The http.Request object.
	//
	// Returns:
	//   string - The response to be sent back to the client.
	Handler(w http.ResponseWriter, r *http.Request) string
}
