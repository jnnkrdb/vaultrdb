package kvs

import (
	"encoding/json"
	"net/http"
	"vrdb-storage/objects"
	"vrdb-storage/server"

	"vrdb.go/logging"
)

// list all results of the kvs table
//
// parameters will be handled in the future
func List(w http.ResponseWriter, r *http.Request) {
	var kvs_list = []objects.KeyValueSet{}
	if err := server.Database.Model(&objects.KeyValueSet{}).Preload("Tags").Find(&kvs_list).Error; err != nil {
		logging.Log.Info("error receiving list of kvs", "response-code", http.StatusInternalServerError, "kvs_list", kvs_list, "err", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		if err = json.NewEncoder(w).Encode(kvs_list); err != nil {
			logging.Log.Info("error parsing result into json", "response-code", http.StatusInternalServerError, "err", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
