package v1

import (
	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/libs/logging"
)

const (
	URI_API_V1 string = "/api/v1"
)

// enables the endpoint for metadata endpoint
//
// Routes:
//   - http://<host>:<port>/api/v1
func EnableEndpoint_ApiV1(r *mux.Router) {

	logging.Log.WithValues(
		"uri.apiv1", URI_API_V1,
	).Info("creating apiv1 endpoints")

	// r.Methods(http.MethodGet).Path(URI_Metadata).Handler(server.DefaultMiddleware.ThenFunc(nil))
}
