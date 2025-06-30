package api_v1_buckets_sink

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/conf"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

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

	key, ok := mux.Vars(r)["key"]
	if !ok {
		log.Warn("key in query is missing",
			"response-code", http.StatusBadRequest,
			"query", mux.Vars(r),
		)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := conf.Vault.DeleteKey(bucket, key); err != nil {
		log.Error("error removing key from bucket",
			"bucket", bucket,
			"key", key,
			"response-code", http.StatusInternalServerError,
			"err", err,
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("OK"))
}
