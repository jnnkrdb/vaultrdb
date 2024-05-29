package kvs

import (
	"net/http"
	"vrdb-storage/objects"
	"vrdb-storage/server"

	"github.com/gorilla/mux"
	"vrdb.go/logging"
)

// remove a specific keyvalueset
func Remove(w http.ResponseWriter, r *http.Request) {

	// get the uid from the query params
	if key, ok := mux.Vars(r)["key"]; !ok {

		logging.Log.Info("key in query is missing", "response-code", http.StatusBadRequest, "query", mux.Vars(r))

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

	} else {

		// get the object id
		var kvs = objects.KeyValueSet{}
		if result := server.Database.First(&kvs, "key = ?", key); result.Error != nil {

			logging.Log.Info("error finding kvs", "response-code", http.StatusInternalServerError, "kvs", kvs, "result.Error", result.Error)

			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			return
		}

		// remove the object
		if result := server.Database.Delete(&kvs); result.Error != nil {

			logging.Log.Info("error removing kvs", "response-code", http.StatusInternalServerError, "kvs", kvs, "result.Error", result.Error)

			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			return
		}
	}

	w.Write([]byte("Deleted"))
}
