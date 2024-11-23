package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server"
	api_v1_fx "github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/api/v1/fx"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

const (
	URI_API_V1 string = "/api/v1"
)

// enables the endpoint for metadata endpoint
//
// Routes:
//   - http://<host>:<port>/api/v1
func EnableEndpoint_ApiV1(r *mux.Router) {

	logging.SLog.Info("creating apiv1 endpoints",
		"uri.apiv1", URI_API_V1,
	)

	// fx functions
	r.Methods(http.MethodPost).Path(URI_API_V1 + "/fx/encrypt").Handler(server.DefaultMiddleware.ThenFunc(api_v1_fx.EncryptValue))
	r.Methods(http.MethodPost).Path(URI_API_V1 + "/fx/decrypt").Handler(server.DefaultMiddleware.ThenFunc(api_v1_fx.DecryptValue))
}
