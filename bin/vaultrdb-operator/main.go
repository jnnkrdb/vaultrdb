package main

import (
	"fmt"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/conf"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

func main() {

	// set the default logger
	logging.Default = logging.GetLogger(conf.YC.Log.FormatJSON, conf.YC.Log.Level)

	// starting vaultrdb operator
	logging.Default.Info("starting vaultrdb operator")

	logging.Default.Error("starting vaultrdb operator not possible, no logic found, shutting down", "error", fmt.Errorf("no operator logic"))
}
