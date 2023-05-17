package vaultset

import (
	"encoding/json"
	"net/http"

	"github.com/jnnkrdb/gomw/handlers/httpfnc"
	structs_v1 "github.com/jnnkrdb/vaultrdb/structs/v1"
)

func put(w http.ResponseWriter, r *http.Request) {
	var vaultset structs_v1.VaultSet
	if err := json.NewDecoder(r.Body).Decode(&vaultset); err != nil {
		httpfnc.AddErrorToHeaderIfAny(&w, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := vaultset.Update(); err != nil {
		httpfnc.AddErrorToHeaderIfAny(&w, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(vaultset.ID))
}
