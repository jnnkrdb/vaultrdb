package api

import (
	"log"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/routines/api/handlers"
	"github.com/jnnkrdb/vaultrdb/routines/api/healthz"

	hndlrs "github.com/jnnkrdb/gomw/handlers"
	mw "github.com/jnnkrdb/gomw/middlewares"
	"github.com/jnnkrdb/gomw/middlewares/security/cors"
)

// default http api port is 8080
func HandleAPI() {
	// checking for errors
	if err := (&http.Server{
		Addr:    ":80",
		Handler: hndlrs.GetHandler(httpHandlers),
	}).ListenAndServe(); err != nil {
		log.Panicf("%#v\n", err)
	}
}

var httpHandlers = hndlrs.HttpFunctionSet{
	// get a new uuidv4
	{Pattern: "/api/v1/uuidv4", MainHandler: http.HandlerFunc(handlers.UUIDv4), Middlewares: mw.New(cors.AddCORSHeaders)},

	// encrypt/decrypt functions
	{Pattern: "/api/v1/encrypt", MainHandler: http.HandlerFunc(handlers.Encrypt), Middlewares: mw.New(cors.AddCORSHeaders)},
	{Pattern: "/api/v1/decrypt", MainHandler: http.HandlerFunc(handlers.Decrypt), Middlewares: mw.New(cors.AddCORSHeaders)},

	// host ui
	{Pattern: "/ui", MainHandler: http.StripPrefix("/ui", http.FileServer(http.Dir("/app/ui"))), Middlewares: mw.New(cors.AddCORSHeaders)},

	// healthz checks
	{Pattern: "/healthz/live", MainHandler: http.HandlerFunc(healthz.Liveness), Middlewares: mw.MiddleWareChain{}},
	{Pattern: "/healthz/ready", MainHandler: http.HandlerFunc(healthz.Readiness), Middlewares: mw.MiddleWareChain{}},
	{Pattern: "/metrics", MainHandler: http.HandlerFunc(healthz.Metrics), Middlewares: mw.MiddleWareChain{}},
}
