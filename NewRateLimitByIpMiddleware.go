package router

import (
	"strconv"
	"time"

	"github.com/go-chi/httprate"
)

func NewRateLimitByIpMiddleware(requestLimit int, seconds int) Middleware {
	m := Middleware{
		Name:    "Limit " + strconv.Itoa(requestLimit) + "r/" + strconv.Itoa(seconds) + "s By IP Middleware",
		Handler: httprate.LimitByIP(requestLimit, time.Duration(seconds)*time.Second),
	}

	return m
}
