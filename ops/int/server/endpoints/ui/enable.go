package ui

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/int/server"
	"github.com/jnnkrdb/vaultrdb/libs/logging"
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

	logging.Log.WithValues(
		"uri.ui", URI_UI,
		"dir.ui", DIR_UI,
	).Info("creating ui endpoints")

	r.Methods(http.MethodGet).Path(URI_UI).Handler(server.DefaultMiddleware.Then(http.FileServer(http.Dir(DIR_UI))))
}
