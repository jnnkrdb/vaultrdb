package api

import (
	"log"
	"net/http"

	hndlrs "github.com/jnnkrdb/gomw/handlers"
)

// default http api port is 8080
func HandleAPI() {
	// checking for errors
	if err := (&http.Server{
		Addr:    ":8080",
		Handler: hndlrs.GetHandler(httpHandlers),
	}).ListenAndServe(); err != nil {
		log.Panicf("%#v\n", err)
	}
}
