package kvs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vrdb-storage/objects"
	"vrdb-storage/server"

	"github.com/gorilla/mux"
	"vrdb.go/logging"
)

// remove a specific keyvalueset
func Remove(w http.ResponseWriter, r *http.Request) {

	// get the uid from the query params
	key, ok := mux.Vars(r)["key"]
	if !ok {
		logging.Log.Info("key in query is missing", "response-code", http.StatusBadRequest, "query", mux.Vars(r))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// get the object id
	var kvs = objects.KeyValueSet{}
	if result := server.Database.Preload("Tags").First(&kvs, "key = ?", key); result.Error != nil {

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

	result := struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf("removed object with key [%s]", key),
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {

		logging.Log.Info("error parsing result into json", "response-code", http.StatusInternalServerError, "result", result, "err", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

}
