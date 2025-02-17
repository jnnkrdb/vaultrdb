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

	logging.SLog.Info("creating swagger endpoints",
		"uri.swagger", URI_Swagger,
	)

	r.Methods(http.MethodHead, http.MethodGet).Path(URI_Swagger).Handler(server.DefaultMiddleware.Then(
		http.StripPrefix(URI_Swagger, http.FileServer(http.Dir(DIR_Swagger)))))
}
