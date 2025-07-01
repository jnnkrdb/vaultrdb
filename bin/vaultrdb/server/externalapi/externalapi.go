package externalapi

import (
	"net/http"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/conf"
	v1 "github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/externalapi/endpoints/api/v1"
	"github.com/jnnkrdb/vaultrdb/pkg/http/helpers"
	"github.com/jnnkrdb/vaultrdb/pkg/http/server"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// http server to handle the external api
var Server server.Server

// starting the http endpoint
func init() {

	logging.Default.Info("initiating the external api", "port", 80)
	Server = server.NewServer(80)

	// if requested, then boot up the swagger ui
	if conf.YC.Setup.ExternalAPI.SwaggerUI.Enabled {
		logging.Default.Debug("enabling swagger ui", "setup.externalapi.swaggerui.enabled", conf.YC.Setup.ExternalAPI.SwaggerUI.Enabled)
		Server.GetRouter().PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("/opt/vaultrdb/swagger"))))
	}

	// enable the ui frontend
	if conf.YC.Setup.ExternalAPI.FrontendUI.Enabled {
		Server.GetRouter().PathPrefix("/").Handler(http.FileServer(http.Dir("/opt/vaultrdb/web")))
	}

	// healthz api
	logging.Default.Info("creating healthz endpoints", "uri-path", "/healthz/<req>")
	Server.GetRouter().Path("/healthz/live").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("OK")) })
	Server.GetRouter().Path("/healthz/ready").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("OK")) })

	// metadata api
	logging.Default.Info("creating metadata endpoints", "uri-path", "/meta/<metadata>")
	Server.GetRouter().Path("/meta/version").Methods(http.MethodGet).Handler(Server.GetMiddleware().ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		helpers.FilesContent(w, r, "/opt/vaultrdb/home/VERSION")
	}))
	Server.GetRouter().Path("/meta/license").Methods(http.MethodGet).Handler(Server.GetMiddleware().ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		helpers.FilesContent(w, r, "/opt/vaultrdb/home/LICENSE")
	}))

	v1.EnableEndpoint_ApiV1(Server.GetRouter(), Server.GetMiddleware())
}
