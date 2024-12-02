package api_v1_configs

import (
	"encoding/json"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/configstore/buckets"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// list all keys in the internaldb
//
// parameters will be handled in the future
func List_Buckets(w http.ResponseWriter, r *http.Request) {

	var result = struct {
		Buckets []string `json:"buckets"`
	}{Buckets: buckets.DefaultBuckets()}

	logging.SLog.Info("sending buckets", "result", result)

	// translate into json and ship
	if err := json.NewEncoder(w).Encode(result); err != nil {

		logging.SLog.Info("error parsing result into json",
			"response-code", http.StatusInternalServerError,
			"err", err,
		)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
