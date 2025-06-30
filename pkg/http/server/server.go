package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/gomw/middlewares"
	mw "github.com/jnnkrdb/vaultrdb/pkg/http/middlewares"
	"github.com/rs/cors"
)

type Server struct {
	httpsrv *http.Server
	router  *mux.Router

	// default middleware used by the endpoint
	mw middlewares.MiddleWareChain
}

func NewServer(port uint) Server {

	rtr := mux.NewRouter()

	var httpserver = &http.Server{
		Addr:                         fmt.Sprint(":%d", port),
		DisableGeneralOptionsHandler: false,
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
		}).Handler(rtr),
	}

	var server = Server{
		httpsrv: httpserver,
		router:  rtr,
		mw: middlewares.New(
			mw.QueryLog,
		),
	}
	return server
}

func (srv Server) GetRouter() *mux.Router {
	return srv.router
}

func (srv Server) GetMiddleware() middlewares.MiddleWareChain {
	return srv.mw
}

func (srv Server) Start() error {
	return srv.httpsrv.ListenAndServe()
}

func (srv Server) StartTLS(certFile string, keyFile string) error {
	return srv.httpsrv.ListenAndServeTLS(certFile, keyFile)
}

func (srv Server) Stop() error {
	return srv.httpsrv.Shutdown(context.TODO())
}
