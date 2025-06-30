package api_v1_buckets

import (
	"encoding/json"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/conf"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// sends a json response in the following format:
//
//	{
//		"buckets": [
//		  "bucket-1",
//		  "bucket-2",
//		  "bucket-3"
//	 ]
//	}
func List(w http.ResponseWriter, r *http.Request) {
	var log = logging.FromContext(r.Context())

	log.Debug("listing buckets")

	res, err := conf.Vault.ListBuckets()
	if err != nil {
		log.Error("error reading buckets",
			"response-code", http.StatusInternalServerError,
			"error", err.Error(),
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(struct {
		Buckets []string `json:"buckets"`
	}{Buckets: res}); err != nil {
		log.Error("error parsing response to json",
			"response-code", http.StatusInternalServerError,
			"error", err.Error(),
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
