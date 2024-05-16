package v1

import (
	"net/http"
	"vrdb-storage/endpoints/v1/configs"
	"vrdb-storage/endpoints/v1/secrets"
	"vrdb-storage/server"

	"github.com/gorilla/mux"
	"vrdb.go/logging"
)

const ENABLE_STORAGEAPI_V1 bool = true

// enables the endpoint for the storage api version 1
//
// routes
//
// - http://<host>:<port>/api/v1
func EnableEndpoint_StorageAPI_V1(r *mux.Router) {

	// if storageapiv1 shouldn't be enabled, then skip the insertion of the endpoint
	if !ENABLE_STORAGEAPI_V1 {
		return
	}

	logging.Log.Info("creating storageapi endpoints", "version", "v1")

	// appending the default routes for secrets -> crud like
	r.Path("/api/v1/secrets").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(secrets.List))
	r.Path("/api/v1/secrets/{uid}").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(secrets.Find))
	r.Path("/api/v1/secrets/{uid}").Methods(http.MethodPost).Handler(server.DefaultMiddleware.ThenFunc(secrets.Insert))
	r.Path("/api/v1/secrets/{uid}").Methods(http.MethodPut, http.MethodPatch).Handler(server.DefaultMiddleware.ThenFunc(secrets.Update))
	r.Path("/api/v1/secrets/{uid}").Methods(http.MethodDelete).Handler(server.DefaultMiddleware.ThenFunc(secrets.Remove))

	r.Path("/api/v1/configs").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(configs.List))
	r.Path("/api/v1/configs/{uid}").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(configs.Find))
	r.Path("/api/v1/configs/{uid}").Methods(http.MethodPost).Handler(server.DefaultMiddleware.ThenFunc(configs.Insert))
	r.Path("/api/v1/configs/{uid}").Methods(http.MethodPut, http.MethodPatch).Handler(server.DefaultMiddleware.ThenFunc(configs.Update))
	r.Path("/api/v1/configs/{uid}").Methods(http.MethodDelete).Handler(server.DefaultMiddleware.ThenFunc(configs.Remove))
}
