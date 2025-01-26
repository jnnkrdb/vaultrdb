package healthz

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

const (
	URI_Liveness  string = "/healthz/live"
	URI_Readiness string = "/healthz/ready"
)

// enables the endpoint for healthz
//
// Routes:
//   - http://<host>:<port>/healthz/liveness
//   - http://<host>:<port>/healthz/readiness
func EnableEndpoint_Healthz(r *mux.Router) {

	logging.SLog.Info("creating healthz endpoints",
		"uri.liveness", URI_Liveness,
		"uri.readiness", URI_Readiness,
	)

	r.Methods(http.MethodGet).Path(URI_Liveness).Handler(
		server.DefaultMiddleware.ThenFunc(live))

	r.Methods(http.MethodGet).Path(URI_Readiness).Handler(
		server.DefaultMiddleware.ThenFunc(ready))
}
