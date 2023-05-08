package api

import (
	"net/http"

	"github.com/jnnkrdb/vaultrdb/routines/api/handlers"
	"github.com/jnnkrdb/vaultrdb/routines/api/healthz"

	hndlrs "github.com/jnnkrdb/gomw/handlers"
	mw "github.com/jnnkrdb/gomw/middlewares"
	"github.com/jnnkrdb/gomw/middlewares/security/authorization/apikey"
	"github.com/jnnkrdb/gomw/middlewares/security/cors"
)

var DefaultMW = mw.New(
	cors.AddCORSHeaders,
	apikey.APIKeyCheck,
)

var httpHandlers = hndlrs.HttpFunctionSet{

	// encrypt/decrypt functions
	{Pattern: "/api/v1/encrypt", MainHandler: http.HandlerFunc(handlers.Encrypt), Middlewares: mw.New(cors.AddCORSHeaders)},
	{Pattern: "/api/v1/decrypt", MainHandler: http.HandlerFunc(handlers.Decrypt), Middlewares: mw.New(cors.AddCORSHeaders)},

	// healthz checks
	{Pattern: "/healthz/live", MainHandler: http.HandlerFunc(healthz.Liveness), Middlewares: mw.MiddleWareChain{}},
	{Pattern: "/healthz/ready", MainHandler: http.HandlerFunc(healthz.Readyness), Middlewares: mw.MiddleWareChain{}},
	{Pattern: "/metrics", MainHandler: http.HandlerFunc(healthz.Metrics), Middlewares: mw.MiddleWareChain{}},
}
