package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server"
	api_v1_fx "github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/api/v1/fx"
	api_v1_buckets "github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/api/v1/vault/buckets"
	api_v1_buckets_sink "github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/api/v1/vault/buckets/sink"
)

// enables the endpoint for metadata endpoint
//
// Routes:
//   - http://<host>:<port>/api/v1
func EnableEndpoint_ApiV1(r *mux.Router) {

	r.Path("/api/v1/fx/encrypt").Methods(http.MethodPost).Handler(server.DefaultMiddleware.ThenFunc(api_v1_fx.EncryptValue))
	r.Path("/api/v1/fx/decrypt").Methods(http.MethodPost).Handler(server.DefaultMiddleware.ThenFunc(api_v1_fx.DecryptValue))

	r.Path("/api/v1/vault/buckets/{bucketpath}").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(api_v1_buckets.List))
	r.Path("/api/v1/vault/buckets/{bucketpath}").Methods(http.MethodPost).Handler(server.DefaultMiddleware.ThenFunc(api_v1_buckets.Create))
	r.Path("/api/v1/vault/buckets/{bucketpath}/bucket/{bucket}").Methods(http.MethodDelete).Handler(server.DefaultMiddleware.ThenFunc(api_v1_buckets.Delete))

	r.Path("/api/v1/vault/buckets/{bucketpath}/sink").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(api_v1_buckets_sink.List))
	r.Path("/api/v1/vault/buckets/{bucketpath}/sink/{key}").Methods(http.MethodGet).Handler(server.DefaultMiddleware.ThenFunc(api_v1_buckets_sink.Find))
	r.Path("/api/v1/vault/buckets/{bucketpath}/sink/{key}").Methods(http.MethodPost, http.MethodPut, http.MethodPatch).Handler(server.DefaultMiddleware.ThenFunc(api_v1_buckets_sink.Write))
	r.Path("/api/v1/vault/buckets/{bucketpath}/sink/{key}").Methods(http.MethodDelete).Handler(server.DefaultMiddleware.ThenFunc(api_v1_buckets_sink.Delete))
}
