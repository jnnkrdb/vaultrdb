package main

import (
	"fmt"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

func main() {

	logging.InitSLOG("Debug")

	logging.SLog.Info("starting injector sidecar for pod", "podname", "pod-xxxx-xxxx")

	logging.SLog.Error("starting vaultrdb sidecar not possible, no logic found, shutting down", "error", fmt.Errorf("no sidecar logic"))
}
