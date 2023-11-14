package frontend

import "net/http"

var EnableUI bool = false

func IsUIEnabled(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if EnableUI {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "frontend is disabled", http.StatusForbidden)
		}
	})
}
