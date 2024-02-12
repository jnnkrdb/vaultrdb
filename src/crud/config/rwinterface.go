package config

import "net/http"

type RWInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
}

func GetHandler(rwi RWInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			rwi.Read(w, r)
		case http.MethodPost:
			rwi.Create(w, r)
		case http.MethodPut:
			rwi.Update(w, r)
		case http.MethodDelete:
			rwi.Delete(w, r)
		default:
			http.Error(w, "method not implemented", http.StatusNotImplemented)
		}
	})
}
