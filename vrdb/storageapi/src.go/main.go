package main

import (
	"flag"
	"vrdb-storage/endpoints"
	v1 "vrdb-storage/endpoints/v1"
	"vrdb-storage/objects"
	"vrdb-storage/server"
	"vrdb-storage/server/configs"

	"vrdb.go/logging"
)

func main() {

	logging.InitLogger("storage-api")

	flag.Parse()

	server.ConnectToDatabase()

	objects.Migrate()

	configs.InitInternalDB()

	server.StartStorageAPI(
		endpoints.EnableEndpoint_Healthz,
		v1.EnableEndpoint_StorageAPI_V1,
	)

	configs.DB.Close()
}
