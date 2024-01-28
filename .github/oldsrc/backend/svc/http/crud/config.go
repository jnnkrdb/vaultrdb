package crud

import (
	"net/http"

	hndlrs "github.com/jnnkrdb/gomw/handlers"
	mw "github.com/jnnkrdb/gomw/middlewares"
	"github.com/jnnkrdb/vaultrdb/svc/http/middlewares"
)

var Routes = []hndlrs.HttpFunction{
	{
		Pattern: "/crud",
		MainHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			switch {
			case r.Method == http.MethodGet:
				Read(w, r)
				return
			case r.Method == http.MethodPost:
				Create(w, r)
				return
			case r.Method == http.MethodPut || r.Method == http.MethodPatch:
				Update(w, r)
				return
			case r.Method == http.MethodDelete:
				Delete(w, r)
				return
			}

			http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		}),
		Middlewares: mw.New(
			middlewares.OptionsResponse,
			middlewares.BasicAuthentication,
		),
	},
}
