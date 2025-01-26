package api_v1_kvs

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/database"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/database/obj"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// find the first object by a given key
func Find(w http.ResponseWriter, r *http.Request) {

	// get the uid from the query params
	key, ok := mux.Vars(r)["key"]
	if !ok {
		logging.SLog.Info("key in query is missing",
			"response-code", http.StatusBadRequest,
			"query", mux.Vars(r),
		)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var kvs = obj.KeyValueSet{}

	if result := database.Database.Preload("Tags").First(&kvs, "key = ?", key); result.Error != nil {
		logging.SLog.Info("error receiving item of kvs",
			"response-code", http.StatusInternalServerError,
			"result.Error", result.Error,
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(kvs); err != nil {
		logging.SLog.Info("error parsing result into json",
			"response-code", http.StatusInternalServerError,
			"kvs", kvs,
			"err", err,
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
