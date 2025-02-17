package server

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/gomw/middlewares"
	mw "github.com/jnnkrdb/vaultrdb/pkg/http/middlewares"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
	"github.com/rs/cors"
)

// middleware used by the frontend http server
var DefaultMiddleware middlewares.MiddleWareChain

// default prot of the http server
const _PORT int = 80

var _SRV *http.Server

var (
	bootSwaggerUI *bool = flag.Bool("swaggerui", false, "If set, then the swagger ui will be activated.")
)

// starting the http endpoint
func StartHTTP(fnc ...func(*mux.Router)) error {

	// initializing the middleware for the http server
	logging.SLog.Info("defining default middlewares for http endpoints")

	DefaultMiddleware = middlewares.New(
		mw.QueryLog,
	)

	// the instance of the http server, serving the frontend files
	var router *mux.Router = mux.NewRouter()

	// append the endpoints to the default router
	logging.SLog.Info("creating the server and adding the required endpoints")

	for _, f := range fnc {

		f(router)
	}

	if *bootSwaggerUI {
		logging.SLog.Info("starting swagger ui")
		router.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("/opt/vaultrdb/swagger"))))
	}

	// booting the frontend http server
	logging.SLog.Info("booting the server", "port", _PORT)

	_SRV = &http.Server{
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
	}

	return _SRV.ListenAndServe()
}

// stop the http backend server
func StopHTTP() {
	logging.SLog.Info("shutting down the server")
	if err := _SRV.Shutdown(context.TODO()); err != nil {
		logging.SLog.Error("error gracefully shutting down http server", "error", err.Error())
	}
}
