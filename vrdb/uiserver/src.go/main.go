package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	mw "github.com/jnnkrdb/gomw/middlewares"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

// config via args
var (
	_ENABLE_SWAGGER bool = false
	_ENABLE_LICENSE bool = true
	_ENABLE_VERSION bool = true
	_ENABLE_UI      bool = true
	_ENABLE_HEALTHZ bool = true
)

func main() {

	flag.BoolVar(&_ENABLE_SWAGGER, "swagger", false, "Enables the swagger ui.")

	var opts = zap.Options{
		Development: true,
	}

	opts.BindFlags(flag.CommandLine)

	var log = zap.New(zap.UseFlagOptions(&opts)).WithName("uiserver")

	log.V(1).Info("initializing http server for ui frontend")

	var router *mux.Router = mux.NewRouter().StrictSlash(true)

	// initialize the middlewares
	var _mw mw.MiddleWareChain = mw.New()

	// append the ui endpoints, which are configured using the args
	switch {

	case _ENABLE_VERSION: // enable the version endpoint for the ui
		router.Handle("/version", _mw.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "/opt/vaultrdb/config/VERSION")
		}))
		log.V(2).Info("enabled version ui", "relative-path", "http://localhost/version")

	case _ENABLE_LICENSE: // enable the license endpoint for the ui
		router.Handle("/license", _mw.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "/opt/vaultrdb/config/LICENSE")
		}))
		log.V(2).Info("enabled license ui", "relative-path", "http://localhost/license")

	case _ENABLE_SWAGGER: // enable the swagger endpoint for the ui
		router.PathPrefix("/swagger/").Handler(_mw.Then(http.StripPrefix("/swagger/", http.FileServer(http.Dir("/opt/vaultrdb/swagger")))))
		log.V(2).Info("enabled swagger ui", "relative-path", "http://localhost/swagger")

	case _ENABLE_UI: // enable the ui endpoint for the ui
		router.PathPrefix("/ui/").Handler(_mw.Then(http.StripPrefix("/ui/", http.FileServer(http.Dir("/opt/vaultrdb/web")))))
		log.V(2).Info("enabled ui", "relative-path", "http://localhost/ui/...")

	case _ENABLE_HEALTHZ: // enable the healthz endpoint for kubernetes
		router.Handle("/healthz/live", _mw.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		}))
		router.Handle("/healthz/ready", _mw.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		}))
		log.V(2).Info("enabled healthz endpoint for kubernetes", "relative-path--liveness", "http://localhost/healthz/live", "relative-path--readiness", "http://localhost/healthz/ready")
	}

	// start the http server
	if e := http.ListenAndServe(":80", router); e != nil {
		log.Error(e, "error keeping up http frontend server")
		os.Exit(1)
	}
}
