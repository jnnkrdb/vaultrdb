package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/gomw/middlewares"
	"vrdb.go/http/mw"
	"vrdb.go/logging"
)

// middleware used by the frontend http server
var DefaultMiddleware middlewares.MiddleWareChain

// edfault prot of the http server
const _PORT int = 80

// starting the Server
func StartStorageAPI(fnc ...func(*mux.Router)) {

	// initializing the middleware for the http server
	logging.Log.Info("setting up the default middleware for http requests")

	DefaultMiddleware = middlewares.New(
		mw.QueryLog,
	)

	// the instance of the http server, serving the frontend files
	var router *mux.Router = mux.NewRouter().StrictSlash(true)

	// append the endpoints to the default router
	logging.Log.Info("creating the server and adding the required endpoints")

	for _, f := range fnc {

		f(router)
	}

	// booting the frontend http server
	logging.Log.Info("booting the server", "port", _PORT)

	if e := (&http.Server{
		Addr:    fmt.Sprintf(":%d", _PORT),
		Handler: router,
	}).ListenAndServe(); e != nil {

		logging.Log.Error(e, "error keeping up http storage api server")

		os.Exit(1)
	}
}
