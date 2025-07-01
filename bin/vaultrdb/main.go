package main

import (
	"context"
	"os"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/conf"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/authstore"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/configstore"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/configstore/initialconfigs"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/externalapi"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
	"github.com/jnnkrdb/vaultrdb/pkg/termination"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// add the context to the termination handler
	termination.AsyncHandle(ctx)
	// testing with the own test handler for Close() error
	termination.AddHandlers(&termination.Test{})

	// set the default logger
	logging.Default = logging.GetLogger(conf.YC.Log.FormatJSON, conf.YC.Log.Level)

	// add logger to context
	logging.IntoContext(ctx, logging.Default)

	// initialize the needed stores
	configstore.InitDB()
	authstore.InitDB()

	termination.AddHandlers(conf.Authz, conf.Configs, conf.Vault)

	// set the initial configs, if not already set
	if err := initialconfigs.SetInitialConfigsIfNotConfiguredAlready(); err != nil {
		logging.Default.Error("starting vaultrdb http backend async", "error", err.Error())
		os.Exit(1)
	}

	// starting the http server for vaultrdb
	logging.Default.Info("starting vaultrdb http backend async")
	termination.AddHandlers(&externalapi.Server)

	if err := externalapi.Server.Start(); err != nil {
		logging.Default.Error("error keeping up the http server", "err", err.Error())
	}
}
