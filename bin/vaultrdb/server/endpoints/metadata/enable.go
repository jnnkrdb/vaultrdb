package metadata

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

const pathPrefix string = "/meta"

// enables the endpoint for metadata endpoint
//
// Routes:
//   - http://<host>:<port>/meta/version
//   - http://<host>:<port>/meta/license
func EnableEndpoint_Metadata(r *mux.Router) {

	logging.SLog.Info("creating metadata endpoints")
	r.PathPrefix(pathPrefix).Path("/version").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(version))
	r.PathPrefix(pathPrefix).Path("/license").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(license))
}
