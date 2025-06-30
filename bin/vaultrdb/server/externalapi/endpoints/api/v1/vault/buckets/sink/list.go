package api_v1_buckets_sink

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/conf"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// sends a json response in the following format:
//
//	{
//		"keys": [
//		  "key-1",
//		  "key-2",
//		  "key-3"
//	 ]
//	}
func List(w http.ResponseWriter, r *http.Request) {
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

	list, err := conf.Vault.ListKeys(bucket)
	if err != nil {
		log.Error("error receiving keys from bucket",
			"bucket", bucket,
			"response-code", http.StatusInternalServerError,
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(struct {
		Keys []string `json:"keys"`
	}{Keys: list}); err != nil {
		log.Error("error parsing response to json",
			"response-code", http.StatusInternalServerError,
			"error", err.Error(),
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
