package main

import (
	"flag"
	"os"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/authstore"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/configstore"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/configstore/initialconfigs"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/vaultrdbstore"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server"
	v1 "github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/api/v1"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/healthz"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/metadata"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/ui"
	"github.com/jnnkrdb/vaultrdb/pkg/http/swagger"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
	"github.com/jnnkrdb/vaultrdb/pkg/termination"
)

// flags for service config
var (
	bootSwaggerUI bool = *flag.Bool("swaggerui", false, "If set, then the swagger ui will be started, with the configured port via --swagger-port.")
	swaggerPort   int  = *flag.Int("swagger-port", 81, "If set, then the swagger ui will be started, with the configured port.")
)

// list of termination funcs
var terminationFuncs []func()

func main() {

	flag.Parse()

	logging.InitSLOG("Debug")
	logging.SLog.Info("received flags", "arguments", os.Args, "non-flags", flag.Args())

	// initialize the needed stores
	configstore.InitDB()
	vaultrdbstore.InitDB()
	authstore.InitDB()

	terminationFuncs = append(terminationFuncs,
		func() { authstore.DB.CloseDB() },
		func() { configstore.DB.CloseDB() },
		func() { vaultrdbstore.DB.CloseDB() },
	)

	// set the initial configs, if not already set
	if err := initialconfigs.SetInitialConfigsIfNotConfiguredAlready(); err != nil {
		logging.SLog.Error("starting vaultrdb http backend async", "error", err.Error())
		os.Exit(1)
	}

	if bootSwaggerUI {
		logging.SLog.Info("starting swagger ui", "port", swaggerPort)
		if err := swagger.StartSwaggerServer("/opt/vaultrdb/swagger", swaggerPort); err != nil {
			logging.SLog.Error("error starting swagger ui", "error", err.Error())
		}
		terminationFuncs = append(terminationFuncs, swagger.StopSwaggerServer)
	}

	// set the termination methods
	termination.HandleTermination(terminationFuncs...)

	// starting the http server for vaultrdb
	logging.SLog.Info("starting vaultrdb http backend async")
	terminationFuncs = append(terminationFuncs, server.StopHTTP)
	if err := server.StartHTTP(
		healthz.EnableEndpoint_Healthz,
		metadata.EnableEndpoint_Metadata,
		ui.EnableEndpoint_UI,
		v1.EnableEndpoint_ApiV1,
	); err != nil {
		logging.SLog.Error("error keeping up the http server", "err", err.Error())
		termination.Shutdown()
	}
}
