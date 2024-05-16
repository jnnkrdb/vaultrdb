package endpoints

import (
	"encoding/json"
	"net/http"
	"os"
	"vrdb-uiserver/server"

	"github.com/gorilla/mux"
	"vrdb.go/logging"
)

const (
	VERSION_FILE string = "/opt/vaultrdb/config/VERSION"
	LICENSE_FILE string = "/opt/vaultrdb/config/LICENSE"
)

const ENABLE_METADATA bool = true

// enables the endpoint for metadata
//
// route - http://<host>:<port>/metadata
func EnableEndpoint_Metadata(r *mux.Router) {

	// if metadata shouldn't be enabled, then skip the insertion of the endpoint
	if !ENABLE_METADATA {
		return
	}

	logging.Log.Info("creating metadata endpoint under relative path [/metadata]")

	// enable the metadata endpoint for meta informations
	r.Methods(http.MethodGet).Path("/metadata").Handler(server.DefaultMiddleware.ThenFunc(
		func(w http.ResponseWriter, r *http.Request) {

			// create the pseudo struct
			logging.Log.Info("creating pseudo struct for meta info")

			var result = struct {
				Version struct {
					UIServer string `json:"uiserver"`
				} `json:"version"`
				License string `json:"license"`
			}{}

			var (
				b   []byte
				err error
			)

			if b, err = os.ReadFile(VERSION_FILE); err != nil {
				logging.Log.Error(err, "error reading bytes from file", "file", VERSION_FILE)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			// setting the version for the ui server
			result.Version.UIServer = string(b)

			// get the license content of the project
			if b, err = os.ReadFile(LICENSE_FILE); err != nil {
				logging.Log.Error(err, "error reading bytes from file", "file", VERSION_FILE)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			// setting the license for the project
			result.License = string(b)

			// parse the result object into json and ship
			if err := json.NewEncoder(w).Encode(result); err != nil {
				logging.Log.Error(err, "error parsing result into json", "result", result)
			}
		}))
}
