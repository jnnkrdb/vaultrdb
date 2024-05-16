package secrets

import (
	"net/http"
)

func Find(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}
