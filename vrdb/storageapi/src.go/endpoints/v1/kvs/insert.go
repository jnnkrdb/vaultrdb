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

	// receiving the required date of the keyvalueset, to create
	// a new dataset in the database and create the associations
	var obj = objects.NewKeyValueSet{}
	if err := obj.FromJSON(w, r); err != nil {
		return
	}

	// translate into db object
	var kvs = objects.KeyValueSet{
		Key:         obj.Key,
		Tags:        obj.Tags,
		Description: obj.Description,
		Value:       obj.Value,
	}

	// insert into database
	if err := server.Database.Create(&kvs).Error; err != nil {
		logging.Log.Info("error creating object in database", "response-code", http.StatusInternalServerError, "kvs", kvs, "err", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// send result
	if err := json.NewEncoder(w).Encode(kvs); err != nil {
		logging.Log.Info("error parsing result into json", "response-code", http.StatusInternalServerError, "kvs", kvs, "err", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
