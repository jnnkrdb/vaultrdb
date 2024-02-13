package api

import (
	v1 "github.com/jnnkrdb/vaultrdb/crud/api/v1"
	"github.com/jnnkrdb/vaultrdb/crud/config"
	"github.com/jnnkrdb/vaultrdb/crud/middlewares"
)

func init() {
	// append basic auth handler to middlewares
	var mw = config.DefaultMW.Append(middlewares.BasicAuth)

	// generic functions listing
	config.RESTSRV.Handle("/crud/v1/{kind}", mw.ThenFunc(v1.LIST_ALL)).Methods("GET")
	config.RESTSRV.Handle("/crud/v1/{kind}/{namespace}", mw.ThenFunc(v1.LIST_NAMESPACE)).Methods("GET")

	// generic functions for one object
	config.RESTSRV.Handle("/crud/v1/{kind}/{namespace}/{name}", mw.ThenFunc(v1.READ)).Methods("GET")
	config.RESTSRV.Handle("/crud/v1/{kind}/{namespace}/{name}", mw.ThenFunc(v1.CREATE)).Methods("POST")
	config.RESTSRV.Handle("/crud/v1/{kind}/{namespace}/{name}", mw.ThenFunc(v1.UPDATE)).Methods("PUT", "PATCH")
	config.RESTSRV.Handle("/crud/v1/{kind}/{namespace}/{name}", mw.ThenFunc(v1.DELETE)).Methods("DELETE")
}
