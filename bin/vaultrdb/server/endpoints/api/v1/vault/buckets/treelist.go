package api_v1_buckets

import (
	"net/http"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

func TreeList(w http.ResponseWriter, r *http.Request) {

	logging.SLog.Warn("getting a treelist is currently not implemented",
		"response-code", http.StatusNotImplemented,
	)

	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}
