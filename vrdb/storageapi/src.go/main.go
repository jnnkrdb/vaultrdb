package main

import (
	"flag"
	"vrdb-storage/endpoints"
	v1 "vrdb-storage/endpoints/v1"
	"vrdb-storage/objects"
	"vrdb-storage/server"
	i "vrdb-storage/server/init"

	"vrdb.go/logging"
)

func main() {

	logging.InitLogger("storage-api")

	flag.Parse()

	server.ConnectToDatabase()

	objects.Migrate()

	i.CreateHASH()

	server.StartStorageAPI(
		endpoints.EnableEndpoint_Healthz,
		v1.EnableEndpoint_StorageAPI_V1,
	)
}
