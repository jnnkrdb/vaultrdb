package kvs

import (
	"encoding/json"
	"net/http"
	"vrdb-storage/objects/keyvalueset"

	"vrdb.go/logging"
)

// insert a new kvs into the database
func Insert(w http.ResponseWriter, r *http.Request) {

	// receiving the required date of the keyvalueset, to create
	// a new dataset in the database and create the associations
	var obj = keyvalueset.NewKeyValueSet{}
	if err := obj.FromJSON(w, r); err != nil {
		return
	}

	insert, err := keyvalueset.InsertIntoDB(obj)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// send result
	if err := json.NewEncoder(w).Encode(insert); err != nil {
		logging.Log.Info("error parsing result into json", "response-code", http.StatusInternalServerError, "kvs", kvs, "err", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
