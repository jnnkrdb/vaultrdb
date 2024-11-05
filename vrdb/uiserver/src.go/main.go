package main

import (
	"flag"

	"vrdb-uiserver/endpoints"
	"vrdb-uiserver/server"

	"vrdb.go/logging"
)

func main() {

	flag.BoolVar(&endpoints.ENABLE_SWAGGER, "swagger", false, "Enables the swagger ui.")
	flag.StringVar(&endpoints.STORAGEAPISERVER, "storageapi-address", "localhost:8080", "Sets the address of the storage api, without the scheme. Port only if neccessary.")

	logging.InitLogger("ui-server")

	flag.Parse()

	server.StartFrontendUI(
		endpoints.EnableEndpoint_Metadata,
		endpoints.EnableEndpoint_SwaggerUI,
		endpoints.EnableEndpoint_ApiProxy,
		endpoints.EnableEndpoint_UI,
		endpoints.EnableEndpoint_Healthz,
	)
}
