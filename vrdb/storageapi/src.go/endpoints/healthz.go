package endpoints

import (
	"net/http"
	"vrdb-storage/server"

	"github.com/gorilla/mux"
	"vrdb.go/logging"
)

const ENABLE_HEALTHZ bool = true

const (
	RELPATH_Liveness  string = "/healthz/live"
	RELPATH_Readiness string = "/healthz/ready"
)

// enables the endpoint for healthz
//
// routes
//
// - http://<host>:<port>/healthz/liveness
//
// - http://<host>:<port>/healthz/readiness
func EnableEndpoint_Healthz(r *mux.Router) {

	// if healthz shouldn't be enabled, then skip the insertion of the endpoint
	if !ENABLE_HEALTHZ {
		return
	}

	logging.Log.Info("creating healthz endpoints under relative path [/healthz]", "liveness", RELPATH_Liveness, "readiness", RELPATH_Readiness)

	r.Methods(http.MethodGet).Path(RELPATH_Liveness).Handler(
		server.DefaultMiddleware.ThenFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Write([]byte("OK"))
		}))

	r.Methods(http.MethodGet).Path(RELPATH_Readiness).Handler(
		server.DefaultMiddleware.ThenFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Write([]byte("OK"))
		}))

}
