package main

import (
	"errors"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

func main() {

	logging.InitSLOG("Debug")

	logging.SLog.Info("starting injector sidecar for pod", "podname", "pod-xxxx-xxxx")

	logging.SLog.Error("starting vaultrdb sidecar not possible, no logic found, shutting down", errors.New("no sidecar logic").Error())
}
