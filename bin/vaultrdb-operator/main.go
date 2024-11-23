package main

import (
	"errors"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

func main() {

	logging.InitSLOG("Debug")

	// starting vaultrdb operator
	logging.SLog.Info("starting vaultrdb operator")

	logging.SLog.Error("starting vaultrdb operator not possible, no logic found, shutting down", errors.New("no operator logic"))
}
