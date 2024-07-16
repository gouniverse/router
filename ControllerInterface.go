package router

import "net/http"

type ControllerInterface interface {
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
