package api_v1_kvs

import (
	"encoding/json"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/database"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/database/obj"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// list all results of the kvs table
//
// parameters will be handled in the future
func List(w http.ResponseWriter, r *http.Request) {

	var kvs_list = []obj.KeyValueSet{}

	if err := database.Database.Model(&obj.KeyValueSet{}).Preload("Tags").Find(&kvs_list).Error; err != nil {
		logging.SLog.Info("error receiving list of kvs",
			"response-code", http.StatusInternalServerError,
			"kvs_list", kvs_list,
			"err", err,
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(kvs_list); err != nil {
		logging.SLog.Info("error parsing result into json",
			"response-code", http.StatusInternalServerError,
			"err", err,
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
