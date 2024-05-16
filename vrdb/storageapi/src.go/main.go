package main

import (
	"flag"
	"vrdb-storage/endpoints"
	v1 "vrdb-storage/endpoints/v1"
	"vrdb-storage/server"

	"vrdb.go/logging"
)

func main() {

	logging.InitLogger("storage-api")

	flag.Parse()

	server.ConnectToDatabase()

	server.StartStorageAPI(
		endpoints.EnableEndpoint_Healthz,
		v1.EnableEndpoint_StorageAPI_V1,
	)
}
