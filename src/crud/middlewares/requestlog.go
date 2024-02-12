package middlewares

import (
	"fmt"
	"hash/fnv"
	"net/http"
	"time"

	"github.com/jnnkrdb/vaultrdb/crud/config"
)

func ReportRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var startTime = time.Now()
		var hash string = func() string {
			var res = fnv.New32a()
			res.Write([]byte(startTime.Format(time.RFC3339Nano)))
			return fmt.Sprintf("%d", res.Sum32())
		}()

		config.CrudLog.WithValues(
			"requestID", hash,
			"method", r.Method,
			"requesturi", r.RequestURI,
			"parameters", r.Form.Encode(),
			"host", r.Host,
		).Info("received request")

		next.ServeHTTP(w, r)

		config.CrudLog.WithValues(
			"requestID", hash,
			"timeSince", time.Since(startTime).Milliseconds(),
		).Info("finished request")
	})
}
