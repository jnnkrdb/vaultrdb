package v1

import (
	"net/http"
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

	// appending the default routes for keyvaluesets -> crud like
	r.Path("/api/v1/kvs").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(KVS_List))
	r.Path("/api/v1/kvs/{key}").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(KVS_Find))
	r.Path("/api/v1/kvs").Methods(http.MethodPost).Handler(server.DefaultMiddleware.ThenFunc(KVS_Insert))
	r.Path("/api/v1/kvs/{key}").Methods(http.MethodPut, http.MethodPatch).Handler(server.DefaultMiddleware.ThenFunc(KVS_Update))
	r.Path("/api/v1/kvs/{key}").Methods(http.MethodDelete).Handler(server.DefaultMiddleware.ThenFunc(KVS_Remove))

	// appending the serviceconfig routes
	r.Path("/api/v1/serviceconfigs").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(SC_List))
}
