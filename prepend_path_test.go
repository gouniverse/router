package router

import (
	"net/http"
	"testing"
)

func TestRoutesPrependPath(t *testing.T) {
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

	// 1. Prepend "/path1"
	routes = RoutesPrependPath(routes, "/path1")

	if len(routes) != 2 {
		t.Error("Expected routes length to be 2, got", len(routes))
	}

	if routes[0].GetPath() != "/path1/route1" {
		t.Error("Expected route 1 path to be /path1/route1, got", routes[0].GetPath())
	}

	if routes[1].GetPath() != "/path1/route2" {
		t.Error("Expected route 1 path to be /path1/route1, got", routes[1].GetPath())
	}

	// 2. Prepend "/path2"
	routes = RoutesPrependPath(routes, "/path2")

	if len(routes) != 2 {
		t.Error("Expected routes length to be 2, got", len(routes))
	}

	if routes[0].GetPath() != "/path2/path1/route1" {
		t.Error("Expected route 1 path to be /path2/path1/route1, got", routes[0].GetPath())
	}

	if routes[1].GetPath() != "/path2/path1/route2" {
		t.Error("Expected route 1 path to be /path2/path1/route1, got", routes[1].GetPath())
	}
}
