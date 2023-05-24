package vaultset

import (
	"encoding/json"
	"net/http"

	"github.com/jnnkrdb/gomw/handlers/httpfnc"
	"github.com/jnnkrdb/vaultrdb/structs/v1/v1_vaultset"
)

func post(w http.ResponseWriter, r *http.Request) {
	var vaultset v1_vaultset.VaultSet
	if err := json.NewDecoder(r.Body).Decode(&vaultset); err != nil {
		httpfnc.AddErrorToHeaderIfAny(&w, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := v1_vaultset.Insert(&vaultset); err != nil {
		httpfnc.AddErrorToHeaderIfAny(&w, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(vaultset.ID))
}
