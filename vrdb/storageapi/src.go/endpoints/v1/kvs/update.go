package kvs

import (
	"encoding/json"
	"net/http"
	"vrdb-storage/objects"
	"vrdb-storage/server"

	"gorm.io/gorm"
	"vrdb.go/logging"
)

// updating the key/value set
func Update(w http.ResponseWriter, r *http.Request) {

	// receiving the required date of the keyvalueset, to update
	// a specific dataset in the database and update the associations
	var obj = objects.NewKeyValueSet{}
	if err := obj.FromJSON(w, r); err != nil {
		return
	}

	var kvs = objects.KeyValueSet{}
	// request the original from the database
	if err := server.Database.Preload("Tags").First(&kvs, "key = ?", obj.Key).Error; err != nil {
		logging.Log.Info("error finding original in database", "response-code", http.StatusNotFound, "obj", obj, "err", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// update the fields
	kvs.Tags = obj.Tags
	kvs.Value = obj.Value
	kvs.Description = obj.Description

	// save the object in database
	if err := server.Database.Session(&gorm.Session{
		FullSaveAssociations: true,
	}).Updates(&kvs).Error; err != nil {
		logging.Log.Info("error updating object in database", "kvs", kvs, "err", err)
		return
	}

	// change the tags from the object
	if err := server.Database.Model(&kvs).Association("Tags").Replace(obj.Tags); err != nil {
		logging.Log.Info("error updating objects tags in database", "obj.Tags", obj.Tags, "err", err)
		return
	}

	// send result
	if err := json.NewEncoder(w).Encode(kvs); err != nil {
		logging.Log.Info("error parsing result into json", "response-code", http.StatusInternalServerError, "kvs", kvs, "err", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
