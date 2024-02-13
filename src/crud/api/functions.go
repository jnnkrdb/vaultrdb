package api

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	mw "github.com/jnnkrdb/gomw/middlewares"
	v1 "github.com/jnnkrdb/vaultrdb/crud/api/v1"
	"github.com/jnnkrdb/vaultrdb/crud/config"
	"github.com/jnnkrdb/vaultrdb/crud/middlewares"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (

	// rest router for http requests
	RESTSRV *mux.Router = mux.NewRouter().StrictSlash(true)

	// default middlewares for the rest api
	DefaultMW = mw.New(middlewares.OptionsResponse)
)

// threadstart te rest api
func Start(c client.Client) {

	// set logger
	config.CrudLog = ctrl.Log.WithName("crud")

	// set the global config for the kubernetes api client
	// to receive the information from the api server
	config.KClient = c

	// adding swagger ui to mux router
	// if environment variable is configured
	if val, ok := os.LookupEnv("ENABLE_SWAGGERUI"); ok && val == "true" {
		//RESTSRV.PathPrefix("/swagger/").Handler(DefaultMW.Then(http.StripPrefix("/swagger/", http.FileServer(http.Dir("/vaultrdb/swagger")))))
		RESTSRV.Handle("/swagger/", DefaultMW.Then(http.StripPrefix("/swagger/", http.FileServer(http.Dir("/vaultrdb/swagger")))))
		config.CrudLog.Info("enabled swagger ui", "uri", "http://localhost:9080/swagger/")
	}

	// serving the license of the container image
	RESTSRV.Handle("/license", DefaultMW.ThenFunc(func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "/vaultrdb/LICENSE") }))
	config.CrudLog.Info("added license to http server", "uri", "http://localhost:9080/license")

	// serving the version of the container image
	RESTSRV.Handle("/version", DefaultMW.ThenFunc(func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "/vaultrdb/VERSION") }))
	config.CrudLog.Info("added version to http server", "uri", "http://localhost:9080/version")

	// serving the ui for the frontend
	RESTSRV.Handle("/ui/", DefaultMW.Then(http.StripPrefix("/ui/", http.FileServer(http.Dir("/vaultrdb/ui")))))
	config.CrudLog.Info("activated frontend ui", "uri", "localhost:9080/ui/")

	// append basic auth handler to middlewares
	var mw = DefaultMW.Append(middlewares.BasicAuth)

	// generic functions listing
	RESTSRV.Handle("/crud/v1/{kind}", mw.ThenFunc(v1.LIST_ALL)).Methods("GET")
	RESTSRV.Handle("/crud/v1/{kind}/{namespace}", mw.ThenFunc(v1.LIST_NAMESPACE)).Methods("GET")

	// generic functions for one object
	RESTSRV.Handle("/crud/v1/{kind}/{namespace}/{name}", mw.ThenFunc(v1.READ)).Methods("GET")
	RESTSRV.Handle("/crud/v1/{kind}/{namespace}/{name}", mw.ThenFunc(v1.CREATE)).Methods("POST")
	RESTSRV.Handle("/crud/v1/{kind}/{namespace}/{name}", mw.ThenFunc(v1.UPDATE)).Methods("PUT", "PATCH")
	RESTSRV.Handle("/crud/v1/{kind}/{namespace}/{name}", mw.ThenFunc(v1.DELETE)).Methods("DELETE")

	// start the rest api in another go routine
	go func() {
		if e := http.ListenAndServe(":9080", RESTSRV); e != nil {
			config.CrudLog.Error(e, "error keeping up api server")
			os.Exit(1)
		}
	}()
}
