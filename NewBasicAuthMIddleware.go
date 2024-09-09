package router

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

func NewBasicAuthentication(expectedUsername, expectedPassword string) Middleware {
	m := Middleware{
		Name: "Basic Authentication Middleware",
		Handler: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				submittedUsername, submittedPassword, ok := r.BasicAuth()

				if !ok {
					w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}

				expectedUsernameHash := sha256.Sum256([]byte(expectedUsername))
				expectedPasswordHash := sha256.Sum256([]byte(expectedPassword))
				submittedUsernameHash := sha256.Sum256([]byte(submittedUsername))
				submittedPasswordHash := sha256.Sum256([]byte(submittedPassword))

				// Use the subtle.ConstantTimeCompare to evaluate both the
				// username and password before checking the return values to
				// avoid leaking information.
				usernameMatch := (subtle.ConstantTimeCompare(submittedUsernameHash[:], expectedUsernameHash[:]) == 1)
				passwordMatch := (subtle.ConstantTimeCompare(submittedPasswordHash[:], expectedPasswordHash[:]) == 1)

				if usernameMatch && passwordMatch {
					next.ServeHTTP(w, r)
					return
				}

				w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			})
		},
	}

	return m
}
