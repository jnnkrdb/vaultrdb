package crud

import (
	"net/http"
	"os"

	hndlrs "github.com/jnnkrdb/gomw/handlers"
	mw "github.com/jnnkrdb/gomw/middlewares"

	"github.com/jnnkrdb/vaultrdb/crud/api"
	"github.com/jnnkrdb/vaultrdb/crud/config"
	"github.com/jnnkrdb/vaultrdb/crud/middlewares"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CRUDServer struct {
	Client client.Client
}

func (csrv *CRUDServer) Start() {

	go func() {

		var functionSet hndlrs.HttpFunctionSet
		var crudLog = ctrl.Log.WithName("crud")

		// if the swaggerui environmentvariable is set to true, the swagger ui will
		// be activated and can be accessed via http://localhost:9080/swagger/
		if val, ok := os.LookupEnv("ENABLE_SWAGGERUI"); ok && val == "true" {
			functionSet = append(functionSet, hndlrs.HttpFunction{
				Pattern:     "/swagger/",
				MainHandler: http.StripPrefix("/swagger/", http.FileServer(http.Dir("/vaultrdb/swagger"))),
				Middlewares: mw.New(middlewares.OptionsResponse),
			})
			crudLog.Info("enabled swagger ui", "uri", "http://localhost:9080/swagger/")
		}

		// add the license to the http webserver
		functionSet = append(functionSet, hndlrs.HttpFunction{
			Pattern: "/license",
			MainHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, "/vaultrdb/LICENSE")
			}),
			Middlewares: mw.New(middlewares.OptionsResponse),
		})
		crudLog.Info("added license to http server", "uri", "http://localhost:9080/license")

		// add the version to the http webserver
		functionSet = append(functionSet, hndlrs.HttpFunction{
			Pattern: "/version",
			MainHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, "/vaultrdb/VERSION")
			}),
			Middlewares: mw.New(middlewares.OptionsResponse),
		})
		crudLog.Info("added version to http server", "uri", "http://localhost:9080/version")

		// add the ui frontend to the httpserver
		functionSet = append(functionSet, hndlrs.HttpFunction{
			Pattern:     "/ui/",
			MainHandler: http.StripPrefix("/ui/", http.FileServer(http.Dir("/vaultrdb/ui"))),
			Middlewares: mw.New(middlewares.OptionsResponse),
		})
		crudLog.Info("activated frontend ui", "uri", "localhost:9080/ui/")

		// append the api functions to the http server with middleware
		// configurations
		functionSet = append(functionSet, api.ApiFunctionSet...)
		crudLog.Info("activated crud api", "uri", "localhost:9080/swagger/")

		// set the global config for the kubernetes api client
		// to receive the information from the api server
		config.KClient = csrv.Client

		if e := http.ListenAndServe(":9080", hndlrs.GetHandler(functionSet)); e != nil {
			crudLog.Error(e, "error keeping up crud server")
			os.Exit(1)
		}
	}()
}
