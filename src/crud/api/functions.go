package api

import (
	hndlrs "github.com/jnnkrdb/gomw/handlers"
	mw "github.com/jnnkrdb/gomw/middlewares"
	v1 "github.com/jnnkrdb/vaultrdb/crud/api/v1"
	"github.com/jnnkrdb/vaultrdb/crud/config"
	"github.com/jnnkrdb/vaultrdb/crud/middlewares"
)

var defaultMiddlewares = mw.New(
	middlewares.OptionsResponse,
	middlewares.BasicAuth,
	middlewares.ReportRequest,
)

var ApiFunctionSet = hndlrs.HttpFunctionSet{
	{Pattern: "/crud/v1/vrdbconfig", MainHandler: config.GetHandler(&v1.VRDBConfigCRUD{}), Middlewares: defaultMiddlewares},
	{Pattern: "/crud/v1/vrdbconfiglist", MainHandler: config.GetHandler(&v1.VRDBConfigListCRUD{}), Middlewares: defaultMiddlewares},
	{Pattern: "/crud/v1/vrdbsecret", MainHandler: config.GetHandler(&v1.VRDBSecretCRUD{}), Middlewares: defaultMiddlewares},
	{Pattern: "/crud/v1/vrdbsecretlist", MainHandler: config.GetHandler(&v1.VRDBSecretListCRUD{}), Middlewares: defaultMiddlewares},
	{Pattern: "/crud/v1/namespacelist", MainHandler: config.GetHandler(&v1.NamespaceListCRUD{}), Middlewares: defaultMiddlewares},
}
