package api_v1_kvs

import (
	"encoding/json"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/database"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/obj"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// insert a new kvs into the database
func Insert(w http.ResponseWriter, r *http.Request) {

	// receiving the required date of the keyvalueset, to create
	// a new dataset in the database and create the associations
	var newKVS = obj.NewKeyValueSet{}
	if err := newKVS.FromJSON(w, r); err != nil {
		return
	}

	// translate into db object
	var kvs = obj.KeyValueSet{
		Key:         newKVS.Key,
		Tags:        newKVS.Tags,
		Description: newKVS.Description,
		Value:       newKVS.Value,
	}

	// insert into database
	if err := database.Database.Create(&kvs).Error; err != nil {
		logging.SLog.Info("error creating object in database",
			"response-code", http.StatusInternalServerError,
			"kvs", kvs,
			"err", err,
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// send result
	if err := json.NewEncoder(w).Encode(kvs); err != nil {
		logging.SLog.Info("error parsing result into json",
			"response-code", http.StatusInternalServerError,
			"kvs", kvs,
			"err", err,
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
