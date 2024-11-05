package v1

import (
	"net/http"
	"vrdb-storage/endpoints/v1/fx"
	"vrdb-storage/endpoints/v1/internaldb"
	"vrdb-storage/endpoints/v1/kvs"
	"vrdb-storage/endpoints/v1/tags"
	"vrdb-storage/server"

	"github.com/gorilla/mux"
	"vrdb.go/logging"
)

const ENABLE_STORAGEAPI_V1 bool = true

// enables the endpoint for the storage api version 1
//
// Routes:
//   - http://<host>:<port>/api/v1
func EnableEndpoint_StorageAPI_V1(r *mux.Router) {

	// if storageapiv1 shouldn't be enabled, then skip the insertion of the endpoint
	if !ENABLE_STORAGEAPI_V1 {
		return
	}

	logging.Log.Info("creating storageapi endpoints", "version", "v1")

	// appending the default routes for keyvaluesets -> crud like
	r.Path("/api/v1/storedb/kvs").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(kvs.List))
	r.Path("/api/v1/storedb/kvs/{key}").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(kvs.Find))
	r.Path("/api/v1/storedb/kvs").Methods(http.MethodPost).Handler(server.DefaultMiddleware.ThenFunc(kvs.Insert))
	r.Path("/api/v1/storedb/kvs/{key}").Methods(http.MethodPut, http.MethodPatch).Handler(server.DefaultMiddleware.ThenFunc(kvs.Update))
	r.Path("/api/v1/storedb/kvs/{key}").Methods(http.MethodDelete).Handler(server.DefaultMiddleware.ThenFunc(kvs.Remove))

	// appending the default routes for tags -> crud like
	r.Path("/api/v1/storedb/tags").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(tags.List))

	// list the contents of the internaldb
	r.Path("/api/v1/internaldb/buckets").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(internaldb.List_Buckets))
	r.Path("/api/v1/internaldb/buckets/{bucket}/sink").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(internaldb.List_Keys))

	// appending the helperfunctions routes
	r.Path("/api/v1/f/decrypt/{key}").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(fx.DecryptValue))
}
