package crud

import (
	"net/http"

	hndlrs "github.com/jnnkrdb/gomw/handlers"
	mw "github.com/jnnkrdb/gomw/middlewares"
)

var Routes = []hndlrs.HttpFunction{
	{
		Pattern: "/crud",
		MainHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("CRUDAPI"))
		}),
		Middlewares: mw.MiddleWareChain{},
	},
}
