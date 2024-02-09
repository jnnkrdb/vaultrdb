package auth

import (
	"net/http"

	hndlrs "github.com/jnnkrdb/gomw/handlers"
	mw "github.com/jnnkrdb/gomw/middlewares"
)

var Routes = []hndlrs.HttpFunction{
	{
		Pattern: "/auth/login",
		MainHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("AUTH-LOGIN"))
		}),
		Middlewares: mw.MiddleWareChain{},
	},
	{
		Pattern: "/auth/refresh",
		MainHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("AUTH-REFRESH"))
		}),
		Middlewares: mw.MiddleWareChain{},
	},
	{
		Pattern: "/auth/logout",
		MainHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("AUTH-LOGOUT"))
		}),
		Middlewares: mw.MiddleWareChain{},
	},
}
