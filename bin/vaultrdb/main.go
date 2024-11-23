package main

import (
	"os"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/configstore"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/database"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server"
	v1 "github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/api/v1"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/healthz"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/metadata"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/swagger"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/ui"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

func main() {

	logging.InitSLOG("Debug")

	configstore.InitConfigStore()

	database.Connect()

	// starting the http server for vaultrdb
	logging.SLog.Info("starting vaultrdb http backend async")
	if err := server.StartHTTP(
		healthz.EnableEndpoint_Healthz,
		metadata.EnableEndpoint_Metadata,
		swagger.EnableEndpoint_Swagger,
		ui.EnableEndpoint_UI,

		v1.EnableEndpoint_ApiV1,
	); err != nil {
		logging.SLog.Error("error keeping up the http server", err)
		os.Exit(1)
	}
}
