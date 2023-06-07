package v2

import "net/http"

func GetPath(w http.ResponseWriter, r *http.Request) {
	Get(w, r)
}
