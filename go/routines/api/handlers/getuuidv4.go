package handlers

import (
	"net/http"

	"github.com/google/uuid"
)

func UUIDv4(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte(uuid.New().String()))
	case http.MethodOptions: // ------------------- default responses
		w.Header().Add("Allow", "GET,OPTIONS")
		w.Write([]byte(http.StatusText(http.StatusOK)))
	default:
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}
}
