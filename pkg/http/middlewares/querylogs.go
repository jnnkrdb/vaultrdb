package mw

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// logging information for an http request
// like time since request, url and method
func QueryLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var start = time.Now()

		// insert logger into the request context
		requestLogger := logging.Default.With(
			"request-uid", uuid.NewString(),
			"request-starttime", start.Format(time.RFC3339Nano),
			"request-url", r.URL.String(),
			"method", r.Method,
		)

		requestLogger.Debug("received request")

		// serve http request
		next.ServeHTTP(w, r.WithContext(logging.IntoContext(r.Context(), requestLogger)))

		requestLogger.Debug("finished request", "time-since", fmt.Sprintf("%dms", time.Since(start).Microseconds()))
	})
}
