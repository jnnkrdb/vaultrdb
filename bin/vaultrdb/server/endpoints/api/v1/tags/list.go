package api_v1_tags

import (
	"encoding/json"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/database"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/obj"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// list all results of the tags table
//
// parameters will be handled in the future
func ListTags(w http.ResponseWriter, r *http.Request) {
	var tag_list = []obj.Tag{}

	if err := database.Database.Distinct("tag").Find(&tag_list).Error; err != nil {
		logging.SLog.Info("error receiving list of tag",
			"response-code", http.StatusInternalServerError,
			"tag_list", tag_list,
			"err", err,
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(tag_list); err != nil {
		logging.SLog.Info("error parsing result into json",
			"response-code", http.StatusInternalServerError,
			"err", err,
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
