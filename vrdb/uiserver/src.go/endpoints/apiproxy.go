package endpoints

import (
	"io"
	"net/http"
	"net/url"
	"vrdb-uiserver/server"

	"github.com/gorilla/mux"
	"vrdb.go/logging"
)

const ENABLE_APIPROXY bool = true

var (
	STORAGEAPISERVER     string   = ""
	_storageApiServerUrl *url.URL = &url.URL{
		Scheme: "http",
		Host:   STORAGEAPISERVER,
	}
)

// enables the proxy endpoint for the api
//
// route - http://<host>:<port>/api/[...]
func EnableEndpoint_ApiProxy(r *mux.Router) {

	// if api proxy shouldn't be enabled, then skip the insertion of the endpoint
	if !ENABLE_APIPROXY {
		return
	}

	logging.Log.V(3).Info("creating proxy endpoint under relative path [/storageapi/...]", "url", *_storageApiServerUrl)

	// enable the proxy endpoint for proxying to storage api server
	r.Methods(
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodHead,
	).Path("/storageapi/{storageapi:.*}").Handler(
		server.DefaultMiddleware.ThenFunc(func(w http.ResponseWriter, r *http.Request) {

			// Constructing  a new url according to
			// the incoming request and the target url
			r.URL = &url.URL{
				Scheme:      _storageApiServerUrl.Scheme,
				Host:        _storageApiServerUrl.Host,
				Path:        "/" + mux.Vars(r)["storageapi"],
				RawPath:     _storageApiServerUrl.RawPath,
				RawQuery:    _storageApiServerUrl.RawQuery,
				Fragment:    _storageApiServerUrl.Fragment,
				RawFragment: _storageApiServerUrl.RawFragment,
			}

			logging.Log.V(5).Info("new request url", "url", *r.URL)

			resp, err := http.DefaultTransport.RoundTrip(r)
			if err != nil {
				logging.Log.V(5).Error(err, "error proxying the request", "response", *resp)
				http.Error(w, err.Error(), http.StatusServiceUnavailable)
				return
			}

			// Copy the incoming headers to the outgoing headers
			for src_key, src_value := range resp.Header {
				for _, val := range src_value {
					w.Header().Add(src_key, val)
				}
			}

			// Set http status code
			w.WriteHeader(resp.StatusCode)

			// copy the incoming response body to the outgoing response writer
			_, _ = io.Copy(w, resp.Body)

			// Logging the progress
			defer func() {
				logging.Log.V(5).WithValues(
					"request_body", r.Body,
					"request_url", r.URL.String(),
					"request_headers", r.Header,
					"method", r.Method,
					"response", resp.Body,
				).Info("proxied request")

				resp.Body.Close()
			}()
		}))
}
