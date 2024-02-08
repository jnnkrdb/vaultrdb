package httpsrcv

import (
	"net/http"

	hndlrs "github.com/jnnkrdb/gomw/handlers"
	mw "github.com/jnnkrdb/gomw/middlewares"

	"github.com/go-logr/logr"

	"github.com/jnnkrdb/vaultrdb/svc/http/auth"
	"github.com/jnnkrdb/vaultrdb/svc/http/crud"
	"github.com/jnnkrdb/vaultrdb/svc/http/middlewares"
)

var (
	EnableSwaggerUI bool = false
	EnableRest      bool = false
	EnableUI        bool = false
	EnableAuth      bool = false
)

func BootHTTP(setupLog logr.Logger) {
	var funcArr hndlrs.HttpFunctionSet

	switch {

	case EnableSwaggerUI:
		funcArr = append(funcArr, hndlrs.HttpFunction{
			Pattern:     "/swaggerui/",
			MainHandler: http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("/usr/share/vaultrdb/swaggerui"))),
			Middlewares: mw.New(middlewares.OptionsResponse),
		})

	case EnableRest || EnableUI:
		funcArr = append(funcArr, crud.Routes...)

	case EnableUI:
		funcArr = append(funcArr, hndlrs.HttpFunction{
			Pattern:     "/ui/",
			MainHandler: http.StripPrefix("/ui/", http.FileServer(http.Dir("/usr/share/vaultrdb/frontend"))),
			Middlewares: mw.New(middlewares.OptionsResponse),
		})

	case EnableAuth || EnableUI || EnableRest:
		funcArr = append(funcArr, auth.Routes...)
	}

	if err := http.ListenAndServe(":80", hndlrs.GetHandler(funcArr)); err != nil {
		setupLog.Error(err, "failed keeping up http server")
	}
}
