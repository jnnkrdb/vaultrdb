package api_v1_buckets_sink

import "net/http"

func List(w http.ResponseWriter, r *http.Request) {

	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}
