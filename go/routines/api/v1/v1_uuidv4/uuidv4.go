package v1_uuidv4

import (
	"net/http"

	"github.com/google/uuid"
)

// uuid function to receive a possible new id
func UUIDv4(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:

		// create a new uuid and write the result into the response body
		w.Write([]byte(uuid.New().String()))

	case http.MethodOptions: // ------------------- default responses

		// add the default allow header
		// uuid does not need any other methods than get or options
		w.Header().Add("Allow", "GET,OPTIONS")

		// write "OK" into the body
		w.Write([]byte(http.StatusText(http.StatusOK)))

	default:

		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}
}
