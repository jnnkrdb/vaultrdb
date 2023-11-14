package httpsrcv

import (
	"net/http"

	hndlrs "github.com/jnnkrdb/gomw/handlers"
	mw "github.com/jnnkrdb/gomw/middlewares"

	"github.com/jnnkrdb/vaultrdb/svc/http/frontend"
	"github.com/jnnkrdb/vaultrdb/svc/http/restapi"
)

func BootHTTP() {

	var funcArr = hndlrs.HttpFunctionSet{
		hndlrs.HttpFunction{
			Pattern: "/rest",
			MainHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("RestAPI"))
			}),
			Middlewares: mw.New(
				restapi.IsRESTApiEnabled,
			),
		},
		hndlrs.HttpFunction{
			Pattern: "/ui",
			MainHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Frontend"))
			}),
			Middlewares: mw.New(
				frontend.IsUIEnabled,
			),
		},
		hndlrs.HttpFunction{
			Pattern: "/auth",
			MainHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("AUTH"))
			}),
			Middlewares: mw.MiddleWareChain{},
		},
	}

	http.ListenAndServe(":80", hndlrs.GetHandler(funcArr))
}
