package main

import (
	"os"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/conf"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/authstore"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/configstore"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/configstore/initialconfigs"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/vaultrdbstore"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
	"github.com/jnnkrdb/vaultrdb/pkg/termination"
)

// list of termination funcs
var terminationFuncs []func()

func main() {

	// set the default logger
	logging.Default = logging.GetLogger(conf.YC.Log.FormatJSON, conf.YC.Log.Level)

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
		logging.Default.Error("starting vaultrdb http backend async", "error", err.Error())
		os.Exit(1)
	}

	// set the termination methods
	termination.HandleTermination(terminationFuncs...)

	// starting the http server for vaultrdb
	logging.Default.Info("starting vaultrdb http backend async")
	terminationFuncs = append(terminationFuncs, server.StopExternalAPI)

	if err := server.StartExternalAPI(); err != nil {
		logging.Default.Error("error keeping up the http server", "err", err.Error())
		termination.Shutdown()
	}
}
