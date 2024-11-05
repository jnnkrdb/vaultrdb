package tags

import (
	"encoding/json"
	"net/http"
	"vrdb-storage/objects"
	"vrdb-storage/server"

	"vrdb.go/logging"
)

// list all results of the tags table
//
// parameters will be handled in the future
func List(w http.ResponseWriter, r *http.Request) {
	var tag_list = []objects.Tag{}
	if err := server.Database.Distinct("tag").Find(&tag_list).Error; err != nil {
		logging.Log.Info("error receiving list of tag", "response-code", http.StatusInternalServerError, "tag_list", tag_list, "err", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		if err = json.NewEncoder(w).Encode(tag_list); err != nil {
			logging.Log.Info("error parsing result into json", "response-code", http.StatusInternalServerError, "err", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
