package router

import (
	"fmt"
	"net/http"
	"strings"
)

// NewNakedDomainToWwwMiddleware will redirect a "www" subdomain to naked (non-www) domain
func NewNakedDomainToWwwMiddleware(hostExcludes []string) Middleware {
	m := Middleware{
		Name: "Naked Domain to WWW Middleware",
		Handler: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				nakedDomainToWww(next, hostExcludes).ServeHTTP(w, r)
			})
		},
	}

	return m
}

// NakedDomainToWWW is http middleware that ensures a naked domain is redirected to "www" subdomain and "https".
// `hostExcludes` is a list of host names to ignore, such as `localhost`.
func nakedDomainToWww(next http.Handler, hostExcludes []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host := strings.ToLower(r.Host)

		if len(host) == 0 {
			next.ServeHTTP(w, r)
			return
		}

		scheme := "https"

		if r.TLS != nil {
			scheme = "http"
		}

		redirectUrl := fmt.Sprintf(scheme+"://www.%s%s", r.Host, r.URL.Path)

		if strings.HasPrefix(host, "www.") {
			next.ServeHTTP(w, r)
			return
		}

		for _, v := range hostExcludes {
			if strings.HasPrefix(host, v) {
				next.ServeHTTP(w, r)
				return
			}
		}

		http.Redirect(w, r, redirectUrl, http.StatusPermanentRedirect)
	})
}
