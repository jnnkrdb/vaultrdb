package restapi

import "net/http"

var EnableRest bool = false

func IsRESTApiEnabled(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if EnableRest {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "restapi is disabled", http.StatusForbidden)
		}
	})
}
