package endpoints

import (
	"net/http"
	"vrdb-uiserver/server"

	"github.com/gorilla/mux"
	"vrdb.go/logging"
)

const (
	SWAGGER_DIR string = "/opt/vaultrdb/swagger"
)

var ENABLE_SWAGGER bool = false

// enables the endpoint for swagger
//
// route - http://<host>:<port>/swagger
func EnableEndpoint_SwaggerUI(r *mux.Router) {

	// if swagger shouldn't be enabled, then skip the insertion of the endpoint
	if !ENABLE_SWAGGER {
		return
	}

	logging.Log.Info("creating swagger endpoint under relative path [/swagger]")

	// enable the metadata endpoint for meta informations
	r.Methods(http.MethodGet).PathPrefix("/swagger/").Handler(
		server.DefaultMiddleware.Then(
			http.StripPrefix("/swagger/", http.FileServer(http.Dir(SWAGGER_DIR)))))
}
