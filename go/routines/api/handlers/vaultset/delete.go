package vaultset

import (
	"fmt"
	"net/http"

	"github.com/jnnkrdb/gomw/handlers/httpfnc"
	structs_v1 "github.com/jnnkrdb/vaultrdb/structs/v1"
)

func delete(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("id") {
		httpfnc.AddErrorToHeaderIfAny(&w, fmt.Errorf("this request needs an id in the query"))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := structs_v1.Delete(r.URL.Query().Get("id")); err != nil {
		httpfnc.AddErrorToHeaderIfAny(&w, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		w.Write([]byte("OK"))
	}
}
