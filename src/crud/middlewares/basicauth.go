package middlewares

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
	"os"
)

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// check for basic auth
		if user, pass, ok := r.BasicAuth(); ok {

			usernameHash := sha256.Sum256([]byte(user))
			passwordHash := sha256.Sum256([]byte(pass))
			expectedUsernameHash := sha256.Sum256([]byte(os.Getenv("BASICAUTH_USER")))
			expectedPasswordHash := sha256.Sum256([]byte(os.Getenv("BASICAUTH_PASS")))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		// when the basic auth checks failed
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
