package swagger

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

const (
	URI_Swagger string = "/swagger/"
	DIR_Swagger string = "/opt/vaultrdb/swagger"
)

// enables the endpoint for swagger ui
//
// Routes:
//   - http://<host>:<port>/swagger/
func EnableEndpoint_Swagger(r *mux.Router) {

	logging.Log.WithValues(
		"uri.swagger", URI_Swagger,
	).Info("creating swagger endpoints")

	r.Methods(http.MethodGet).Path(URI_Swagger).Handler(server.DefaultMiddleware.Then(
		http.StripPrefix("/swagger/", http.FileServer(http.Dir(DIR_Swagger)))))
}
