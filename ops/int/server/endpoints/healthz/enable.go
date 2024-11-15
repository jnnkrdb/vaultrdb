package healthz

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/int/server"
	"github.com/jnnkrdb/vaultrdb/libs/logging"
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

	logging.Log.WithValues(
		"uri.liveness", URI_Liveness,
		"uri.readiness", URI_Readiness,
	).Info("creating healthz endpoints")

	r.Methods(http.MethodGet).Path(URI_Liveness).Handler(
		server.DefaultMiddleware.ThenFunc(live))

	r.Methods(http.MethodGet).Path(URI_Readiness).Handler(
		server.DefaultMiddleware.ThenFunc(ready))
}
