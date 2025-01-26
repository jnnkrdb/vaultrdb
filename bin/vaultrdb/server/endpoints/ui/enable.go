package ui

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

const (
	URI_UI string = "/"
	DIR_UI string = "/opt/vaultrdb/web"
)

// enables the endpoint for ui hosting
//
// Routes:
//   - http://<host>:<port>/
func EnableEndpoint_UI(r *mux.Router) {

	logging.SLog.Info("creating ui endpoints",
		"uri.ui", URI_UI,
		"dir.ui", DIR_UI,
	)

	r.Methods(http.MethodGet).Path(URI_UI).Handler(server.DefaultMiddleware.Then(http.FileServer(http.Dir(DIR_UI))))
}
