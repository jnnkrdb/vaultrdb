package mw

import (
	"net/http"
	"time"

	"vrdb.go/logging"
)

// logging information for an http request
// like time since request, url and method
func QueryLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var start = time.Now()
		logging.Log.WithValues("request-url", r.URL.String(), "method", r.Method).Info("received request")

		next.ServeHTTP(w, r)

		logging.Log.WithValues("request-url", r.URL.String(), "method", r.Method).Info("finished request", "timeSince", time.Since(start))
	})
}
