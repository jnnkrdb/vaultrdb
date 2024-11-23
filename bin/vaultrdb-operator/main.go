package main

import (
	"errors"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

func main() {

	logging.InitLogger("vaultrdb-operator")

	// starting vaultrdb operator
	logging.Log.Info("starting vaultrdb operator")

	logging.Log.Error(errors.New("no operator logic"), "starting vaultrdb operator not possible, no logic found, shutting down")
}
