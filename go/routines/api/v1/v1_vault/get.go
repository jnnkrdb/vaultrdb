package v1_vault

import (
	"encoding/json"
	"net/http"

	"github.com/jnnkrdb/gomw/handlers/httpfnc"
	"github.com/jnnkrdb/vaultrdb/structs/v1/v1_vaultset"
)

func get(w http.ResponseWriter, r *http.Request) {

	// select by id
	if r.URL.Query().Has("id") {
		if result, err := v1_vaultset.SelectByID(r.URL.Query().Get("id")); err != nil {
			httpfnc.AddErrorToHeaderIfAny(&w, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		} else {
			if err = json.NewEncoder(w).Encode(result); err != nil {
				httpfnc.AddErrorToHeaderIfAny(&w, err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}
		return
	}

	if result, err := v1_vaultset.Select(); err != nil {
		httpfnc.AddErrorToHeaderIfAny(&w, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	} else {
		if err = json.NewEncoder(w).Encode(result); err != nil {
			httpfnc.AddErrorToHeaderIfAny(&w, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}
