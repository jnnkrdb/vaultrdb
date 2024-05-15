package main

import (
	"flag"
	"net/http"

	"vrdb-uiserver/server"

	"vrdb.go/logging"
)

// config via args
var (
	ENABLE_SWAGGER bool = false
	ENABLE_LICENSE bool = true
	ENABLE_VERSION bool = true
	ENABLE_UI      bool = true
	ENABLE_HEALTHZ bool = true
)

func main() {

	flag.BoolVar(&server.ENABLE_SWAGGER, "swagger", false, "Enables the swagger ui.")
	flag.Parse()

	logging.InitLogger("ui-server")

	server.StartFrontendUI()

	// append the ui endpoints, which are configured using the args
	switch {

	case _ENABLE_VERSION: // enable the version endpoint for the ui
		router.Handle("/version", _mw.ThenFunc(func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "/opt/vaultrdb/config/VERSION") }))
		log.V(2).Info("enabled version ui", "relative-path", "http://localhost/version")

	case _ENABLE_LICENSE: // enable the license endpoint for the ui
		router.Handle("/license", _mw.ThenFunc(func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "/opt/vaultrdb/config/LICENSE") }))
		log.V(2).Info("enabled license ui", "relative-path", "http://localhost/license")

	case _ENABLE_SWAGGER: // enable the swagger endpoint for the ui
		router.PathPrefix("/swagger/").Handler(_mw.Then(http.StripPrefix("/swagger/", http.FileServer(http.Dir("/opt/vaultrdb/swagger")))))
		log.V(2).Info("enabled swagger ui", "relative-path", "http://localhost/swagger")

	case _ENABLE_UI: // enable the ui endpoint for the ui
		router.PathPrefix("/ui/").Handler(_mw.Then(http.StripPrefix("/ui/", http.FileServer(http.Dir("/opt/vaultrdb/web")))))
		log.V(2).Info("enabled ui", "relative-path", "http://localhost/ui/...")

	case _ENABLE_HEALTHZ: // enable the healthz endpoint for kubernetes
		router.Handle("/healthz/live", _mw.ThenFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("OK")) }))
		router.Handle("/healthz/ready", _mw.ThenFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("OK")) }))
		log.V(2).Info("enabled healthz endpoint for kubernetes", "relative-path--liveness", "http://localhost/healthz/live", "relative-path--readiness", "http://localhost/healthz/ready")
	}

}
