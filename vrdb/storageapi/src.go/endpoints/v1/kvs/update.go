package kvs

import (
	"encoding/json"
	"net/http"
	"vrdb-storage/objects"
	"vrdb-storage/server"

	"github.com/lib/pq"
	"vrdb.go/logging"
)

// updating the key/value set
func Update(w http.ResponseWriter, r *http.Request) {

	var obj = struct {
		Key         string         `json:"key"`
		Value       string         `json:"value"`
		Tags        pq.StringArray `json:"tags"`
		Description string         `json:"description"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {

		logging.Log.Info("error parsing body into struct", "response-code", http.StatusBadRequest, "obj", obj, "err", err)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	// get the object id
	var kvs = objects.KeyValueSet{}
	if result := server.Database.First(&kvs, "key = ?", obj.Key); result.Error != nil {

		logging.Log.Info("error finding kvs", "response-code", http.StatusInternalServerError, "kvs", kvs, "result.Error", result.Error)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	kvs.Value = obj.Value
	kvs.Description = obj.Description

	// create object in database
	if result := server.Database.Save(&kvs); result.Error != nil {

		logging.Log.Info("error creating object in database", "response-code", http.StatusInternalServerError, "obj", obj, "result.Error", result.Error)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	// send result
	if err := json.NewEncoder(w).Encode(kvs); err != nil {

		logging.Log.Info("error parsing result into json", "response-code", http.StatusInternalServerError, "kvs", kvs, "err", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
