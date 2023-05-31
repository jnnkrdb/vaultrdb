package healthz

import "net/http"

// healthz check -> liveness
func Liveness(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}
