package swagger

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
	"github.com/rs/cors"
)

var swaggerserver *http.Server

func StartSwaggerServer(swaggerdir string, port int) error {

	logging.SLog.Info("booting swagger ui server")

	http.DefaultServeMux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir(swaggerdir))))
	swaggerserver = &http.Server{

		Addr: fmt.Sprintf(":%d", port),
		// adding the cors options
		Handler: cors.New(cors.Options{
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodOptions,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
			AllowedOrigins: []string{
				"*",
			},
			AllowCredentials: true,
		}).Handler(http.DefaultServeMux),
	}

	return swaggerserver.ListenAndServe()
}

func StopSwaggerServer() {
	logging.SLog.Info("shutting down the swagger server")
	if err := swaggerserver.Shutdown(context.TODO()); err != nil {
		logging.SLog.Error("error gracefully shutting down swagger http server", "error", err.Error())
	}
}
