package api_v1_kvs

import (
	"encoding/json"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/database"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/database/obj"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
	"gorm.io/gorm"
)

// updating the key/value set
func Update(w http.ResponseWriter, r *http.Request) {

	// receiving the required date of the keyvalueset, to update
	// a specific dataset in the database and update the associations
	var newKVS = obj.NewKeyValueSet{}
	if err := newKVS.FromJSON(w, r); err != nil {
		return
	}

	var kvs = obj.KeyValueSet{}
	// request the original from the database
	if err := database.Database.Preload("Tags").First(&kvs, "key = ?", newKVS.Key).Error; err != nil {
		logging.SLog.Info("error finding original in database",
			"response-code", http.StatusNotFound,
			"newKVS", newKVS,
			"err", err,
		)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// update the fields
	kvs.Value = newKVS.Value
	kvs.Description = newKVS.Description

	// save the object in database
	if err := database.Database.Session(&gorm.Session{
		FullSaveAssociations: true,
	}).Updates(&kvs).Error; err != nil {
		logging.SLog.Info("error updating object in database",
			"kvs", kvs,
			"err", err,
		)
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
