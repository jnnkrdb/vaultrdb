package v1

import (
	"net/http"
	"vrdb-storage/objects"
	"vrdb-storage/server"
	"vrdb-storage/server/funcs"

	"github.com/gorilla/mux"
	"vrdb.go/logging"
)

// list all results of the kvs table
//
// parameters will be handled in the future
func DecryptValue(w http.ResponseWriter, r *http.Request) {

	if key, ok := mux.Vars(r)["key"]; !ok {

		logging.Log.Info("key in query is missing", "response-code", http.StatusBadRequest, "query", mux.Vars(r))

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

	} else {

		logging.Log.Info("requested encryption for key/value", "key", key)

		var kvs objects.KeyValueSet

		if result := server.Database.First(&kvs, "key = ?", key); result.Error != nil {

			logging.Log.Info("error receiving item of kvs", "response-code", http.StatusInternalServerError, "result.Error", result.Error)

			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		} else {

			if value, err := funcs.Decrypt(kvs.Value); err != nil {

				logging.Log.WithValues(
					"response-code", http.StatusInternalServerError,
					"kvs.Key", kvs.Key,
					"err", err,
				).Info("error decrypting result")

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			} else {

				w.Write([]byte(value))
			}
		}
	}
}
