package api_v1_buckets

import "net/http"

func Create(w http.ResponseWriter, r *http.Request) {

	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}
