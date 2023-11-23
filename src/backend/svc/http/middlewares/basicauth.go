package middlewares

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
	"os"
)

var (
	BasicAuth_Username string = func() string {
		if user, ok := os.LookupEnv("BASICAUTH_USERNAME"); ok {
			return user
		}
		return "ADMIN"
	}()
	BasicAuth_Password string = func() string {
		if pass, ok := os.LookupEnv("BASICAUTH_PASSWORD"); ok {
			return pass
		}
		return "ADMIN"

	}()
)

func BasicAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(BasicAuth_Username))
			expectedPasswordHash := sha256.Sum256([]byte(BasicAuth_Password))

			// check the username
			if subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1 {
				// check the password
				if subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1 {
					next.ServeHTTP(w, r)
					return
				}
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
