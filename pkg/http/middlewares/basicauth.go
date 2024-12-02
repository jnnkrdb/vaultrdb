package mw

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

// TODO: implement loading basicauth configs from external source, e.g. ENV Var or config.yaml
var (
	expectedUsernameHash [32]byte = sha256.Sum256([]byte("admin")) // os.Getenv("BASICAUTH_USER")
	expectedPasswordHash [32]byte = sha256.Sum256([]byte("admin")) // os.Getenv("BASICAUTH_PASS")
)

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// check for basic auth
		if user, pass, ok := r.BasicAuth(); ok {

			var (
				usernameHash [32]byte = sha256.Sum256([]byte(user))
				passwordHash [32]byte = sha256.Sum256([]byte(pass))
			)

			if (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1) &&
				(subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1) {

				// serve next request step
				next.ServeHTTP(w, r)
				return
			}
		}

		// when the basic auth checks failed
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
