package healthz

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

const pathPrefix string = "/healthz"

// enables the endpoint for healthz
//
// Routes:
//   - http://<host>:<port>/healthz/liveness
//   - http://<host>:<port>/healthz/readiness
func EnableEndpoint_Healthz(r *mux.Router) {

	logging.Default.Info("creating healthz endpoints")

	r.PathPrefix(pathPrefix).Path("/live").Methods(http.MethodGet).HandlerFunc(live)

	r.PathPrefix(pathPrefix).Path("/ready").Methods(http.MethodGet).HandlerFunc(ready)
}
