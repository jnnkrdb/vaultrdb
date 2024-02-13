package config

import (
	"net/http"
	"os"

	"github.com/go-logr/logr"
	"github.com/gorilla/mux"
	mw "github.com/jnnkrdb/gomw/middlewares"
	"github.com/jnnkrdb/vaultrdb/crud/middlewares"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	// default logger for http requests
	CrudLog logr.Logger = ctrl.Log.WithName("crud")

	// default client for requests against the kubernets
	// api server, used for vrdb structs, etc.
	KClient client.Client

	// rest router for http requests
	RESTSRV *mux.Router = mux.NewRouter()

	// default middlewares for the rest api
	DefaultMW = mw.New(middlewares.OptionsResponse)
)

// threadstart te rest api
func Start(c client.Client) {

	// set the global config for the kubernetes api client
	// to receive the information from the api server
	KClient = c

	// start the rest api in another go routine
	go func() {
		if e := http.ListenAndServe(":9080", RESTSRV); e != nil {
			CrudLog.Error(e, "error keeping up api server")
			os.Exit(1)
		}
	}()
}

// initialize the http rest router, to enable it
// with the default routes
func init() {

	// adding swagger ui to mux router
	// if environment variable is configured
	if val, ok := os.LookupEnv("ENABLE_SWAGGERUI"); ok && val == "true" {
		RESTSRV.Handle("/swagger/", DefaultMW.Then(http.StripPrefix("/swagger/", http.FileServer(http.Dir("/vaultrdb/swagger")))))
		CrudLog.Info("enabled swagger ui", "uri", "http://localhost:9080/swagger/")
	}

	// serving the license of the container image
	RESTSRV.Handle("/license", DefaultMW.ThenFunc(func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "/vaultrdb/LICENSE") }))
	CrudLog.Info("added license to http server", "uri", "http://localhost:9080/license")

	// serving the version of the container image
	RESTSRV.Handle("/version", DefaultMW.ThenFunc(func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "/vaultrdb/VERSION") }))
	CrudLog.Info("added version to http server", "uri", "http://localhost:9080/version")

	// serving the ui for the frontend
	RESTSRV.Handle("/ui/", DefaultMW.Then(http.StripPrefix("/ui/", http.FileServer(http.Dir("/vaultrdb/ui")))))
	CrudLog.Info("activated frontend ui", "uri", "localhost:9080/ui/")

}
