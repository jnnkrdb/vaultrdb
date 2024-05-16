package configs

import (
	"net/http"
)

func Insert(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}
