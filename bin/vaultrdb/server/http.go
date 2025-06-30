package server

import (
	"net/http"

	"github.com/jnnkrdb/gomw/middlewares"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/conf"
	"github.com/jnnkrdb/vaultrdb/pkg/http/server"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// middleware used by the frontend http server
var DefaultMiddleware middlewares.MiddleWareChain

// http server to handle the external api
var ExternalAPI server.Server

// starting the http endpoint
func StartExternalAPI() error {

	logging.Default.Info("initiating the external api", "port", 80)
	ExternalAPI = server.NewServer(80)

	// if requested, then boot up the swagger ui
	if conf.YC.Setup.ExternalAPI.SwaggerUI.Enabled {
		logging.Default.Debug("enabling swagger ui", "setup.externalapi.swaggerui.enabled", conf.YC.Setup.ExternalAPI.SwaggerUI.Enabled)
		ExternalAPI.GetRouter().PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("/opt/vaultrdb/swagger"))))
	}

	// enable the ui frontend
	if conf.YC.Setup.ExternalAPI.FrontendUI.Enabled {
		ExternalAPI.GetRouter().PathPrefix("/").Handler(http.FileServer(http.Dir("/opt/vaultrdb/web")))
	}

	return ExternalAPI.Start()
}

// stop the http backend server
func StopExternalAPI() {
	logging.Default.Info("shutting down the server")
	if err := ExternalAPI.Stop(); err != nil {
		logging.Default.Error("error gracefully shutting down http server", "error", err.Error())
	}
}
