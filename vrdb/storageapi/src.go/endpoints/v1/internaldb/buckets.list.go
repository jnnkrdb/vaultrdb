package internaldb

import (
	"encoding/json"
	"net/http"
	"vrdb-storage/server/configs"

	"vrdb.go/logging"
)

// list all keys in the internaldb
//
// parameters will be handled in the future
func List_Buckets(w http.ResponseWriter, r *http.Request) {

	var result = struct {
		Buckets []string `json:"buckets"`
	}{Buckets: configs.Buckets}

	logging.Log.Info("sending buckets", "result", result)

	// translate into json and ship
	if err := json.NewEncoder(w).Encode(result); err != nil {

		logging.Log.Info("error parsing result into json", "response-code", http.StatusInternalServerError, "err", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
