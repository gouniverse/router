package router

import (
	"net/http"
	"testing"
)

func TestRoutesPrependMiddleware(t *testing.T) {
	routes := []RouteInterface{
		&Route{
			Path: "/route1",
			HTMLHandler: func(w http.ResponseWriter, r *http.Request) string {
				return "Hello, World 1!"
			},
		},
		&Route{
			Path: "/route2",
			HTMLHandler: func(w http.ResponseWriter, r *http.Request) string {
				return "Hello, World 2!"
			},
		},
	}

	// 1. Prepend 2 middlewares
	routes = RoutesPrependMiddlewares(routes, []Middleware{
		{
			Name: "Middleware 1",
		},
		{
			Name: "Middleware 2",
		},
	})

	if len(routes) != 2 {
		t.Error("Expected routes length to be 2, got", len(routes))
	}

	for _, route := range routes {
		if len(route.GetMiddlewares()) != 2 {
			t.Error("Expected route middleware length to be 2, got", len(routes[0].GetMiddlewares()))
		}

		if route.GetMiddlewares()[0].Name != "Middleware 1" {
			t.Error("Expected route middleware 1 name to be Middleware 1, got", route.GetMiddlewares()[0].Name)
		}

		if route.GetMiddlewares()[1].Name != "Middleware 2" {
			t.Error("Expected route middleware 2 name to be Middleware 2, got", route.GetMiddlewares()[1].Name)
		}
	}

	// 2. Prepend 2 more middlewares
	routes = RoutesPrependMiddlewares(routes, []Middleware{
		{
			Name: "Middleware 3",
		},
		{
			Name: "Middleware 4",
		},
	})

	for _, route := range routes {
		if len(route.GetMiddlewares()) != 4 {
			t.Error("Expected route middleware length to be 4, got", len(routes[0].GetMiddlewares()))
		}

		if route.GetMiddlewares()[0].Name != "Middleware 3" {
			t.Error("Expected route middleware 1 name to be Middleware 3, got", route.GetMiddlewares()[0].Name)
		}

		if route.GetMiddlewares()[1].Name != "Middleware 4" {
			t.Error("Expected route middleware 2 name to be Middleware 4, got", route.GetMiddlewares()[1].Name)
		}

		if route.GetMiddlewares()[2].Name != "Middleware 1" {
			t.Error("Expected route middleware 3 name to be Middleware 1, got", route.GetMiddlewares()[2].Name)
		}

		if route.GetMiddlewares()[3].Name != "Middleware 2" {
			t.Error("Expected route middleware 4 name to be Middleware 2, got", route.GetMiddlewares()[3].Name)
		}
	}
}
