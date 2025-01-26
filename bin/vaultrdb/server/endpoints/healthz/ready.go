package healthz

import "net/http"

// check the readyness of the service
func ready(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))
}
