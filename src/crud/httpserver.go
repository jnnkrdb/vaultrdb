package crud

import (
	"net/http"
	"os"

	hndlrs "github.com/jnnkrdb/gomw/handlers"
	mw "github.com/jnnkrdb/gomw/middlewares"

	"github.com/jnnkrdb/vaultrdb/crud/middlewares"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// configuration of the http server startup
var (
	ARG_SwaggerUI bool = false
	ARG_CrudAPI   bool = false
	ARG_UI        bool = false
	ARG_Auth      bool = false
)

type CRUDServer struct {
	Client client.Client
}

func (csrv *CRUDServer) Start() {

	if ARG_Auth ||
		ARG_CrudAPI ||
		ARG_SwaggerUI ||
		ARG_UI {

		go func() {

			var functionSet hndlrs.HttpFunctionSet
			var crudLog = ctrl.Log.WithName("crud")

			switch {

			case ARG_SwaggerUI:
				functionSet = append(functionSet, hndlrs.HttpFunction{
					Pattern:     "/swagger/",
					MainHandler: http.StripPrefix("/swagger/", http.FileServer(http.Dir("/vaultrdb/swagger"))),
					Middlewares: mw.New(middlewares.OptionsResponse),
				})

				crudLog.Info("enabled swagger ui", "uri", "localhost:80/swagger/")

			case ARG_UI:
				ARG_CrudAPI = true

				functionSet = append(functionSet, hndlrs.HttpFunction{
					Pattern:     "/ui/",
					MainHandler: http.StripPrefix("/ui/", http.FileServer(http.Dir("/vaultrdb/ui"))),
					Middlewares: mw.New(middlewares.OptionsResponse),
				})

				crudLog.Info("activated frontend ui", "uri", "localhost:80/ui/")
			case ARG_CrudAPI:

				crudLog.Info("activated crud api", "uri", "localhost:80/swagger/")
			case ARG_Auth:

				crudLog.Info("added auth endpoint", "uri", "localhost:80/swagger/")
			}

			if e := http.ListenAndServe(":80", hndlrs.GetHandler(functionSet)); e != nil {
				crudLog.Error(e, "error keeping up crud server")
				os.Exit(1)
			}
		}()
	}
}
