package vaultset

import "net/http"

func HandleHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		get(w, r)
	case http.MethodPost:
		post(w, r)
	case http.MethodPut:
		put(w, r)
	case http.MethodDelete:
		delete(w, r)

	case http.MethodOptions: // ------------------- default responses
		w.Header().Add("Allow", "GET,POST,PUT,DELETE,OPTIONS")
		w.Write([]byte(http.StatusText(http.StatusOK)))
	default:
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}
}
