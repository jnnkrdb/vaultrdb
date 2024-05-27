package v1

import (
	"encoding/json"
	"net/http"
	"vrdb-storage/objects"
	"vrdb-storage/server"

	"vrdb.go/logging"
)

// list all results of the sc table
//
// parameters will be handled in the future
func SC_List(w http.ResponseWriter, r *http.Request) {

	var serviceconfigs []objects.ServiceConfig
	if result := server.Database.Find(&serviceconfigs); result.Error != nil {

		logging.Log.Info("error receiving list of serviceconfigs", "response-code", http.StatusInternalServerError, "serviceconfigs", serviceconfigs, "result.Error", result.Error)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

	} else {

		if err := json.NewEncoder(w).Encode(serviceconfigs); err != nil {

			logging.Log.Info("error parsing result into json", "response-code", http.StatusInternalServerError, "err", err)

			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
