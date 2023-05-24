package vaultset

import (
	"fmt"
	"net/http"

	"github.com/jnnkrdb/gomw/handlers/httpfnc"
	"github.com/jnnkrdb/vaultrdb/structs/v1/v1_vaultset"
)

func delete(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("id") {
		httpfnc.AddErrorToHeaderIfAny(&w, fmt.Errorf("this request needs an id in the query"))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := v1_vaultset.Delete(r.URL.Query().Get("id")); err != nil {
		httpfnc.AddErrorToHeaderIfAny(&w, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		w.Write([]byte("OK"))
	}
}
