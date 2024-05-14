package v1

import (
	"net/http"
)

func UPDATE(w http.ResponseWriter, r *http.Request) {
	// ------------------------------------------------
	// currently this function is deactivated
	http.Error(w, "not implemented", http.StatusNotImplemented)
	// ------------------------------------------------
}
