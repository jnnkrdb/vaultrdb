package healthz

import "net/http"

// check the liveness of the service
func live(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}
