package endpoints

import (
	"net/http"
	"vrdb-uiserver/server"

	"github.com/gorilla/mux"
	"vrdb.go/logging"
)

const (
	UI_Dir string = "/opt/vaultrdb/web"
)

const ENABLE_UI bool = true

// enables the endpoint for the main ui
//
// route - http://<host>:<port>/ui/[...]
func EnableEndpoint_UI(r *mux.Router) {

	// if swagger shouldn't be enabled, then skip the insertion of the endpoint
	if !ENABLE_UI {
		return
	}

	logging.Log.Info("creating ui endpoint under relative path [/ui/...]")

	// enable the metadata endpoint for meta informations
	r.Methods(http.MethodGet).PathPrefix("/ui/").Handler(
		server.DefaultMiddleware.Then(
			http.StripPrefix("/ui/", http.FileServer(http.Dir(UI_Dir)))))
}
