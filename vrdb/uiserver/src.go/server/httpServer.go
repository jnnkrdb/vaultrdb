package server

import (
	"net/http"
	"os"

	"vrdb.go/http/mw"
	"vrdb.go/logging"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/gomw/middlewares"
)

// middleware used by the frontend http server
var DefaultMiddleware middlewares.MiddleWareChain

// starting the Server
func StartFrontendUI(fnc ...func(*mux.Router)) {

	// initializing the middleware for the http server
	logging.Log.V(1).Info("setting up the default middleware for http requests")

	DefaultMiddleware = middlewares.New(
		mw.QueryLog,
	)

	// the instance of the http server, serving the frontend files
	var router *mux.Router = mux.NewRouter().StrictSlash(true)

	// append the endpoints to the default router
	logging.Log.V(1).Info("creating the server and adding the required endpoints")

	for _, f := range fnc {

		f(router)
	}

	// booting the frontend http server
	logging.Log.V(1).Info("booting the server under port 80")

	if e := http.ListenAndServe(":80", router); e != nil {

		logging.Log.Error(e, "error keeping up http frontend server")

		os.Exit(1)
	}
}
