package main

import "github.com/jnnkrdb/vaultrdb/pkg/logging"

func main() {

	logging.InitLogger("vaultrdb-injector")

	logging.Log.Info("starting injector sidecar for pod", "podname", "pod-xxxx-xxxx")

}
