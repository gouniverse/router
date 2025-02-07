package router

import (
	"github.com/go-chi/chi/v5/middleware"
)

// NewHeartbeatMiddleware endpoint middleware useful to setting up a path like
// `/ping` that load balancers or uptime testing external services
// can make a request before hitting any routes. It's also convenient
// to place this above ACL middlewares as well.
func NewHeartbeatMiddleware(endpoint string) Middleware {
	m := Middleware{
		Name:    "Heartbeat at " + endpoint + " Middleware",
		Handler: middleware.Heartbeat(endpoint),
	}

	return m
}
