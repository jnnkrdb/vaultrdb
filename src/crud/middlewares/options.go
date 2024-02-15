package middlewares

import (
	"net/http"
)

func OptionsResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// setting the default response headers
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))

		if r.Method == http.MethodOptions { // if options, the send the response with the headers
			w.Header().Set("Allow", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			w.Write([]byte(http.StatusText(http.StatusOK)))
		} else {
			// if else, then process the next handler
			next.ServeHTTP(w, r)
		}
	})
}
