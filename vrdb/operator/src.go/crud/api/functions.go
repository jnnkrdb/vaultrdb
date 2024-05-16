package api

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	mw "github.com/jnnkrdb/gomw/middlewares"
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

	// serving the license of the container image
	RESTSRV.Handle("/crud/alive", DefaultMW.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))
	config.CrudLog.Info("added health startup check to http server", "uri", "http://localhost:9080/crud/alive")

	//	// generic functions listing
	//	RESTSRV.Handle("/crud/v1/{kind}", mw.ThenFunc(v1.LIST_ALL)).Methods("GET", "OPTIONS")
	//	RESTSRV.Handle("/crud/v1/{kind}/{namespace}", mw.ThenFunc(v1.LIST_NAMESPACE)).Methods("GET", "OPTIONS")
	//
	//	// generic functions for one object
	//	RESTSRV.Handle("/crud/v1/{kind}/{namespace}/{name}", mw.ThenFunc(v1.READ)).Methods("GET", "OPTIONS")
	//	RESTSRV.Handle("/crud/v1/{kind}/{namespace}/{name}", mw.ThenFunc(v1.CREATE)).Methods("POST", "OPTIONS")
	//	RESTSRV.Handle("/crud/v1/{kind}/{namespace}/{name}", mw.ThenFunc(v1.UPDATE)).Methods("PUT", "PATCH", "OPTIONS")
	//	RESTSRV.Handle("/crud/v1/{kind}/{namespace}/{name}", mw.ThenFunc(v1.DELETE)).Methods("DELETE", "OPTIONS")

	// start the rest api in another go routine
	go func() {
		if e := http.ListenAndServe(":9080", RESTSRV); e != nil {
			config.CrudLog.Error(e, "error keeping up api server")
			os.Exit(1)
		}
	}()
}
