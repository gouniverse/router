package router

import (
	"net/http"
	"strings"
)

// NewWwwToNakedDomainMiddleware will redirect a "www" subdomain to naked (non-www) domain
func NewWwwToNakedDomainMiddleware() Middleware {
	m := Middleware{
		Name:    "WWW to Naked Domain Middleware",
		Handler: wwwToNakedDomain,
	}

	return m
}

// wwwToNakedDomain is http middleware that ensures a "www" sub domain is redirected to naked domain
func wwwToNakedDomain(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host := strings.ToLower(r.Host)
		scheme := strings.ToLower(r.URL.Scheme)

		if strings.HasPrefix(host, "www") {
			if scheme == "" || scheme == "/" {
				scheme = "https"
			}
			redirectURl := (scheme + "://" + r.Host[4:] + r.URL.Path)
			// TODO. Add pub sub slot
			// models.LogStore.InfoWithContext("Redirecting www to non www", map[string]string{
			// 	"host": r.Host,
			// 	"uri":  r.RequestURI,
			// })
			http.Redirect(w, r, redirectURl, http.StatusTemporaryRedirect)
			return
		}

		next.ServeHTTP(w, r)
	})
}
