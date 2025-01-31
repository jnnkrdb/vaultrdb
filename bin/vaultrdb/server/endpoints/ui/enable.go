package ui

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

const (
	DIR_UI string = "/opt/vaultrdb/web"
)

// enables the endpoint for ui hosting
//
// Routes:
//   - http://<host>:<port>/
func EnableEndpoint_UI(r *mux.Router) {

	logging.SLog.Info("creating ui endpoints")

	r.Path("/").Methods(http.MethodGet).Handler(server.DefaultMiddleware.Then(http.FileServer(http.Dir(DIR_UI))))
}
