package router

import (
	"strconv"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func NewTimeoutMiddleware(seconds int) Middleware {
	secondsStr := strconv.Itoa(seconds)

	m := Middleware{
		Name:    "Timeout " + secondsStr + "s Middleware",
		Handler: middleware.Timeout(time.Second * time.Duration(seconds)),
	}

	return m
}
