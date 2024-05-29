package kvs

import (
	"encoding/json"
	"net/http"
	"vrdb-storage/objects"
	"vrdb-storage/server"

	"vrdb.go/logging"
)

// insert a new kvs into the database
func Insert(w http.ResponseWriter, r *http.Request) {

	var (
		obj struct {
			Key         string `json:"key"`
			Value       string `json:"value"`
			Description string `json:"description"`
		}
	)

	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {

		logging.Log.Info("error parsing body into struct", "response-code", http.StatusBadRequest, "obj", obj, "err", err)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	// create object in database
	if result := server.Database.Create(&objects.KeyValueSet{
		Key:         obj.Key,
		Value:       obj.Value,
		Description: obj.Description}); result.Error != nil {

		logging.Log.Info("error creating object in database", "response-code", http.StatusInternalServerError, "obj", obj, "result.Error", result.Error)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	// get the object id
	var kvs = objects.KeyValueSet{}
	if result := server.Database.First(&kvs, "key = ?", obj.Key); result.Error != nil {

		logging.Log.Info("error finding kvs", "response-code", http.StatusInternalServerError, "kvs", kvs, "result.Error", result.Error)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	// send result
	if err := json.NewEncoder(w).Encode(kvs); err != nil {

		logging.Log.Info("error parsing result into json", "response-code", http.StatusInternalServerError, "kvs", kvs, "err", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
