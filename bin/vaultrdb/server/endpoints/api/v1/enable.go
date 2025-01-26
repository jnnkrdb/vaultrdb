package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server"
	api_v1_buckets "github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/api/v1/buckets"
	api_v1_buckets_sink "github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/api/v1/buckets/sink"
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
	r.Path(URI_API_V1 + "/fx/encrypt").Methods(http.MethodPost).Handler(server.DefaultMiddleware.ThenFunc(api_v1_fx.EncryptValue))
	r.Path(URI_API_V1 + "/fx/decrypt").Methods(http.MethodPost).Handler(server.DefaultMiddleware.ThenFunc(api_v1_fx.DecryptValue))

	// bucket functions
	r.Path(URI_API_V1 + "/vault/buckets").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(api_v1_buckets.List))
	r.Path(URI_API_V1 + "/vault/buckets/{bucketpath}").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(api_v1_buckets.Find))
	r.Path(URI_API_V1 + "/vault/buckets/{bucketpath}").Methods(http.MethodPost).Handler(server.DefaultMiddleware.ThenFunc(api_v1_buckets.Create))
	r.Path(URI_API_V1 + "/vault/buckets/{bucketpath}").Methods(http.MethodDelete).Handler(server.DefaultMiddleware.ThenFunc(api_v1_buckets.Delete))

	// keyvalue sinks functions
	r.Path(URI_API_V1 + "/vault/buckets/{bucketpath}/sink").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(api_v1_buckets_sink.List))
	r.Path(URI_API_V1 + "/vault/buckets/{bucketpath}/sink/{key}").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(api_v1_buckets_sink.Find))
	r.Path(URI_API_V1 + "/vault/buckets/{bucketpath}/sink").Methods(http.MethodPost).Handler(server.DefaultMiddleware.ThenFunc(api_v1_buckets_sink.Create))
	r.Path(URI_API_V1+"/vault/buckets/{bucketpath}/sink/{key}").Methods(http.MethodPut, http.MethodPatch).Handler(server.DefaultMiddleware.ThenFunc(api_v1_buckets_sink.Update))
	r.Path(URI_API_V1 + "/vault/buckets/{bucketpath}/sink/{key}").Methods(http.MethodDelete).Handler(server.DefaultMiddleware.ThenFunc(api_v1_buckets_sink.Delete))
}
