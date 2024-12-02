package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server"
	api_v1_configs "github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/api/v1/configs"
	api_v1_fx "github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/api/v1/fx"
	api_v1_kvs "github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/api/v1/kvs"
	api_v1_tags "github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/api/v1/tags"
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
	r.Path(URI_API_V1 + "/fx/encrypt").Methods(http.MethodPost).Handler(server.DefaultMiddleware.ThenFunc(api_v1_fx.EncryptValue))
	r.Path(URI_API_V1 + "/fx/decrypt").Methods(http.MethodPost).Handler(server.DefaultMiddleware.ThenFunc(api_v1_fx.DecryptValue))

	// configs functions
	r.Path(URI_API_V1 + "/configs/buckets").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(api_v1_configs.List_Buckets))
	r.Path(URI_API_V1 + "/configs/buckets/{bucket}/sink").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(api_v1_configs.List_Keys))

	// tag functions
	r.Path(URI_API_V1 + "/tags").Methods(http.MethodPost).Handler(server.DefaultMiddleware.ThenFunc(api_v1_tags.ListTags))

	// kvs functions
	r.Path(URI_API_V1 + "/kvs").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(api_v1_kvs.List))
	r.Path(URI_API_V1 + "/kvs/{key}").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(api_v1_kvs.Find))
	r.Path(URI_API_V1 + "/kvs").Methods(http.MethodPost).Handler(server.DefaultMiddleware.ThenFunc(api_v1_kvs.Insert))
	r.Path(URI_API_V1+"/kvs/{key}").Methods(http.MethodPut, http.MethodPatch).Handler(server.DefaultMiddleware.ThenFunc(api_v1_kvs.Update))
	r.Path(URI_API_V1 + "/kvs/{key}").Methods(http.MethodDelete).Handler(server.DefaultMiddleware.ThenFunc(api_v1_kvs.Remove))

}
