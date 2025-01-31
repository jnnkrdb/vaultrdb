package healthz

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

const pathPrefix string = "/health"

// enables the endpoint for healthz
//
// Routes:
//   - http://<host>:<port>/healthz/liveness
//   - http://<host>:<port>/healthz/readiness
func EnableEndpoint_Healthz(r *mux.Router) {

	logging.SLog.Info("creating healthz endpoints")
	r.PathPrefix(pathPrefix).Path("/live").Methods(http.MethodGet).Handler(
		server.DefaultMiddleware.ThenFunc(live))

	r.PathPrefix(pathPrefix).Path("/ready").Methods(http.MethodGet).Handler(
		server.DefaultMiddleware.ThenFunc(ready))
}
