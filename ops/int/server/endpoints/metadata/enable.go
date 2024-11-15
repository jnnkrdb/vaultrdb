package metadata

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/int/server"
	"github.com/jnnkrdb/vaultrdb/libs/logging"
)

// enables the endpoint for metadata endpoint
//
// Routes:
//   - http://<host>:<port>/meta/version
//   - http://<host>:<port>/meta/license
func EnableEndpoint_Metadata(r *mux.Router) {

	logging.Log.WithValues(
		"uri.metadata.version", URI_Metadata_Version,
		"uri.metadata.license", URI_Metadata_License,
	).Info("creating metadata endpoints")

	r.Methods(http.MethodGet).Path(URI_Metadata_Version).Handler(server.DefaultMiddleware.ThenFunc(version))
	r.Methods(http.MethodGet).Path(URI_Metadata_License).Handler(server.DefaultMiddleware.ThenFunc(license))
}
