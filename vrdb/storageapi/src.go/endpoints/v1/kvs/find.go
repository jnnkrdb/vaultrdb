package kvs

import (
	"encoding/json"
	"net/http"
	"vrdb-storage/objects"
	"vrdb-storage/server"

	"github.com/gorilla/mux"
	"vrdb.go/logging"
)

// find the first object by a given key
func Find(w http.ResponseWriter, r *http.Request) {

	// get the uid from the query params
	if key, ok := mux.Vars(r)["key"]; !ok {

		logging.Log.Info("key in query is missing", "response-code", http.StatusBadRequest, "query", mux.Vars(r))

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

	} else {

		var kvs objects.KeyValueSet

		if result := server.Database.First(&kvs, "key = ?", key); result.Error != nil {

			logging.Log.Info("error receiving item of kvs", "response-code", http.StatusInternalServerError, "result.Error", result.Error)

			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		} else {

			if err := json.NewEncoder(w).Encode(kvs); err != nil {

				logging.Log.Info("error parsing result into json", "response-code", http.StatusInternalServerError, "kvs", kvs, "err", err)

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}
	}
}
