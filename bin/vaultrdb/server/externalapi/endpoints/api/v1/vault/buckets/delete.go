package api_v1_buckets

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/conf"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// returns 200 - OK when successful
func Delete(w http.ResponseWriter, r *http.Request) {
	var log = logging.FromContext(r.Context())

	bucket, ok := mux.Vars(r)["bucket"]
	if !ok {
		log.Warn("bucket in query is missing",
			"response-code", http.StatusBadRequest,
			"query", mux.Vars(r),
		)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := conf.Vault.DeleteBucket(bucket); err != nil {
		log.Error("deleting bucket failed",
			"response-code", http.StatusInternalServerError,
			"error", err.Error(),
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("OK"))
}
