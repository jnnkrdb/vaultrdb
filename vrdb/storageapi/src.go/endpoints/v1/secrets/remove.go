package secrets

import (
	"net/http"
)

func Remove(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}
