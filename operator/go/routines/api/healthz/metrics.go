package healthz

import "net/http"

// healthz check -> metrics
func Metrics(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
