package main

import (
	"fmt"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/conf"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

func main() {

	// set the default logger
	logging.Default = logging.GetLogger(conf.YC.Log.FormatJSON, conf.YC.Log.Level)

	logging.Default.Info("starting injector sidecar for pod", "podname", "pod-xxxx-xxxx")

	logging.Default.Error("starting vaultrdb sidecar not possible, no logic found, shutting down", "error", fmt.Errorf("no sidecar logic"))
}
