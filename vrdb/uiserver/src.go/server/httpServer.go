package server

import (
	"fmt"
	"net/http"
	"os"

	"vrdb.go/http/mw"
	"vrdb.go/logging"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/gomw/middlewares"
	"github.com/rs/cors"
)

// middleware used by the frontend http server
var DefaultMiddleware middlewares.MiddleWareChain

// edfault prot of the http server
const _PORT int = 80

// starting the Server
func StartFrontendUI(fnc ...func(*mux.Router)) {

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
		Addr: fmt.Sprintf(":%d", _PORT),
		// adding the cors options
		Handler: cors.New(cors.Options{
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodOptions,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
			AllowedOrigins: []string{
				"*",
			},
			AllowCredentials: true,
		}).Handler(router),
	}).ListenAndServe(); e != nil {

		logging.Log.Error(e, "error keeping up http frontend server")

		os.Exit(1)
	}
}
